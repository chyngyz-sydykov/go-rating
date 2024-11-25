package handlers

import (
	"context"
	"fmt"
	"log"

	"github.com/chyngyz-sydykov/go-rating/internal/rating"
	pb "github.com/chyngyz-sydykov/go-rating/proto/rating"
	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
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
	ratings, err := handler.service.GetByBookID(int(bookId))
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, "rating must be between 1 and 5")
	}

	fmt.Println(ratings)

	// Convert []pb.Rating to []*pb.Rating
	ratingsPtr := make([]*pb.Rating, len(ratings))
	for i := range ratings {
		ratingsPtr[i] = &pb.Rating{
			RatingId: ratings[i].ID.String(),
			BookId:   int32(ratings[i].BookId),
			Rating:   int32(ratings[i].Rating),
			Comment:  ratings[i].Comment,
		}
	}
	return &pb.GetRatingsResponse{Ratings: ratingsPtr}, nil
}
