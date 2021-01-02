package storage

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

type Repository struct {
	Client *mongo.Client
}

func GetMyOpenData(repo Repository) (MyOpenData, error) {
	lab2Collection := repo.Client.Database("Otvoreno").Collection("lab2")

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*15)
	defer cancel()

	cursor, err := lab2Collection.Find(ctx, bson.M{})
	if err != nil {
		return MyOpenData{}, err
	}

	myOpenDataStructArray := make([]MyOpenDataStruct, 0)
	for cursor.Next(ctx) {
		var data MyOpenDataStruct

		// decoding data
		err := cursor.Decode(&data)
		if err != nil {
			return MyOpenData{}, err
		}

		myOpenDataStructArray = append(myOpenDataStructArray, data)
	}

	return MyOpenData{OpenData: myOpenDataStructArray}, nil

}
