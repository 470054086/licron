package router

import (
	"github.com/gin-gonic/gin"
	"licron.com/api/deamon"
)

// 定时任务的路由
func DeamonRouter(r *gin.RouterGroup) {
	group := r.Group("/deamon")
	{
		// 添加
		group.POST("/add", deamon.Add)
		// 杀死任务
		group.POST("/kill", deamon.Kill)
	}
}
