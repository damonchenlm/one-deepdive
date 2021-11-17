package models

import "server/global"

type User struct {
	global.CommonModel
	Username string
	Password string
}
