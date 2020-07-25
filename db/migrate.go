package db

import "context"

func Migrate() error {
	_, err := db.Exec(context.Background(),
		`CREATE TABLE IF NOT EXISTS users(
		id SERIAL PRIMARY KEY, 
		username VARCHAR NOT NULL,
		salt VARCHAR NOT NULL,
		password VARCHAR NOT NULL
	);`)
	if err != nil {
		return err
	}

	_, err = db.Exec(context.Background(),
		`CREATE TABLE IF NOT EXISTS transfers(
		id SERIAL PRIMARY KEY,
		owner INTEGER,
		size BIGINT,
		final_size BIGINT,
		path VARCHAR,
		status VARCHAR,
		expires TIMESTAMP WITH TIME ZONE
		);`)
	if err != nil {
		return err
	}

	return err
}
