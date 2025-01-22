package rating

import (
	"github.com/chyngyz-sydykov/go-rating/infrastructure/db/models"

	"gorm.io/gorm"
)

type RatingRepositoryInterface interface {
	GetByBookID(bookId int) ([]models.Rating, error)
	Create(rating *models.Rating) error
}

type RatingRepository struct {
	db *gorm.DB
}

func NewRatingRepository(db *gorm.DB) *RatingRepository {
	return &RatingRepository{db: db}
}

func (repository *RatingRepository) GetByBookID(bookId int) ([]models.Rating, error) {
	var ratings []models.Rating
	if err := repository.db.Where("book_id = ?", bookId).Order("ID desc").Find(&ratings).Error; err != nil {
		return ratings, err
	}
	return ratings, nil
}

func (repository *RatingRepository) Create(rating *models.Rating) error {
	return repository.db.Create(&rating).Error
}
