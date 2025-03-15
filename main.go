package main

import (
	"LotteryService/api/v1/auth"
	"LotteryService/internal/config"
	"LotteryService/internal/handler"
	"LotteryService/internal/service"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	// 加載配置
	if err := config.LoadConfig(); err != nil {
		log.Fatal("加載配置失敗:", err)
	}

	// 初始化數據庫連接
	db, err := config.InitDB()
	if err != nil {
		log.Fatal("數據庫連接失敗:", err)
	}
	defer db.Close()

	r := gin.Default()

	// 初始化服務
	authService := service.NewAuthService(db)
	authHandler := handler.NewAuthHandler(authService)

	// 註冊路由
	auth.RegisterRoutes(r, authHandler)

	// 啟動服務器
	r.Run(":8080")
}
