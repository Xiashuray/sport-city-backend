package create

import (
	"fmt"
	"net/http"

	"github.com/sport-city-backend/app/db"
	"github.com/sport-city-backend/app/token"
)

func Response_create_item(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	_, err := token.TokenValid(r)
	if err == nil {
		db.Request_db(r)
		fmt.Fprintf(w, "platfrom create")
	} else {
		fmt.Fprintf(w, "token not valid")
	}
}
