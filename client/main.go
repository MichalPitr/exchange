package main

import (
	"context"
	"log"
	"math/rand"
	"sync"
	"time"

	pb "github.com/MichalPitr/exchange/protos"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	var wg sync.WaitGroup
	clientCount := 1000      // Number of concurrent clients
	requestsPerClient := 100 // Number of requests per client

	t0 := time.Now().UnixMilli()

	for i := 0; i < clientCount; i++ {
		wg.Add(1)
		go func(clientID int) {
			defer wg.Done()
			if clientID%2 == 0 {
				simulateClient(clientID, requestsPerClient, "BUY")
			} else {
				simulateClient(clientID, requestsPerClient, "SELL")
			}
		}(i)
	}

	wg.Wait()
	t1 := time.Now().UnixMilli()
	log.Printf("Processing took %d milliseconds\n", t1-t0)
}

func simulateClient(clientID, numRequests int, side string) {
	conn, err := grpc.Dial("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithBlock())
	if err != nil {
		log.Fatalf("Client %d did not connect: %v", clientID, err)
	}
	defer conn.Close()

	c := pb.NewOrderServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Minute)
	defer cancel()

	rand.Seed(time.Now().UnixNano()) // Seed for each client

	for i := 0; i < numRequests; i++ {
		orderType := "LIMIT"
		amount := rand.Int31n(20) + 1 // Random amount between 1 and 20
		price := rand.Int63n(500) + 1 // Random price between 1 and 500

		_, err := c.SendOrder(ctx, &pb.OrderRequest{
			UserId:    int32(clientID*numRequests + i),
			Type:      side,
			OrderType: orderType,
			Amount:    amount,
			Price:     price,
		})

		if err != nil {
			log.Printf("Client %d could not send order: %v", clientID, err)
			continue
		}
	}
}
