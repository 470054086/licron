package service

import (
	"context"
	"licron.com/model"
	pb "licron.com/rpc/protoc"
	"licron.com/rpc/protoc/cron"
)

type CronRpc struct {
	m model.Cron
}

func (c *CronRpc) CronLists(ctx context.Context, in *cron.SearchPage) (*pb.CronList, error) {
	//获取全部的cronList数据
	all, _ := c.m.GetAll()
	data := &pb.CronList{Items: nil}
	for _, v := range all {
		d := cron.CronListResponse{
			Id:        int32(v.ID),
			Name:      v.Name,
			Exp:       v.Exp,
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
