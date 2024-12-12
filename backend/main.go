package main

import (
	"log"
	"lunchking/backend/api"

	"github.com/joho/godotenv"
)

func main() {
	// 加載 .env 文件（如果有）
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}

	// 啟動路由器並運行伺服器
	router := api.SetupRouter()
	if err := router.Run(":8080"); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
