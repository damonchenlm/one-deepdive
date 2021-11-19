package middlewares

import (
	"github.com/gin-gonic/gin"
	"server/pkg/e"
	"server/pkg/helper"
	"strings"
)

func JwtCheck() gin.HandlerFunc {
	return func(c *gin.Context) {
		header := c.GetHeader("Authorization")
		split := strings.Split(header, " ")
		if split[0] != "Bearer" && split[1] == "" {
			c.JSON(422, e.ApiError{
				Status:  422,
				Code:    40007,
				Message: "不允许访问",
			})
			c.Abort()
			return
		}
		decode, err := helper.Decode(split[1], helper.Key)
		if err != nil {
			c.JSON(422, e.ApiError{
				Status:  422,
				Code:    40008,
				Message: "解析token失败",
			})
			c.Abort()
			return
		}

		c.Set("X-ID", decode.Wid)
		c.Set("X-USERNAME", decode.Username)
		c.Next()
	}
}
