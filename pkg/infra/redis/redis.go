/*
 * @author serod
 */

package redis

import (
	"github.com/go-redis/redis/v8"
	"github.com/serod11/gofiber-boilerplate/pkg/utils/config"
)

func NewRedisClient(c config.Config) *redis.Client {
	cache := redis.NewClient(&redis.Options{
		Addr:     c.Redis.Url,
		Password: c.Redis.Password,
	})

	return cache
}
