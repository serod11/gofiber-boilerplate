/*
 * @author serod
 */

package db

import (
	"github.com/go-redis/redis/v8"
	"github.com/serod11/gofiber-boilerplate/pkg/config"
)

func NewRedisClient(c config.Config) *redis.Client {
	cache := redis.NewClient(&redis.Options{
		Addr:     c.Redis.Url,
		Password: c.Redis.Password,
	})

	return cache
}
