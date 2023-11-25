package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	proto "github.com/emjakobsen1/dsys52023/c2fe"
	"google.golang.org/grpc"
)

func main() {
	// Parse command-line arguments to determine the frontend port
	if len(os.Args) != 2 {
		fmt.Println("Usage: client <frontend_port>")
		os.Exit(1)
	}

	frontendPort, err := strconv.Atoi(os.Args[1])
	if err != nil {
		log.Fatalf("Invalid frontend_port: %v", err)
	}

	// Create a gRPC connection to the frontend server
	conn, err := grpc.Dial(fmt.Sprintf("localhost:%d", frontendPort), grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Could not connect to frontend server on port %d: %v", frontendPort, err)
	}
	defer conn.Close()

	// Create a client for the frontend server
	client := proto.NewAuctionServiceClient(conn)

	// Start an infinite loop to continuously prompt for user input and send requests
	for {
		action, value := getUserInput()

		switch action {
		case "bid":
			sendBidRequest(client, int32(value))
		case "result":
			fmt.Println("Not implemented yet!")
		default:
			fmt.Println("Invalid action. Supported actions: bid, result")
		}
	}
}

func getUserInput() (string, int) {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter action and value (e.g., bid 100 or result): ")
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)
	parts := strings.Fields(input)

	if len(parts) == 1 {
		return parts[0], 0
	}
	if len(parts) < 2 {
		fmt.Println("Invalid input format. Supported format: action value")
		return "", 0
	}

	action := parts[0]
	value, err := strconv.Atoi(parts[1])
	if err != nil {
		fmt.Println("Invalid value. Please enter a valid integer.")
		return "", 0
	}

	return action, value
}
func sendBidRequest(client proto.AuctionServiceClient, value int32) {
	// Create a context
	ctx := context.Background()

	// Create a Bid request with the specified value
	req := &proto.Amount{
		Id:     123, // Provide an appropriate ID
		Amount: value,
	}

	// Send the Bid request to the frontend
	ack, err := client.Bid(ctx, req)
	if err != nil {
		log.Fatalf("Error sending Bid request: %v", err)
	}

	// Process the Ack response
	fmt.Printf("Received Ack response: State=%d, ID=%d\n", ack.State, ack.Id)
}
