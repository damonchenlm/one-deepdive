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
