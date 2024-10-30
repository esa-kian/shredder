package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	"github.com/esa-kian/shredder/pkg/api"
	"github.com/esa-kian/shredder/pkg/db"
	"github.com/gorilla/mux"
)

func main() {
	entity := flag.String("entity", "", "The name of the entity for CRUD operations")
	flag.Parse()

	if *entity == "" {
		log.Fatal("Entity name is required")
	}

	// Mock database config - would be configured through env variables in a real app
	dbConfig := db.DBConfig{
		Driver:   "postgres", // or "mysql"
		Host:     "localhost",
		Port:     5432,
		User:     "youruser",
		Password: "yourpassword",
		DBName:   "yourdbname",
	}

	dbConn, err := db.NewConnection(dbConfig)
	if err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
	}
	defer dbConn.Close()

	router := mux.NewRouter()
	handler := api.CRUDHandler{
		EntityName: *entity,
		DB:         dbConn,
	}
	handler.RegisterRoutes(router)

	fmt.Printf("Starting Shredder server for %s entity...\n", *entity)
	log.Fatal(http.ListenAndServe(":8080", router))
}
