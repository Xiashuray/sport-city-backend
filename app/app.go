package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/sport-city-backend/app/routes"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/", DoHealthCheck)
	router.HandleFunc("/test", routes.Test)
	log.Fatal(http.ListenAndServe(":8080", router))
}

func DoHealthCheck(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, i'm a golang microservice test")
}
