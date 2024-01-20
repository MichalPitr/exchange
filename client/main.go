package main

import (
	"context"
	"log"
	"time"

	pb "github.com/MichalPitr/exchange/protos"

	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewOrderServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.SendOrder(ctx, &pb.OrderRequest{UserId: 1, Type: "SELL", OrderType: "LIMIT", Amount: 10, Price: 110})
	r, err = c.SendOrder(ctx, &pb.OrderRequest{UserId: 1, Type: "SELL", OrderType: "LIMIT", Amount: 10, Price: 150})

	r, err = c.SendOrder(ctx, &pb.OrderRequest{UserId: 1, Type: "BUY", OrderType: "LIMIT", Amount: 30, Price: 200})

	if err != nil {
		log.Fatalf("could not send order: %v", err)
	}
	log.Printf("Response: %s", r.GetStatus())
}
