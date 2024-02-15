package config

import (
	"log"
	"os"
	"strconv"

	"github.com/dotenv-org/godotenvvault"
)

type MongoDB struct {
	URI      string
	DBName   string
	CollName string
}
type Redis struct {
	URI           string
	CACHE_TIMEOUT int
}

type Config struct {
	MongoDB MongoDB
	Redis   Redis
}

func Load() *Config {
	err := godotenvvault.Load()
	if err != nil {
		log.Fatal("Не удалось загрузить файл .env")
	}

	timeout, err := strconv.Atoi(os.Getenv("CACHE_TIMEOUT"))
	if err != nil {
		timeout = 0
	}
	return &Config{
		MongoDB: MongoDB{
			URI:      os.Getenv("DB_URI"),
			DBName:   os.Getenv("MONGO_INITDB_ROOT_USERNAME"),
			CollName: os.Getenv("DB_COLL_NAME"),
		},
		Redis: Redis{
			URI:           os.Getenv("REDIS_URI"),
			CACHE_TIMEOUT: timeout,
		},
	}
}
