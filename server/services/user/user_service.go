package user

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
	"server/dto"
	"server/pkg/hash"
	"server/pkg/helper"
	"server/repositories"
	"strconv"
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
func (s Service) Login(userDto dto.UserDto) (string, error) {
	model := repositories.GetUserByUsername(userDto.Username)
	if model.ID == 0 {
		return "", errors.New("账号不存在")
	}
	b := hash.Bcrypt{Cost: bcrypt.DefaultCost}
	err := b.Check([]byte(model.Password), []byte(userDto.Password))
	if err != nil {
		return "", errors.New("密码错误")
	}

	// 整合 jwt
	claims := helper.Claims{
		Username:       model.Username,
		Wid:            strconv.Itoa(int(model.ID)),
		StandardClaims: jwt.StandardClaims{},
	}
	// 生成 token
	token, err := helper.Encode(claims, helper.Key)
	if err != nil {
		return "", errors.New("token 生成失败")
	}
	return token, nil
}
