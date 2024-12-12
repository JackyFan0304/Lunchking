package api

import (
	"lunchking/backend/api/handlers"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	// 定義路由
	r.GET("/candidates", handlers.GetCandidatesHandler()) // 獲取候選人列表
	r.POST("/vote", handlers.VoteHandler())               // 投票功能

	return r
}
