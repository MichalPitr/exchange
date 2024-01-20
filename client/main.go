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
	conn, err := grpc.Dial("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewOrderServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Minute)
	defer cancel()

	rand.New(rand.NewSource(0)) // Set fixed seed for replicability.

	var r *pb.OrderResponse
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		defer wg.Done()
		for userID := 1; userID <= 100; userID++ {
			orderType := "LIMIT"
			// if rand.Intn(2) == 0 {
			// 	orderType = "MARKET"
			// }

			// amount := rand.Int31n(100) + 1 // Random amount between 1 and 100
			amount := int32(10)
			price := rand.Int63n(500) + 1 // Random price between 1 and 500

			r, err = c.SendOrder(ctx, &pb.OrderRequest{
				UserId:    int32(userID),
				Type:      "BUY",
				OrderType: orderType,
				Amount:    amount,
				Price:     price,
			})

			if err != nil {
				// Handle error
				log.Fatalf("could not send order: %v", err)
			}
			time.Sleep(1 * time.Millisecond)
		}
	}()

	go func() {
		defer wg.Done()
		for userID := 101; userID <= 200; userID++ {
			orderType := "LIMIT"
			// if rand.Intn(2) == 0 {
			// 	orderType = "MARKET"
			// }

			// amount := rand.Int31n(100) + 1 // Random amount between 1 and 100
			amount := int32(10)
			price := rand.Int63n(500) + 1 // Random price between 1 and 500

			r, err = c.SendOrder(ctx, &pb.OrderRequest{
				UserId:    int32(userID),
				Type:      "SELL",
				OrderType: orderType,
				Amount:    amount,
				Price:     price,
			})

			if err != nil {
				// Handle error
				log.Fatalf("could not send order: %v", err)
			}
			time.Sleep(1 * time.Millisecond)
		}
	}()

	wg.Wait()
	log.Printf("Response: %s", r.GetStatus())
}
