package handlers

import (
	"chat-platform/database"
	"chat-platform/models"
	"chat-platform/response"

	"github.com/gin-gonic/gin"
)

// GetUsers 取得所有用戶
// @Summary Get all users
// @Description Get all users
// @Tags users
// @Accept  json
// @Produce  json
// @Success 200 {array} models.User
// @Failure 500 {object} gin.H
// @Router /users [get]
func GetUsers(c *gin.Context) {
	var users []models.User
	result := database.DB.Find(&users)
	if result.Error != nil {
		response.Error(c, response.ResJSON{
			Msg: result.Error.Error(),
		})
		return
	}
	response.Success(c, response.ResJSON{
		Data: users,
	})
}

// CreateUsers 創建用戶
func CreateUsers(user models.User) (models.User, error) {
	result := database.DB.Create(&user)
	return user, result.Error
}

// FindUserByEmail 透過 email 查找用戶
func FindUserByEmail(email string) (models.User, error) {
	var user models.User
	result := database.DB.Where("email = ?", email).First(&user)
	return user, result.Error
}
