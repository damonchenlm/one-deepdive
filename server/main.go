package main

import (
	"fmt"
	"server/global"
	"server/initialize"
	"server/router"
)

func main() {
	// 初始化配置文件
	initialize.InitConfig()

	// 初始化数据库
	initialize.InitDB()

	// 初始化路由
	engine := router.SetupRouter()

	addr := fmt.Sprintf("%s:%d", global.Settings.Host, global.Settings.Port)
	_ = engine.Run(addr)
}
