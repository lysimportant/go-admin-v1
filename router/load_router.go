package router

import "github.com/gin-gonic/gin"

var Router *gin.Engine

func LoadRouter() {
	Router = gin.Default()
	UserRouter()
	Router.GET("/", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"code": 0,
			"data": "",
			"msg":  "现在测试的Gin已正常运行~",
		})
	})
	Router.Run()
}
