package api

import (
	"github.com/gin-gonic/gin"
	"server/models"
)

func GetAllNews(context *gin.Context) {
	news, err := models.GetAllNews()
	if err != nil {
		context.JSON(500, err)
	}
	context.JSON(200, news)
}
