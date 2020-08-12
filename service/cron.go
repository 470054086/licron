package service

import (
	"licron.com/global"
	"licron.com/model"
	"licron.com/params/request"
	"licron.com/schedule/constans"
	"licron.com/schedule/types"
)

type Cron struct {
	cronModel model.Cron
}

// 获取全部的数据
func (c *Cron) Lists(r *request.CronListRequest) ([]*request.CronListResponse, error) {
	lists, err := c.cronModel.GetAll()
	if err != nil {
		return nil, err
	}
	var data []*request.CronListResponse
	for _, v := range lists {
		d := request.CronListResponse{
			ID:        v.ID,
			Name:      v.Name,
			Exp:       v.Exp,
			Command:   v.Command,
			Desc:      v.Desc,
			CreatedAt: v.CreatedAt,
			UpdatedAt: v.UpdatedAt,
			IsKiller:  v.IsKiller,
		}
		data = append(data, &d)
	}
	return data, nil
}

func (c *Cron) Add(r *request.CronRequest) error {
	m := types.Cron{
		Name:     r.Name,
		Exp:      r.Exp,
		Command:  r.Command,
		Desc:     r.Desc,
		IsKiller: 0,
	}
	if err := global.G_DB.Save(&m).Error; err != nil {
		return err
	}
	global.G_Schedule.AddNotify(&m)
	return nil
}

func (d *Cron) Killer(r *request.DeamonKillRequest) error {
	// 判断是否存在
	id := r.ID
	if cron, err := d.cronModel.GetFirstById(id); err != nil {
		return err
	} else {
		t := d.TransFrom(cron)
		global.G_Schedule.KillNotify(t)
	}
	return nil
}

func (d *Cron) TransFrom(cron *model.Cron) *types.Cron {
	return &types.Cron{
		ID:       cron.ID,
		Name:     cron.Name,
		Exp:      cron.Exp,
		Command:  cron.Command,
		Desc:     cron.Desc,
		IsKiller: cron.IsKiller,
		Types:    constans.NormalType,
	}
}
