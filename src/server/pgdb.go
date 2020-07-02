package server

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq" // sql driver
)

func openPosgresDB(connectionString string) (*sql.DB, error) {
	log.Println("Try connect to PG...")

	db, err := sql.Open("postgres", connectionString)
	if err != nil {
		return nil, err
	}

	_, err = db.Exec(`CREATE EXTENSION IF NOT EXISTS pgcrypto;`)
	if err != nil {
		return nil, err
	}

	log.Println("PG pool connected!")
	return db, nil
}
