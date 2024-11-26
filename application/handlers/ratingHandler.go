package handlers

import (
	"context"
	"errors"
	"fmt"
	"log"

	my_error "github.com/chyngyz-sydykov/go-rating/error"
	"github.com/chyngyz-sydykov/go-rating/infrastructure/db/models"
	"github.com/chyngyz-sydykov/go-rating/internal/rating"
	pb "github.com/chyngyz-sydykov/go-rating/proto/rating"
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
	log.Printf("SaveRating request")
	rating := &models.Rating{
		BookId:  int64(req.BookId),
		Rating:  int(req.Rating),
		Comment: req.Comment,
	}

	err := handler.service.Create(rating)
	if err != nil {
		handler.commonHandler.HandleError(codes.InvalidArgument, err)
		return nil, status.Error(codes.InvalidArgument, "resource value(s) is invalid")
	}

	ratingResponse := &pb.Rating{
		RatingId: rating.ID.String(),
		BookId:   int32(rating.BookId),
		Rating:   int32(rating.Rating),
		Comment:  rating.Comment,
	}

	return &pb.SaveRatingResponse{Rating: ratingResponse}, nil
}
func (handler *RatingHandler) GetRatings(ctx context.Context, req *pb.GetRatingsRequest) (*pb.GetRatingsResponse, error) {
	log.Printf("GetRatings request")
	bookId := req.BookId
	ratings, err := handler.service.GetByBookID(int(bookId))
	if err != nil {
		if errors.Is(err, my_error.ErrInvalidArgument) {
			return nil, status.Errorf(codes.InvalidArgument, "book with ID %d is invalid", bookId)
		}
		return nil, status.Errorf(codes.Unknown, "unknown error %s", err)
	}

	fmt.Println(ratings)

	// Convert []pb.Rating to []*pb.Rating
	ratingList := handler.mapGormToGrpcObject(ratings)
	return &pb.GetRatingsResponse{Ratings: ratingList}, nil
}

func (*RatingHandler) mapGormToGrpcObject(ratings []models.Rating) []*pb.Rating {
	ratingsPtr := make([]*pb.Rating, len(ratings))
	for i := range ratings {
		ratingsPtr[i] = &pb.Rating{
			RatingId: ratings[i].ID.String(),
			BookId:   int32(ratings[i].BookId),
			Rating:   int32(ratings[i].Rating),
			Comment:  ratings[i].Comment,
		}
	}
	return ratingsPtr
}
