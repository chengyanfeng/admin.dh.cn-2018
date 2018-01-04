package routers

import (
	"admin.dh.cn/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{},"get:Get")
	beego.Router("/user", &controllers.UserController{},"get:List")
	beego.Router("/user/create", &controllers.UserController{},"get:Create")
	beego.Router("/user/edit", &controllers.UserController{},"get:Edit")
	beego.Router("/user/add", &controllers.UserController{},"post:Add")
	beego.Router("/user/update", &controllers.UserController{},"post:Update")
	beego.Router("/user/remove", &controllers.UserController{},"get:Remove")
}
