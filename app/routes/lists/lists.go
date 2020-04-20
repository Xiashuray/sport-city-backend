package lists

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/sport-city-backend/app/model"
	"github.com/sport-city-backend/app/token"
	"go.mongodb.org/mongo-driver/bson"
)

type get_platfroms interface {
	Get_data_platfroms(model.Page, bson.M) []model.Platfroms
	Set_page_and_tagging() (model.Page, bson.M)
}

func Put_platfrom(g get_platfroms, w http.ResponseWriter, r *http.Request) {
	_, err := token.TokenValid(r)
	if err == nil {
		page, filtre := g.Set_page_and_tagging()
		data := g.Get_data_platfroms(page, filtre)
		json.NewEncoder(w).Encode(data)
	} else {
		fmt.Fprintf(w, "token not valid")
	}
}
