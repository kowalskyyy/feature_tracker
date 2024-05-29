package internal

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var Client *mongo.Client


func InitMongoDB() {
    var err error
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel() 
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27019")
    log.Println("Connecting to MongoDB...")
	Client, err = mongo.Connect(ctx, clientOptions)

    if err != nil {
        log.Fatal(err)
    }

    err = Client.Ping(ctx, nil)
    if err != nil {
        log.Fatal(err)
    }

    log.Println("Connected to MongoDB!")
}