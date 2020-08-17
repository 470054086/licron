package service

import (
	"context"
	"licron.com/global"
	"licron.com/model"
	"licron.com/params/request"
	"licron.com/rpc/protoc/daemon"
	"licron.com/schedule/constans"
	"licron.com/schedule/types"
)

type DeamonService struct {
	deamonModel model.Daemon
}

// 获取全部的数据
func (c *DeamonService) Lists(r *request.DeamonListRequest) ([]*request.DaemonListResponse, error) {
	var l daemon.ListRequest
	lists, err := global.G_RPC_DAEMON.Lists(context.Background(), &l)
	if err != nil {
		return nil, err
	}
	var data []*request.DaemonListResponse
	for _, v := range lists.GetItems() {
		d := request.DaemonListResponse{
			ID:        int(v.Id),
			Name:      v.Name,
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

func (d *DeamonService) Add(r *request.DeamonRequest) error {

	m := &daemon.AddRequest{R: &daemon.CronBase{
		Name:     r.Name,
		Command:  r.Command,
		Desc:     r.Desc,
		IsKiller: 0,
	},}
	if _, err := global.G_RPC_DAEMON.Add(context.Background(), m); err != nil {
		return nil
	}
	// 发送channel
	t := &types.Cron{
		ID:       int(m.R.Id),
		Name:     m.R.Name,
		Exp:      "",
		Command:  m.R.Command,
		Desc:     m.R.Desc,
		IsKiller: 0,
	}
	global.G_Deamon_Schedule.AddNotify(t)
	return nil
}

func (d *DeamonService) Killer(r *request.DeamonKillRequest) error {
	// 判断是否存在
	id := r.ID
	p := &daemon.IdRequest{Id: int32(id)}
	daemon, err := global.G_RPC_DAEMON.GetFirstById(context.Background(), p)
	if err != nil {
		return nil
	}

	t := d.TransFrom(daemon)
	global.G_Deamon_Schedule.KillNotify(t)
	return nil
}

func (d *DeamonService) TransFrom(daemon *daemon.CronBase) *types.Cron {
	return &types.Cron{
		ID:       int(daemon.Id),
		Name:     daemon.Name,
		Exp:      "",
		Command:  daemon.Command,
		Desc:     daemon.Desc,
		IsKiller: int(daemon.IsKiller),
		Types:    constans.DaemonType,
	}
}
