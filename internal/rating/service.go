package rating

import (
	"github.com/chyngyz-sydykov/go-rating/infrastructure/db/models"
	"gorm.io/gorm"
)

type RatingServiceInterface interface {
	GetByID(id int) (models.Rating, error)
	GetByBookID(bookId int) ([]models.Rating, error)
	Create(rating *models.Rating) error
}

type RatingService struct {
	repository RatingRepositoryInterface
	validator  RatingValidator
}

func NewRatingService(db *gorm.DB) *RatingService {

	repository := NewRatingRepository(db)
	validator := NewRatingValidator()

	return &RatingService{
		repository: repository,
		validator:  *validator,
	}
}

func (service *RatingService) GetByID(id int) (models.Rating, error) {
	return models.Rating{}, nil
}

func (service *RatingService) GetByBookID(bookId int) ([]models.Rating, error) {
	err := service.validator.validateBookId(bookId)
	if err != nil {
		return nil, err
	}

	ratings, err := service.repository.GetByBookID(bookId)
	if err != nil {
		return ratings, err
	}
	return ratings, err
}

func (service *RatingService) Create(rating *models.Rating) error {
	err := service.validator.validateRating(rating.Rating)
	if err != nil {
		return err
	}
	return service.repository.Create(rating)
}
