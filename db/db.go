package db

import (
	"context"

	"github.com/jackc/pgx/v4/pgxpool"
)

var db *pgxpool.Pool

func Connect(conStr string) (err error) {
	db, err = pgxpool.Connect(context.Background(), conStr)

	return
}

func Migrate() error {
	_, err := db.Exec(context.Background(),
		`CREATE TABLE IF NOT EXISTS users(
		id SERIAL PRIMARY KEY, 
		username VARCHAR NOT NULL,
		salt VARCHAR NOT NULL,
		password VARCHAR NOT NULL
	);`)

	return err
}
