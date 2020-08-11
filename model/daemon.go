package model

import (
	"licron.com/global"
	"licron.com/schedule/types"
)

type Daemon struct {
}

// 获取全部的常驻内存服务
func (c Daemon) GetAll() ([]*types.Deamon, error) {
	var r []*types.Deamon
	if err := global.G_DB.Where("is_del=?", 0).Find(&r).Error; err != nil {
		return nil, err
	}
	return r, nil
}

func (c Daemon) GetFirstById(id int) (*types.Deamon, error) {
	var r types.Deamon
	if err := global.G_DB.Where("id = ? and is_del = ?", id, 0).First(&r).Error; err != nil {
		return nil, err
	}
	return &r, nil
}
