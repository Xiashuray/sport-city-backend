package routes

import (
	"fmt"
	"net/http"

	"github.com/sport-city-backend/app/auth"
)

func Test(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	res := auth.Auth_google("6e98oC70UGVpKB7RIEK3WnRCmYP2")
	fmt.Fprintf(w, "%v", res)
}
