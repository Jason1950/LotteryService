package main

import (
	"LotteryService/api/v1/auth"
	"LotteryService/internal/handler"
	"LotteryService/internal/service"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// 初始化服務
	authService := service.NewAuthService()
	authHandler := handler.NewAuthHandler(authService)

	// 註冊路由
	auth.RegisterRoutes(r, authHandler)

	// 啟動服務器
	r.Run(":8080")
}
