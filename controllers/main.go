package controllers

import (
	"fmt"

	"github.com/astaxie/beego"
)

type IndexController struct {
	beego.Controller
}

func (c *IndexController) init(i int) {
	c.Layout = "common/layout.html"
	c.LayoutSections = make(map[string]string)
	c.LayoutSections["HtmlHead"] = "common/header.html"
	c.LayoutSections["HtmlFooter"] = "common/footer.html"
	for k, v := range Menu {
		if k != i {
			v["On"] = 0
		} else {
			Menu[i]["On"] = 1
		}
	}
	Authname := c.GetSession("Authname")
	c.Data["Authname"] = Authname
	c.Data["Menu"] = Menu
}
func (c *IndexController) Get() {
	fmt.Print("进去main了" +
		"")
	c.init(0)
	c.TplName = "index/index.html"
}
