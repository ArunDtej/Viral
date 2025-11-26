package common

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

// Exported variables (Capitalized = public)
var (
	KVAddr       string
	KVPassword   string
	KVDB         int
	KVPoolSize   int
	KVIdleConns  int
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Println("⚠️ No .env file found, using system environment variables.")
	}

	// KV DB
	KVAddr = getEnv("KVRocks_ADDR", "localhost:6666")
	KVPassword = getEnv("KVRocks_PASSWORD", "")
	KVDB = getEnvAsInt("KVRocks_DB", 0)
	KVPoolSize = getEnvAsInt("KVPoolSize", 100)
	KVIdleConns = getEnvAsInt("KVIdleConns", 25)

	// QuickWit DB
}

func getEnv(key, fallback string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return fallback
}

func getEnvAsInt(key string, fallback int) int {
	valStr := getEnv(key, "")
	if val, err := strconv.Atoi(valStr); err == nil {
		return val
	}
	return fallback
}
