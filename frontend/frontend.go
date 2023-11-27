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

	proto "github.com/emjakobsen1/dsys52023/proto"
	"google.golang.org/grpc"
)

var ID int
var l, sl *log.Logger

type frontendServer struct {
	proto.UnimplementedAuctionServiceServer
	replicas      map[int32]proto.AuctionServiceClient
	sendSeq       int32
	delivered     map[int32]int
	ackbuffer     []proto.Ack
	outcomebuffer []proto.Outcome
	mutex         sync.Mutex
}

func main() {

	var err error
	ID, err = strconv.Atoi(os.Args[1])
	if err != nil {
		log.Fatalf("Invalid client ID: %v", err)
	}
	// frontend's log, shared log
	l, sl = setLog(ID)

	replicaPorts := []int32{5000, 5001, 5002}

	replicaClients := make(map[int32]proto.AuctionServiceClient)

	for _, port := range replicaPorts {
		conn, err := grpc.Dial(fmt.Sprintf("localhost:%d", port), grpc.WithInsecure())
		if err != nil {
			l.Fatalf("Could not connect to replica server on port %d: %v", port, err)
		}
		defer conn.Close()

		replicaClients[port] = proto.NewAuctionServiceClient(conn)
	}

	frontend := &frontendServer{
		replicas:      replicaClients,
		sendSeq:       0,
		delivered:     make(map[int32]int),
		ackbuffer:     []proto.Ack{},
		outcomebuffer: []proto.Outcome{},
	}

	grpcServer := grpc.NewServer()
	proto.RegisterAuctionServiceServer(grpcServer, frontend)

	listener, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", ID))
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}

}
func (s *frontendServer) Bid(ctx context.Context, req *proto.Amount) (*proto.Ack, error) {

	req.SendSeq = int32(s.sendSeq)
	rep, err := s.BroadcastBid(ctx, req)

	if err != nil {
		l.Printf("Failed to forward request to replica: %v\n", err)
		sl.Printf("Failed to forward request to replica: %v\n", err)
	}

	s.sendSeq++

	return rep, nil

}

func (s *frontendServer) BroadcastBid(ctx context.Context, req *proto.Amount) (*proto.Ack, error) {

	var ack *proto.Ack
	for _, replicaClient := range s.replicas {
		rep, err := replicaClient.Bid2Replicas(ctx, req)
		if err != nil {
			l.Printf("Failed to forward request to replica: %v\n", err)
			sl.Printf("Failed to forward request to replica: %v\n", err)
			continue
		}
		s.ackbuffer = append(s.ackbuffer, *rep)
		ack = rep
	}

	var replies = false
	for ids, msg := range s.ackbuffer {
		if s.sendSeq == msg.SendSeq {
			replies = true
			l.Printf("Ack: %s req by %d\n", msg.State, msg.ReqId)
			sl.Printf("Ack: %s req by %d\n", msg.State, msg.ReqId)
			s.delivered[int32(ids)]++
		}
	}

	if replies {
		return ack, nil
	}

	// If the above failed, something went wrong reply with exception.
	return &proto.Ack{State: "EXCEPTION"}, nil

}

func (s *frontendServer) Result(ctx context.Context, req *proto.Void) (*proto.Outcome, error) {

	req.SendSeq = int32(s.sendSeq)
	rep, err := s.BroadcastResult(ctx, req)

	if err != nil {
		l.Printf("Failed to forward request to replica: %v\n", err)
		sl.Printf("Failed to forward request to replica: %v\n", err)
	}

	s.sendSeq++

	return rep, nil

}

func (s *frontendServer) BroadcastResult(ctx context.Context, req *proto.Void) (*proto.Outcome, error) {

	var outcome *proto.Outcome
	for _, replicaClient := range s.replicas {
		rep, err := replicaClient.Result2Replicas(ctx, req)
		if err != nil {
			l.Printf("Failed to forward request to replica: %v\n", err)
			sl.Printf("Failed to forward request to replica: %v\n", err)
			continue
		}
		s.outcomebuffer = append(s.outcomebuffer, *rep)
		outcome = rep
	}

	var replies = false
	for ids, msg := range s.outcomebuffer {
		if s.sendSeq == msg.SendSeq {
			replies = true
			l.Printf("Outcome: %d req by: %d \n", msg.HighestBid, msg.ReqId)
			sl.Printf("Outcome: %d req by: %d \n", msg.HighestBid, msg.ReqId)
			s.delivered[int32(ids)]++

		}
	}

	if replies {
		return outcome, nil
	}

	return &proto.Outcome{}, nil
}

func setLog(id int) (*log.Logger, *log.Logger) {

	clearLog(fmt.Sprintf("../logs/frontend-%v.txt", id))

	frontendLogFile, err := os.OpenFile(fmt.Sprintf("../logs/frontend-%v.txt", id), os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("Error opening replica log file: %v", err)
	}

	sharedLogFile, err := os.OpenFile("../logs/combined.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("Error opening shared log file: %v", err)
	}

	frontendLogger := log.New(io.MultiWriter(os.Stdout, frontendLogFile), fmt.Sprintf("Frontend (%v): ", id), log.Flags())
	sharedLogger := log.New(sharedLogFile, fmt.Sprintf("Frontend (%v): ", id), log.Flags())
	frontendLogger.SetFlags(0)
	sharedLogger.SetFlags(0)

	return frontendLogger, sharedLogger
}

func clearLog(filename string) {
	if err := os.Truncate(filename, 0); err != nil {
		log.Printf("Failed to truncate log file %v: %v\n", filename, err)
	}
}
