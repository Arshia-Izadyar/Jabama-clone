package cache

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"

	"github.com/Arshia-Izadyar/Jabama-clone/src/config"
	"github.com/Arshia-Izadyar/Jabama-clone/src/pkg/logger"
)

var redisClient *redis.Client
var logging = logger.NewLogger(config.GetConfig())

func InitRedis(cfg *config.Config) error {
	redisClient = redis.NewClient(&redis.Options{
		Addr:         fmt.Sprintf("%s:%d", cfg.Redis.Host, cfg.Redis.Port),
		Password:     cfg.Redis.Password,
		DB:           0,
		DialTimeout:  cfg.Redis.DialTimeout * time.Second,
		ReadTimeout:  cfg.Redis.ReadTimeout * time.Second,
		WriteTimeout: cfg.Redis.WriteTimeout * time.Second,
		PoolSize:     cfg.Redis.PoolSize,
	})

	_, err := redisClient.Ping(context.Background()).Result()
	if err != nil {
		return err
	}
	return nil
}
func GetRedis() *redis.Client {
	return redisClient
}
func CloseRedis() {

	err := redisClient.Close()
	if err != nil {
		logging.Fatal(logger.Redis, logger.Close, err, nil)
	}
}

func Set[T any](key string, value T, ttl time.Duration) error {
	v, err := json.Marshal(value)
	if err != nil {
		return err
	}
	_, err = redisClient.Set(context.Background(), key, v, ttl*time.Minute).Result()
	if err != nil {
		return err
	}
	return nil
}

func Get[T any](key string, c *redis.Client) (*T, error) {
	v, err := c.Get(context.Background(), key).Result()
	dest := *new(T)
	if err != nil {
		return nil, err
	}
	json.Unmarshal([]byte(v), &dest)
	return &dest, nil
}
