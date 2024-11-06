package middleware

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Response 是一個通用的 API 回應結構
type Response struct {
	StateCode int         `json:"state_code"`     // 狀態碼
	Message   string      `json:"message"`        // 回應訊息
	Data      interface{} `json:"data,omitempty"` // 可選的資料欄位
}

// NewSuccessResponse 創建一個成功的回應
func NewSuccessResponse(c *gin.Context) {
	c.JSON(http.StatusOK, Response{
		StateCode: c.Writer.Status(),
		Message:   "Request successful",
		Data:      c.Keys["data"],
	})
	return
}

// NewErrorResponse 創建一個錯誤的回應
func NewErrorResponse(c *gin.Context, message string) {
	//印出c.Writer.Status()
	fmt.Println(c.Writer.Status())

	c.JSON(http.StatusOK, Response{
		StateCode: c.Writer.Status(),
		Message:   message,
		Data:      c.Keys["data"],
	})
	return
}

// ResponseHandlingMiddleware 是一個處理 API 回應的中間件
func ResponseHandlingMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next() // 處理請求

		// 如果有錯誤回應或狀態碼不為2xx，視為失敗
		isSuccess := c.Writer.Status() >= 200 && c.Writer.Status() < 300

		// 處理請求後檢查是否有錯誤
		if !isSuccess {
			NewErrorResponse(c, "Internal Server Error")
			c.Abort()
		} else {
			NewSuccessResponse(c)
		}
	}
}
