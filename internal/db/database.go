package db

import (
	"context"
	"fmt"
	"github.com/ardaorkun/go-todo-app/config"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"time"
)

var tasksCollection *mongo.Collection

func InitializeDatabase() {
	client, err := connectToMongoDB()
	if err != nil {
		log.Fatal(err)
	}

	checkMongoDBConnection(client)

	setupCollections(client)
}

func connectToMongoDB() (*mongo.Client, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(config.GetConfig().DatabaseURL))
	if err != nil {
		return nil, err
	}

	return client, nil
}

func checkMongoDBConnection(client *mongo.Client) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err := client.Ping(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to MongoDB!")
}

func setupCollections(client *mongo.Client) {
	database := client.Database("golang-todo-backend")
	collection := database.Collection("tasks")
	tasksCollection = collection
}
