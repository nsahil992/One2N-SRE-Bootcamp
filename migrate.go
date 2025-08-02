package main

import (
	"fmt"
	"log"
	"os" // <-- Use os, not io/ioutil
)

// runMigrations runs the SQL statements in the given schema file on the DB configured by cfg
func runMigrations(cfg Config, schemaFile string) error {
	log.Println("Connecting to database...")
	db := ConnectDB(cfg)
	defer func() {
		if err := db.Close(); err != nil {
			log.Printf("Error closing DB during migration: %v", err)
		}
	}()

	log.Printf("Reading schema file: %s", schemaFile)
	schema, err := os.ReadFile(schemaFile) // <-- Use os.ReadFile
	if err != nil {
		return fmt.Errorf("failed to read schema file: %w", err)
	}

	log.Printf("Executing schema SQL...")
	_, err = db.Exec(string(schema))
	if err != nil {
		return fmt.Errorf("failed to execute schema: %w", err)
	}

	log.Println("Schema executed successfully")
	return nil
}
