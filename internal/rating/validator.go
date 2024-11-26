package rating

import my_error "github.com/chyngyz-sydykov/go-rating/error"

type RatingValidator struct {
}

func NewRatingValidator() *RatingValidator {
	return &RatingValidator{}
}

func (ratingValidator *RatingValidator) validateBookId(bookId int) error {
	if bookId <= 0 {
		return my_error.ErrInvalidArgument
	}
	return nil
}

func (ratingValidator *RatingValidator) validateRating(rating int) error {
	if rating < 1 || rating > 5 {
		return my_error.ErrInvalidArgument
	}
	return nil
}
