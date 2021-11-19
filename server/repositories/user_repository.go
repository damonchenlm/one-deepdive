package repositories

import (
	"server/dto"
	"server/global"
	"server/models"
)

func CreateUser(dto dto.UserDto) error {
	user := models.User{}
	user.Username = dto.Username
	user.Password = dto.Password

	err := global.DB.Create(&user).Error
	return err
}

// GetUserByUsername 根据用户名查询用户
func GetUserByUsername(username string) models.User {
	user := models.User{}
	global.DB.Find(&user, global.DB.Where("username = ?", username))
	return user
}
