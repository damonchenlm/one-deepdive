package handler

import (
	"github.com/gin-gonic/gin"
	"server/model"
)

func GetAllNews(context *gin.Context) {
	news, err := model.GetAllNews()
	if err != nil {
		context.JSON(500, err)
	}
	context.JSON(200, news)
}
