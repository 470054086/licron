package service

import (
	"context"
	"licron.com/global"
	"licron.com/model"
	pb "licron.com/rpc/protoc/daemon"
)

// CronRpc 数据服务
type DaemonRpc struct {
	m model.Daemon
}

func (c *DaemonRpc) Lists(ctx context.Context, r *pb.ListRequest) (*pb.ListResponse, error) {
	//获取全部的cronList数据
	all, _ := c.m.GetAll()
	data := &pb.ListResponse{Items: nil}
	for _, v := range all {
		d := pb.DaemonList{
			Id:        int32(v.ID),
			Name:      v.Name,
			Command:   v.Command,
			Desc:      v.Desc,
			CreatedAt: v.CreatedAt.String(),
			UpdatedAt: v.UpdatedAt.String(),
			IsKiller:  int32(v.IsKiller),
		}
		data.Items = append(data.Items, &d)
	}
	return data, nil
}
func (c *DaemonRpc) GetFirstById(ctx context.Context, in *pb.IdRequest) (*pb.CronBase, error) {
	cron, err := c.m.GetFirstById(int(in.Id))
	if err != nil {
		return nil, err
	}
	p := &pb.CronBase{
		Name:      cron.Name,
		Command:   cron.Command,
		Desc:      cron.Desc,
		IsKiller:  int32(cron.IsKiller),
		RunAt:     "",
		CreatedAt: cron.CreatedAt.String(),
		UpdatedAt: cron.UpdatedAt.String(),
		Id:        int32(cron.ID),
	}
	return p, nil
}
func (c *DaemonRpc) Add(ctx context.Context, r *pb.AddRequest) (*pb.Error, error) {
	m := model.Daemon{
		Name:     r.R.Name,
		Command:  r.R.Command,
		Desc:     r.R.Desc,
		IsKiller: 0,
	}
	if err := global.G_DB.Save(&m).Error; err != nil {
		return nil, err
	}
	p := &pb.Error{}
	return p, nil
}

func (c *DaemonRpc) Update(ctx context.Context, r *pb.UpdateRequest) (*pb.Error, error) {
	var m = model.Daemon{
		Name:     r.R.Name,
		Command:  r.R.Command,
		Desc:     r.R.Desc,
		IsKiller: int(r.R.IsKiller),
	}
	err := c.m.UpdateById(int(r.Id), &m)
	p := &pb.Error{}
	return p, err
}

func (c *DaemonRpc) Killer(ctx context.Context, r *pb.KillerRequest) (*pb.Error, error) {
	id := r.Id
	if _, err := c.m.GetFirstById(int(id)); err != nil {
		return nil, err
	}
	err := c.m.Killer(int(id))
	p := &pb.Error{}
	return p, err
}

func (c *DaemonRpc) OnLine(ctx context.Context, r *pb.KillerRequest) (*pb.Error, error) {
	id := r.Id
	if _, err := c.m.GetFirstById(int(id)); err != nil {
		return nil, err
	}
	err := c.m.Online(int(id))
	p := &pb.Error{}
	return p, err
}
