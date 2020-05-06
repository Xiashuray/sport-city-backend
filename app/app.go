package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/sport-city-backend/app/routes"

	"github.com/sport-city-backend/app/routes/item/create"
	"github.com/sport-city-backend/app/routes/item/deletion"
	"github.com/sport-city-backend/app/routes/item/interactions"
	lists_private "github.com/sport-city-backend/app/routes/lists/private"
	lists_public "github.com/sport-city-backend/app/routes/lists/public"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/", DoHealthCheck)
	router.HandleFunc("/token/new/{id}", routes.GetTokenServer).Methods("POST")
	router.HandleFunc("/add", create.Response_create_item).Methods("POST")
	router.HandleFunc("/del", deletion.Response_delete_item).Methods("POST")
	router.HandleFunc("/ptafrom/sub", interactions.Response_follow_item).Methods("POST")
	router.HandleFunc("/ptafrom/view", interactions.Response_view_item).Methods("POST")
	router.HandleFunc("/ptafrom/rating", interactions.Response_Assessment_item).Methods("POST")
	router.HandleFunc("/platfrom/news", lists_public.Response_News).Methods("GET")
	router.HandleFunc("/platfrom/best", lists_public.Response_best).Methods("GET")
	router.HandleFunc("/platfrom/near", lists_public.Response_Near).Methods("GET")
	router.HandleFunc("/platfrom/follow", lists_private.Response_Follow).Methods("GET")
	router.HandleFunc("/platfrom/create_user", lists_private.Response_Create_user).Methods("GET")
	log.Fatal(http.ListenAndServe(":8080", router))
}

func DoHealthCheck(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, i'm a golang microservice test")
}
