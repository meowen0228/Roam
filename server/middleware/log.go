package middleware

import (
	"chat-platform/handlers/log"
	"chat-platform/models"
	"chat-platform/response"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

const swaggerAPIName = "github.com/swaggo/gin-swagger.CustomWrapHandler.func1"

// LoggingMiddleware is a middleware that logs API requests
func LoggingMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()

		apiName := c.HandlerName()

		// 如果是 Swagger API，則不記錄
		if apiName == swaggerAPIName {
			return
		}

		// gin.Context header取得userid 如果沒有就是空
		userID := c.GetHeader("userID")
		ip := c.ClientIP()

		// Create log entry
		logEntry := models.Log{
			APIName:    apiName,
			APIUrl:     c.Request.URL.Path,
			APIMethod:  c.Request.Method,
			APISuccess: c.Writer.Status() == http.StatusOK,
			IP:         ip,
			UserID:     userID,
			CreatedAt:  time.Now(),
		}

		// 將 log id 放進 context
		c.Set("logID", logEntry.ID)

		log.CreateLog(c, logEntry)
	}

}

// ErrorHandlingMiddleware is a middleware that handles errors
func ErrorHandlingMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()
		// 如果有未知錯誤，則回傳錯誤訊息，並記錄錯誤
		if len(c.Errors) > 0 {
			response.ServerFail(c, response.ResJSON{
				Msg: c.Errors[0].Error(),
			})
			logEntry := models.LogError{
				ID:        uint(c.GetInt("logID")),
				ErrorCode: c.Errors[0].Error(),
				Msg:       c.Errors[0].Error(),
			}
			log.CreateErrorLog(c, logEntry)
		}
	}

}
