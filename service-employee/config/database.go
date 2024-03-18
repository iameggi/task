package config

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	_ "github.com/lib/pq"
)

var db *sql.DB

func init() {
	connectionString := "postgres://postgres:1@localhost:5432/service?sslmode=disable"
	var err error
	db, err = sql.Open("postgres", connectionString)
	if err != nil {
		panic(err)
	}

	if err = db.Ping(); err != nil {
		panic(err)
	}

	fmt.Println("Connected to PostgreSQL database")
}

func NewPostgresContext() (context.Context, context.CancelFunc) {
	return context.WithTimeout(context.Background(), 10*time.Second)
}

func GetPostgresDB() *sql.DB {
	return db
}
