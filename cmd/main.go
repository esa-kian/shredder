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

	dbConfig := db.DBConfig{
		Driver:   "mysql",
		Host:     "localhost",
		Port:     3306,
		User:     "root",
		Password: "",
		DBName:   "mydb",
	}

	dbConn, err := db.NewConnection(dbConfig)
	if err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
	}
	defer dbConn.Close()

	router := mux.NewRouter()
	router.Use(api.LoggingMiddleware)

	handler := api.CRUDHandler{
		EntityName: *entity,
		DB:         dbConn,
	}
	handler.RegisterRoutes(router)

	fmt.Printf("Starting Shredder server for %s entity...\n", *entity)
	log.Fatal(http.ListenAndServe(":8080", router))
}
