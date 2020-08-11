package router

import (
	"github.com/gin-gonic/gin"
	"licron.com/api/oncecron"
)

// 定时任务的路由
func OnceRouter(r *gin.RouterGroup) {
	group := r.Group("/once")
	{
		group.POST("/lists", oncecron.Lists)
	}
}
