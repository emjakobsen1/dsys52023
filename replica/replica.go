package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"strconv"
	"sync"
	"time"

	proto "github.com/emjakobsen1/dsys52023/proto"
	"google.golang.org/grpc"
)

type replicaServer struct {
	proto.UnimplementedAuctionServiceServer
	id              int32
	bidders         map[int32]int32
	mutex           sync.Mutex
	auctionDuration time.Time
	started         bool
}

var l, sl *log.Logger

// First call registers the bidder and starts the auction.
func (r *replicaServer) Bid2Replicas(ctx context.Context, req *proto.Amount) (*proto.Ack, error) {

	r.mutex.Lock()
	defer r.mutex.Unlock()

	if !r.started {
		r.auctionDuration = time.Now().Add(1 * time.Minute)
		r.started = true
	}

	if time.Until(r.auctionDuration) <= 0 {
		return &proto.Ack{State: "FAIL", ReqId: req.Id, SendSeq: req.SendSeq}, nil
	}

	l.Printf("Bid: %d from %d (sendSeq %d)\n", req.Amount, req.Id, req.SendSeq)
	sl.Printf("Bid: %d from %d (sendSeq %d)\n", req.Amount, req.Id, req.SendSeq)

	if req.Amount > findHighestBid(r.bidders) {
		r.bidders[req.Id] = int32(req.Amount)
		return &proto.Ack{State: "SUCCESS", ReqId: req.Id, SendSeq: req.SendSeq}, nil
	}

	return &proto.Ack{State: "FAIL", ReqId: req.Id, SendSeq: req.SendSeq}, nil
}

func (r *replicaServer) Result2Replicas(ctx context.Context, req *proto.Void) (*proto.Outcome, error) {

	r.mutex.Lock()
	defer r.mutex.Unlock()

	highestBid := findHighestBid(r.bidders)

	l.Printf("Result from %d (sendSeq %d)\n", req.Id, req.SendSeq)
	sl.Printf("Result from %d (sendSeq %d)\n", req.Id, req.SendSeq)
	rep := &proto.Outcome{HighestBid: highestBid, ReqId: req.Id, SendSeq: req.SendSeq}
	return rep, nil
}

func findHighestBid(bidders map[int32]int32) int32 {

	var highestBid int32

	for _, bid := range bidders {
		if bid > highestBid {
			highestBid = bid
		}
	}

	return highestBid
}

func main() {

	port, err := strconv.Atoi(os.Args[1])
	if err != nil {
		log.Fatalf("Invalid port: %v", err)
	}

	list, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		log.Fatalf("Failed to listen on port %d: %v", port, err)
	}

	//replica's log, shared log
	l, sl = setLog(port)

	grpcServer := grpc.NewServer()
	r := &replicaServer{
		id:      int32(port),
		bidders: make(map[int32]int32),
		started: false,
	}

	proto.RegisterAuctionServiceServer(grpcServer, r)

	if err := grpcServer.Serve(list); err != nil {
		log.Fatalf("Failed to serve on port %d: %v", port, err)
	}
}

func setLog(id int) (*log.Logger, *log.Logger) {

	clearLog(fmt.Sprintf("../logs/replica-%v.txt", id))

	replicaLogFile, err := os.OpenFile(fmt.Sprintf("../logs/replica-%v.txt", id), os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("Error opening replica log file: %v", err)
	}

	sharedLogFile, err := os.OpenFile("../logs/combined.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("Error opening shared log file: %v", err)
	}

	replicaLogger := log.New(io.MultiWriter(os.Stdout, replicaLogFile), fmt.Sprintf("Replica (%v): ", id), log.Flags())
	sharedLogger := log.New(sharedLogFile, fmt.Sprintf("Replica (%v): ", id), log.Flags())
	replicaLogger.SetFlags(0)
	sharedLogger.SetFlags(0)

	return replicaLogger, sharedLogger
}

func clearLog(filename string) {
	if err := os.Truncate(filename, 0); err != nil {
		log.Printf("Failed to truncate log file %v: %v\n", filename, err)
	}
}
