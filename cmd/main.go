package main

import (
	"flag"
	"log"

	"github.com/esa-kian/shredder/pkg/db"
	"github.com/esa-kian/shredder/pkg/models"
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

	userModel := models.Model{
		EntityName: "User",
		Fields: []models.Field{
			{Name: "id", DataType: "INT", IsPrimaryKey: true},
			{Name: "name", DataType: "VARCHAR(255)", IsRequired: true},
			{Name: "email", DataType: "VARCHAR(255)", IsRequired: true},
			{Name: "age", DataType: "INT", IsRequired: true},
		},
	}

	if err := db.CreateTableFromModel(dbConn, userModel); err != nil {
		log.Fatal("Error creating table:", err)
	}
	log.Println("Table created successfully!")
}
