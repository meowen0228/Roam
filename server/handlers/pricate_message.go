package handlers

import (
	"chat-platform/database"
	"chat-platform/models"
	"chat-platform/response"
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetPrivateMessages 獲取所有私人消息
// @Summary Get all private messages
// @Description Get all private messages
// @Tags messages
// @Accept  json
// @Produce  json
// @Success 200 {array} models.PrivateMessage
// @Failure 500 {object} gin.H
// @Router /messages [get]
func GetPrivateMessages(c *gin.Context) {
	var messages []models.PrivateMessage
	result := database.DB.Find(&messages)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}
	c.JSON(http.StatusOK, messages)
}

// PostPrivateMessages 獲取所有私人消息
// @Summary Get all private messages
// @Description Get all private messages
// @Tags messages
// @Accept  json
// @Produce  json
// @Success 200 {array} models.PrivateMessage
// @Failure 500 {object} gin.H
// @Router /messages [post]
func PostPrivateMessages(c *gin.Context) {
	var messages []models.PrivateMessage
	result := database.DB.Find(&messages)
	if result.Error != nil {
		response.Error(c, response.ResJSON{
			Msg: "Internal Server Error",
		})
		return
	}
	c.JSON(http.StatusOK, messages)
}
