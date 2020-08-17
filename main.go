package main

import (
	"fmt"
	_ "licron.com/core"
	"licron.com/global"
	"licron.com/initialize"
	"licron.com/log"
)

func main() {
	//启动日志
	log.New()

	// 启动数据库
	initialize.InitMysql()
	defer global.G_DB.Close()
	// 启动redis
	client := initialize.InitRpcClient()
	// 程序退出 退出所有的链接
	defer func() {
		for _, v := range client {
			v.Close()
		}
	}()
	global.G_LOG.Info(fmt.Sprintf("rpc server is success"))
	// 启动ectd
	initialize.InitSchedule()

	// 启动路由 Http服务
	initialize.InitRouter()
}
