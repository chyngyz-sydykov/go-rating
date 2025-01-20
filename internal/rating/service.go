package rating

import (
	"fmt"
	"time"

	"github.com/chyngyz-sydykov/go-rating/infrastructure/db/models"
	"github.com/chyngyz-sydykov/go-rating/infrastructure/messagebroker"
	"gorm.io/gorm"
)

const BOOKRATED_EVENT_NAME = "bookRated"

type RatingServiceInterface interface {
	GetByID(id int) (models.Rating, error)
	GetByBookID(bookId int) ([]models.Rating, error)
	Create(rating *models.Rating) error
}

type RatingService struct {
	repository    RatingRepositoryInterface
	messageBroker messagebroker.MessageBrokerInterface
	validator     RatingValidator
}

func NewRatingService(
	db *gorm.DB,
	messageBrokerPublisher messagebroker.MessageBrokerInterface,
) *RatingService {

	repository := NewRatingRepository(db)
	validator := NewRatingValidator()

	return &RatingService{
		repository:    repository,
		validator:     *validator,
		messageBroker: messageBrokerPublisher,
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
	err = service.repository.Create(rating)
	if err != nil {
		return err
	}
	return service.publishMessage(rating, BOOKRATED_EVENT_NAME)
}

func (service *RatingService) publishMessage(rating *models.Rating, event string) error {

	bookMessage := BookMessage{
		BookId:   int(rating.BookId),
		EditedAt: time.Now(),
		Event:    event,
	}

	if err := service.messageBroker.Publish(bookMessage); err != nil {
		err := fmt.Errorf("failed to publish event: %v", err)
		return err

	}
	return nil
}
