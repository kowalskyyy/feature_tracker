package api

import (
	context_ "context"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/kowalskyyy/feature_tracker.git/internal"
	"go.mongodb.org/mongo-driver/bson"
)

func submitFeature(context *gin.Context) {

	var data internal.Feature

	if err := context.BindJSON(&data); err != nil {
		context.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Your order data is incorrect"})
		return
	}

	collection := internal.Client.Database("FeatureTracker").Collection("Features")

	ctx, cancel := context_.WithTimeout(context_.Background(), 10*time.Second)
	defer cancel()

	_, err := collection.InsertOne(ctx, data)
	if err != nil {
		context.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Couldnt save data to mongodb", "error": err})
		return
	}

	context.IndentedJSON(http.StatusOK, gin.H{"message": "correct data received", "data": data})
}

func getFeaturesWithFilter(context *gin.Context) {
	var data []internal.Feature

	var filter []internal.Filter

	var filterBson bson.D

	if err := context.ShouldBindJSON(&filter); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	for _, f := range filter {
		filterBson = append(filterBson, bson.E{Key: f.Key, Value: f.Value})
	}

	collection := internal.Client.Database("FeatureTracker").Collection("Features")

	ctx, cancel := context_.WithTimeout(context_.Background(), 10*time.Second)
	defer cancel()

	cursor, err := collection.Find(ctx, filterBson)

	if err != nil {
		context.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Could not find any objects matching the query", "error": err})
		return
	}

	for cursor.Next(context_.TODO()) {
		var result internal.Feature
		if err := cursor.Decode(&result); err != nil {
			log.Fatal(err)
		}
		data = append(data, result)
	}

	if err := cursor.Err(); err != nil {
		log.Fatal(err)
	}

	context.IndentedJSON(http.StatusOK, gin.H{"message": "Data matching the filter:", "data": data})

}

func getFeaturesWithQuery(context *gin.Context) {
	var data []internal.Feature

	var filterBson bson.M

	if err := context.ShouldBindJSON(&filterBson); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	collection := internal.Client.Database("FeatureTracker").Collection("Features")

	ctx, cancel := context_.WithTimeout(context_.Background(), 10*time.Second)
	defer cancel()

	cursor, err := collection.Find(ctx, filterBson)

	if err != nil {
		context.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Could not find any objects matching the query", "error": err})
		return
	}

	for cursor.Next(context_.TODO()) {
		var result internal.Feature
		if err := cursor.Decode(&result); err != nil {
			log.Fatal(err)
		}
		data = append(data, result)
	}

	if err := cursor.Err(); err != nil {
		log.Fatal(err)
	}

	context.IndentedJSON(http.StatusOK, gin.H{"message": "Data matching the filter:", "data": data})

}
