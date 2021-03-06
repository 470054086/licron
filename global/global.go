package global

import (
	"github.com/jinzhu/gorm"
	"github.com/sirupsen/logrus"
	"licron.com/config"
	"licron.com/rpc/protoc/cron"
	"licron.com/rpc/protoc/daemon"
	"licron.com/schedule/inter"
)

// 设置全局变量 单例模式
var (
	G_Config          *config.Config
	G_LOG             *logrus.Entry
	G_DB              *gorm.DB
	G_Schedule        inter.Schedule
	G_Deamon_Schedule inter.Schedule
	G_Oonce_Schedule  inter.Schedule
	G_RPC_CRON        cron.CronClient
	G_RPC_DAEMON      daemon.DaemonClient
)
