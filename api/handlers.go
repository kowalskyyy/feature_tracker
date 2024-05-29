package api

import (
	"net/http"

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



	context.IndentedJSON(http.StatusOK, gin.H{"message": "correct data received", "data": data})
}
}