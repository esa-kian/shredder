
		package routes

		import ( 
			"net/http" 
			"github.com/esa-kian/shredder/controllers" 
		)

		func RegisterusersRoutes() { 
			controller := controllers.NewusersController() 
			http.HandleFunc("/users/create", controller.Create) 
			// TODO: Add other CRUD routes 
		} 
		