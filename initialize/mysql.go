package initialize

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"licron.com/global"
)

func InitMysql() {
	config := global.G_Config.Mysql
	if &config == nil {
		panic(fmt.Sprintf("mysql not config"))
	}
	mysqlConfig := config.Username + ":" + config.Password + "@(" + config.Path + ")/" + config.Dbname + "?" + config.Config
	fmt.Println(mysqlConfig)
	if db, err := gorm.Open("mysql", mysqlConfig); err != nil {
		panic(fmt.Sprintf("mysql start is error %v", err))
	} else {
		db.DB().SetMaxOpenConns(config.MaxOpenConns)
		db.DB().SetMaxIdleConns(config.MaxIdleConns)
		db.LogMode(config.LogMode)
		global.G_DB = db
	}
	global.G_LOG.Info("mysql start success")
}
