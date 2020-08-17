package model

import (
	"licron.com/global"
	"time"
)

// cron数据库表达式
type Cron struct {
	ID        int
	Name      string
	Exp       string
	Command   string
	Desc      string
	IsKiller  int
	CreatedAt time.Time `gorm:"column:created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at"`
}

func (c Cron) TableName() string {
	return "cron"
}

func (c Cron) GetAll() ([]*Cron, error) {
	var r []*Cron
	if err := global.G_DB.Where("is_killer=?", 0).Find(&r).Error; err != nil {
		return nil, err
	}
	return r, nil
}

// 获取单个任务
func (c Cron) GetFirstById(id int) (*Cron, error) {
	var r Cron
	if err := global.G_DB.Where("id = ? and is_killer = ?", id, 0).First(&r).Error; err != nil {
		return nil, err
	}
	return &r, nil
}

// 修改某些字段
func (c Cron) UpdateById(id int32, update *Cron) error {
	err := global.G_DB.Where("id = ? and is_killer = ?", id, 0).Update(&update).Error
	return err
}

// 杀死任务
func (c Cron) Killer(id int) error {
	r := Cron{IsKiller: 1}
	return global.G_DB.Where("id = ? and is_killer = ?", id, 0).Update(&r).Error
}

// 上线某个任务
func (c Cron) Online(id int) error {
	r := Cron{IsKiller: 0}
	return global.G_DB.Where("id = ? and is_killer = ?", id, 0).Update(&r).Error
}
