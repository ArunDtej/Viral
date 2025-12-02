package db

import (
	"Viral/common"
	"context"
	"database/sql" // New import for SQL database operations
	"fmt"
	"log"
	"os" // New import for os.Getenv
	"time"

	_ "github.com/lib/pq" // New import for PostgreSQL driver
	"github.com/redis/go-redis/v9"
)

var (
	ctx = context.Background()
	rdb *redis.Client
	Pdb *sql.DB // New global variable for PostgreSQL database connection
)

func InitDB() error {
	// if err := initKV(common.KVAddr, common.KVPassword, common.KVDB); err != nil {
	// 	return err
	// }
	if err := initPostgres(); err != nil {
		return err
	}
	return nil
}

func initKV(addr, password string, db int) error {
	rdb = redis.NewClient(&redis.Options{
		Addr:         addr,               // "localhost:6666" or your Kvrocks host
		Password:     password,           // "" if none
		DB:           db,                 // typically 0
		PoolSize:     common.KVPoolSize,  // number of connections in pool
		MinIdleConns: common.KVIdleConns, // keep a few always ready
		PoolTimeout:  30 * time.Second,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 5 * time.Second,
	})

	// Test connection
	if err := rdb.Ping(ctx).Err(); err != nil {
		return fmt.Errorf("❌ Redis connection failed: %v", err)
	}
	fmt.Println("✅ Redis connection pool ready")
	return nil
}

func GetKVConnection() *redis.Client {
	return rdb
}

func initPostgres() error {
	dbUser := os.Getenv("POSTGRES_USER")
	if dbUser == "" {
		dbUser = "user" // Default user
	}
	dbPassword := os.Getenv("POSTGRES_PASSWORD")
	if dbPassword == "" {
		dbPassword = "password" // Default password
	}
	dbName := os.Getenv("POSTGRES_DB")
	if dbName == "" {
		dbName = "viral_db" // Default database name
	}
	dbHost := os.Getenv("POSTGRES_HOST")
	if dbHost == "" {
		dbHost = "localhost:5432" // Default host and port
	}

	postgresURL := fmt.Sprintf("postgres://%s:%s@%s/%s?sslmode=disable", dbUser, dbPassword, dbHost, dbName)
	log.Printf("Attempting to connect to PostgreSQL with URL: %s", postgresURL) // Added logging

	var err error
	Pdb, err = sql.Open("postgres", postgresURL)
	if err != nil {
		return fmt.Errorf("❌ Error opening PostgreSQL database connection: %v", err)
	}

	// Ping PostgreSQL to verify connection
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	err = Pdb.PingContext(ctx)
	if err != nil {
		return fmt.Errorf("❌ Could not connect to PostgreSQL: %v", err)
	}
	fmt.Println("✅ PostgreSQL connection pool ready")

	// Create predictions table and indexes if they don't exist
	createTableAndIndexesSQL := `
	CREATE TABLE IF NOT EXISTS predictions (
		id SERIAL PRIMARY KEY,
		uid TEXT UNIQUE, -- Application will generate random unique IDs
		page_type TEXT NOT NULL,
		user_data JSONB,
		predictions JSONB,
		created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
	);
	CREATE UNIQUE INDEX IF NOT EXISTS idx_predictions_uid ON predictions (uid);
	CREATE INDEX IF NOT EXISTS idx_predictions_created_at ON predictions (created_at);`

	_, err = Pdb.Exec(createTableAndIndexesSQL)
	if err != nil {
		return fmt.Errorf("❌ Error creating predictions table or indexes in PostgreSQL: %v", err)
	}
	fmt.Println("✅ PostgreSQL predictions table and indexes checked/created successfully!")

	return nil
}
