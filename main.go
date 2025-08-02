package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	// Load .env file into environment variables (only errors if .env is missing)
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found or could not load, proceeding with environment variables")
	}

	// CLI subcommand parsing: migrate or run server
	if len(os.Args) > 1 && os.Args[1] == "migrate" {
		migrateCmd := flag.NewFlagSet("migrate", flag.ExitOnError)
		schemaPath := migrateCmd.String("c", "schema.sql", "Path to schema.sql")
		if err := migrateCmd.Parse(os.Args[2:]); err != nil { // <-- Fix: check error
			log.Fatalf("failed to parse migrate subcommand: %v", err)
		}

		cfg := LoadConfig()
		log.Println("Starting database migration...")
		if err := runMigrations(cfg, *schemaPath); err != nil {
			log.Fatalf("Migration failed: %v", err)
			os.Exit(1)
		}
		log.Println("Migration completed successfully")
		os.Exit(0) // Explicit exit after migration is done
	}

	// Run the API server normally
	cfg := LoadConfig()

	db := ConnectDB(cfg)
	defer func() {
		if cerr := db.Close(); cerr != nil {
			log.Printf("Error closing DB: %v", cerr)
		}
	}()

	router := SetupRouter(db)
	addr := fmt.Sprintf(":%s", cfg.Port)
	log.Printf("Starting server on %s", addr)
	if err := router.Run(addr); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
