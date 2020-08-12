package cron

import (
	"github.com/gin-gonic/gin"
	cronexpr "github.com/gorhill/cronexpr"
	"licron.com/params/request"
	"licron.com/service"
	"licron.com/utils"
	"net/http"
)

// 获取service类
var (
	cronService *service.Cron
)

func init() {
	cronService = &service.Cron{}
}

// 获取所有的列表
func Lists(c *gin.Context) {
	var r request.CronListRequest
	if err := c.ShouldBindJSON(&r); err != nil {
		paramError := utils.ResponseParamError()
		c.JSON(http.StatusOK, paramError)
		return
	}
	lists, err := cronService.Lists(&r)
	if err != nil {
		paramError := utils.ResponseActionError()
		c.JSON(http.StatusOK, paramError)
		return
	}
	c.JSON(http.StatusOK, lists)
}

// 添加某个任务
func Add(c *gin.Context) {
	var r request.CronRequest
	if err := c.ShouldBindJSON(&r); err != nil {
		paramError := utils.ResponseParamError()
		c.JSON(http.StatusOK, paramError)
		return
	}
	exp := r.Exp
	if _, err := cronexpr.Parse(exp); err != nil {
		paramError := utils.ResponseMessageError("命令行解析错误")
		c.JSON(http.StatusOK, paramError)
		return
	}
	// 传递参数给service端调用
	err := cronService.Add(&r)
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
	err := cronService.Killer(&r)
	if err != nil {
		paramError := utils.ResponseActionError()
		c.JSON(http.StatusOK, paramError)
		return
	}
	c.JSON(http.StatusOK, utils.ResponseSuccess())
}
