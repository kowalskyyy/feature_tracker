package main

import (
	"context"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/kowalskyyy/feature_tracker.git/api"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var client *mongo.Client


func initMongoDB() {
    var err error
    clientOptions := options.Client().ApplyURI("mongodb://localhost:27018")
    client, err = mongo.Connect(context.TODO(), clientOptions)
    if err != nil {
        log.Fatal(err)
    }

    err = client.Ping(context.TODO(), nil)
    if err != nil {
        log.Fatal(err)
    }

    log.Println("Connected to MongoDB!")
}

func main() {
	initMongoDB()
	fmt.Println("Starting server on localhost:8080")
	router := gin.Default()
	api.RegisterRoutes(router, client)
	router.Run("localhost:8080")
}