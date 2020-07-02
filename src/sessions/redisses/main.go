package redisses

import (
	"audit/src/config"
	"audit/src/sessions"
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/go-redis/redis/v7"
)

type redisStorage struct {
	client *redis.Client
}

// NewStorage create redis storage
func NewStorage(cfg *config.AppConfig) (sessions.IStorage, error) {
	client := redis.NewClient(&redis.Options{
		Addr: cfg.Redis.ConnectionString,
	})

	log.Println("Try to ping redis...")

	_, err := client.Ping().Result()
	if err != nil {
		return nil, fmt.Errorf("redis ping error %w", err)
	}

	log.Println("Redis client connected!")

	return &redisStorage{client}, nil
}

// Has returns false and empty error if value does not exist.
func (r *redisStorage) Has(key string) (bool, error) {
	_, err := r.client.Get(key).Result()

	if err == redis.Nil {
		return false, nil
	}

	if err != nil {
		return false, err
	}

	return true, nil
}

// Get returns empty string and empty error if value does not exist.
func (r *redisStorage) Get(key string) (string, error) {
	val, err := r.client.Get(key).Result()

	if err == redis.Nil {
		return "", nil
	}

	if err != nil {
		return "", err
	}

	return val, nil
}

// GetJSON returns nil if value does not exist.
func (r *redisStorage) GetJSON(key string) (*map[string]interface{}, error) {
	str, err := r.client.Get(key).Result()

	if err == redis.Nil {
		return nil, nil
	}

	if err != nil {
		return nil, err
	}

	res := make(map[string]interface{})
	err = json.Unmarshal([]byte(str), &res)

	if err != nil {
		return nil, err
	}

	return &res, err
}

// SetJSON set value by key and expiration in seconds
func (r *redisStorage) SetJSON(key string, data interface{}, expiration int) error {
	bytes, err := json.Marshal(data)
	if err != nil {
		return err
	}
	return r.Set(key, string(bytes), expiration)
}

//Set value by key and expiration in seconds
func (r *redisStorage) Set(key string, value string, expiration int) error {
	return r.client.Set(key, value, time.Duration(expiration)*time.Second).Err()
}

// Delete returns false and empty error if valuse does not exist.
func (r *redisStorage) Delete(key string) (bool, error) {
	err := r.client.Del(key).Err()

	if err == redis.Nil {
		return false, nil
	}

	if err != nil {
		return false, err
	}

	return true, nil
}
