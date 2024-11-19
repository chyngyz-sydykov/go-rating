package main

import (
	"log"
	"net"

	"google.golang.org/grpc"
)

func main() {
	server := grpc.NewServer()
	//ratingHandler := handlers.NewRatingHandler()

	//services.RegisterRatingServiceServer(server, ratingHandler)

	listener, err := net.Listen("tcp", ":8001")
	if err != nil {
		log.Fatalf("Failed to listen on port 8001: %v", err)
	}

	log.Println("Rating service is running on port 8001")
	if err := server.Serve(listener); err != nil {
		log.Fatalf("Failed to serve gRPC server: %v", err)
	}
}
