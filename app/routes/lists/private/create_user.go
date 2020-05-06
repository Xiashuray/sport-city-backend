package private

import (
	"encoding/json"
	"net/http"

	"github.com/sport-city-backend/app/db"
	"github.com/sport-city-backend/app/model"
	"github.com/sport-city-backend/app/routes/lists"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Create_user_platfroms struct {
	d model.Default_auth
}

func (g Create_user_platfroms) Set_page_and_tagging() (model.Page, bson.M) {
	var page model.Page
	_ = json.NewDecoder(g.d.R.Body).Decode(&page)

	filtre := bson.M{}
	if len(page.CreateId) != 0 {

		filtre = bson.M{"createId": page.CreateId}

	}
	return page, filtre
}

func (g Create_user_platfroms) Get_data_platfroms(page model.Page, filtre bson.M) []model.Platfroms {

	options := options.Find()
	options.SetSort(bson.D{{"_id", 1}})
	options.SetLimit(int64(page.End))
	options.SetSkip(int64(page.Start))

	return db.Response_db(filtre, options)
}

func Response_Create_user(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	res := Create_user_platfroms{model.Default_auth{R: r}}

	lists.Put_platfrom(res, w, r)
}
