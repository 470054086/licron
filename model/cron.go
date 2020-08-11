package model

import (
	"licron.com/global"
)

// cron数据库表达式
type Cron struct {
	ID      int
	Name    string
	Exp     string
	Command string
	Desc    string
	IsDel   int
}

func (c Cron) TableName() string {
	return "cron"
}

func (c Cron) GetAll() ([]*Cron, error) {
	var r []*Cron
	if err := global.G_DB.Where("is_del=?", 0).Find(&r).Error; err != nil {
		return nil, err
	}
	return r, nil
}
