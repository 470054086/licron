package model

import (
	"licron.com/global"
	"time"
)

// model类
type Daemon struct {
	ID       int
	Name     string
	Command  string
	Desc     string
	IsKiller int
	CreatedAt time.Time `gorm:"column:created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at"`
}

func (c Daemon) TableName() string {
	return "daemon_cron"
}

// 获取全部的常驻内存服务
func (c Daemon) GetAll() ([]*Daemon, error) {
	var r []*Daemon
	if err := global.G_DB.Where("is_killer=?", 0).Find(&r).Error; err != nil {
		return nil, err
	}
	return r, nil
}

func (c Daemon) UpdateById(id int,daemon *Daemon) error {
	err := global.G_DB.Where("id = ? and is_killer = ?", id, 0).Update(&daemon).Error
	return err
}

// 获取单个任务
func (c Daemon) GetFirstById(id int) (*Daemon, error) {
	var r Daemon
	if err := global.G_DB.Where("id = ? and is_killer = ?", id, 0).First(&r).Error; err != nil {
		return nil, err
	}
	return &r, nil
}

// 杀死任务
func (c Daemon) Killer(id int) error {
	r := Daemon{IsKiller: 1}
	return global.G_DB.Where("id = ? and is_killer = ?", id, 0).Update(&r).Error
}

// 上线某个任务
func (c Daemon) Online(id int) error {
	r := Daemon{IsKiller: 0}
	return global.G_DB.Where("id = ? and is_killer = ?", id, 0).Update(&r).Error
}
