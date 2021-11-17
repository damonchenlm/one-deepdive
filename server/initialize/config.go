package initialize

import (
	"github.com/spf13/viper"
	"server/conf"
	"server/global"
)

func InitConfig() {
	// 实例化 Viper
	v := viper.New()
	// 读取配置文件
	v.SetConfigFile("./conf/config.yaml")
	if err := v.ReadInConfig(); err != nil {
		panic(err)
	}
	// 解析配置文件
	serverConfig := conf.ServerConfig{}
	if err := v.Unmarshal(&serverConfig); err != nil {
		panic(err)
	}
	global.Settings = serverConfig
}
