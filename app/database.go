package app

import (
	"database/sql"
	"time"

	_ "github.com/lib/pq"
)

func NewDb() *sql.DB {
	db, err := sql.Open("postgres", "host=localhost user=han port=5432 password=solo dbname=pos1 sslmode=disable")
	if err != nil {
		panic(err)
	}
	db.SetMaxIdleConns(5)
	db.SetMaxOpenConns(20)
	db.SetConnMaxIdleTime(10 * time.Minute)
	db.SetConnMaxLifetime(60 * time.Minute)
	return db
}
