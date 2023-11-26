package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	proto "github.com/emjakobsen1/dsys52023/proto"
	"google.golang.org/grpc"
)

var ID int

func main() {
	// Usage: go run client.go <id> 5003 from a terminal in /client. e.g. go run client.go 1 5003
	// Starts a client sending to the frontend at 5003 with ID 1.
	// It supports the commandline actions, bid and result.
	// bid takes an argument on the form: bid <integer>, with integer beingthe actual bid, e.g. bid 200
	var err error
	ID, err = strconv.Atoi(os.Args[1])
	if err != nil {
		log.Fatalf("Invalid client ID: %v", err)
	}

	frontendPort, err := strconv.Atoi(os.Args[2])
	if err != nil {
		log.Fatalf("Invalid frontend port: %v", err)
	}

	conn, err := grpc.Dial(fmt.Sprintf("localhost:%d", frontendPort), grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Could not connect to frontend port %d: %v", frontendPort, err)
	}
	defer conn.Close()

	client := proto.NewAuctionServiceClient(conn)

	for {
		action, value := getUserInput()

		switch action {
		case "bid":
			sendBidRequest(client, int32(value))
		case "result":
			sendResultRequest(client)
		default:
			fmt.Println("Invalid action. Supported actions: bid, result")
		}
	}
}

func getUserInput() (string, int) {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Write result to query the auction, bid <integer> to bid")
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)
	parts := strings.Fields(input)

	if len(parts) == 1 {
		return parts[0], 0
	}
	if len(parts) < 2 {
		fmt.Println("Invalid input format")
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

// Send bid request to the frontend
func sendBidRequest(client proto.AuctionServiceClient, value int32) {

	ctx := context.Background()

	req := &proto.Amount{
		Id:     int32(ID),
		Amount: value,
	}

	ack, err := client.Bid(ctx, req)
	if err != nil {
		log.Fatalf("Error sending Bid request: %v", err)
	}

	fmt.Printf("Received Ack response: State=%s, ID=%d\n", ack.State, ack.Id)
}

// Sends result request to the frontend
func sendResultRequest(client proto.AuctionServiceClient) {

	ctx := context.Background()

	req := &proto.Void{
		Id: int32(ID),
	}

	outcome, err := client.Result(ctx, req)
	if err != nil {
		log.Fatalf("Error sending Bid request: %v", err)
	}

	// Process the Ack response
	fmt.Printf("Received outcome response: Highest bid: %d, ID: %d\n", outcome.HighestBid, outcome.Id)
}
