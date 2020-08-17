package main

import (
	"fmt"
	"google.golang.org/grpc"
	_ "licron.com/core"
	"licron.com/global"
	"licron.com/initialize"
	"licron.com/log"
	"licron.com/rpc/protoc/daemon"
	"licron.com/rpc/server/daemon/service"
	"net"
)

func main() {
	//启动日志
	log.New()
	// 启动数据库
	initialize.InitMysql()
	// 捕捉错误
	defer func() {
		if err := recover(); err != nil {
			global.G_LOG.Info(fmt.Sprintf("internal server error is %+v", err))
		}
	}()
	port := global.G_Config.App.DaemonRpcPort
	listen, err := net.Listen("tcp", fmt.Sprintf(":%s", port))
	if err != nil {
		global.G_LOG.Info("failed to listen: %v", err)
	}
	global.G_LOG.Info(fmt.Sprintf("cron rpc server is start port is %s", port))

	server := grpc.NewServer()
	daemon.RegisterDaemonServer(server,&service.DaemonRpc{} )
	err = server.Serve(listen)
	if err != nil {
		global.G_LOG.Info("fail server start is %v", err)
	}
	global.G_LOG.Info(fmt.Sprintf("cron rpc service start"))
}
