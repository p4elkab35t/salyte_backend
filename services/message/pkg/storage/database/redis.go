package database

import (
	"context"
	"fmt"
	"sync"

	"github.com/go-redis/redis/v8"
)

type redisDB struct {
	client *redis.Client
}

var (
	redisInstance *redisDB
	redisOnce     sync.Once
)

func NewRedis(ctx context.Context, addr, password string, db int) (*redisDB, error) {
	redisOnce.Do(func() {
		client := redis.NewClient(&redis.Options{
			Addr:     addr,
			Password: password,
			DB:       db,
		})

		_, err := client.Ping(ctx).Result()
		if err != nil {
			redisInstance = nil
			return
		}

		redisInstance = &redisDB{client}
	})

	if redisInstance == nil {
		return nil, fmt.Errorf("unable to create redis client")
	}

	return redisInstance, nil
}

func (r *redisDB) Close() {
	r.client.Close()
}

func (r *redisDB) GetClient() *redis.Client {
	return r.client
}
