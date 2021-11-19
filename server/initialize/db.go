package initialize

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"server/global"
	"server/models"
)

func InitDB() {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		global.Settings.MySQLInfo.Name,
		global.Settings.MySQLInfo.Password,
		global.Settings.MySQLInfo.Host,
		global.Settings.MySQLInfo.Port,
		global.Settings.MySQLInfo.DBName)
	var err error
	global.DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	_ = global.DB.AutoMigrate(&models.News{})
	_ = global.DB.AutoMigrate(&models.User{})
}
