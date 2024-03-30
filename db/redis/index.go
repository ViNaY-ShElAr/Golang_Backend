package redis

import (
	"context"
	"github.com/redis/go-redis/v9"
	"os"
	"time"

	"GO_PROJECT/logger"
	"GO_PROJECT/model"
)

var RedisClient *redis.Client

func ConnectRedisDb() {

	cfg := model.RedisCfg{
		Host:     os.Getenv("REDIS_HOST"),
		Port:     os.Getenv("REDIS_PORT"),
		Password: os.Getenv("REDIS_PASSWORD"),
	}

	RedisClient = redis.NewClient(&redis.Options{
		Addr:     cfg.Host + ":" + cfg.Port,
		Password: cfg.Password,
		DB:       0,
		PoolSize: 10,
	})

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if _, err := RedisClient.Ping(ctx).Result(); err != nil {
		logger.Log.Fatal("Error in Connecting Redis Db :", err)
	}
	logger.Log.Info("Redis Db Connected Successfully")
}
