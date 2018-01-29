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

func (c *LoginController) init(i int) {
	c.Layout = "common/layout.html"
	c.LayoutSections = make(map[string]string)
	c.LayoutSections["HtmlHead"] = "common/header.html"
	c.LayoutSections["HtmlFooter"] = "common/footer.html"
	for k, v := range Menu {
		if k != i {
			v["On"] = 0
		} else {
			Menu[i]["On"] = 1
			if Menu[i]["Sub"] != nil {
				a := Menu[i]["Sub"].(interface{})
				b := a.([]utils.P)
				for _, v := range b {
					v["On"] = 1
				}
			}
		}
	}
	Authname := c.Ctx.GetCookie("Authname")
	c.Data["Authname"] = Authname
	c.Data["Menu"] = Menu
}

func (c *LoginController) Get() {
	c.TplName = "index/login.html"
}

func (c *LoginController) Login() {
	c.init(0)
	username := c.GetString("name")
	password := c.GetString("password")
	userfilter := map[string]interface{}{}
	userfilter["name"] = username
	userfilter["password"] = utils.Md5(password, def.Md5Salt)
	fmt.Println(userfilter["password"],"---------------------password----------")
	user := new(models.DhUser).Find(userfilter)
	fmt.Print(user)
	fmt.Print(user,"---------------------------user--------------------")
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
	c.Ctx.Redirect(302, "/login")
}
