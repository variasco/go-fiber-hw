package config

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	Database DatabaseConfig
	Server   ServerConfig
}

type DatabaseConfig struct {
	Url string
}

type ServerConfig struct {
	Port int
}

func LoadConfig() *Config {
	err := godotenv.Load()
	if err != nil {
		log.Println("Can't load config from .env file: ", err)
	}

	return &Config{
		Database: DatabaseConfig{
			Url: getString("DB_URL", ""),
		},
		Server: ServerConfig{
			Port: getInt("PORT", 3000),
		},
	}
}

func getString(key string, def string) string {
	value := os.Getenv(key)

	if len(value) == 0 {
		return def
	}

	return value
}

func getInt(key string, def int) int {
	value := os.Getenv(key)

	if len(value) == 0 {
		return def
	}

	integer, err := strconv.Atoi(value)
	if err != nil {
		return def
	}

	return integer
}
