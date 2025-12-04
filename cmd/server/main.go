package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/jlau-ice/collect/internal/config"
	"github.com/jlau-ice/collect/internal/database"
	"github.com/jlau-ice/collect/internal/router"
)

func main() {
	// 加载配置
	if err := config.LoadConfig(); err != nil {
		log.Fatalf("加载配置失败: %v", err)
	}

	// 设置Gin模式
	gin.SetMode(config.AppConfig.Server.Mode)

	// 初始化数据库
	if err := database.InitDB(); err != nil {
		log.Fatalf("数据库初始化失败: %v", err)
	}
	defer database.CloseDB()

	// 创建Gin引擎
	r := gin.Default()

	// 设置路由
	router.SetupRoutes(r)

	// 启动服务器
	port := ":" + config.AppConfig.Server.Port
	log.Printf("服务器启动在端口 %s", port)
	if err := r.Run(port); err != nil {
		log.Fatalf("服务器启动失败: %v", err)
	}
}
