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
func (handler *RatingHandler) SaveRating(ctx context.Context, req *pb.SaveRatingRequest) (*pb.SaveRatingResponse, error) {

	log.Printf("Received SaveRating request")
	rating := &pb.Rating{
		RatingId: uuid.New().String(),
		BookId:   req.BookId,
		Rating:   req.Rating,
		Comment:  "SaveRating!",
	}

	return &pb.SaveRatingResponse{Rating: rating}, nil
}
func (handler *RatingHandler) GetRatings(ctx context.Context, req *pb.GetRatingsRequest) (*pb.GetRatingsResponse, error) {
	bookId := req.BookId
	//_ := handler.service.GetByBookID(int(bookId))

	ratings := []pb.Rating{
		{
			RatingId: uuid.New().String(),
			BookId:   bookId,
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
