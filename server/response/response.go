package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// ResJSON 返回 json 結構體
type ResJSON struct {
	Status int         `json:"status"`
	Msg    string      `json:"msg"`
	Data   interface{} `json:"data"`
}

// HTTPResponse 返回 JSON
func HTTPResponse(ctx *gin.Context, status int, resp ResJSON) {
	ctx.AbortWithStatusJSON(status, ResJSON{
		Status: status,
		Msg:    resp.Msg,
		Data:   resp.Data,
	})
}

// 構建狀態碼，如果沒有設置狀態碼，則使用默認狀態碼
func buildStatus(resp ResJSON, defaultStatus int) int {
	if resp.Status == 0 {
		return defaultStatus
	}
	return resp.Status
}

// Success 成功返回
func Success(ctx *gin.Context, resp ResJSON) {
	HTTPResponse(ctx, buildStatus(resp, http.StatusOK), resp)
}

// Error 錯誤返回
func Error(ctx *gin.Context, resp ResJSON) {
	HTTPResponse(ctx, buildStatus(resp, http.StatusBadRequest), resp)
}

// ServerFail 服务器错误返回
func ServerFail(ctx *gin.Context, resp ResJSON) {
	HTTPResponse(ctx, buildStatus(resp, http.StatusInternalServerError), resp)
}
