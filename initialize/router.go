package initialize

// 路由初始化
import (
	"fmt"
	"github.com/gin-gonic/gin"
	"licron.com/global"
	"licron.com/router"
	"net/http"
)

// 使用gin的路由
func InitRouter() {
	e := gin.New()
	// 添加路由
	addRouter(e.Group(""))
	port := global.G_Config.App.Port
	err := http.ListenAndServe(fmt.Sprintf(":%s", port), e)
	if err != nil {
		panic(fmt.Sprintf("http start is erros %v", err))
	}
	global.G_LOG.Info("http port %s is running", port)
}

// 添加路由的操作
func addRouter(group *gin.RouterGroup) {
	router.CronRouter(group)
	router.DeamonRouter(group)
	router.OnceRouter(group)
}
