package interactions

import (
	"context"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/sport-city-backend/app/db"
	"github.com/sport-city-backend/app/model"
	"github.com/sport-city-backend/app/routes/item"
	"go.mongodb.org/mongo-driver/bson"
)

type Assessment_Platform struct {
	d model.Default_auth
}

func (i Assessment_Platform) Response_item() model.Platfroms_update {

	platfrom := model.Platfroms_update{}
	_ = json.NewDecoder(i.d.R.Body).Decode(&platfrom)
	return platfrom
}

func (i Assessment_Platform) Verification_request(p model.Platfroms_update) (string, bson.A, error) {
	coll := db.Request_update()
	cur, err := coll.Find(context.TODO(), bson.M{"_id": p.ID})

	indexArray := 0

	for cur.Next(context.TODO()) {
		item := model.Platfroms{}
		cur.Decode(&item)
		for index, el := range item.Rating {

			if el.CreateId == p.Rating[0].CreateId {
				indexArray = index
			}
		}
	}

	if err != nil {
		return "null", bson.A{}, err
	}

	return p.Follow.Method,
		bson.A{
			bson.M{"_id": p.ID},
			bson.M{"$set": bson.M{
				"rating." + strconv.Itoa(indexArray) + ".info.eval": p.Rating[0].Info.Evaluation,
			},
			},
		},
		nil

}

func (i Assessment_Platform) Push_response(method string, filter bson.A) {
	coll := db.Request_update()
	coll.UpdateOne(
		context.TODO(),
		filter[0],
		filter[1],
	)
}
func Response_Assessment_item(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	res := Assessment_Platform{model.Default_auth{R: r}}
	item.Push_subscription_platfrom(res, w, r)
}
