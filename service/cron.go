package service

import (
	"context"
	"licron.com/global"
	"licron.com/model"
	"licron.com/params/request"
	"licron.com/rpc/protoc/cron"
	"licron.com/schedule/constans"
	"licron.com/schedule/types"
)

type Cron struct {
	cronModel model.Cron
}

// 获取全部的数据
func (c *Cron) Lists(r *request.CronListRequest) ([]*request.CronListResponse, error) {
	l := &cron.ListRequest{}
	lists, err := global.G_RPC_CRON.Lists(context.Background(), l)
	if err != nil {
		return nil, err
	}
	var data []*request.CronListResponse
	for _, v := range lists.GetItems() {
		d := request.CronListResponse{
			ID:        int(v.Id),
			Name:      v.Name,
			Exp:       v.Exp,
			Command:   v.Command,
			Desc:      v.Desc,
			CreatedAt: v.CreatedAt,
			UpdatedAt: v.UpdatedAt,
			IsKiller:  int(v.IsKiller),
		}
		data = append(data, &d)
	}
	return data, nil
}

func (c *Cron) Add(r *request.CronRequest) error {
	// 调取rpc进行处理
	rpcRequrst := &cron.AddRequest{R: &cron.CronBase{
		Name:     r.Name,
		Exp:      r.Exp,
		Command:  r.Command,
		Desc:     r.Desc,
		IsKiller: 0,
	}}
	_, err := global.G_RPC_CRON.Add(context.Background(), rpcRequrst)
	if err != nil {
		return nil
	}
	// 生成数据 进行通知
	m := types.Cron{
		Name:     r.Name,
		Exp:      r.Exp,
		Command:  r.Command,
		Desc:     r.Desc,
		IsKiller: 0,
	}
	global.G_Schedule.AddNotify(&m)
	return nil
}

func (d *Cron) Killer(r *request.DeamonKillRequest) error {
	// 判断是否存在
	var rpcRequrst = cron.KillerRequest{Id: int32(r.ID)}
	if _, err := global.G_RPC_CRON.Killer(context.Background(), &rpcRequrst); err != nil {
		return err
	} else {
		var idRequest = cron.IdRequest{Id: int32(r.ID)}
		cron, _ := global.G_RPC_CRON.GetFirstById(context.Background(), &idRequest)
		t := d.TransFrom(cron)
		global.G_Schedule.KillNotify(t)
	}
	return nil
}

func (d *Cron) TransFrom(cron *cron.CronBase) *types.Cron {
	return &types.Cron{
		ID:       int(cron.Id),
		Name:     cron.Name,
		Exp:      cron.Exp,
		Command:  cron.Command,
		Desc:     cron.Desc,
		IsKiller: int(cron.IsKiller),
		Types:    constans.NormalType,
	}
}
