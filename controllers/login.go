package controllers

import (
	"fmt"

	"common.dh.cn/controllers"
	"common.dh.cn/def"
	"common.dh.cn/models"
	"common.dh.cn/utils"
)

type LoginController struct {
	controllers.BaseController
}

func (c *LoginController) Get() {
	c.TplName = "index/login.html"
}

func (c *LoginController) Login() {
	username := c.GetString("name")
	password := c.GetString("password")
	userfilter := map[string]interface{}{}
	userfilter["name"] = username
	userfilter["password"] = utils.Md5(password, def.Md5Salt)
	fmt.Println(userfilter["password"], "---------------------password----------")
	user := new(models.DhUser).Find(userfilter)
	if user == nil {
		c.EchoJsonErr("账号或密码错误")
		c.StopRun()
	} else if user.IsAdmin != 1 {
		c.EchoJsonErr("您不是管理员")
		c.StopRun()
	}
	c.SetSession("auth", user.Auth)
	c.SetSession("Authname", user.Name)
	c.EchoJsonOk()
}
func (c *LoginController) Quit() {
	c.DelSession("auth")
	c.DelSession("Authname")
	c.Ctx.Redirect(302, "/admin/login")
}
