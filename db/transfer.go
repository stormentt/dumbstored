package db

import (
	"context"
	"fmt"
	"time"
)

type Transfer struct {
	ID        int
	Owner     int
	Size      int64
	FinalSize int64
	Status    string
	Expires   time.Time
}

func (t Transfer) String() string {
	return fmt.Sprintf("Transfer<%d,%d,%d/%d,%s,%s>",
		t.ID,
		t.Owner,
		t.Size,
		t.FinalSize,
		t.Status,
		t.Expires)
}

func CreateTransfer(owner int, finalSize int64) (int, error) {
	row := db.QueryRow(context.Background(),
		`INSERT INTO transfers(owner, final_size, expires) VALUES($1, $2, $3) RETURNING id`,
		owner,
		finalSize,
		time.Now().Add(time.Hour*24*7))
	id := 0
	err := row.Scan(&id)

	return id, err
}
