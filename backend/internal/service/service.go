package service

import (
	"finance/pkg/logger"
)

type BaseService struct {
	Logger logger.Logger
}

func NewService(l logger.Logger) *BaseService {
	return &BaseService{
		Logger: l,
	}
}
