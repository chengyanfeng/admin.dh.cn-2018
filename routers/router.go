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
	beego.Router("/user/updateStatusAva", &controllers.UserController{},"get:UpdateStatusAva")
	beego.Router("/user/getCorp", &controllers.UserController{},"get:GetCorp")
	beego.Router("/user/delectAndAddCorp", &controllers.UserController{},"get:DelectAndAddCorp")
	beego.Router("/user/listremove", &controllers.UserController{},"post:Listremove")
	beego.Router("/user/listChangeType", &controllers.UserController{},"post:ListChangeType")
	beego.Router("/user/changeType", &controllers.UserController{},"get:ChangeType")
	beego.Router("/user/getUserScreen", &controllers.UserController{},"get:GetUserScreen")
	beego.Router("/user/delectUserScreen", &controllers.UserController{},"get:DelectUserScreen")

	//自动注入路由
	beego.AutoRouter(&controllers.CorpController{})
	beego.AutoRouter(&controllers.SourcetypeController{})
	beego.AutoRouter(&controllers.ScreenController{})
	beego.AutoRouter(&controllers.UserscreenController{})
	beego.AutoRouter(&controllers.InvitationCodeController{})





}
