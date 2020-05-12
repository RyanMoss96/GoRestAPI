package database

import (
	"log"

	"github.com/go-redis/redis"
)

type RedisClient struct {
	client *redis.Client
}

var defaultClient *redis.Client

func NewRedisClient(client *redis.Client) *redis.Client {
	defaultClient = client
	return defaultClient
}

func GetRedisClient() *redis.Client {
	if defaultClient == nil {
		log.Fatal("No RedisClient has been created.")
	}

	return defaultClient
}
