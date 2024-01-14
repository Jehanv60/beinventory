package app

import (
	"database/sql"
	"time"

	_ "github.com/lib/pq"
)

var DB *sql.DB

func NewDb() *sql.DB {
	var err error
	DB, err = sql.Open("postgres", "host=localhost user=han port=5432 password=solo dbname=pos1 sslmode=disable")
	if err != nil {
		panic(err)
	}
	DB.SetMaxIdleConns(5)
	DB.SetMaxOpenConns(20)
	DB.SetConnMaxIdleTime(10 * time.Minute)
	DB.SetConnMaxLifetime(60 * time.Minute)
	return DB
}
