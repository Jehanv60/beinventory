package app

import (
	"database/sql"
	"fmt"
	"os"
	"time"

	"github.com/Jehanv60/helper"
	_ "github.com/lib/pq"
)

var DB *sql.DB

func NewDb() *sql.DB {
	var err error
	helper.GoDoEnv()
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Jakarta", os.Getenv("DB_HOST"), os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_NAME"), os.Getenv("DB_PORT"))
	DB, err = sql.Open("postgres", dsn)
	if err != nil {
		panic(err)
	}
	DB.SetMaxIdleConns(5)
	DB.SetMaxOpenConns(20)
	DB.SetConnMaxIdleTime(10 * time.Minute)
	DB.SetConnMaxLifetime(60 * time.Minute)
	return DB
}
