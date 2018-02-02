package controllers

import (
	"fmt"

	"common.dh.cn/controllers"
	"common.dh.cn/def"
	"common.dh.cn/utils"
)

type IndexController struct {
	controllers.BaseController
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
	c.init(0)
	c.TplName = "index/index.html"
}

func (c *IndexController) FileUploader() {
	data := c.GetString("data")
	bin := utils.Base64Decode(data)
	name := c.GetString("name")
	index, _ := c.GetInt64("index")
	size, _ := c.GetInt64("size")
	total, _ := c.GetInt64("total")
	ext := utils.ToString(utils.Pathinfo(name)["extension"])
	if size > def.MAX_UPLOAD_SIZE {
		c.EchoJsonErr("文件尺寸不得大于", def.MAX_UPLOAD_SIZE)
	}
	md5 := utils.Md5(name)
	file_path := fmt.Sprintf("upload/%v.%v", md5, ext)
	utils.WriteFile(file_path, bin)
	real_url := c.Hostname() + "/" + file_path
	result := utils.P{
		"result":     1,
		"index":      index,
		"total":      total,
		"name":       "name",
		"real_url":   real_url,
		"file_index": c.GetString("file_index"),
	}
	c.EchoJson(result)
}
