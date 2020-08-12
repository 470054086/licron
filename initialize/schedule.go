package initialize

import (
	"licron.com/global"
	"licron.com/schedule"
)

func InitSchedule() {
	// 初始化
	if global.G_Config.App.Schedule == "true" {
		s := schedule.NewSimple(10)
		go s.Run()
		global.G_Schedule = s
		// 初始化
		d := schedule.DaemonNew(10)
		go d.Run()
		global.G_Deamon_Schedule = d
		// 初始化
		once := schedule.NewOnce(10)
		go once.Run()
		global.G_Oonce_Schedule = once
	}
}
