package routes

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
	"golang.org/x/tools/go/analysis/passes/defers"
)

var DB *pgxpool.Pool

func ConnectDB(){
	dbURL := "postgres://postgres:password@localhost:5432/book_db"

	ctx , cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	pool , err := pgxpool.New(ctx,dbURL)
	if err != nil  {
		log.Fatal("Failed to connect to DB:", err)
	}

	if err := pool.Ping(ctx); err != nil {
		log.Fatal("DB not reachable:", err)
	}

	DB = pool
	log.Println("PostgreSQL connected ✅")
}
