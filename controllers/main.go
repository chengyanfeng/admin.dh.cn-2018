package controllers

import (
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) init(i int) {
	c.Layout = "common/layout.html"
    c.LayoutSections = make(map[string]string)
	c.LayoutSections["HtmlHead"] = "common/header.html"
	c.LayoutSections["HtmlFooter"] = "common/footer.html"
	for   k,v :=range Menu{
		if k!=i{
			v["On"]=0
		}else {
			Menu[i]["On"]=1
		}
	}
	c.Data["Menu"]=Menu
	Authname,_:=c.GetSecureCookie("2rdsfada3@#$%^&*","Authname")
	c.Data["Authname"]=Authname
}

func (c *MainController) Get() {
	c.init(0)
    c.TplName = "index/index.html"
}
