package main

import (
	"lianghj.top/admin/goadminv1/components"
	"lianghj.top/admin/goadminv1/router"
)

func main() {
	components.InitAll()
	router.LoadRouter()
}
