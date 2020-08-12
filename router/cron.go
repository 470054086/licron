package router

import (
	"github.com/gin-gonic/gin"
	"licron.com/api/cron"
)

// 定时任务的路由
func CronRouter(r *gin.RouterGroup) {
	group := r.Group("/cron")
	{
		group.POST("/lists", cron.Lists)
		group.POST("/add", cron.Add)
		group.POST("/kill", cron.Kill)
	}
}
