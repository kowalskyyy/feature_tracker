package api

import (
	context_ "context"
	"net/http"

	"time"

	"github.com/gin-gonic/gin"
	"github.com/kowalskyyy/feature_tracker.git/internal"
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
