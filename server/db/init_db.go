package db

import (
	"Viral/common"
	"context"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
)

var (
	ctx = context.Background()
	rdb *redis.Client
)

func InitDB() error {
	if err := initKV(common.KVAddr, common.KVPassword, common.KVDB); err != nil {
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
