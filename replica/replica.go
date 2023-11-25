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
	id int32
}

func (s *replicaServer) Bid(ctx context.Context, req *proto.Amount) (*proto.Ack, error) {
	// Implement your Bid logic here for the replica server
	// You can access req.Id and req.Amount for processing
	// Return an Ack message based on your logic
	fmt.Printf("Received Bid request from frontend: ID=%d, Amount=%d\n", req.Id, req.Amount)
	rep := &proto.Ack{State: 0, Id: s.id}
	return rep, nil
}

func (s *replicaServer) Bid2(ctx context.Context, req *proto.Amount) (*proto.Ack, error) {
	// Implement your Bid logic here for the replica server
	// You can access req.Id and req.Amount for processing
	// Return an Ack message based on your logic
	fmt.Printf("Received Bid request from frontend: ID=%d, Amount=%d\n", req.Id, req.Amount)
	rep := &proto.Ack{State: 0, Id: s.id}
	return rep, nil
}

func main() {
	// Parse command-line arguments to determine the port for this replica server
	if len(os.Args) != 2 {
		fmt.Println("Usage: replica <port>")
		os.Exit(1)
	}

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

	// Register the replicaServer with the gRPC server
	proto.RegisterAuctionServiceServer(grpcServer, &replicaServer{id: int32(port)})

	// Start serving requests
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve on port %d: %v", port, err)
	}
}
