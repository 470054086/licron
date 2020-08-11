package service

import "licron.com/model"

var (
	daemodModel *model.Daemon
)

func init() {
	daemodModel = &model.Daemon{}
}
