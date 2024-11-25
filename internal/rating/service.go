package rating

import (
	"github.com/chyngyz-sydykov/go-rating/infrastructure/db/models"
	"gorm.io/gorm"
)

type RatingServiceInterface interface {
	GetByID(id int) (models.Rating, error)
	GetByBookID(bookId int) ([]models.Rating, error)
	Create(book *models.Rating) error
}

type RatingService struct {
	repository RatingRepository
}

func NewRatingService(db *gorm.DB) *RatingService {
	repository := NewRatingRepository(db)
	return &RatingService{
		repository: *repository,
	}
}

func (service *RatingService) GetByID(id int) (models.Rating, error) {
	return models.Rating{}, nil
}

func (service *RatingService) GetByBookID(bookId int) ([]models.Rating, error) {
	return nil, nil
}
func (service *RatingService) Create(book *models.Rating) error {
	return nil
}
