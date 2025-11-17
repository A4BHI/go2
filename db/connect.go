package db

import (
	"context"
	"log"

	"github.com/jackc/pgx/v5/pgxpool"
)

var Pool *pgxpool.Pool

func Connect() {
	var err error
	Pool, err = pgxpool.New(context.Background(), "postgres://a4bhi:a4bhi@localhost:5432/go2")
	if err != nil {
		log.Fatal("DB connection error:", err)
	}

	log.Println("Connected to PostgreSQL")
}
