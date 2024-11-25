package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/chyngyz-sydykov/go-rating/application"
	"github.com/chyngyz-sydykov/go-rating/application/handlers"
	"github.com/chyngyz-sydykov/go-rating/infrastructure/config"
	"github.com/chyngyz-sydykov/go-rating/infrastructure/db/models"
	"github.com/chyngyz-sydykov/go-rating/infrastructure/logger"
	pb "github.com/chyngyz-sydykov/go-rating/proto/rating"
	"github.com/stretchr/testify/assert"

	"github.com/stretchr/testify/mock"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func (suite *IntegrationSuite) TestSaveRatingMethod_ShouldReturnSuccessResponseWithNewRating() {
	// Create an in-memory gRPC server
	listener, server := createInMemoryGrpcServer()

	defer server.Stop()

	//Create a gRPC client connected to the server
	fmt.Println("aaaa", listener.Addr().String())

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

func (suite *IntegrationSuite) TestGetRatingMethod_ShouldReturnSuccessResponseWithListOfRatings() {
	// Create an in-memory gRPC server
	listener, server := createInMemoryGrpcServer()

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
		BookId: 111,
	}
	res, err := client.GetRatings(context.Background(), req)
	returnedRating := res.Ratings[0]
	// Assert
	assert.NoError(suite.T(), err)
	assert.NotNil(suite.T(), res)
	assert.Equal(suite.T(), req.BookId, returnedRating.BookId)
	assert.Equal(suite.T(), int32(5), returnedRating.Rating)
	assert.Equal(suite.T(), returnedRating.Comment, "GetRatings!")
	assert.NotEmpty(suite.T(), returnedRating.RatingId)
}

func createInMemoryGrpcServer() (net.Listener, *grpc.Server) {
	app := provideDependencies()
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
			log.Fatalf("Failed to serve gRPC server:dddd %v", err)
		}
	}()
	return listener, server
}

func provideDependencies() *application.App {
	logger := logger.NewLogger()
	commonHandler := handlers.NewCommonHandler(logger)

	ratingServiceMock := newRatingServiceMock()

	ratingHandler := handlers.NewRatingHandler(ratingServiceMock, *commonHandler)

	app := &application.App{
		RatingHandler: *ratingHandler,
	}
	return app
}

type RatingServiceMock struct {
	mock.Mock
}

func newRatingServiceMock() *RatingServiceMock {
	return &RatingServiceMock{}
}

func (service *RatingServiceMock) GetByID(id int) (models.Rating, error) {
	return models.Rating{}, nil
}

func (service *RatingServiceMock) GetByBookID(bookId int) ([]models.Rating, error) {
	return nil, nil
}
func (service *RatingServiceMock) Create(book *models.Rating) error {
	return nil
}
