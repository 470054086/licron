package service

import (
	"licron.com/global"
	"licron.com/params/request"
	"licron.com/schedule/types"
)

type Cron struct {
}

func (c *Cron) Add(r *request.CronRequest) error {
	m := types.Cron{
		Name:    r.Name,
		Exp:     r.Exp,
		Command: r.Command,
		Desc:    r.Desc,
		IsDel:   0,
	}
	if err := global.G_DB.Save(&m).Error; err != nil {
		return err
	}
	//global.G_Schedule.AddNotify(&m)
	return nil
}
