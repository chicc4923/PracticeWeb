package routers

import (
	jwt "PracticeWeb/middleware"
	"PracticeWeb/pkg/setting"
	"PracticeWeb/routers/api"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	r := gin.New()

	r.Use(gin.Logger())

	r.Use(gin.Recovery())

	gin.SetMode(setting.RunMode)

	//r.GET("/test", func(c *gin.Context) {
	//	c.JSON(200, gin.H{
	//		"message": "test",
	//	})
	//})
	r.POST("/user/add", api.AddUser)
	r.GET("/auth", api.GetAuth)
	r.Use(jwt.JWT())

	return r
}
