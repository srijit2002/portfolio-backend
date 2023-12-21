package routers

import (
	"portfolio-backend/controllers"
	"github.com/gorilla/mux"
)

func Router() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/projects", controllers.GetAllProjects).Methods("GET")
	router.HandleFunc("/projects/new", controllers.CreateProject).Methods("POST")
	router.HandleFunc("/projects/{id}", controllers.DeleteProject).Methods("DELETE")
	return router
}
