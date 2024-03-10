package database

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

func Create() {
	// Устанавливаем соединение с MongoDB
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	// Проверяем, что соединение установлено успешно
	err = client.Ping(context.Background(), nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected to MongoDB!")

	// Выбираем базу данных
	database := client.Database("mydatabase")

	// Создаем коллекцию в выбранной базе данных
	collection := database.Collection("mycollection")

	// Здесь вы можете выполнить операции с вашей коллекцией, например, вставку документов
	// Ниже приведен пример вставки одного документа в коллекцию
	_, err = collection.InsertOne(context.Background(), map[string]interface{}{
		"name": "John Doe",
		"age":  30,
	})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Document inserted successfully!")
}
