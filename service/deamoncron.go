package service

import (
	"licron.com/global"
	"licron.com/model"
	"licron.com/params/request"
	"licron.com/schedule/constans"
	"licron.com/schedule/types"
)

type DeamonService struct {
}

func (d *DeamonService) Add(r *request.DeamonRequest) error {
	m := &model.Daemon{
		Name:     r.Name,
		Command:  r.Command,
		Desc:     r.Desc,
		IsOnline: 0,
		IsDel:    0,
	}
	if err := global.G_DB.Save(m).Error; err != nil {
		return err
	}

	// 发送channel
	t := &types.Cron{
		ID:       m.ID,
		Name:     m.Name,
		Exp:      "",
		Command:  m.Command,
		Desc:     m.Desc,
		IsDel:    0,
		IsOnline: 0,
	}
	global.G_Deamon_Schedule.AddNotify(t)
	return nil
}

func (d *DeamonService) Killer(r *request.DeamonKillRequest) error {
	// 判断是否存在
	id := r.ID
	if daemon, err := daemodModel.GetFirstById(id); err != nil {
		return err
	} else {
		t := d.TransFrom(daemon)
		global.G_Deamon_Schedule.KillNotify(t)
	}
	return nil
}

func (d *DeamonService) TransFrom(daemon *model.Daemon) *types.Cron {
	return &types.Cron{
		ID:       daemon.ID,
		Name:     daemon.Name,
		Exp:      "",
		Command:  daemon.Command,
		Desc:     daemon.Desc,
		IsDel:    daemon.IsDel,
		IsOnline: daemon.IsOnline,
		Types:    constans.DaemonType,
	}
}
