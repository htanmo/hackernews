package database

import (
	"context"
	"log"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
)

var Pool *pgxpool.Pool

func ConnectDB() {
	dsn := os.Getenv("DATABASE_URL")
	if dsn == "" {
		log.Fatal("DATABASE_URL need to be set.")
	}
	ctx := context.Background()
	var err error
	Pool, err = pgxpool.New(ctx, dsn)
	if err != nil {
		log.Fatalf("Unable to create connection pool: %v\n", err)
	}

	err = Pool.Ping(ctx)
	if err != nil {
		log.Fatalf("Database ping failed: %v\n", err)
	}

	log.Println("Successfully connected to the database.")
}

func CloseDB() {
	if Pool != nil {
		Pool.Close()
	}
	log.Println("Database connection pool closed.")
}
