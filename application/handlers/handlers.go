package handlers

import (
	"github.com/chyngyz-sydykov/go-rating/infrastructure/logger"
	"google.golang.org/grpc/codes"
)

type CommonHandler struct {
	logger logger.LoggerInterface
}

const INVALID_REQUEST string = "INVALID_REQUEST"
const RESOURCE_NOT_FOUND string = "RESOURCE_NOT_FOUND"
const SERVER_ERROR string = "SERVER_ERROR"

func NewCommonHandler(logger logger.LoggerInterface) *CommonHandler {

	return &CommonHandler{logger: logger}
}

func (c *CommonHandler) HandleError(statusCode codes.Code, err error) {
	c.logger.LogError(statusCode, err)
}
