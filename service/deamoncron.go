package service

import (
	"licron.com/global"
	"licron.com/params/request"
	"licron.com/schedule/types"
)

type DeamonService struct {
}

func (d *DeamonService) Add(r *request.DeamonRequest) error {
	m := &types.Deamon{
		Name:     r.Name,
		Command:  r.Command,
		Desc:     r.Desc,
		IsOnline: 0,
		IsDel:    0,
	}
	if err := global.G_DB.Save(m).Error; err != nil {
		return err
	}
	global.G_Deamon_Schedule.AddNotify(m)
	return nil
}

func (d *DeamonService) Killer(r *request.DeamonKillRequest) error {
	// 判断是否存在
	id := r.ID
	if daemon, err := daemodModel.GetFirstById(id); err != nil {
		return err
	} else {
		global.G_Deamon_Schedule.KillNotify(daemon)
	}
	return nil
}
