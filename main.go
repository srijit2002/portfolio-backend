package main

import (
	"log"
	"net/http"
	"portfolio-backend/routers"
)

func main() {
	r:=routers.Router()
	log.Fatal(http.ListenAndServe(":4000",r))
}
