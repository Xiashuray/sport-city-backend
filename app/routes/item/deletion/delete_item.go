package deletion

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/sport-city-backend/app/db"
	"github.com/sport-city-backend/app/model"
	"github.com/sport-city-backend/app/routes/item"
	"go.mongodb.org/mongo-driver/bson"
)

type Delete_item struct {
	d model.Default_auth
}

func (i Delete_item) Response_item() model.Platfroms_update {

	platfrom := model.Platfroms_update{}
	_ = json.NewDecoder(i.d.R.Body).Decode(&platfrom)
	return platfrom
}

func (i Delete_item) Verification_request(p model.Platfroms_update) (bson.M, error) {
	coll := db.Request_update()
	_, err := coll.Find(context.TODO(), bson.M{"_id": p.ID, "createId": p.CreateId})
	if err != nil {
		return bson.M{}, err
	}
	return bson.M{"_id": p.ID}, nil
}

func (i Delete_item) Push_response(filter bson.M) {
	coll := db.Request_update()
	coll.DeleteOne(context.TODO(), filter)
}

func Response_delete_item(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	res := Delete_item{model.Default_auth{R: r}}
	item.Push_changes_one(res, w, r)
}
