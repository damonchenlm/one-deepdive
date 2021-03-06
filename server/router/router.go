package router

import (
	"server/api"
	"server/middlewares"
	"server/pkg/e"

	"github.com/gin-gonic/gin"
)

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}

func SetupRouter() *gin.Engine {
	router := gin.Default()
	router.Use(CORSMiddleware())

	router.GET("/news/getAllNews", api.GetAllNews).Use(middlewares.JwtCheck())
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, "pong")
	})

	// user
	router.POST("/register", e.ErrorWrapper(api.RegisterHandle))
	router.POST("/login", e.ErrorWrapper(api.LoginHandle))

	return router
}
