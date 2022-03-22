package drivers

import (
	"fmt"
	"runtime"

	"com.thebeachmaster/entexample/config"
	"github.com/go-redis/redis/v8"
)

func NewRedisDBClient(cfg *config.Config) (*redis.Client, error) {

	host := cfg.Cache.RediscloudURL
	if host == "" {
		return nil, fmt.Errorf("RedisURL not found")
	}

	redisOpt, err := redis.ParseURL(host)
	if err != nil {
		return nil, err
	}

	redisOpt.PoolSize = runtime.NumCPU()

	redisClient := redis.NewClient(redisOpt)

	return redisClient, nil

}
