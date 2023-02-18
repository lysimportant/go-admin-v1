package controller

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"lianghj.top/admin/goadminv1/service"
)

func GetUserAllController(ctx *gin.Context) {
	users, err := service.GetUserAll()
	if err != nil {
		fmt.Println("db err: ", err)
		ctx.JSON(http.StatusOK, gin.H{
			"code": 1,
			"msg":  "系统错误: " + err.Error(),
			"data": err,
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"code": 0,
		"msg":  "查找成功~",
		"data": users,
	})
}
