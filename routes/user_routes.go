package routes

import (
	"net/http"

	"github.com/esa-kian/shredder/controllers"
)

func RegisteruserRoutes() {
	controller := controllers.NewuserController()
	http.HandleFunc("/user/create", controller.Create)
	// TODO: Add other CRUD routes
}
