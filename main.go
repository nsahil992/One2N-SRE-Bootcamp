package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

func main() {
	cfg := LoadConfig()

	db, err := sql.Open("postgres", BuildDBConnectionString(cfg))
	if err != nil {
		log.Fatalf("failed to connect to DB: %v", err)
	}
	// Check error when closing
	defer func() {
		if err := db.Close(); err != nil {
			log.Printf("failed to close db: %v", err)
		}
	}()

	r := SetupRouter(db)

	// Check error when running server
	if err := r.Run(":" + cfg.Port); err != nil {
		log.Fatalf("server failed to start: %v", err)
	}
}

// BuildDBConnectionString builds PostgreSQL connection string from Config
func BuildDBConnectionString(cfg Config) string {
	return fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		cfg.DBHost, cfg.DBPort, cfg.DBUser, cfg.DBPassword, cfg.DBName, cfg.DBSSLMode)
}
