package handlers

import (
	"context"
	"log"

	"github.com/chyngyz-sydykov/go-rating/internal/rating"
	pb "github.com/chyngyz-sydykov/go-rating/proto/rating"
	"github.com/google/uuid"
)

type RatingHandler struct {
	pb.RatingServiceServer
	service       rating.RatingServiceInterface
	commonHandler CommonHandler
}

func NewRatingHandler(service rating.RatingServiceInterface, commonHandler CommonHandler) *RatingHandler {
	return &RatingHandler{
		service:       service,
		commonHandler: commonHandler,
	}
}
func (handler *RatingHandler) SaveRating(context.Context, *pb.SaveRatingRequest) (*pb.SaveRatingResponse, error) {
	log.Printf("Received SaveRating request")
	rating := &pb.Rating{
		RatingId: uuid.New().String(),
		BookId:   101,
		Rating:   5,
		Comment:  "SaveRating!",
	}

	return &pb.SaveRatingResponse{Rating: rating}, nil
}
func (handler *RatingHandler) GetRatings(ctx context.Context, req *pb.GetRatingsRequest) (*pb.GetRatingsResponse, error) {
	log.Printf("Received GetRatings request for book_id: %d", req.BookId)

	ratings := []pb.Rating{
		{
			RatingId: uuid.New().String(),
			BookId:   101,
			Rating:   5,
			Comment:  "GetRatings!",
		},
	}

	// Convert []pb.Rating to []*pb.Rating
	ratingsPtr := make([]*pb.Rating, len(ratings))
	for i := range ratings {
		ratingsPtr[i] = &ratings[i]
	}
	return &pb.GetRatingsResponse{Ratings: ratingsPtr}, nil
}
