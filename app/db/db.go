package db

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/sport-city-backend/app/model"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const URL = "mongodb://127.0.0.1:27017"

func auth_db() *mongo.Client {
	clientOpts := options.Client().ApplyURI(URL)
	client, _ := mongo.Connect(context.TODO(), clientOpts)
	// fmt.Println(err)
	return client
}

func Response_db() *mongo.Collection {
	clietn := auth_db()
	coll := clietn.Database("sportcity").Collection("platfroms")

	return coll
}

func Request_db(r *http.Request) {
	clietn := auth_db()
	coll := clietn.Database("sportcity").Collection("platfroms")

	var platfrom model.Platfroms

	_ = json.NewDecoder(r.Body).Decode(&platfrom)
	coll.InsertOne(context.TODO(), platfrom)
}
