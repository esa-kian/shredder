package main

import (
	"database/sql"
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/esa-kian/shredder/pkg/api"
	"github.com/esa-kian/shredder/pkg/db"
	"github.com/esa-kian/shredder/pkg/migration"
	"github.com/joho/godotenv"
)

func main() {
	entity := flag.String("entity", "", "The name of the entity for CRUD operations")
	flag.Parse()

	if *entity == "" {
		log.Fatal("Entity name is required")
	}

	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}
	dbConfig := db.DBConfig{
		Driver:   "mysql",
		Host:     "localhost",
		Port:     3306,
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
		DBName:   os.Getenv("DB_NAME"),
	}

	dbConn, err := db.NewConnection(dbConfig)
	if err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
	}
	defer dbConn.Close()

	// Create table for model
	if err := RunMigrations(dbConn); err != nil {
		log.Fatal("Migration error:", err)
	}

	if err := ensureDir("./controllers"); err != nil {
		log.Fatal("Failed to ensure controllers directory:", err)
	}
	// Generate controller and routes
	if err := api.GenerateControllerFile(*entity); err != nil {
		log.Fatal("Failed to generate controller file:", err)
	}

	if err := ensureDir("./routes"); err != nil {
		log.Fatal("Failed to ensure routes directory:", err)
	}
	if err := api.GenerateRoutesFile(*entity); err != nil {
		log.Fatal("Failed to generate routes file:", err)
	}

	fmt.Printf("Successfully created table, controller, and routes for %s entity.\n", *entity)

}

func RunMigrations(dbConn *sql.DB) error {
	modelsList, err := migration.LoadModelsFromMigrationDir()
	if err != nil {
		return fmt.Errorf("failed to load models: %w", err)
	}

	for _, model := range modelsList {
		exists, err := db.TableExists(dbConn, model.EntityName)
		if err != nil {
			return fmt.Errorf("error checking table existence: %w", err)
		}

		if !exists {
			if err := db.CreateTableFromModel(dbConn, model); err != nil {
				return fmt.Errorf("failed to create table for %s: %w", model.EntityName, err)
			}
			log.Printf("Table created for entity %s", model.EntityName)
		} else {
			log.Printf("Table for entity %s already exists, skipping", model.EntityName)
		}
	}

	return nil
}

func ensureDir(dirName string) error {
	if _, err := os.Stat(dirName); os.IsNotExist(err) {
		err = os.MkdirAll(dirName, 0755) // Creates the directory with read/write permissions
		if err != nil {
			return fmt.Errorf("failed to create directory %s: %w", dirName, err)
		}
	}
	return nil
}
