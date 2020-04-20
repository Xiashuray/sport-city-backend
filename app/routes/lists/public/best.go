package public

import (
	"encoding/json"
	"net/http"

	"github.com/sport-city-backend/app/db"
	"github.com/sport-city-backend/app/model"
	"github.com/sport-city-backend/app/routes/lists"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Best_platfroms struct {
	d model.Default_auth
}

func (g Best_platfroms) Set_page_and_tagging() (model.Page, bson.M) {
	var page model.Page
	_ = json.NewDecoder(g.d.R.Body).Decode(&page)

	filtre := bson.M{
		// "$match": bson.M{
		// 	"rating.eval": bson.M{
		// 		"$ne": 0,
		// 	},
		// },
	}
	// if len(page.Tag) != 0 {
	// 	items := bson.A{}
	// 	for _, item := range page.Tag {
	// 		items = append(items, item)
	// 	}

	// 	filtre = bson.M{"tag": items}
	// }

	return page, filtre
}

func (g Best_platfroms) Get_data_platfroms(page model.Page, filtre bson.M) []model.Platfroms {

	options := options.Aggregate()
	// options.SetSort(bson.D{{"rating", bson.D{{"eval", 1}}}})
	// options.SetLimit(int64(page.End))
	// options.SetSkip(int64(page.Start))

	return db.Request_db_avg(page, filtre, options)
}

func Response_best(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	res := Best_platfroms{model.Default_auth{R: r}}

	lists.Put_platfrom(res, w, r)
}
