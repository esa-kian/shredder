package api

import (
	"database/sql"
	"net/http"

	"github.com/gorilla/mux"
)

// CRUDHandler is a template for handling CRUD requests
type CRUDHandler struct {
	EntityName string
	DB         *sql.DB
}

// RegisterRoutes registers CRUD routes for an entity
func (h *CRUDHandler) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/"+h.EntityName, h.Create).Methods("POST")
	router.HandleFunc("/"+h.EntityName+"/{id}", h.Read).Methods("GET")
	router.HandleFunc("/"+h.EntityName+"/{id}", h.Update).Methods("PUT")
	router.HandleFunc("/"+h.EntityName+"/{id}", h.Delete).Methods("DELETE")
}

func (h *CRUDHandler) Create(w http.ResponseWriter, r *http.Request) {
	// TODO: Implement Create logic
	w.Write([]byte("Create endpoint"))
}

func (h *CRUDHandler) Read(w http.ResponseWriter, r *http.Request) {
	// TODO: Implement Read logic
	w.Write([]byte("Read endpoint"))
}

func (h *CRUDHandler) Update(w http.ResponseWriter, r *http.Request) {
	// TODO: Implement Update logic
	w.Write([]byte("Update endpoint"))
}

func (h *CRUDHandler) Delete(w http.ResponseWriter, r *http.Request) {
	// TODO: Implement Delete logic
	w.Write([]byte("Delete endpoint"))
}
