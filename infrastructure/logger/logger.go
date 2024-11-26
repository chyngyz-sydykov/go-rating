package logger

import (
	"log"
	"os"

	"google.golang.org/grpc/codes"
)

type Logger struct {
}

type LoggerInterface interface {
	LogError(statusCode codes.Code, err error)
}

func NewLogger() *Logger {
	return &Logger{}
}

func (l *Logger) LogError(statusCode codes.Code, err error) {
	errorLog := log.New(os.Stderr, "ERROR: ", log.Ldate|log.Ltime)
	errorLog.Printf("status code: %d, error: %s", statusCode, err.Error())
}
