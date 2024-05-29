package api

import (
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.Engine) {
	router.POST("/submit-feature", submitFeature)
	router.POST("/get-features-filter", getFeaturesWithFilter)
	router.POST("/get-features-query", getFeaturesWithQuery)

}
