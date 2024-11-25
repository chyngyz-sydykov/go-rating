package main

import (
	"context"
	"log"
	"net"

	"github.com/chyngyz-sydykov/go-rating/application"
	"github.com/chyngyz-sydykov/go-rating/application/handlers"
	"github.com/chyngyz-sydykov/go-rating/infrastructure/config"
	"github.com/chyngyz-sydykov/go-rating/infrastructure/db/models"
	"github.com/chyngyz-sydykov/go-rating/infrastructure/logger"
	"github.com/chyngyz-sydykov/go-rating/internal/rating"
	pb "github.com/chyngyz-sydykov/go-rating/proto/rating"
	"github.com/stretchr/testify/assert"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func (suite *IntegrationSuite) TesssstSaveRatingMethod_ShouldReturnSuccessResponseWithNewRating() {
	// Create an in-memory gRPC server
	listener, server := createInMemoryGrpcServer(suite)
	defer server.Stop()

	//Create a gRPC client connected to the server
	conn, err := grpc.NewClient(listener.Addr().String(), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed to dial server: %v", err)
	}
	defer conn.Close()

	client := pb.NewRatingServiceClient(conn)

	// Act
	req := &pb.SaveRatingRequest{
		BookId:  111,
		Rating:  5,
		Comment: "SaveRating!",
	}
	res, err := client.SaveRating(context.Background(), req)

	// Assert
	assert.NoError(suite.T(), err)
	assert.NotNil(suite.T(), res)
	assert.Equal(suite.T(), req.BookId, res.Rating.BookId)
	assert.Equal(suite.T(), req.Rating, res.Rating.Rating)
	assert.Equal(suite.T(), req.Comment, "SaveRating!")
	assert.NotEmpty(suite.T(), res.Rating.RatingId)
}

func (suite *IntegrationSuite) TestGetRatingsMethod_ShouldReturnSuccessResponseWithListOfRatings() {
	expectedRating := models.Rating{Rating: 1, Comment: "test comment 1", BookId: 9}
	suite.db.Create(&expectedRating)

	// Create an in-memory gRPC server
	listener, server := createInMemoryGrpcServer(suite)

	defer server.Stop()

	//Create a gRPC client connected to the server
	conn, err := grpc.NewClient(listener.Addr().String(), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed to dial server: %v", err)
	}
	defer conn.Close()

	client := pb.NewRatingServiceClient(conn)

	// Act
	req := &pb.GetRatingsRequest{
		BookId: 9,
	}
	res, err := client.GetRatings(context.Background(), req)

	returnedRating := res.Ratings[0]
	// // Assert
	assert.NoError(suite.T(), err)
	assert.NotNil(suite.T(), res)
	assert.Equal(suite.T(), int32(9), returnedRating.BookId)
	assert.Equal(suite.T(), int32(1), returnedRating.Rating)
	assert.Equal(suite.T(), "test comment 1", returnedRating.Comment)
	assert.NotEmpty(suite.T(), returnedRating.RatingId)

	suite.db.Unscoped().Where("1 = 1").Delete(&models.Rating{})
}

func createInMemoryGrpcServer(suite *IntegrationSuite) (net.Listener, *grpc.Server) {
	app := provideDependencies(suite)
	server := grpc.NewServer()
	pb.RegisterRatingServiceServer(server, &app.RatingHandler)

	config, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Could not config: %v", err)
	}
	listener, err := net.Listen("tcp", "localhost:"+config.ApplicationPort)
	if err != nil {
		log.Fatalf("Failed to start listener: %v", err)
	}
	go func() {
		if err := server.Serve(listener); err != nil {
			log.Fatalf("Failed to serve gRPC server: %v", err)
		}
	}()
	return listener, server
}

func provideDependencies(suite *IntegrationSuite) *application.App {
	logger := logger.NewLogger()
	commonHandler := handlers.NewCommonHandler(logger)

	ratingService := rating.NewRatingService(suite.db)

	ratingHandler := handlers.NewRatingHandler(ratingService, *commonHandler)

	app := &application.App{
		RatingHandler: *ratingHandler,
	}
	return app
}
