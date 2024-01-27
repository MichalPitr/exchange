package main

import (
	"log"
	"net"
	"os"
	"os/signal"
	"sync"
	"syscall"

	pb "github.com/MichalPitr/exchange/protos"
	"google.golang.org/grpc"

	"github.com/MichalPitr/exchange/engine"
)

func main() {
	// Run your program here
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	s := grpc.NewServer()
	e, err := engine.New(1000)
	if err != nil {
		log.Fatalf("Failed to start engine: %v", err)
	}

	go engine.ProcessOrders(e)
	pb.RegisterOrderServiceServer(s, e)

	// Setting up signal handling for graceful shutdown
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM)

	var wg sync.WaitGroup

	// Start the server in a separate goroutine
	go func() {
		log.Printf("Server listening at %v", lis.Addr())
		if err := s.Serve(lis); err != nil {
			log.Fatalf("Failed to serve: %v", err)
		}
	}()

	// Goroutine to handle graceful shutdown
	wg.Add(1)
	go func() {
		<-sigChan // Wait for interrupt signal
		log.Println("Shutting down the server...")

		s.GracefulStop()
		e.Close()

		wg.Done()
	}()

	wg.Wait()
	log.Println("Server successfully stopped")
}
