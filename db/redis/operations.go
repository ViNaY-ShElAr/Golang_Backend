package redis

import (
	"time"
	"context"
)

func GetKey(key string) (interface{}, error) {
    val, err := RedisClient.Get(context.Background(), key).Result()
    if err != nil {
        return nil, err
    }
    return val, nil
}

func SetKey(key string, value interface{}) error {
    err := RedisClient.Set(context.Background(), key, value, 0).Err()
    if err != nil {
        return err
    }
    return nil
}

func SetKeyWithExpiry(key string, value interface{}, ttl time.Duration) error {
    err := RedisClient.Set(context.Background(), key, value, time.Second * ttl).Err()
    if err != nil {
        return err
    }
    return nil
}

func DeleteKey(key string) error {
    err := RedisClient.Del(context.Background(), key).Err()
    if err != nil {
        return err
    }
    return nil
}

func IncrementKey(key string) (int64, error) {
    value, err := RedisClient.Incr(context.Background(), key).Result()
    if err != nil {
        return 0, err
    }
    return value, nil
}

func DecrementKey(key string) (int64, error) {
    value, err := RedisClient.Decr(context.Background(), key).Result()
    if err != nil {
        return 0, err
    }
    return value, nil
}

func IsKeyExist(key string) (bool, error) {
    exists, err := RedisClient.Exists(context.Background(), key).Result()
    if err != nil {
        return false, err
    }
    return exists > 0, nil
}