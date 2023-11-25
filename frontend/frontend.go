package main

import (
	"context"
	"fmt"
	"log"
	"net"

	proto "github.com/emjakobsen1/dsys52023/c2fe"
	"google.golang.org/grpc"
)

// Define a struct to represent the frontend server
type frontendServer struct {
	proto.UnimplementedAuctionServiceServer
	replicas map[int32]proto.AuctionServiceClient
}

func main() {
	// Initialize the frontend server and create client connections to replicas

	// Replica ports to connect to
	replicaPorts := []int32{5000, 5001, 5002}

	// Initialize a map to store client connections to replicas
	replicaClients := make(map[int32]proto.AuctionServiceClient)

	// Connect to each replica server
	for _, port := range replicaPorts {
		conn, err := grpc.Dial(fmt.Sprintf("localhost:%d", port), grpc.WithInsecure())
		if err != nil {
			log.Fatalf("Could not connect to replica server on port %d: %v", port, err)
		}
		defer conn.Close()

		// Create a client for each replica server and store it in the map
		replicaClients[port] = proto.NewAuctionServiceClient(conn)
	}

	// Initialize the frontend server with the replica clients
	frontend := &frontendServer{
		replicas: replicaClients,
	}

	// Start the frontend server to handle incoming requests
	grpcServer := grpc.NewServer()

	// Register the frontendServer with the gRPC server
	proto.RegisterAuctionServiceServer(grpcServer, frontend)

	// Create a listener on the desired port for the frontend server
	listener, err := net.Listen("tcp", "localhost:5003")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	// Print a message indicating that the frontend server is listening
	fmt.Println("Frontend server is listening on port 5003...")

	// Start serving incoming requests
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
func (s *frontendServer) Bid(ctx context.Context, req *proto.Amount) (*proto.Ack, error) {
	// Implement your Bid logic here for the replica server
	// You can access req.Id and req.Amount for processing
	// Return an Ack message based on your logic

	rep := &proto.Ack{State: 0, Id: 5003}
	s.ForwardToReplicas(ctx, req)
	return rep, nil
}
func (s *frontendServer) Bid2(ctx context.Context, req *proto.Amount) (*proto.Ack, error) {
	// Implement your Bid logic here for the replica server
	// You can access req.Id and req.Amount for processing
	// Return an Ack message based on your logic

	rep := &proto.Ack{State: 0, Id: 5003}

	return rep, nil
}

func (s *frontendServer) ForwardToReplicas(ctx context.Context, req *proto.Amount) (*proto.Ack, error) {
	// Forward the request to each replica
	for _, replicaClient := range s.replicas {
		_, err := replicaClient.Bid2(ctx, req)
		if err != nil {
			fmt.Printf("Failed to forward request to replica: %v\n", err)
		}
	}

	// Return an Ack indicating success or failure
	return &proto.Ack{State: proto.State_SUCCESS}, nil
}

// Implement methods for the frontendServer struct as needed
