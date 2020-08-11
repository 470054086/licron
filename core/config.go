package core

// 加载配置文件类
import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"licron.com/global"
	"os"
)
// 配置文件名称
const (
	defaultConfigName       = "config-local.yaml"
	defaultConfigSimName    = "config-sim.yaml"
	defaultConfigOnlineName = "config-online.yaml"
)

func init() {
	v := viper.New()
	// 根据环境变量 设置配置文件
	switch os.Getenv("mode") {
	case "online":
		v.SetConfigFile(defaultConfigOnlineName)
	case "sim":
		v.SetConfigFile(defaultConfigSimName)
	default:
		v.SetConfigFile(defaultConfigName)
	}
	// 尝试读取 是否报错
	err := v.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
	// 将文件内容写到变量中
	if err := v.Unmarshal(&global.G_Config); err != nil {
		panic(fmt.Errorf("config changge error: %s \n", err))
	}
	fmt.Printf("%+v",global.G_Config)
	// 监听文件
	v.WatchConfig()
	// 文件变化 重新加载内容
	v.OnConfigChange(func(e fsnotify.Event) {
		if err := v.Unmarshal(&global.G_Config); err != nil {
			panic(fmt.Errorf("config changge error: %s \n", err))
		}
	})
}
