package item

import (
	"fmt"
	"net/http"

	"github.com/sport-city-backend/app/model"
	"github.com/sport-city-backend/app/token"
	"go.mongodb.org/mongo-driver/bson"
)

type changes_platfrom interface {
	Response_item() model.Platfroms_update
	Verification_request(model.Platfroms_update) (bson.M, error)
	Push_response(bson.M)
}

func Push_changes_one(p changes_platfrom, w http.ResponseWriter, r *http.Request) {
	_, err := token.TokenValid(r)
	if err == nil {
		bson, err := p.Verification_request(p.Response_item())
		if err == nil {
			p.Push_response(bson)
			fmt.Fprintf(w, "data changes")
		} else {
			fmt.Fprintf(w, "error data %v", err)
		}
	} else {
		fmt.Fprintf(w, "token not valid")
	}
}

type subscription_platfrom interface {
	Response_item() model.Platfroms_update
	Verification_request(model.Platfroms_update) (string, bson.A, error)
	Push_response(string, bson.A)
}

func Push_subscription_platfrom(p subscription_platfrom, w http.ResponseWriter, r *http.Request) {
	_, err := token.TokenValid(r)
	if err == nil {
		method, bson, err := p.Verification_request(p.Response_item())
		if err == nil {
			p.Push_response(method, bson)
			fmt.Fprintf(w, "data changes")
		} else {
			fmt.Fprintf(w, "error data %v", err)
		}
	} else {
		fmt.Fprintf(w, "token not valid")
	}
}
