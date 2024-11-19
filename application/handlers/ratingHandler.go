package handlers

type RatingHandler struct {
	//ratings []*proto.Rating
}

func NewRatingHandler() *RatingHandler {
	//return &RatingHandler{ratings: []*proto.Rating{}}

	return &RatingHandler{}
}

// func (h *RatingHandler) CreateRating(ctx context.Context, req *proto.CreateRatingRequest) (*proto.CreateRatingResponse, error) {
// 	rating := &proto.Rating{
// 		RatingId: "rating123", // Generate a unique ID here
// 		UserId:   req.UserId,
// 		BookId:   req.BookId,
// 		Score:    req.Score,
// 	}
// 	h.ratings = append(h.ratings, rating)
// 	return &proto.CreateRatingResponse{RatingId: rating.RatingId}, nil
// }

// func (h *RatingHandler) ListRatings(ctx context.Context, req *proto.ListRatingsRequest) (*proto.ListRatingsResponse, error) {
// 	var filteredRatings []*proto.Rating
// 	for _, rating := range h.ratings {
// 		if rating.BookId == req.BookId {
// 			filteredRatings = append(filteredRatings, rating)
// 		}
// 	}
// 	return &proto.ListRatingsResponse{Ratings: filteredRatings}, nil
// }
