package server

import (
	"audit/src/config"
	"context"
	"log"

	"github.com/jackc/pgx/v4/pgxpool"
)

func createPgxPool(cfg *config.AppConfig) (*pgxpool.Pool, error) {
	pgCfg, err := pgxpool.ParseConfig(cfg.PG.ConnectionString)
	if err != nil {
		return nil, err
	}

	pool, err := pgxpool.ConnectConfig(context.Background(), pgCfg)
	if err != nil {
		return nil, err
	}

	log.Println("Try connect to PG...")
	_, err = pool.Exec(context.Background(), `CREATE EXTENSION IF NOT EXISTS pgcrypto;`)
	if err != nil {
		return nil, err
	}
	log.Println("PG pool connected!")

	return pool, nil
}
