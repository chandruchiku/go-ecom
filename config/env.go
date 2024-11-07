package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type DBConfig struct {
	User   string
	Passwd string
	Addr   string
	Name   string
}

type Config struct {
	PublicHost string
	Port       string
	DBConfig   DBConfig
}

var Envs = initConfig()

func initConfig() Config {

	godotenv.Load()
	DBConfig := DBConfig{
		User:   getEnv("DB_USER", "root"),
		Passwd: getEnv("DB_PASSWD", "mypassword"),
		Addr:   fmt.Sprintf("%s:%s", getEnv("DB_HOST", "localhost"), getEnv("DB_PORT", "3306")),
		Name:   getEnv("DB_NAME", "ecom"),
	}
	return Config{
		PublicHost: getEnv("PUBLIC_HOST", "http://localhost"),
		Port:       getEnv("PORT", "8080"),
		DBConfig:   DBConfig,
	}
}

func getEnv(key string, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}
