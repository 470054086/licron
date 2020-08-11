package model

import (
	"licron.com/global"
)

// model类
type Daemon struct {
	ID       int
	Name     string
	Command  string
	Desc     string
	IsOnline int
	IsDel    int
}

func (c Daemon) TableName() string {
	return "daemon_cron"
}

// 获取全部的常驻内存服务
func (c Daemon) GetAll() ([]*Daemon, error) {
	var r []*Daemon
	if err := global.G_DB.Where("is_del=?", 0).Find(&r).Error; err != nil {
		return nil, err
	}
	return r, nil
}

// 获取单个任务
func (c Daemon) GetFirstById(id int) (*Daemon, error) {
	var r Daemon
	if err := global.G_DB.Where("id = ? and is_del = ?", id, 0).First(&r).Error; err != nil {
		return nil, err
	}
	return &r, nil
}
