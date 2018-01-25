package main

import (
	_ "admin.dh.cn/routers"
	"github.com/astaxie/beego"


)

func main() {
	beego.BConfig.WebConfig.Session.SessionOn = true
	beego.Run()
}
