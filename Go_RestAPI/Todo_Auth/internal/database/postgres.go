package database

import (
	"context"
	"log"

	"github.com/jackc/pgx/v5/pgxpool"
)

func Connect(databaseURL string) (*pgxpool.Pool, error) {
	var ctx context.Context = context.Background()
	var config *pgxpool.Config
	var err error

	config, err = pgxpool.ParseConfig(databaseURL)
	if err != nil {
		return nil, err
	}

	var pool *pgxpool.Pool
	pool, err = pgxpool.NewWithConfig(ctx, config)
	if err != nil {
		log.Println("Error while connecting to database", err)
		return nil, err
	}

	err = pool.Ping(ctx)
	if err != nil {
		log.Println("Error while pinging database", err)
		pool.Close()
		return nil, err
	}

	return pool, nil
}
