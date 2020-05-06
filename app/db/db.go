package db

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/sport-city-backend/app/model"
	"go.mongodb.org/mongo-driver/bson"
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

func Response_db(filtre bson.M, options *options.FindOptions) []model.Platfroms {
	clietn := auth_db()
	coll := clietn.Database("sportcity").Collection("platfroms")

	cur, _ := coll.Find(context.TODO(), filtre, options)

	data := []model.Platfroms{}

	for cur.Next(context.TODO()) {
		var names model.Platfroms
		cur.Decode(&names)
		data = append(data, names)
	}

	return data
}

func Request_db_avg(page model.Page, filtre bson.M, options *options.AggregateOptions) []model.Platfroms {

	client := auth_db()

	coll := client.Database("sportcity").Collection("platfroms")
	testavg := bson.A{}
	items := bson.A{}
	testavg = bson.A{
		bson.M{
			"$match": bson.M{
				"rating.info.eval": bson.M{
					"$gte": 0,
				},
			},
		},
		bson.M{
			"$sort": bson.M{
				"rating.info.eval": -1,
			},
		},

		bson.M{
			"$skip": page.Start,
		},
		bson.M{
			"$limit": page.End,
		},
	}
	if len(page.Tag) != 0 {
		for _, item := range page.Tag {
			items = append(items, item)
		}
		testavg = bson.A{
			bson.M{
				"$match": bson.M{
					"rating.info.eval": bson.M{
						"$gte": 0,
					},
					"tag": bson.M{"$all": items},
				},
			},
			bson.M{
				"$sort": bson.M{
					"rating.info.eval": -1,
				},
			},

			bson.M{
				"$skip": page.Start,
			},
			bson.M{
				"$limit": page.End,
			},
		}
	}

	cur, _ := coll.Aggregate(context.TODO(), testavg, options)

	data := []model.Platfroms{}

	for cur.Next(context.TODO()) {
		var names model.Platfroms
		cur.Decode(&names)
		data = append(data, names)
	}

	return data
}

func Request_db(r *http.Request) {
	clietn := auth_db()
	coll := clietn.Database("sportcity").Collection("platfroms")

	var platfrom model.Platfroms

	_ = json.NewDecoder(r.Body).Decode(&platfrom)
	coll.InsertOne(context.TODO(), platfrom)
}

func Request_update() *mongo.Collection {
	clietn := auth_db()
	return clietn.Database("sportcity").Collection("platfroms")
}
