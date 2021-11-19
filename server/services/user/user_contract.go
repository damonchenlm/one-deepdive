package user

import "server/dto"

type Contract interface {
	Register(dto dto.UserDto) error
	Login(dto dto.UserDto) error
}
