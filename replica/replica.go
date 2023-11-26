// replica.go
package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"
	"strconv"

	proto "github.com/emjakobsen1/dsys52023/proto"
	"google.golang.org/grpc"
)

type replicaServer struct {
	proto.UnimplementedAuctionServiceServer
	id      int32
	bidders map[int32]int32
}

func (r *replicaServer) Bid2Replicas(ctx context.Context, req *proto.Amount) (*proto.Ack, error) {
	// Implement your Bid logic here for the replica server
	// You can access req.Id and req.Amount for processing
	// Return an Ack message based on your logic

	// First call registers the bidder.

	fmt.Printf("Received Bid request from frontend: ID=%d, Amount=%d sendSeq %d\n", req.Id, req.Amount, req.SendSeq)
	r.bidders[req.Id] = int32(req.Amount)
	rep := &proto.Ack{State: "SUCCESS", Id: r.id, SendSeq: req.SendSeq}
	return rep, nil
}

func (r *replicaServer) Result2Replicas(ctx context.Context, req *proto.Void) (*proto.Outcome, error) {

	var highestBid int32
	//var highestBidder int32

	for _, bid := range r.bidders {
		if bid > highestBid {
			highestBid = bid
			//highestBidder = bidderID
		}
	}

	fmt.Printf("Received result request from frontend: ID=%d,sendSeq %d\n", req.Id, req.SendSeq)
	rep := &proto.Outcome{HighestBid: highestBid, Id: r.id, SendSeq: req.SendSeq}
	return rep, nil
}

func main() {
	// Parse command-line arguments to determine the port for this replica server

	port, err := strconv.Atoi(os.Args[1])
	if err != nil {
		log.Fatalf("Invalid port: %v", err)
	}

	// Create a listener on the specified port
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		log.Fatalf("Failed to listen on port %d: %v", port, err)
	}

	// Initialize a new gRPC server
	grpcServer := grpc.NewServer()
	r := &replicaServer{
		id:      int32(port),
		bidders: make(map[int32]int32),
	}
	// Register the replicaServer with the gRPC server
	proto.RegisterAuctionServiceServer(grpcServer, r)

	// Start serving requests
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve on port %d: %v", port, err)
	}
}
