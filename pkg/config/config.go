package config

import (
	"os"
	"strconv"
)

type Config struct {
	AppPort      string
	DBHost       string
	DBPort       string
	DBUser       string
	DBPassword   string
	DBName       string
	RedisAddr    string
	AWSAccessKey string
	AWSSecretKey string
	AWSBucket    string
	JWTSecret    string
	RateLimit    int
}

func LoadConfig() *Config {
	return &Config{
		AppPort:      getEnv("APP_PORT", "8080"),
		DBHost:       getEnv("DB_HOST", "localhost"),
		DBPort:       getEnv("DB_PORT", "5432"),
		DBUser:       getEnv("DB_USER", "postgres"),
		DBPassword:   getEnv("DB_PASSWORD", "secret"),
		DBName:       getEnv("DB_NAME", "file_sharing"),
		RedisAddr:    getEnv("REDIS_ADDR", "localhost:6379"),
		AWSAccessKey: getEnv("AWS_ACCESS_KEY_ID", "minioadmin"),
		AWSSecretKey: getEnv("AWS_SECRET_ACCESS_KEY", "minioadmin"),
		AWSBucket:    getEnv("AWS_BUCKET_NAME", "uploads"),
		JWTSecret:    getEnv("JWT_SECRET", "supersecretkey"),
		RateLimit:    getIntEnv("RATE_LIMIT", 100),
	}
}

func getEnv(key, fallback string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return fallback
}

func getIntEnv(key string, fallback int) int {
	if value, exists := os.LookupEnv(key); exists {
		return atoi(value)
	}
	return fallback
}

func atoi(s string) int {
	v, _ := strconv.Atoi(s)
	return v
}
