package routers

import (
	"admin.dh.cn/controllers"
	. "admin.dh.cn/filter"
	"github.com/astaxie/beego"
)

func init() {
	beego.InsertFilter("admin/", beego.BeforeRouter, BaseFilter)
	beego.InsertFilter("admin/index", beego.BeforeRouter, BaseFilter)
	beego.InsertFilter("admin/user/*", beego.BeforeRouter, BaseFilter)
	beego.InsertFilter("admin/corp/*", beego.BeforeRouter, BaseFilter)
	beego.InsertFilter("admin/datasource_type/*", beego.BeforeRouter, BaseFilter)
	beego.InsertFilter("admin/datasource_pub/*", beego.BeforeRouter, BaseFilter)
	beego.InsertFilter("admin/screen_template/*", beego.BeforeRouter, BaseFilter)
	beego.InsertFilter("admin/screen/*", beego.BeforeRouter, BaseFilter)
	beego.InsertFilter("admin/invitationcode/*", beego.BeforeRouter, BaseFilter)
	//上传文件
	beego.Router("admin/file_upload", &controllers.IndexController{}, "post:FileUploader")
	//登录
	beego.Router("admin/login", &controllers.LoginController{}, "get:Get")
	beego.Router("admin/login", &controllers.LoginController{}, "post:Login")
	beego.Router("admin/logout", &controllers.LoginController{}, "get:Quit")
	//首页
	beego.Router("admin/", &controllers.IndexController{}, "get:Get")
	beego.Router("admin/index", &controllers.IndexController{}, "get:Get")
	//用户
	beego.Router("admin/user/list", &controllers.UserController{}, "get:List")
	beego.Router("admin/user/getuserdata", &controllers.UserController{}, "get:GetUserData")
	beego.Router("admin/user/create", &controllers.UserController{}, "get:Create")
	beego.Router("admin/user/edit", &controllers.UserController{}, "get:Edit")
	beego.Router("admin/user/add", &controllers.UserController{}, "post:Add")
	beego.Router("admin/user/update", &controllers.UserController{}, "post:Update")
	beego.Router("admin/user/remove", &controllers.UserController{}, "get:Remove")
	beego.Router("admin/user/updateStatusAva", &controllers.UserController{}, "get:UpdateStatusAva")
	beego.Router("admin/user/getCorp", &controllers.UserController{}, "get:GetCorp")
	beego.Router("admin/user/delectAndAddCorp", &controllers.UserController{}, "get:DelectAndAddCorp")
	beego.Router("admin/user/listremove", &controllers.UserController{}, "post:Listremove")
	beego.Router("admin/user/listChangeType", &controllers.UserController{}, "post:ListChangeType")
	beego.Router("admin/user/changeType", &controllers.UserController{}, "get:ChangeType")
	beego.Router("admin/user/getUserScreen", &controllers.UserController{}, "get:GetUserScreen")
	beego.Router("admin/user/delectUserScreen", &controllers.UserController{}, "get:DelectUserScreen")
	//团队
	beego.Router("admin/corp/list", &controllers.CorpController{}, "get:List")
	beego.Router("admin/corp/create", &controllers.CorpController{}, "get:Create")
	beego.Router("admin/corp/edit", &controllers.CorpController{}, "get:Edit")
	beego.Router("admin/corp/add", &controllers.CorpController{}, "post:Add")
	beego.Router("admin/corp/update", &controllers.CorpController{}, "post:Update")
	beego.Router("admin/corp/remove", &controllers.CorpController{}, "get:Remove")
	beego.Router("admin/corp/getusercorp", &controllers.CorpController{}, "get:GetUserCorp")
	beego.Router("admin/corp/removeanduser", &controllers.CorpController{}, "get:RemoveAndUser")
	beego.Router("admin/corp/changeuserrole", &controllers.CorpController{}, "get:ChangeUserRole")
	beego.Router("admin/corp/listremove", &controllers.CorpController{}, "get:ListRemove")
	//数据源分类
	beego.Router("admin/datasource_type/list", &controllers.DatasourceTypeController{}, "get:List")
	beego.Router("admin/datasource_type/create", &controllers.DatasourceTypeController{}, "get:Create")
	beego.Router("admin/datasource_type/edit", &controllers.DatasourceTypeController{}, "get:Edit")
	beego.Router("admin/datasource_type/add", &controllers.DatasourceTypeController{}, "post:Add")
	beego.Router("admin/datasource_type/update", &controllers.DatasourceTypeController{}, "post:Update")
	beego.Router("admin/datasource_type/remove", &controllers.DatasourceTypeController{}, "get:Remove")
	beego.Router("admin/datasource_type/listremove", &controllers.DatasourceTypeController{}, "get:ListRemove")
	//公共数据源管理
	beego.Router("admin/datasource_pub/list", &controllers.DatasourcePubController{}, "get:List")
	beego.Router("admin/datasource_pub/create", &controllers.DatasourcePubController{}, "get:Create")
	beego.Router("admin/datasource_pub/edit", &controllers.DatasourcePubController{}, "get:Edit")
	beego.Router("admin/datasource_pub/add", &controllers.DatasourcePubController{}, "post:Add")
	beego.Router("admin/datasource_pub/update", &controllers.DatasourcePubController{}, "post:Update")
	beego.Router("admin/datasource_pub/remove", &controllers.DatasourcePubController{}, "get:Remove")
	beego.Router("admin/datasource_pub/listremove", &controllers.DatasourcePubController{}, "get:ListRemove")
	//大屏模版管理
	beego.Router("admin/screen_template/list", &controllers.ScreenTemplateController{}, "get:List")
	beego.Router("admin/screen_template/create", &controllers.ScreenTemplateController{}, "get:Create")
	beego.Router("admin/screen_template/edit", &controllers.ScreenTemplateController{}, "get:Edit")
	beego.Router("admin/screen_template/add", &controllers.ScreenTemplateController{}, "post:Add")
	beego.Router("admin/screen_template/update", &controllers.ScreenTemplateController{}, "post:Update")
	beego.Router("admin/screen_template/remove", &controllers.ScreenTemplateController{}, "get:Remove")
	beego.Router("admin/screen_template/listremove", &controllers.ScreenTemplateController{}, "get:ListRemove")
	//大屏管理
	beego.Router("admin/screen/list", &controllers.ScreenController{}, "get:List")
	beego.Router("admin/screen/create", &controllers.ScreenController{}, "get:Create")
	beego.Router("admin/screen/edit", &controllers.ScreenController{}, "get:Edit")
	beego.Router("admin/screen/add", &controllers.ScreenController{}, "post:Add")
	beego.Router("admin/screen/update", &controllers.ScreenController{}, "post:Update")
	beego.Router("admin/screen/remove", &controllers.ScreenController{}, "get:Remove")
	beego.Router("admin/screen/listremove", &controllers.ScreenController{}, "get:ListRemove")
	//邀请码管理
	beego.Router("admin/invitationcode/list", &controllers.InvitationCodeController{}, "get:List")
	beego.Router("admin/invitationcode/create", &controllers.InvitationCodeController{}, "get:Create")
	beego.Router("admin/invitationcode/edit", &controllers.InvitationCodeController{}, "get:Edit")
	beego.Router("admin/invitationcode/add", &controllers.InvitationCodeController{}, "post:Add")
	beego.Router("admin/invitationcode/update", &controllers.InvitationCodeController{}, "post:Update")
	beego.Router("admin/invitationcode/remove", &controllers.InvitationCodeController{}, "get:Remove")
	beego.Router("admin/invitationcode/listremove", &controllers.InvitationCodeController{}, "get:ListRemove")
	//自动注入路由
	//beego.AutoRouter(&controllers.CorpController{})
	//beego.AutoRouter(&controllers.SourcetypeController{})
	//beego.AutoRouter(&controllers.ScreenController{})
	//beego.AutoRouter(&controllers.UserscreenController{})
	//beego.AutoRouter(&controllers.InvitationCodeController{})
	//beego.AutoRouter(&controllers.SourceController{})

}
