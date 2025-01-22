package main

import (
	"encoding/json"
	"fmt"
	"log"
	"strings"
	"sync"
	"time"

	"github.com/chyngyz-sydykov/go-rating/application"
	"github.com/chyngyz-sydykov/go-rating/infrastructure/config"
	"github.com/chyngyz-sydykov/go-rating/infrastructure/db/models"
	"github.com/chyngyz-sydykov/go-rating/internal/rating"
	"github.com/stretchr/testify/mock"
	"google.golang.org/grpc/codes"

	amqp "github.com/rabbitmq/amqp091-go"
)

func (suite *IntegrationSuite) TestSavingRating_ShouldSuccessfullyPublishToRabbitMQ() {

	// Arrange
	config, err := config.LoadMessageBrokerConfig()
	if err != nil {
		log.Fatalf("Could not config: %v", err)
	}

	db := application.InitializeDatabase()
	publisher := application.InitializeRabbitMqPublisher(config, &LoggerMock{})

	ratingService := rating.NewRatingService(db, publisher)

	// Act
	err = ratingService.Create(&models.Rating{
		BookId: 1,
		Rating: 5,
	})
	if err != nil {
		log.Fatalf("Could not create rating: %v", err)
	}

	// consume message from rabbitmq -> confirm it is correct one
	rabbitMQURL := "amqp://" + config.RabbitMqUser + ":" + config.RabbitMqPassword + "@" + config.RabbitMqContainerName + ":5672/"
	conn, ch := newConsumer(rabbitMQURL, config.RabbitMqQueueName)
	defer closeConsumer(conn, ch)

	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		fmt.Println("Waiting for 2 seconds")
		time.Sleep(2 * time.Second)
		closeConsumer(conn, ch)
		publisher.Close()
		defer wg.Done()
	}()

	msgs, err := ch.Consume(
		config.RabbitMqQueueName, // queue
		"",                       // consumer
		true,                     // auto-ack
		false,                    // exclusive
		false,                    // no-local
		false,                    // no-wait
		nil,                      // args
	)
	if err != nil {
		log.Fatalf("failed to register a consumer: %v", err)
	}

	for msg := range msgs {
		var bookMessage rating.BookMessage
		if err := json.Unmarshal(msg.Body, &bookMessage); err != nil {
			log.Fatalf("failed to unmarshal message: %v", err)
		}

		fmt.Println("bookMessage: ", bookMessage)

		// Assert
		suite.Suite.Assert().Equal("bookRated", bookMessage.Event)
		suite.Suite.Assert().Equal(1, bookMessage.BookId)
	}

	wg.Wait()

	suite.db.Unscoped().Where("1 = 1").Delete(&models.Rating{})
}
func (suite *IntegrationSuite) TestRabbitMQ_ShouldLogErrorMessageIfCannotConnectToRabbitMQ() {
	// Arrange
	config, err := config.LoadMessageBrokerConfig()
	if err != nil {
		log.Fatalf("Could not config: %v", err)
	}
	config.RabbitMqContainerName = "invalid"
	loggerMock := &LoggerMock{}

	// Assert
	loggerMock.On("LogError",
		mock.Anything,
		mock.MatchedBy(func(err error) bool {
			return err != nil && strings.Contains(err.Error(), "failed to initialize message publisher")
		}),
	).Once()

	//act
	_ = application.InitializeRabbitMqPublisher(config, loggerMock)
}

func newConsumer(url, queueName string) (*amqp.Connection, *amqp.Channel) {
	conn, err := amqp.Dial(url)
	if err != nil {
		log.Fatalf("failed to connect to RabbitMQ: %v", err)
	}

	ch, err := conn.Channel()
	if err != nil {
		log.Fatalf("failed to open a channel: %v", err)
	}

	// Declare the queue
	_, err = ch.QueueDeclare(
		queueName, // queue name
		true,      // durable
		false,     // auto-delete
		false,     // exclusive
		false,     // no-wait
		nil,       // arguments
	)
	if err != nil {
		log.Fatalf("failed to declare queue: %v", err)
	}
	return conn, ch
}

func closeConsumer(conn *amqp.Connection, ch *amqp.Channel) {
	if ch != nil {
		ch.Close()
	}
	if conn != nil {
		conn.Close()
	}
}

type LoggerMock struct {
	mock.Mock
}

func (m *LoggerMock) LogError(statusCode codes.Code, err error) {
	fmt.Printf("statusCode %T\n", statusCode)
	fmt.Println("statusCode: ", (statusCode == codes.Aborted), statusCode)
	fmt.Println("LogError", statusCode, err)
	m.Called(statusCode, err)
}
