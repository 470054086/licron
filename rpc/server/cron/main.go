package main

import (
	"google.golang.org/grpc"
	_ "licron.com/core"
	"licron.com/global"
	"licron.com/initialize"
	"licron.com/log"
	"licron.com/rpc/protoc/cron"
	"licron.com/rpc/server/service"
	"net"
)

func main() {
	//启动日志
	log.New()
	// 启动数据库
	initialize.InitMysql()

	listen, err := net.Listen("tcp", ":9090")
	if err != nil {
		global.G_LOG.Info("failed to listen: %v", err)
	}
	server := grpc.NewServer()
	cron.RegisterCronServer(server, &service.CronRpc{})

	server.Serve(listen)
}
