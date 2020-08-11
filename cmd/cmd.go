package main

import (
	_ "licron.com/core"
	"licron.com/global"
	"licron.com/initialize"
	"licron.com/log"
	"licron.com/schedule"
)

func main() {
	// 启动数据库
	log.New()
	if global.G_Config.App.DbType == "mysql" {
		initialize.InitMysql()
		defer global.G_DB.Close()
	}


	d := schedule.DaemonNew(10, 10, 10, 10, 10)
	go d.Run()
}
