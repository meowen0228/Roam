package main

import (
	"chat-platform/database"
	"chat-platform/handlers"
	_ "chat-platform/handlers"
	"chat-platform/middleware"
	"chat-platform/routes"
	"chat-platform/ws"

	"fmt"

	_ "chat-platform/docs"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title Chat Platform API
// @version 1.0
// @description This is a sample server for a chat platform.
// @termsOfService http://swagger.io/terms/
// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io
// @host localhost:8080
// @BasePath /api
func main() {
	port := 8080

	// 連接資料
	database.ConnectDB()
	database.MigratedDB()

	r := gin.Default()

	// 設定 proxies
	r.SetTrustedProxies([]string{"127.0.0.1"})

	// 設置Gin路由
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	r.Use(middleware.Cors())

	// 設置白名單
	// whitelist := []string{"127.0.0.1", "localhost", "::1"}

	// middleware
	// r.Use(middleware.LoggingMiddleware(), middleware.IPWhiteList(whitelist), middleware.ErrorHandlingMiddleware())
	r.Use(middleware.LoggingMiddleware(), middleware.ErrorHandlingMiddleware())

	// Swagger
	if mode := gin.Mode(); mode == gin.DebugMode {
		url := ginSwagger.URL(fmt.Sprintf("http://localhost:%d/swagger/doc.json", port))
		r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))
	}

	// 註冊路由和處理器
	routes.Auth(r)

	r.GET("/api/users", handlers.GetUsers)
	r.GET("/api/messages", handlers.GetPrivateMessages)
	r.POST("/api/messages", handlers.PostPrivateMessages)

	// chat
	hub := ws.NewHub()
	go hub.Run()
	r.GET("/ws", ws.ServeWs(hub))

	// 啟動 HTTP 服務器
	err := r.Run(fmt.Sprintf(":%d", port))
	if err != nil {
		panic("Failed to start the server: " + err.Error())
	}
}
