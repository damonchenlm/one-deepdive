package global

import (
	"gorm.io/gorm"
	"server/conf"
)

var (
	DB       *gorm.DB
	Settings conf.ServerConfig
)
