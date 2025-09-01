package server

import (
	"os"
	"net/http"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/mana-sg/kv-log-store/pkg/wal"
)


func HandleCompaction(c *gin.Context) {
	savings, err := wal.Compact()
	if err != nil {
		c.Error(fmt.Errorf("Error compacting file: %s", err.Error()))
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": "true",
		"savings": fmt.Sprintf("Succesfully compacted log file, Savings: %.2f%%", savings * 100),
	})
}

func HandleGetSize(c *gin.Context) {
	home, err := os.UserHomeDir() 			
	filePath := home + "/kls/log.bin"
	fileInfo, err := os.Stat(filePath)
	if err != nil {
		c.Error(fmt.Errorf("Error getting file size: %s", err.Error()))
		return
	}

	fileSize := fileInfo.Size()

	c.JSON(http.StatusOK, gin.H {
		"success": "true",
		"message": fmt.Sprintf("Log file size: %d Bytes", fileSize),
	})
}
