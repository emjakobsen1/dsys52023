package main

import (
	"context"
	"fmt"
	"log"
	"net"

	proto "github.com/emjakobsen1/dsys52023/proto"
	"google.golang.org/grpc"
)

// Define a struct to represent the frontend server
type frontendServer struct {
	proto.UnimplementedAuctionServiceServer
	replicas  map[int32]proto.AuctionServiceClient
	sendSeq   int32
	delivered map[int32]int
	buffer    []proto.Ack
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
		replicas:  replicaClients,
		sendSeq:   0,
		delivered: make(map[int32]int),
		buffer:    []proto.Ack{},
	}
	frontend.delivered[5000] = 0
	frontend.delivered[5001] = 0
	frontend.delivered[5002] = 0
	frontend.delivered[5003] = 0

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
	// Implement your Bid logic here for the client requewst
	// You can access req.Id and req.Amount for processing
	// Return an Ack message based on your logic

	rep := &proto.Ack{State: 0, Id: 5003}

	req.SendSeq = int32(s.sendSeq)
	s.Broadcast(ctx, req)
	s.sendSeq++
	return rep, nil
}
func (s *frontendServer) Bid2replicas(ctx context.Context, req *proto.Amount, sendSeq int) (*proto.Ack, error) {
	// Implement your Bid logic here for the frontends request to the replica server
	// You can access req.Id and req.Amount for processing
	// Return an Ack message based on your logic

	rep := &proto.Ack{State: 0, Id: 5003}

	return rep, nil
}

func (s *frontendServer) Broadcast(ctx context.Context, req *proto.Amount) (*proto.Ack, error) {
	// Forward the request to each replica

	for _, replicaClient := range s.replicas {
		rep, err := replicaClient.Bid2Replicas(ctx, req)
		if err != nil {
			fmt.Printf("Failed to forward request to replica: %v\n", err)
		}
		s.buffer = append(s.buffer, *rep)

		//fmt.Printf("Received ack from replica: ID=%d, State=%d\n", rep.Id, rep.State)
	}

	for _, msg := range s.buffer {
		if s.sendSeq == msg.SendSeq {
			fmt.Printf("Received ack from replica: ID=%d, State=%d\n", msg.Id, msg.State)
			s.delivered[msg.Id]++
		}

	}

	// Return an Ack indicating success or failure
	//fmt.Printf("Received ack from replicas?: ID=%d, Amount=%d sendSeq %d\n", req.Id, req.Amount, req.SendSeq)
	return &proto.Ack{State: proto.State_SUCCESS}, nil
}

// Implement methods for the frontendServer struct as needed
