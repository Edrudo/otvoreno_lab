package storage

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
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

func GetMyOpenDataWithLimit(repo Repository, limit int64) (MyOpenData, error) {
	lab2Collection := repo.Client.Database("Otvoreno").Collection("lab2")

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*15)
	defer cancel()

	cursor, err := lab2Collection.Find(ctx, bson.M{}, options.Find().SetLimit(limit))
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

func GetMyOpenDataFromSomeYear(repo Repository, year int64) (MyOpenData, error) {
	lab2Collection := repo.Client.Database("Otvoreno").Collection("lab2")

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*15)
	defer cancel()

	cursor, err := lab2Collection.Find(ctx, bson.M{"year": year})
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

func InsertCar(repo Repository, data MyOpenDataStructDIO) error {
	lab2Collection := repo.Client.Database("Otvoreno").Collection("lab2")

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*15)
	defer cancel()

	_, err := lab2Collection.InsertOne(ctx, data)

	return err
}

func DeleteCar(repo Repository, id primitive.ObjectID) error {
	lab2Collection := repo.Client.Database("Otvoreno").Collection("lab2")

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*15)
	defer cancel()

	_, err := lab2Collection.DeleteOne(ctx, bson.M{"_id": id})
	if err != nil {
		log.Fatal(err)
	}
	return err
}

func UpdateCar(repo Repository, id primitive.ObjectID, newData MyOpenDataStructDIO) error {
	lab2Collection := repo.Client.Database("Otvoreno").Collection("lab2")

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*15)
	defer cancel()

	_, err := lab2Collection.UpdateOne(ctx, bson.M{"_id": id},
		bson.M{
			"$set": bson.M{
				"acceleration":   newData.Acceleration,
				"awards":         newData.Awards,
				"bootSpace":      newData.BootSpace,
				"engine":         newData.Engine,
				"model":          newData.Model,
				"powerOutput":    newData.PowerOutput,
				"topSpeed":       newData.TopSpeed,
				"torque":         newData.Torque,
				"vehicleBrand":   newData.VehicleBrand,
				"year":           newData.Year,
				"weight":         newData.Weight,
				"wikipediaSufix": newData.WikipediaSufix,
			},
		})
	return err
}
