package service

import (
	"github.com/bhankey/BD_lab/backend/pkg/logger"
)

type BaseService struct {
	Logger logger.Logger
}

func NewService(l logger.Logger) *BaseService {
	return &BaseService{
		Logger: l,
	}
}
