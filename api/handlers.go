package api

import (
	context_ "context"
	"net/http"

	"time"

	"github.com/gin-gonic/gin"
	"github.com/kowalskyyy/feature_tracker.git/internal"
	"go.mongodb.org/mongo-driver/mongo"
)


func getData(client *mongo.Client) gin.HandlerFunc {
	return func(context *gin.Context) {

	var data internal.Feature

	if err := context.BindJSON(&data); err != nil {
		context.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Your order data is incorrect"})
		return
	}

	collection := client.Database("FeatureTracker").Collection("Features")

        // Create a context with a 10-second timeout
        ctx, cancel := context_.WithTimeout(context_.Background(), 10*time.Second)
        defer cancel() // Ensure the context is canceled to free resources

        // Insert the item into the collection using the context
        _, err := collection.InsertOne(ctx, data)
        if err != nil {
			context.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Couldnt save data to mongodb", "error": err})
            return
        }



	context.IndentedJSON(http.StatusOK, gin.H{"message": "correct data received", "data": data})
}
}