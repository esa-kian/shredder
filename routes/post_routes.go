package routes

import (
	"net/http"

	"github.com/esa-kian/shredder/controllers"
)

func RegisterpostRoutes() {
	controller := controllers.NewpostController()
	http.HandleFunc("/post/create", controller.Create)
}
