package handlers

import (
	"github.com/chyngyz-sydykov/go-rating/infrastructure/logger"
)

type AppHandlerInterface interface {
}

type ErrorDetail struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

type ErrorResponse struct {
	Error ErrorDetail `json:"error"`
}

type CommonHandler struct {
	logger *logger.Logger
}

const INVALID_REQUEST string = "INVALID_REQUEST"
const RESOURCE_NOT_FOUND string = "RESOURCE_NOT_FOUND"
const SERVER_ERROR string = "SERVER_ERROR"

func NewCommonHandler(logger *logger.Logger) *CommonHandler {

	return &CommonHandler{logger: logger}
}

func (c *CommonHandler) HandleError(err error, statusCode int) {
	c.logger.LogError(statusCode, err)
}
