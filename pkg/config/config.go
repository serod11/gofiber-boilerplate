/*
 * @author serod
 */

package config

import (
	"github.com/joho/godotenv"
	"os"
)

type Config struct {
	Port  string
	Pgsql Pgsql
	Redis Redis
}

func LoadConfig() (c Config, err error) {
	if os.Getenv("GO_ENV") != "PRODUCTION" {
		if e := godotenv.Load("pkg/config/envs/.env"); e != nil {
			return Config{}, e
		}
	}

	return Config{
		Port: os.Getenv("PORT"),
		Pgsql: Pgsql{
			Host:     os.Getenv("PG_HOST"),
			Port:     os.Getenv("PG_PORT"),
			Dbname:   os.Getenv("PG_DBNAME"),
			Username: os.Getenv("PG_USERNAME"),
			Password: os.Getenv("PG_PASSWORD"),
		},
		Redis: Redis{
			Url:      os.Getenv("REDIS_URL"),
			Password: os.Getenv("REDIS_PASS"),
		},
	}, nil
}
