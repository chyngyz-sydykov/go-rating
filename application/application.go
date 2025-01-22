package application

import (
	"fmt"
	"log"

	"github.com/chyngyz-sydykov/go-rating/application/handlers"
	"github.com/chyngyz-sydykov/go-rating/infrastructure/config"
	"github.com/chyngyz-sydykov/go-rating/infrastructure/db"
	"github.com/chyngyz-sydykov/go-rating/infrastructure/logger"
	"github.com/chyngyz-sydykov/go-rating/infrastructure/messagebroker"
	"github.com/chyngyz-sydykov/go-rating/internal/rating"
	"google.golang.org/grpc/codes"
	"gorm.io/gorm"
)

type App struct {
	RatingHandler handlers.RatingHandler
}

func InitializeApplication() *App {
	config, err := config.LoadMessageBrokerConfig()
	if err != nil {
		log.Fatalf("Could not config: %v", err)
	}
	logger := logger.NewLogger()
	db := InitializeDatabase()
	rabbitMqPublisher := InitializeRabbitMqPublisher(config, logger)

	ratingService := rating.NewRatingService(db, rabbitMqPublisher)

	commonHandler := handlers.NewCommonHandler(logger)

	ratingHandler := handlers.NewRatingHandler(ratingService, *commonHandler)

	app := &App{
		RatingHandler: *ratingHandler,
	}
	return app
}
func InitializeRabbitMqPublisher(config *config.MessageBrokerConfig, logger logger.LoggerInterface) messagebroker.MessageBrokerInterface {
	rabbitMQURL := "amqp://" + config.RabbitMqUser + ":" + config.RabbitMqPassword + "@" + config.RabbitMqContainerName + ":5672/"
	publisher, err := messagebroker.NewRabbitMQPublisher(rabbitMQURL, config.RabbitMqQueueName)

	if err != nil {
		err = fmt.Errorf("failed to initialize message publisher: %v", err)
		logger.LogError(codes.Aborted, err)

	} else {
		publisher.InitializeMessageBroker()
	}
	return publisher
}

func InitializeDatabase() *gorm.DB {
	dbConfig, err := config.LoadDBConfig()
	if err != nil {
		log.Fatalf("Could not load database config: %v", err)
	}
	dbInstance, err := db.InitializeDatabase(dbConfig)
	if err != nil {
		log.Fatalf("Coult not initialize db connection %v", err)
	}
	db.Migrate()
	return dbInstance

}
