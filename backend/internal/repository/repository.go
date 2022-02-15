package repository

import (
	"errors"

	"github.com/bhankey/BD_lab/backend/pkg/logger"
)

var ErrNoEntity = errors.New("no entity was found")

type BaseRepository struct {
	Logger logger.Logger
}

func NewRepository(l logger.Logger) *BaseRepository {
	return &BaseRepository{
		Logger: l,
	}
}
