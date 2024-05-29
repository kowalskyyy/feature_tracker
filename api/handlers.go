package api

import (
	"net/http"
	"sync"

	"github.com/gin-gonic/gin"
)

var mutex sync.Mutex

func getData(context *gin.Context) {
	mutex.Lock()
	defer mutex.Unlock()

	context.IndentedJSON(http.StatusOK, gin.H{"message": "elo"})
}