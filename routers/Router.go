package routers

import (
	"github.com/gorilla/mux"
	"net/http"
	"portfolio-backend/controllers"
)

func CORS(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Access-Control-Allow-Origin", "*")
		w.Header().Add("Access-Control-Allow-Credentials", "true")
		w.Header().Add("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		w.Header().Add("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")

		if r.Method == "OPTIONS" {
			http.Error(w, "No Content", http.StatusNoContent)
			return
		}

		next(w, r)
	}
}
func Router() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/projects", controllers.GetAllProjects).Methods("GET")
	router.HandleFunc("/projects/new", controllers.CreateProject).Methods("POST")
	router.HandleFunc("/projects/{id}", controllers.UpdateProject).Methods("PUT")
	router.HandleFunc("/projects/{id}", controllers.DeleteProject).Methods("DELETE")
	router.HandleFunc("/post-form", CORS(controllers.CreateForm)).Methods("POST")
	return router
}
