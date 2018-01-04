package controllers

import (
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) init() {
	c.Layout = "common/layout.html"
    c.LayoutSections = make(map[string]string)
	c.LayoutSections["HtmlHead"] = "common/header.html"
	c.LayoutSections["HtmlFooter"] = "common/footer.html"
	c.Data["Menu"]=Menu
}

func (c *MainController) Get() {
	c.init()
    c.TplName = "index/index.html"
}
