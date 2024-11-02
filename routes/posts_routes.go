
		package routes

		import ( 
			"net/http" 
			"github.com/esa-kian/shredder/controllers" 
		)

		func RegisterpostsRoutes() { 
			controller := controllers.NewpostsController() 
			http.HandleFunc("/posts/create", controller.Create) 
			// TODO: Add other CRUD routes 
		} 
		