package app

import (
	"context"
	"finance/internal/config"
	"finance/pkg/postgresdb"
	"fmt"
	"github.com/go-redis/redis/v8"
	"github.com/jmoiron/sqlx"
)

type dataSources struct {
	db          *sqlx.DB
	redisClient *redis.Client
}

func newDataSource(config *config.Config) (*dataSources, error) {
	db, err := postgresdb.NewClient(config.Postgres.Host, config.Postgres.Port, config.Postgres.User, config.Postgres.Password, config.Postgres.DBName)
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
		db:          db,
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
