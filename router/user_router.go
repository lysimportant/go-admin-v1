package router

import "lianghj.top/admin/goadminv1/controller"

func UserRouter() {
	userRouter := Router.Group("/user")
	userRouter.GET("/", controller.GetUserAllController)
}
