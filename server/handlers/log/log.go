package log

import (
	"chat-platform/database"
	"chat-platform/models"
	"log"

	"github.com/gin-gonic/gin"
)

// Request 請求結構體
type Request struct {
	Credential string `json:"credential" binding:"required" msg:"Credential is required"`
}

// CreateLog 創建日誌
func CreateLog(c *gin.Context, logEntry models.Log) {
	// Save log entry to database
	if err := database.DB.Create(&logEntry).Error; err != nil {
		log.Println("Failed to save log entry:", err)
	}
}

// CreateErrorLog 創建錯誤日誌
func CreateErrorLog(c *gin.Context, logEntry models.LogError) {
	// Save log entry to database
	if err := database.DB.Create(&logEntry).Error; err != nil {
		log.Println("Failed to save log entry:", err)
	}
}
