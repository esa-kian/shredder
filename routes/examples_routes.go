
		package routes

		import ( 
			"github.com/gorilla/mux"
			
			"github.com/esa-kian/shredder/controllers" 
		)

		func RegisterexamplesRoutes() { 
			controller := controllers.NewexamplesController() 
			r := mux.NewRouter()

			r.HandleFunc("/examples", controller.Create).Methods("POST")
			r.HandleFunc("/examples/{id}", controller.Update).Methods("PUT")
			r.HandleFunc("/examples/{id}", controller.Read).Methods("GET")
			r.HandleFunc("/examples/{id}", controller.Delete).Methods("DELETE")

		} 
		