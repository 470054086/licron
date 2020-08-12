package model

import (
	"licron.com/global"
	"time"
)

// model类
type OnceCron struct {
	ID       int
	Name     string
	Runat    time.Time `gorm:"column:run_at"`
	Command  string
	Desc     string
	IsKiller int
}

func (c OnceCron) TableName() string {
	return "once_cron"
}

// 获取全部的常驻内存服务
func (c OnceCron) GetAll() ([]*OnceCron, error) {
	var r []*OnceCron
	if err := global.G_DB.Where("is_killer=?", 0).Find(&r).Error; err != nil {
		return nil, err
	}
	return r, nil
}

// 获取单个任务
func (c OnceCron) GetFirstById(id int) (*OnceCron, error) {
	var r OnceCron
	if err := global.G_DB.Where("id = ? and is_killer = ?", id, 0).First(&r).Error; err != nil {
		return nil, err
	}
	return &r, nil
}

// 杀死任务
func (c OnceCron) Killer(id int) error {
	r := OnceCron{IsKiller: 1}
	return global.G_DB.Where("id = ? and is_killer = ?", id, 0).Update(&r).Error
}

// 上线某个任务
func (c OnceCron) Online(id int) error {
	r := Daemon{IsKiller: 0}
	return global.G_DB.Where("id = ? and is_killer = ?", id, 0).Update(&r).Error
}
