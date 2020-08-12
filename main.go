package main

import (
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

	// 启动ectd
	initialize.InitSchedule()

	// 启动路由 Http服务
	initialize.InitRouter()
}
