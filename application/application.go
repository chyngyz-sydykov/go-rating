package application

import (
	"log"

	"github.com/chyngyz-sydykov/go-rating/infrastructure/config"
	"github.com/chyngyz-sydykov/go-rating/infrastructure/db"
	"gorm.io/gorm"
)

type App struct {
	//BookHandler handlers.BookHandler
}

func InitializeApplication() *App {
	// db := initializeDatabase()

	// logger := logger.NewLogger()

	// commonHandler := handlers.NewCommonHandler(logger)

	// bookService := book.NewBookService(db)
	// bookHandler := handlers.NewBookHandler(*bookService, *commonHandler)

	app := &App{
		//BookHandler: *bookHandler,
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
