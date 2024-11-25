package application

import (
	"log"

	"github.com/chyngyz-sydykov/go-rating/application/handlers"
	"github.com/chyngyz-sydykov/go-rating/infrastructure/config"
	"github.com/chyngyz-sydykov/go-rating/infrastructure/db"
	"github.com/chyngyz-sydykov/go-rating/infrastructure/logger"
	"github.com/chyngyz-sydykov/go-rating/internal/rating"
	"gorm.io/gorm"
)

type App struct {
	RatingHandler handlers.RatingHandler
}

func InitializeApplication() *App {

	db := initializeDatabase()
	ratingService := rating.NewRatingService(db)

	logger := logger.NewLogger()

	commonHandler := handlers.NewCommonHandler(logger)

	ratingHandler := handlers.NewRatingHandler(ratingService, *commonHandler)

	app := &App{
		RatingHandler: *ratingHandler,
	}
	return app
}

func initializeDatabase() *gorm.DB {
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
