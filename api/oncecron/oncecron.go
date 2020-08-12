package oncecron

import (
	"github.com/gin-gonic/gin"
	"licron.com/params/request"
	"licron.com/service"
	"licron.com/utils"
	"net/http"
)

// 获取service类
var (
	onceService *service.OnceService
)

// 添加某个任务
func Add(c *gin.Context) {
	var r request.OnceRequest
	if err := c.ShouldBindJSON(&r); err != nil {
		paramError := utils.ResponseParamError()
		c.JSON(http.StatusOK, paramError)
		return
	}

	// 传递参数给service端调用
	err := onceService.Add(&r)
	if err != nil {
		paramError := utils.ResponseActionError()
		c.JSON(http.StatusOK, paramError)
		return
	}
	c.JSON(http.StatusOK, utils.ResponseSuccess())
}

// 杀死任务
func Kill(c *gin.Context) {
	var r request.DeamonKillRequest
	if err := c.ShouldBindJSON(&r); err != nil {
		paramError := utils.ResponseParamError()
		c.JSON(http.StatusOK, paramError)
		return
	}
	// 传递参数给service端调用
	err := onceService.Killer(&r)
	if err != nil {
		paramError := utils.ResponseActionError()
		c.JSON(http.StatusOK, paramError)
		return
	}
	c.JSON(http.StatusOK, utils.ResponseSuccess())
}
