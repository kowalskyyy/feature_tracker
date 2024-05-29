package main

import (

	"fmt"

	"github.com/kowalskyyy/feature_tracker.git/internal"

	"github.com/gin-gonic/gin"
	"github.com/kowalskyyy/feature_tracker.git/api"

)



func main() {
	internal.InitMongoDB()
	fmt.Println("Starting server on localhost:8080")
	router := gin.Default()
	api.RegisterRoutes(router)
	router.Run("localhost:8080")
}