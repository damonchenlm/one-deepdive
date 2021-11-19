package api

import (
	"github.com/gin-gonic/gin"
	"server/dto"
	"server/pkg/e"
	"server/requests"
	"server/services/user"
)

// RegisterHandle 注册控制器
func RegisterHandle(c *gin.Context) (interface{}, error) {
	request := requests.RegisterRequest{}
	err := c.ShouldBind(&request)
	if err != nil {
		return nil, e.ApiError{
			Status:  422,
			Code:    40004,
			Message: "参数校验失败",
		}
	}
	userDto := dto.UserDto{
		Username: request.Username,
		Password: request.Password,
	}
	service := user.Service{}
	err = service.Register(userDto)
	if err != nil {
		return nil, e.ApiError{
			Status:  422,
			Code:    40005,
			Message: err.Error(),
		}
	}
	return "注册成功", nil
}

// LoginHandle 登录控制器
func LoginHandle(c *gin.Context) (interface{}, error) {
	request := requests.LoginRequest{}
	err := c.ShouldBind(&request)
	if err != nil {
		return nil, e.ApiError{
			Status:  422,
			Code:    40004,
			Message: "参数校验失败",
		}
	}
	userDto := dto.UserDto{
		Username: request.Username,
		Password: request.Password,
	}
	service := user.Service{}
	err = service.Login(userDto)
	if err != nil {
		return nil, e.ApiError{
			Status:  422,
			Code:    40005,
			Message: err.Error(),
		}
	}
	return "登陆成功", nil
}
