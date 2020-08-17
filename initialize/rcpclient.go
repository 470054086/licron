package initialize

import (
	"fmt"
	"google.golang.org/grpc"
	"licron.com/global"
	"licron.com/rpc/protoc/cron"
	"licron.com/rpc/protoc/daemon"
)

// 初始化rpc服务
func InitRpcClient() []*grpc.ClientConn {
	client, cronService := cronClient()
	dclient, daemonService := daemonClient()
	var g []*grpc.ClientConn
	//单例模式
	global.G_RPC_CRON = cronService
	global.G_RPC_DAEMON = daemonService
	g = append(g, client)
	g = append(g, dclient)
	return g
}

func cronClient() (*grpc.ClientConn, cron.CronClient) {
	address := fmt.Sprintf("%s:%s", global.G_Config.App.CronRpcHost, global.G_Config.App.CronRpcPort)
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		global.G_LOG.Error(fmt.Sprintf("rpc client error is %+v", err))
		fmt.Println(err)
	}
	cronService := cron.NewCronClient(conn)
	return conn, cronService
}

func daemonClient() (*grpc.ClientConn, daemon.DaemonClient) {
	address := fmt.Sprintf("%s:%s", global.G_Config.App.DaemonRpcHost, global.G_Config.App.DaemonRpcPort)
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		global.G_LOG.Error(fmt.Sprintf("rpc client error is %+v", err))
		fmt.Println(err)
	}
	cronService := daemon.NewDaemonClient(conn)
	return conn, cronService
}
