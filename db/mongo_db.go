package db

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

func Open_bd() mongo.Client {
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

func Close_db(client mongo.Client) {
	err := client.Disconnect(context.TODO())

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connection to MongoDB closed.")
}

// Заполняет БД (c нуля)
func Start_database(client mongo.Client) {

	loc_col := client.Database("main_db").Collection("LocationNode")
	cat_col := client.Database("main_db").Collection("CategoryNode")

	Clean_all(*loc_col)
	Clean_all(*cat_col)

	GetLocationsTree(*loc_col)  // Загрузка стартовых даннных в бд (запускать один раз!)
	GetCategoriesTree(*cat_col) // Загрузка стартовых даннных в бд (запускать один раз!)

}

func read_db(collection mongo.Collection) {
	options := options.Find()
	filter := bson.D{}

	var results []*Node

	cur, err := collection.Find(context.TODO(), filter, options)
	if err != nil {
		log.Fatal(err)
	}

	for cur.Next(context.TODO()) {
		var elem Node
		err := cur.Decode(&elem)
		if err == nil {
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

func Find_node_in_mongo(id int64, collection mongo.Collection) *Node {
	var locationNode *Node
	filter := bson.D{{"id", id}}
	err := collection.FindOne(context.TODO(), filter).Decode(&locationNode)

	if err != nil {
		log.Fatal(err)
	}
	return locationNode
}

func Clean_all(collection mongo.Collection) {
	filter := bson.D{}
	deleteResult, err := collection.DeleteMany(context.TODO(), filter)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("All have been cleaned", deleteResult.DeletedCount)
}
