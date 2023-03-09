package pgsql

import (
	"context"
	"database/sql"
	"time"

	//have to add this import cause postgres driver will not be recognized
	_ "github.com/lib/pq"
)

func NewClient(ctx context.Context, dsn string, maxOpenCons int, maxIdleCons int, maxIdleTime string) (*sql.DB, error) {
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, err
	}

	db.SetMaxOpenConns(maxOpenCons)

	db.SetMaxIdleConns(maxIdleCons)

	duration, err := time.ParseDuration(maxIdleTime)
	if err != nil {
		return nil, err
	}

	db.SetConnMaxIdleTime(duration)

	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	if err = db.PingContext(ctx); err != nil {
		return nil, err
	}

	return db, nil
}
