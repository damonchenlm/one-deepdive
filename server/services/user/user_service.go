package user

import (
	"errors"
	"golang.org/x/crypto/bcrypt"
	"server/dto"
	"server/pkg/hash"
	"server/repositories"
)

type UserService struct {
}

// Register 用户注册
func (s UserService) Register(userDto dto.UserDto) error {
	// 密码加密
	// 密码加密
	//实例化结构体
	b := hash.Bcrypt{Cost: bcrypt.DefaultCost}
	bytes, err := b.Make([]byte(userDto.Password))
	if err != nil {
		return errors.New(err.Error())
	}
	userDto.Password = string(bytes)
	return repositories.CreateUser(userDto)
}
