package user

import (
	"errors"
	"golang.org/x/crypto/bcrypt"
	"server/dto"
	"server/pkg/hash"
	"server/repositories"
)

type Service struct {
}

// Register 用户注册
func (s Service) Register(userDto dto.UserDto) error {
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

// Login 用户登录
func (s Service) Login(userDto dto.UserDto) error {
	model := repositories.GetUserByUsername(userDto.Username)
	if model.ID == 0 {
		return errors.New("账号不存在")
	}
	b := hash.Bcrypt{Cost: bcrypt.DefaultCost}
	err := b.Check([]byte(model.Password), []byte(userDto.Password))
	if err != nil {
		return errors.New("密码错误")
	}
	return nil
}
