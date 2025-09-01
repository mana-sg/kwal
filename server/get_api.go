package server

import (
	"net/http"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/mana-sg/kv-log-store/pkg/storage"
)

type GetRequest struct {
	Key string      `json:"key"`
}

func HandleGet(c *gin.Context) {
	var newRequest GetRequest

	if err := c.ShouldBindJSON(&newRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": "false",
			"message": "Invalid request format",
		})
		return
	}
	
	if newRequest.Key == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": "false",
			"message": "Request key cannot be empty",
		})
		return
	}

	
	val, err := storage.Kv_store.Get(newRequest.Key)
	if err != nil {
		c.Error(fmt.Errorf("Error fetching value: %s", err.Error()))
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"success": "true",
		"value": val,
	})
}
