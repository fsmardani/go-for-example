package database

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var MessagesCollection *mongo.Collection

func Connect() {
	clientOptions := options.Client().ApplyURI("mongodb://mongodb:27017")
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Connected to MongoDB!")

	MessagesCollection = client.Database("chatbot").Collection("messages")
}

func SaveMessage(message string) error {
	_, err := MessagesCollection.InsertOne(context.TODO(), map[string]interface{}{
		"message":   message,
		"timestamp": time.Now(),
	})
	return err
}
