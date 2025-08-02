package main

import (
	"log"
	"os"
)

type Config struct {
	DBHost     string
	DBPort     string
	DBUser     string
	DBPassword string
	DBName     string
	DBSSLMode  string
	Port       string
}

// LoadConfig loads config values from environment variables
func LoadConfig() Config {
	cfg := Config{
		DBHost:     getEnv("DB_HOST", "localhost"),
		DBPort:     getEnv("DB_PORT", "5432"),
		DBUser:     getEnv("DB_USER", "postgres"),
		DBPassword: getEnv("DB_PASSWORD", ""),
		DBName:     getEnv("DB_NAME", "students"),
		DBSSLMode:  getEnv("DB_SSLMODE", "disable"),
		Port:       getEnv("PORT", "8080"),
	}

	if cfg.DBPassword == "" {
		log.Println("Warning: DB_PASSWORD is empty")
	}

	return cfg
}

func getEnv(key, fallback string) string {
	if val, ok := os.LookupEnv(key); ok && val != "" {
		return val
	}
	return fallback
}
