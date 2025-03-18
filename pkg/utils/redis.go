package utils

import (
	"context"
	"encoding/json"
	"file-sharing-backend/pkg/config"
	"time"

	"github.com/go-redis/redis/v9"
)

type RedisClient struct {
	Client *redis.Client
}

func InitRedis(cfg *config.Config) *RedisClient {
	return &RedisClient{
		Client: redis.NewClient(&redis.Options{
			Addr:     cfg.RedisAddr,
			Password: "",
			DB:       0,
		}),
	}
}

func (r *RedisClient) Get(c *gin.Context, key string) interface{} {
	val, err := r.Client.Get(context.Background(), key).Result()
	if err != nil {
		return nil
	}
	var data interface{}
	json.Unmarshal([]byte(val), &data)
	return data
}

func (r *RedisClient) Set(c *gin.Context, key string, value interface{}, ttl int) {
	data, _ := json.Marshal(value)
	r.Client.Set(context.Background(), key, data, time.Duration(ttl)*time.Second)
}

func (r *RedisClient) Delete(c *gin.Context, key string) {
	r.Client.Del(context.Background(), key)
}
