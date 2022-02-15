package app

import (
	"context"
	"fmt"

	"github.com/bhankey/BD_lab/backend/internal/config"
	"github.com/bhankey/BD_lab/backend/pkg/postgresdb"
	"github.com/go-redis/redis/v8"
	"github.com/jmoiron/sqlx"
)

type dataSources struct {
	db          *sqlx.DB
	redisClient *redis.Client
}

func newDataSource(config *config.Config) (*dataSources, error) {
	postgresDB, err := postgresdb.NewClient(
		config.Postgres.Host,
		config.Postgres.Port,
		config.Postgres.User,
		config.Postgres.Password,
		config.Postgres.DBName)
	if err != nil {
		return nil, err
	}

	rdb := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", config.Redis.Host, config.Redis.Port),
		Password: config.Redis.Password,
		DB:       0,
	})

	if _, err := rdb.Ping(context.Background()).Result(); err != nil {
		return nil, err
	}

	return &dataSources{
		db:          postgresDB,
		redisClient: rdb,
	}, nil
}

func (ds *dataSources) close() error {
	if err := ds.db.Close(); err != nil {
		return err
	}

	if err := ds.redisClient.Close(); err != nil {
		return err
	}

	return nil
}
