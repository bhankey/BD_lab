package repository

import (
	"errors"
	"finance/pkg/logger"
)

var NoEntity = errors.New("no entity was found")

type BaseRepository struct {
	Logger logger.Logger
}

func NewRepository(l logger.Logger) *BaseRepository {
	return &BaseRepository{
		Logger: l,
	}
}
