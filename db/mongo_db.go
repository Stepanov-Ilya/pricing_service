package db

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

func open_bd() mongo.Client {
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://127.0.0.1:27017"))
	if err != nil {
		log.Fatal(err)
	}

	err = client.Connect(context.TODO())
	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}

	return *client
}

func close_db(client mongo.Client) {
	err := client.Disconnect(context.TODO())

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connection to MongoDB closed.")
}

// Запускать один раз !!
func Start_database() {
	client := open_bd()

	loc_col := client.Database("main_db").Collection("LocationNode")
	cat_col := client.Database("main_db").Collection("CategoryNode")
	GetLocationsTree(*loc_col)  // Загрузка стартовых даннных в бд (запускать один раз!)
	GetCategoriesTree(*cat_col) // Загрузка стартовых даннных в бд (запускать один раз!)

	close_db(client)
}

func read_db(collection mongo.Collection) {
	options := options.Find()
	//options.SetLimit(2)
	filter := bson.D{}

	var results []*LocationNode

	cur, err := collection.Find(context.TODO(), filter, options)
	if err != nil {
		log.Fatal(err)
	}

	for cur.Next(context.TODO()) {
		var elem LocationNode
		err := cur.Decode(&elem)
		if err != nil {
			log.Fatal(err)
		}

		results = append(results, &elem)
		if len(elem.Children) > 0 {
			fmt.Println(elem.Name)
			fmt.Print("-> ")
			for _, loc := range elem.Children {
				fmt.Print(loc.Name, " ")
			}
			fmt.Println("\n")
		}
	}

	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}

	cur.Close(context.TODO())
}

func find(filter bson.D, collection mongo.Collection) {

	cursor, _ := collection.Find(context.TODO(), filter)
	var results []LocationNode

	var err error
	if err = cursor.All(context.TODO(), &results); err != nil {
		panic(err)
	}

	for _, result := range results {
		res, _ := bson.MarshalExtJSON(result, false, false)
		fmt.Println(string(res))
	}
}
