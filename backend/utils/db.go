package utils

import (
    "context"
    "log"
    "os"
    "github.com/jackc/pgx/v5/pgxpool"
    "github.com/joho/godotenv"
)

// ConnectDB initializes and returns a PostgreSQL connection pool.
func ConnectDB() *pgxpool.Pool {
    // Load environment variables from .env file
    err := godotenv.Load()
    if err != nil {
        log.Fatal("Error loading .env file")
    }

    // Construct the database URL
    dbURL := "postgres://" + os.Getenv("DB_USER") + ":" + os.Getenv("DB_PASSWORD") + "@" + os.Getenv("DB_HOST") + ":" + os.Getenv("DB_PORT") + "/" + os.Getenv("DB_NAME")

    // Create a connection pool
    db, err := pgxpool.New(context.Background(), dbURL)
    if err != nil {
        log.Fatal("Unable to connect to database:", err)
    }
    return db
}
