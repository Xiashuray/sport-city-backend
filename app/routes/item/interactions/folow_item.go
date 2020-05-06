package interactions

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/sport-city-backend/app/db"
	"github.com/sport-city-backend/app/model"
	"github.com/sport-city-backend/app/routes/item"
	"go.mongodb.org/mongo-driver/bson"
)

type Follow_item struct {
	d model.Default_auth
}

func (i Follow_item) Response_item() model.Platfroms_update {

	platfrom := model.Platfroms_update{}
	_ = json.NewDecoder(i.d.R.Body).Decode(&platfrom)
	return platfrom
}

func (i Follow_item) Verification_request(p model.Platfroms_update) (string, bson.A, error) {
	coll := db.Request_update()
	_, err := coll.Find(context.TODO(), bson.M{"_id": p.ID})
	if err != nil {
		return "null", bson.A{}, err
	}
	return p.Follow.Method,
		bson.A{
			bson.M{"_id": p.ID},
			bson.M{"fovorites": p.Follow.Follow},
		},
		nil

}

func (i Follow_item) Push_response(method string, filter bson.A) {
	coll := db.Request_update()
	if method == "push" {
		coll.UpdateOne(
			context.TODO(),
			filter[0],
			bson.M{"$push": filter[1]},
		)
	} else if method == "remove" {
		coll.UpdateOne(
			context.TODO(),
			filter[0],
			bson.M{
				"$pull": filter[1],
			},
		)
	}

}

func Response_follow_item(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	res := Follow_item{model.Default_auth{R: r}}
	item.Push_subscription_platfrom(res, w, r)
}
