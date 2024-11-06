package routes

import (
	"chat-platform/handlers/auth"

	"github.com/gin-gonic/gin"
)

// Auth 註冊Auth路由
func Auth(r *gin.Engine) {
	r.POST("/api/auth/google", auth.GoogleLoginHandler)
}
