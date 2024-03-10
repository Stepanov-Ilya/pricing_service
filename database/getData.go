package database

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

type Node struct {
}

func GetRawCategories() map[string]interface{} {
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		log.Fatal(err)
	}
	err = client.Connect(context.TODO())
	if err != nil {
		log.Fatal(err)
	}
	defer client.Disconnect(context.TODO())

	collection := client.Database("mydatabase").Collection("rawCategories")

	filter := bson.D{}

	cur, err := collection.Find(context.TODO(), filter)
	if err != nil {
		log.Fatal(err)
	}
	defer cur.Close(context.TODO())

	var bsonDocs []bson.D
	for cur.Next(context.TODO()) {
		var elem bson.D
		err := cur.Decode(&elem)
		if err != nil {
			log.Fatal(err)
		}
		bsonDocs = append(bsonDocs, elem)
	}

	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}

	rawCategories := BsonToArray(bsonDocs)
	fmt.Print(rawCategories)
	return rawCategories
}

func GetRawLocation() map[string]interface{} {
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		log.Fatal(err)
	}
	err = client.Connect(context.TODO())
	if err != nil {
		log.Fatal(err)
	}
	defer client.Disconnect(context.TODO())

	collection := client.Database("mydatabase").Collection("rawLocations")

	filter := bson.D{}

	cur, err := collection.Find(context.TODO(), filter)
	if err != nil {
		log.Fatal(err)
	}
	defer cur.Close(context.TODO())

	var bsonDocs []bson.D
	for cur.Next(context.TODO()) {
		var elem bson.D
		err := cur.Decode(&elem)
		if err != nil {
			log.Fatal(err)
		}
		bsonDocs = append(bsonDocs, elem)
	}

	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}

	rawCategories := BsonToArray(bsonDocs)
	fmt.Print(rawCategories)
	return rawCategories
}

func BsonToArray(bsonData []bson.D) map[string][]string {
	var arrayOfMaps map[string][]string

	for _, doc := range bsonData {
		docMap := make(map[string]interface{})
		for index, pair := range doc {
			if index == 0 {
				continue
			}
			arr := make([]string, 0)
			for _, value := range pair.Value {
				arr = append(arr, value)
			}
			docMap[pair.Key] = pair.Value
		}
		arrayOfMaps = docMap
	}

	return arrayOfMaps
}
