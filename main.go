package main

import (
	_ "admin.dh.cn/routers"
	"github.com/astaxie/beego"
)

func main() {
	beego.SetStaticPath("/upload", "upload")
	beego.SetStaticPath("/admin/static", "static")
	beego.SetStaticPath("/admin/views", "views/datasource_share")
	beego.BConfig.WebConfig.Session.SessionOn = true

	beego.Run()
}
