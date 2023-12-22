package routers

import (
	"github.com/gorilla/mux"
	"portfolio-backend/controllers"
)


func Router() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/projects", controllers.GetAllProjects).Methods("GET")
	router.HandleFunc("/projects/new", controllers.CreateProject).Methods("POST")
	router.HandleFunc("/projects/{id}", controllers.UpdateProject).Methods("PUT")
	router.HandleFunc("/projects/{id}", controllers.DeleteProject).Methods("DELETE")
	router.HandleFunc("/post-form", controllers.CreateForm).Methods("POST")
	return router
}
