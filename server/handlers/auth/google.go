package auth

import (
	"chat-platform/handlers"
	"chat-platform/models"
	"chat-platform/response"
	"os"

	"context"

	"github.com/gin-gonic/gin"
	"google.golang.org/api/idtoken"
)

// Request 請求結構體
type Request struct {
	Credential string `json:"credential" binding:"required" msg:"Credential is required"`
}

// Res 回應結構體
type Res struct {
	Email string `json:"email"`
	Name  string `json:"name"`
	Token string `json:"token"`
}

// GoogleLoginHandler 處理 Google 登入請求
// @Summary Google Login
// @Description Verify Google ID Token and login
// @Tags auth
// @Accept  json
// @Produce  json
// @Param request body Request true "Google ID Token"
// @Success 200 {object} response.ResJSON
// @Failure 400 {object} response.ResJSON
// @Router /auth/google [post]
func GoogleLoginHandler(c *gin.Context) {
	var tokenReq Request
	if err := c.ShouldBindJSON(&tokenReq); err != nil {
		response.Error(c, response.ResJSON{
			Msg: "Invalid request",
		})
		return
	}

	// 從環境變數中讀取 Google Client ID
	clientID := os.Getenv("GOOGLE_CLIENT_ID")
	ctx := context.Background()

	// 驗證 ID Token
	payload, err := idtoken.Validate(ctx, tokenReq.Credential, clientID)
	if err != nil {
		response.Error(c, response.ResJSON{
			Msg: "Invalid token",
		})
		return
	}

	email := payload.Claims["email"].(string)

	_, err = handlers.FindUserByEmail(email)
	if err != nil && err.Error() == "record not found" {
		// 如果 email 不存在，新增使用者
		_, err = handlers.CreateUsers(models.User{
			Username: payload.Claims["name"].(string),
			Email:    payload.Claims["email"].(string),
		})
		if err != nil {
			response.ServerFail(c, response.ResJSON{
				Msg: "Failed to create user",
			})
			return
		}
	}

	// 如果驗證成功，回傳 email, name, token
	response.Success(c, response.ResJSON{
		Msg: "Login success",
		Data: Res{
			Email: payload.Claims["email"].(string),
			Name:  payload.Claims["name"].(string),
			Token: tokenReq.Credential,
		},
	})
}
