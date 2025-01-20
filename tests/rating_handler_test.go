package main

import (
	"context"
	"log"
	"net"
	"testing"
	"time"

	"github.com/chyngyz-sydykov/go-rating/application"
	"github.com/chyngyz-sydykov/go-rating/application/handlers"
	"github.com/chyngyz-sydykov/go-rating/infrastructure/config"
	"github.com/chyngyz-sydykov/go-rating/infrastructure/db/models"
	"github.com/chyngyz-sydykov/go-rating/infrastructure/logger"
	"github.com/chyngyz-sydykov/go-rating/internal/rating"
	pb "github.com/chyngyz-sydykov/go-rating/proto/rating"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/status"
)

func (suite *IntegrationSuite) TestSaveRating_InvalidRating() {
	testCases := []struct {
		name         string
		rating       int32
		expectedCode codes.Code
	}{

		{"RatingIsNegative", -5, codes.InvalidArgument},
		{"RatingTooLow", 0, codes.InvalidArgument},
		{"RatingIs6", 6, codes.InvalidArgument},
		{"RatingTooHigh", 10, codes.InvalidArgument},
	}

	// Create an in-memory gRPC server
	listener, server := createInMemoryGrpcServer(suite)
	defer server.Stop()

	for _, tc := range testCases {
		suite.T().Run(tc.name, func(t *testing.T) {

			//Create a gRPC client connected to the server
			conn, err := grpc.NewClient(listener.Addr().String(), grpc.WithTransportCredentials(insecure.NewCredentials()))
			if err != nil {
				log.Fatalf("Failed to dial server: %v", err)
			}
			defer conn.Close()

			client := pb.NewRatingServiceClient(conn)

			// Act
			req := &pb.SaveRatingRequest{
				BookId:  101,
				Rating:  tc.rating,
				Comment: "Test comment",
			}
			_, err = client.SaveRating(context.Background(), req)

			st, _ := status.FromError(err)

			if tc.expectedCode == codes.OK {
				assert.NoError(t, err)
			} else {
				assert.Error(t, err)
				assert.Equal(t, tc.expectedCode, st.Code())
			}
		})
	}
}

func (suite *IntegrationSuite) TestGetRatingsMethod_ShouldReturnInvalidArgumentWhenRatingIsNotValid() {
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
		BookId: -1,
	}
	_, err = client.GetRatings(context.Background(), req)

	// Assert
	assert.Error(suite.T(), err)
	st, ok := status.FromError(err)
	assert.True(suite.T(), ok)
	assert.Equal(suite.T(), codes.InvalidArgument, st.Code())
}

func (suite *IntegrationSuite) TestSaveRatingMethod_ShouldReturnSuccessResponseWithNewRating() {
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
		BookId:  1,
		Rating:  5,
		Comment: "test comment 5",
	}
	res, err := client.SaveRating(context.Background(), req)

	// Assert
	assert.NoError(suite.T(), err)
	assert.NotNil(suite.T(), res)
	assert.Equal(suite.T(), req.BookId, res.Rating.BookId)
	assert.Equal(suite.T(), req.Rating, res.Rating.Rating)
	assert.Equal(suite.T(), req.Comment, "test comment 5")
	assert.NotEmpty(suite.T(), res.Rating.RatingId)

	err = suite.db.Where("book_id = ? and rating = ? and comment = ?", 1, 5, "test comment 5").First(&models.Rating{}).Error
	suite.Suite.Assert().Nil(err)

	suite.db.Unscoped().Where("1 = 1").Delete(&models.Rating{})
}

func (suite *IntegrationSuite) TestGetRatingsMethod_ShouldReturnInvalidResponseWhenBookIdIsNegative() {
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
		BookId: -1,
	}
	_, err = client.GetRatings(context.Background(), req)

	// Assert
	assert.Error(suite.T(), err)
	st, ok := status.FromError(err)
	assert.True(suite.T(), ok)
	assert.Equal(suite.T(), codes.InvalidArgument, st.Code())
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

func provideDependencies(suite *IntegrationSuite) *application.App {
	logger := logger.NewLogger()
	commonHandler := handlers.NewCommonHandler(logger)

	var messageBrokerMock MessageBrokerMock
	expectedMessage := rating.BookMessage{
		BookId:   1,
		EditedAt: time.Now(),
		Event:    "bookRated",
	}

	messageBrokerMock.On("Publish", mock.MatchedBy(func(msg rating.BookMessage) bool {
		return msg.Event == expectedMessage.Event &&
			msg.BookId == 1
	})).Return(nil)

	ratingService := rating.NewRatingService(suite.db, &messageBrokerMock)
	ratingHandler := handlers.NewRatingHandler(ratingService, *commonHandler)

	app := &application.App{
		RatingHandler: *ratingHandler,
	}
	return app
}

type MessageBrokerMock struct {
	mock.Mock
}

func (m *MessageBrokerMock) Publish(message interface{}) error {
	args := m.Called(message)
	return args.Error(0)
}

func (m *MessageBrokerMock) InitializeMessageBroker() {
	m.Called()
}

func (m *MessageBrokerMock) Close() {
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
