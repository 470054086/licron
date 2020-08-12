package service

import (
	"licron.com/global"
	"licron.com/model"
	"licron.com/params/request"
	"licron.com/schedule/constans"
	"licron.com/schedule/types"
)

type OnceService struct {
	onceModel model.OnceCron
}

func (d *OnceService) Add(r *request.OnceRequest) error {
	m := &model.OnceCron{
		Name:     r.Name,
		Command:  r.Command,
		Desc:     r.Desc,
		Runat:    r.Runat,
		IsKiller: 0,
	}
	if err := global.G_DB.Save(m).Error; err != nil {
		return err
	}

	// 发送channel
	t := &types.Cron{
		ID:       m.ID,
		Name:     m.Name,
		Command:  m.Command,
		Desc:     m.Desc,
		Runat:    m.Runat,
		IsKiller: 0,
		Types:    constans.OnceType,
	}
	global.G_Oonce_Schedule.AddNotify(t)
	return nil
}

func (d *OnceService) Killer(r *request.DeamonKillRequest) error {
	// 判断是否存在
	id := r.ID
	if daemon, err := d.onceModel.GetFirstById(id); err != nil {
		return err
	} else {
		t := d.TransFrom(daemon)
		global.G_Deamon_Schedule.KillNotify(t)
	}
	return nil
}

func (d *OnceService) TransFrom(c *model.OnceCron) *types.Cron {
	return &types.Cron{
		ID:       c.ID,
		Name:     c.Name,
		Command:  c.Command,
		Desc:     c.Desc,
		Runat:    c.Runat,
		IsKiller: c.IsKiller,
		Types:    constans.DaemonType,
	}
}
