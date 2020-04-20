package routes

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/sport-city-backend/app/auth"
	"github.com/sport-city-backend/app/token"
)

func GetTokenServer(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	w.WriteHeader(http.StatusOK)
	user := auth.Auth_google(vars["id"])

	fmt.Println(user)
	if user != "error user" {
		ss, err := token.CreateToken(user)
		if err != nil {
			println(err)
		}
		res := `{"token":"` + ss + `"}`
		fmt.Fprintf(w, "%v", res)
	}
}
