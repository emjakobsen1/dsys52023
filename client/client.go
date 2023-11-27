package main

import (
	"bufio"
	"context"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"

	proto "github.com/emjakobsen1/dsys52023/proto"
	"google.golang.org/grpc"
)

var ID int
var l, sl *log.Logger

func main() {
	// Usage: go run client.go <id> 5003 from a terminal in /client. e.g. go run client.go 1 5003
	// Starts a client sending to the frontend at 5003 with ID 1.
	// It supports the commandline actions, bid and result.
	// bid takes an argument on the form: bid <integer>, with integer being the actual bid, e.g. bid 200

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

	// client's log, shared log
	l, sl = setLog(ID)

	fmt.Printf("Client (%d) - supported actions: \n 	bid <integer argument>\n 	result \n First bid register the bidder and starts the auction.\n", ID)
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

	//fmt.Print("Write result to query the auction, bid <integer> to bid \n")
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

// Sends bid request to the frontend
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

	l.Printf("-> Ack: %s\n", ack.State)
	sl.Printf("-> Ack: %s\n", ack.State)
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

	l.Printf("-> Result: %d\n", outcome.HighestBid)
	sl.Printf("-> Result: %d\n", outcome.HighestBid)
}

func setLog(id int) (*log.Logger, *log.Logger) {

	clearLog(fmt.Sprintf("../logs/client-%v.txt", id))

	clientLogFile, err := os.OpenFile(fmt.Sprintf("../logs/client-%v.txt", id), os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("Error opening replica log file: %v", err)
	}

	sharedLogFile, err := os.OpenFile("../logs/combined.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("Error opening shared log file: %v", err)
	}

	clientLogger := log.New(io.MultiWriter(os.Stdout, clientLogFile), fmt.Sprintf("Client (%v): ", id), log.Flags())
	sharedLogger := log.New(sharedLogFile, fmt.Sprintf("Client (%v): ", id), log.Flags())

	clientLogger.SetFlags(0)
	sharedLogger.SetFlags(0)

	return clientLogger, sharedLogger
}

func clearLog(filename string) {
	if err := os.Truncate(filename, 0); err != nil {
		log.Printf("Failed to truncate log file %v: %v\n", filename, err)
	}
}
