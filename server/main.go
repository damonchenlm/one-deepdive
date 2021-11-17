package main

import (
	"server/global"
	"server/models"
	"server/router"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	// 数据库连接 暂时写在这儿
	dsn := "damon:Cyl851106@tcp(rm-bp1z6653s2t65e4774o.mysql.rds.aliyuncs.com:3306)/one_deepdive?charset=utf8mb4&parseTime=True&loc=Local"
	var err error
	global.DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	_ = global.DB.AutoMigrate(&models.News{})

	engine := router.SetupRouter()
	_ = engine.Run("0.0.0.0:8081")
}
