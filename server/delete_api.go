package server
 
import (
	"net/http"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/mana-sg/kv-log-store/pkg/storage"
)

type DeleteRequest struct {
	Key string `json: key`
}

func HandleDelete(c *gin.Context) {
	var newRequest DeleteRequest

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

	err := storage.Kv_store.Remove(newRequest.Key)

	if err != nil {
		c.Error(fmt.Errorf("Error deleting key: %s", err.Error()))
		return 
	}

	c.JSON(http.StatusOK, gin.H{
		"success": "true",
		"message": "Key deleted succesfully",
	})
}
