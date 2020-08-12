package deamon

import (
	"github.com/gin-gonic/gin"
	"licron.com/params/request"
	"licron.com/service"
	"licron.com/utils"
	"net/http"
)

// 获取service类
var (
	daemonService *service.DeamonService
)

func init() {
	daemonService = &service.DeamonService{}
}
// 添加任务
func Add(c *gin.Context) {
	var r request.DeamonRequest
	if err := c.ShouldBindJSON(&r); err != nil {
		paramError := utils.ResponseParamError()
		c.JSON(http.StatusOK, paramError)
		return
	}
	// 传递参数给service端调用
	err := daemonService.Add(&r)
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
	err := daemonService.Killer(&r)
	if err != nil {
		paramError := utils.ResponseActionError()
		c.JSON(http.StatusOK, paramError)
		return
	}
	c.JSON(http.StatusOK, utils.ResponseSuccess())
}
