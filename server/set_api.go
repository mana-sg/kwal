package server

import (
	"net/http"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/mana-sg/kv-log-store/pkg/storage"
)

type SetRequest struct {
	Key string      `json:"key"`
	Value string    `json:"value"`
}

func HandleSet(c *gin.Context){
		var newRequest SetRequest
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
				"message": "Key cannot be empty",
			})
			return
		}

		if newRequest.Value == "" {
			c.JSON(http.StatusBadRequest, gin.H{
				"success": "false",
				"message": "Value cannot be empty",
			})
			return
		}

		err := storage.Kv_store.Set(newRequest.Key, newRequest.Value)

		if err != nil {
			c.Error(fmt.Errorf("Error setting key: %s", err.Error()))
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"success": "true",
		})
	}
