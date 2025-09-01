
package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mana-sg/kv-log-store/pkg/storage"
	"github.com/mana-sg/kv-log-store/server"
)


func ErrorHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()

		if len(c.Errors) > 0{
			err := c.Errors.Last().Err

			c.JSON(http.StatusInternalServerError, gin.H{
				"success": "false",
				"message": err.Error(),
			})
		}
	}
}

func setUpRouter() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	r := gin.New()
	r.Use(ErrorHandler())

	r.POST("/set", server.HandleSet)
	r.GET("/get", server.HandleGet) 
	r.POST("/delete", server.HandleDelete)
	{
		compactionRoute := r.Group("size")
		compactionRoute.POST("/compact", server.HandleCompaction)
		compactionRoute.GET("/get", server.HandleGetSize)
	}

	return r
}

func main(){
	storage.Kv_store.BuildStore()
	r := setUpRouter()
	
	r.Run(":8080")
}
