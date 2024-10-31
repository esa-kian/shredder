package api

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

// CRUDHandler is a template for handling CRUD requests
type CRUDHandler struct {
	EntityName string
	DB         *sql.DB
}

// Entity represents a general structure for JSON request bodies
type Entity map[string]interface{}

// RegisterRoutes registers CRUD routes for an entity
func (h *CRUDHandler) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/"+h.EntityName, h.Create).Methods("POST")
	router.HandleFunc("/"+h.EntityName+"/{id}", h.Read).Methods("GET")
	router.HandleFunc("/"+h.EntityName+"/{id}", h.Update).Methods("PUT")
	router.HandleFunc("/"+h.EntityName+"/{id}", h.Delete).Methods("DELETE")
}

// validateEntity ensures required fields are present in the entity
func validateEntity(entity Entity) error {
	if _, ok := entity["name"]; !ok {
		return fmt.Errorf("name is required")
	}
	if _, ok := entity["email"]; !ok {
		return fmt.Errorf("email is required")
	}
	return nil
}

func (h *CRUDHandler) Create(w http.ResponseWriter, r *http.Request) {
	var entity Entity
	if err := json.NewDecoder(r.Body).Decode(&entity); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid JSON format")
		return
	}

	if err := validateEntity(entity); err != nil {
		respondWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	// Construct the SQL query dynamically
	query := "INSERT INTO " + h.EntityName + " ("
	values := "VALUES ("
	params := []interface{}{}
	index := 1

	for key, value := range entity {
		if index > 1 {
			query += ", "
			values += ", "
		}
		query += key
		values += "?"
		params = append(params, value)
		index++
	}

	query += ") " + values + ")"

	// Execute the insert statement
	result, err := h.DB.Exec(query, params...)
	if err != nil {
		log.Println("Error creating entity:", err)
		respondWithError(w, http.StatusInternalServerError, "Failed to create entity")
		return
	}

	// Get the last inserted ID
	id, err := result.LastInsertId()
	if err != nil {
		log.Println("Error fetching last insert id:", err)
		respondWithError(w, http.StatusInternalServerError, "Failed to retrieve entity ID")
		return
	}

	// Respond with the new entity's ID
	response := map[string]interface{}{
		"id": id,
	}
	respondWithSuccess(w, response)
}

func (h *CRUDHandler) Read(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	query := "SELECT * FROM " + h.EntityName + " WHERE id = ?"
	rows, err := h.DB.Query(query, id)
	if err != nil {
		http.Error(w, "Entity not found", http.StatusNotFound)
		return
	}
	defer rows.Close()

	if !rows.Next() {
		http.Error(w, "Entity not found", http.StatusNotFound)
		return
	}

	// Get column names from the rows
	columns, err := rows.Columns()
	if err != nil {
		http.Error(w, "Failed to fetch columns", http.StatusInternalServerError)
		return
	}

	// Prepare to scan values for each column
	values := make([]interface{}, len(columns))
	valuePointers := make([]interface{}, len(columns))
	for i := range values {
		valuePointers[i] = &values[i]
	}

	// Scan row into value pointers
	if err := rows.Scan(valuePointers...); err != nil {
		http.Error(w, "Failed to fetch entity", http.StatusInternalServerError)
		return
	}

	// Map column names to their values with proper string conversion
	entity := make(Entity)
	for i, col := range columns {
		val := values[i]
		// If the value is []byte (common for strings in MySQL), convert to string
		if b, ok := val.([]byte); ok {
			entity[col] = string(b)
		} else {
			entity[col] = val
		}
	}

	json.NewEncoder(w).Encode(entity)
}

func (h *CRUDHandler) Update(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	var entity Entity
	if err := json.NewDecoder(r.Body).Decode(&entity); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid JSON format")
		return
	}

	if err := validateEntity(entity); err != nil {
		respondWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	// Dynamically construct the update query
	query := "UPDATE " + h.EntityName + " SET "
	params := []interface{}{}
	index := 1

	for key, value := range entity {
		if index > 1 {
			query += ", "
		}
		query += key + " = ?"
		params = append(params, value)
		index++
	}

	query += " WHERE id = ?"
	params = append(params, id)

	// Execute the update statement
	result, err := h.DB.Exec(query, params...)
	if err != nil {
		log.Println("Error updating entity:", err)
		respondWithError(w, http.StatusInternalServerError, "Failed to update entity")
		return
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		log.Println("Error fetching rows affected:", err)
		respondWithError(w, http.StatusInternalServerError, "Failed to retrieve update status")
		return
	}

	if rowsAffected == 0 {
		respondWithError(w, http.StatusNotFound, "Entity not found")
		return
	}

	respondWithSuccess(w, map[string]string{"status": "updated"})
}

func (h *CRUDHandler) Delete(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	query := "DELETE FROM " + h.EntityName + " WHERE id = ?"
	result, err := h.DB.Exec(query, id)
	if err != nil {
		log.Println("Error deleting entity:", err)
		respondWithError(w, http.StatusInternalServerError, "Failed to delete entity")
		return
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		log.Println("Error fetching rows affected:", err)
		respondWithError(w, http.StatusInternalServerError, "Failed to retrieve delete status")
		return
	}

	if rowsAffected == 0 {
		respondWithError(w, http.StatusNotFound, "Entity not found")
		return
	}

	respondWithSuccess(w, map[string]string{"status": "deleted"})
}

func GenerateControllerFile(entityName string) error {
	controllerTemplate := fmt.Sprintf(`
		package controllers

		import ( "net/http" )

		// %sController handles CRUD operations for the %s model type %sController struct { // Add necessary fields here }

		// New%sController initializes the controller for %s func New%sController() *%sController { return &%sController{} }

		// Example of a Create handler func (c *%sController) Create(w http.ResponseWriter, r *http.Request) { // TODO: Implement Create } `, entityName, entityName, entityName, entityName, entityName, entityName, entityName, entityName, entityName)
	// Write controller file
	fileName := fmt.Sprintf("./controllers/%s_controller.go", entityName)
	return os.WriteFile(fileName, []byte(controllerTemplate), 0644)
}

func GenerateRoutesFile(entityName string) error {
	routesTemplate := fmt.Sprintf(`
		package routes

		import ( "net/http" "%s/controllers" )

		func Register%sRoutes() { controller := controllers.New%sController() http.HandleFunc("/%s/create", controller.Create) // TODO: Add other CRUD routes } `, entityName, entityName, entityName, entityName)

	// Write routes file
	fileName := fmt.Sprintf("./routes/%s_routes.go", entityName)
	return os.WriteFile(fileName, []byte(routesTemplate), 0644)
}
