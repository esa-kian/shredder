package routes

import (
	"net/http"
	"user/controllers"
)

func RegisteruserRoutes() {
	controller := controllers.NewuserController()
	http.HandleFunc("/user/create", controller.Create) // TODO: Add other CRUD routes
}
