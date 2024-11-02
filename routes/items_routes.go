
		package routes

		import ( 
			"net/http" 
			"github.com/esa-kian/shredder/controllers" 
		)

		func RegisteritemsRoutes() { 
			controller := controllers.NewitemsController() 
			http.HandleFunc("/items/create", controller.Create) 
			// TODO: Add other CRUD routes 
		} 
		