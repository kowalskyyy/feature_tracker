package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.Engine) {
	router.LoadHTMLGlob("templates/*")
	router.POST("/submit-feature", submitFeature)
	router.POST("/get-features-filter", getFeaturesWithFilter)
	router.POST("/get-features-query", getFeaturesWithQuery)
	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.tmpl", gin.H{
			"title": "Main website",
		})})
}
