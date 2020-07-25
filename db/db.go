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
