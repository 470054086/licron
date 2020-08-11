package model

import (
	"licron.com/global"
	"licron.com/schedule/types"
)

// cron数据库表达式
type Cron struct {
	ID      int `gorm:"primary_key"`
	Name    string
	Exp     string
	Command string
	Desc    string
	IsDel   int
}

func (c Cron) GetAll() ([]*types.Cron, error) {
	var r []*types.Cron
	if err := global.G_DB.Where("is_del=?", 0).Find(&r).Error; err != nil {
		return nil, err
	}
	return r, nil
}
