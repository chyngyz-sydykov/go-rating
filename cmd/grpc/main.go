package main

import (
	"log"
	"net"

	"github.com/chyngyz-sydykov/go-rating/application"
	pb "github.com/chyngyz-sydykov/go-rating/proto/rating"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	grpcServer := grpc.NewServer()
	reflection.Register(grpcServer)

	app := application.InitializeApplication()

	pb.RegisterRatingServiceServer(grpcServer, &app.RatingHandler)

	listener, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen on port 50051: %v", err)
	}

	log.Println("Rating service is running on port 50051")
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("Failed to serve gRPC server: %v", err)
	}
}
