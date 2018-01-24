package main

import (

	_ "admin.dh.cn/routers"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"

)

func main() {


	// 自动建表
	orm.RunSyncdb("default", false, true)
	beego.Run()
}

