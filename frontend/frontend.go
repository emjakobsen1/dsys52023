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
	buffer2   []proto.Outcome
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

	grpcServer := grpc.NewServer()
	proto.RegisterAuctionServiceServer(grpcServer, frontend)

	// Create a listener on the desired port for the frontend server
	listener, err := net.Listen("tcp", "localhost:5003")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	fmt.Println("Frontend server is listening on port 5003...")

	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
func (s *frontendServer) Bid(ctx context.Context, req *proto.Amount) (*proto.Ack, error) {
	// Implement your Bid logic here for the client request
	// You can access req.Id and req.Amount for processing
	// Return an Ack message based on your logic

	rep := &proto.Ack{State: "SUCCES", Id: 5003}

	req.SendSeq = int32(s.sendSeq)
	if s.BroadcastBid(ctx, req) {
		s.sendSeq++
		return rep, nil
	}
	return &proto.Ack{State: "FAIL", Id: 5003}, nil
}
func (s *frontendServer) Bid2Replicas(ctx context.Context, req *proto.Amount) (*proto.Ack, error) {
	// Implement your Bid logic here for the frontends request to the replica server
	// You can access req.Id and req.Amount for processing
	// Return an Ack message based on your logic

	rep := &proto.Ack{State: "SUCCES", Id: 5003}

	return rep, nil
}

func (s *frontendServer) BroadcastBid(ctx context.Context, req *proto.Amount) bool {
	// Forward the request to each replica

	for _, replicaClient := range s.replicas {
		rep, err := replicaClient.Bid2Replicas(ctx, req)
		if err != nil {
			fmt.Printf("Failed to forward request to replica: %v\n", err)
		}
		s.buffer = append(s.buffer, *rep)
		fmt.Println(s.buffer)
		//fmt.Printf("Received ack from replica: ID=%d, State=%d\n", rep.Id, rep.State)
	}
	var replies = false
	for ids, msg := range s.buffer {

		if s.sendSeq == msg.SendSeq {
			replies = true
			fmt.Printf("Received ack from replica: ID=%d, State=%s\n", msg.Id, msg.State)
			s.delivered[int32(ids)]++

		}
	}
	if replies {
		return true
	}

	return false
	// Return an Ack indicating success or failure
	//fmt.Printf("Received ack from replicas?: ID=%d, Amount=%d sendSeq %d\n", req.Id, req.Amount, req.SendSeq)
}

func (s *frontendServer) Result(ctx context.Context, req *proto.Void) (*proto.Outcome, error) {
	// Implement your Bid logic here for the client request
	// You can access req.Id and req.Amount for processing
	// Return an Ack message based on your logic

	req.SendSeq = int32(s.sendSeq)

	rep, err := s.BroadcastResult(ctx, req)
	if err != nil {
		fmt.Printf("Failed to forward request to replica: %v\n", err)
	}
	s.sendSeq++
	return rep, nil
	// if s.BroadcastResult(ctx, req) {
	// 	rep := &proto.Outcome{HighestBid: 1, Id: 5003}
	// 	s.sendSeq++
	// 	return rep, nil
	// }
	return &proto.Outcome{HighestBid: 0, Id: 5003}, nil
}

func (s *frontendServer) Result2Replicas(ctx context.Context, req *proto.Void) (*proto.Outcome, error) {
	// Implement your Bid logic here for the frontends request to the replica server
	// You can access req.Id and req.Amount for processing
	// Return an Ack message based on your logic

	rep := &proto.Outcome{HighestBid: 0, Id: 5003}

	return rep, nil
}

func (s *frontendServer) BroadcastResult(ctx context.Context, req *proto.Void) (*proto.Outcome, error) {
	// Forward the request to each replica
	var outcome *proto.Outcome
	for _, replicaClient := range s.replicas {
		rep, err := replicaClient.Result2Replicas(ctx, req)
		if err != nil {
			fmt.Printf("Failed to forward request to replica: %v\n", err)
		}
		s.buffer2 = append(s.buffer2, *rep)
		outcome = rep
		//fmt.Printf("Received ack from replica: ID=%d, State=%d\n", rep.Id, rep.State)
	}
	var replies = false
	for ids, msg := range s.buffer2 {

		if s.sendSeq == msg.SendSeq {
			replies = true
			fmt.Printf("Received outcome from replica: ID=%d, Highest bid=%d\n", msg.Id, msg.HighestBid)
			s.delivered[int32(ids)]++

		}
	}
	if replies {
		return outcome, nil
	}

	return &proto.Outcome{}, nil
	// Return an Ack indicating success or failure
	//fmt.Printf("Received ack from replicas?: ID=%d, Amount=%d sendSeq %d\n", req.Id, req.Amount, req.SendSeq)
}
