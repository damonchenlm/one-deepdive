package main

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"server/global"
	"server/model"
	"server/router"
)

func main() {
	// 数据库连接 暂时写在这儿
	dsn := "damon:Cyl851106@tcp(rm-bp1z6653s2t65e4774o.mysql.rds.aliyuncs.com:3306)/one_deepdive?charset=utf8mb4&parseTime=True&loc=Local"
	var err error
	global.DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	_ = global.DB.AutoMigrate(&model.News{})

	engine := router.SetupRouter()
	_ = engine.Run("localhost:8081")
}