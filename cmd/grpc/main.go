package main

import (
	"context"
	"log"
	"net"

	pb "github.com/chyngyz-sydykov/go-web/proto/rating"
	"github.com/google/uuid"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	grpcServer := grpc.NewServer()
	reflection.Register(grpcServer)

	//ratingHandler := handlers.NewRatingHandler()

	//services.RegisterRatingServiceServer(server, ratingHandler)
	pb.RegisterRatingServiceServer(grpcServer, &RatingServiceServer{})

	listener, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen on port 50051: %v", err)
	}

	log.Println("Rating service is running on port 50051")
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("Failed to serve gRPC server: %v", err)
	}
}

type RatingServiceServer struct {
	pb.RatingServiceServer
}

func (s *RatingServiceServer) GetRatings(ctx context.Context, req *pb.GetRatingsRequest) (*pb.GetRatingsResponse, error) {
	log.Printf("Received GetRatings request for book_id: %d", req.BookId)

	ratings := []pb.Rating{
		{
			RatingId: uuid.New().String(),
			BookId:   101,
			Rating:   5,
			Comment:  "Great book!",
		},
	}

	// Convert []pb.Rating to []*pb.Rating
	ratingsPtr := make([]*pb.Rating, len(ratings))
	for i := range ratings {
		ratingsPtr[i] = &ratings[i]
	}
	return &pb.GetRatingsResponse{Ratings: ratingsPtr}, nil
}
