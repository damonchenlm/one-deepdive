package model

import "server/global"

// News 结构体
type News struct {
	global.CommonModel
	Title   string `json:"title" form:"title" gorm:"column:title;comment:标题;type:varchar;"`
	Author  string `json:"author" form:"author" gorm:"column:author;comment:作者;type:varchar;"`
	Content string `json:"content" form:"content" gorm:"column:content;comment:新闻;type:text;"`
}

// TableName News 表名
func (News) TableName() string {
	return "odd_news"
}

func GetAllNews() (res []News, err error) {
	result := global.DB.Model(&News{}).Find(&res)
	return res, result.Error
}
