package routes

import (
	"github.com/gorilla/mux"

	"github.com/esa-kian/shredder/controllers"
)

func RegisteritemsRoutes() {
	controller := controllers.NewitemsController()
	r := mux.NewRouter()

	r.HandleFunc("/items", controller.Create).Methods("POST")
	r.HandleFunc("/items/{id}", controller.Update).Methods("PUT")
	r.HandleFunc("/items/{id}", controller.Read).Methods("GET")
	r.HandleFunc("/items/{id}", controller.Delete).Methods("DELETE")

}
