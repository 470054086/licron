package initialize

import (
	"licron.com/global"
	"licron.com/schedule"
)

func InitSchedule() {
	if global.G_Config.App.Schedule == "simple" {
		//s := schedule.New(10, 10, 10, 10, 10)
		//go s.Run()
		//global.G_Schedule = s

		d := schedule.DaemonNew(10, 10, 10, 10, 10)
		go d.Run()
		global.G_Deamon_Schedule = d
	}

}
