package sessions

import (
	"audit/src/config"
	"fmt"
	"log"
	"time"

	"github.com/go-redis/redis/v7"
)

type redisStorage struct {
	client *redis.Client
}

// NewRedisStorage create redis storage
func NewRedisStorage(cfg *config.AppConfig) (IStorage, error) {
	addr := fmt.Sprintf("%s:%v", cfg.Redis.Host, cfg.Redis.Port)

	client := redis.NewClient(&redis.Options{
		Addr: addr,
	})

	log.Println("Try to ping redis...")

	_, err := client.Ping().Result()
	if err != nil {
		return nil, fmt.Errorf("redis ping error %w", err)
	}

	log.Println("Redis client connected!")

	return &redisStorage{client}, nil
}

func (r *redisStorage) Get(key string) (string, error) {
	val, err := r.client.Get(key).Result()
	if err != nil {
		return "", err
	}
	return val, nil
}

func (r *redisStorage) Set(key string, value string, expiration int) error {
	return r.client.Set(key, value, time.Duration(expiration)*time.Second).Err()
}

func (r *redisStorage) Delete(key string) error {
	_, err := r.client.Del(key).Result()
	return err
}
