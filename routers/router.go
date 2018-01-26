package routers

import (
	"admin.dh.cn/controllers"
	. "admin.dh.cn/filter"
	"github.com/astaxie/beego"
)

func init() {

	//登录
	beego.Router("/login", &controllers.LoginController{}, "get:Get")
	beego.Router("/login", &controllers.LoginController{}, "post:Login")
	beego.Router("/login/quit/", &controllers.LoginController{}, "get:Quit")

	beego.InsertFilter("/index", beego.BeforeRouter, BaseFilter)
	beego.InsertFilter("/user/*", beego.BeforeRouter, BaseFilter)
	beego.InsertFilter("/corp/*", beego.BeforeRouter, BaseFilter)
	beego.InsertFilter("/sourcetype/*", beego.BeforeRouter, BaseFilter)
	beego.InsertFilter("/source/*", beego.BeforeRouter, BaseFilter)
	beego.InsertFilter("/screen/*", beego.BeforeRouter, BaseFilter)
	beego.InsertFilter("/userscreen/*", beego.BeforeRouter, BaseFilter)
	beego.InsertFilter("/invitationcode/*", beego.BeforeRouter, BaseFilter)

	//首页
	beego.Router("/index", &controllers.IndexController{}, "get:Get")

	//用户
	beego.Router("/user", &controllers.UserController{}, "get:List")
	beego.Router("/user/getuserdata", &controllers.UserController{}, "get:GetUserData")
	beego.Router("/user/create", &controllers.UserController{}, "get:Create")
	beego.Router("/user/edit", &controllers.UserController{}, "get:Edit")
	beego.Router("/user/add", &controllers.UserController{}, "post:Add")
	beego.Router("/user/update", &controllers.UserController{}, "post:Update")
	beego.Router("/user/remove", &controllers.UserController{}, "get:Remove")
	beego.Router("/user/updateStatusAva", &controllers.UserController{}, "get:UpdateStatusAva")
	beego.Router("/user/getCorp", &controllers.UserController{}, "get:GetCorp")
	beego.Router("/user/delectAndAddCorp", &controllers.UserController{}, "get:DelectAndAddCorp")
	beego.Router("/user/listremove", &controllers.UserController{}, "post:Listremove")
	beego.Router("/user/listChangeType", &controllers.UserController{}, "post:ListChangeType")
	beego.Router("/user/changeType", &controllers.UserController{}, "get:ChangeType")
	beego.Router("/user/getUserScreen", &controllers.UserController{}, "get:GetUserScreen")
	beego.Router("/user/delectUserScreen", &controllers.UserController{}, "get:DelectUserScreen")

	//自动注入路由
	beego.AutoRouter(&controllers.CorpController{})
	beego.AutoRouter(&controllers.SourcetypeController{})
	beego.AutoRouter(&controllers.ScreenController{})
	beego.AutoRouter(&controllers.UserscreenController{})
	beego.AutoRouter(&controllers.InvitationCodeController{})
	beego.AutoRouter(&controllers.SourceController{})

}
