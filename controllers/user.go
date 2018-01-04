package controllers

import (
	"math/rand"
	"common.dh.cn/def"
	"common.dh.cn/utils"
	"common.dh.cn/models"
	"common.dh.cn/controllers"
)

type UserController struct {
	controllers.BaseController
}

func (c *UserController) init() {
	c.Layout = "common/layout.html"
    c.LayoutSections = make(map[string]string)
	c.LayoutSections["HtmlHead"] = "common/header.html"
	c.LayoutSections["HtmlFooter"] = "common/footer.html"
	c.Data["Menu"]=Menu
}

func (c *UserController) List() {
	c.init()
	c.TplName = "user/index.html"
	page,_ := c.GetInt64("page",1)
	page_size,_ := c.GetInt64("page_size",10)
	filters := map[string]interface{}{}
	filters["status__gte"] = 0
	total,total_page,list := new(models.DhUser).OrderPager(page, page_size, filters,"-create_time")
	data := []utils.P{}
	if len(list) > 0 {
		for _, info := range list {
			_user := utils.P{}
			_user["ObjectId"] = info.ObjectId
			_user["Name"] = info.Name
			_user["Avatar"] = info.Avatar
			_user["Email"] = info.Email
			_user["Mobile"] = info.Mobile
			_user["CreateTime"] = info.CreateTime.Format("2006-01-02 15:04:05")
			data = append(data, _user)
		}
	}
	c.Data["List"] = data
	c.Data["Pagination"] = PagerHtml(int(total), int(total_page), int(page_size), int(page),"/user")
}

func (c *UserController) Create() {
	c.TplName = "user/create.html"
}

func (c *UserController) Edit() {
	c.Require("id")
	id := c.GetString("id")
	user := new(models.DhUser).Find(id)
	if user == nil {
		c.EchoJsonErr("用户不存在")
		c.StopRun()
	}
	c.Data["object"] = &user
	c.TplName = "user/edit.html"	
}

func (c *UserController) Add() {
	user := new(models.DhUser)
	user.Name = c.GetString("name")
	user.Corp = c.GetString("corp")
	user.Email = c.GetString("email")
	user.Mobile = c.GetString("mobile")
	user.Password = utils.Md5(c.GetString("password"), def.Md5Salt)
	user.Auth = utils.Md5(c.GetString("email"), c.GetString("password"), rand.Intn(1000)*rand.Intn(1000))
	if c.GetString("status") == "1" {
		user.Status = 1 
	} else {
		user.Status = 0
	}
	if c.GetString("is_dataI_user") == "1" {
		user.IsDataIUser = 1 
	} else {
		user.IsDataIUser = 0
	}
	if c.GetString("is_dataX_user") == "1" {
		user.IsDataXUser = 1 
	} else {
		user.IsDataXUser = 0
	}
	result := user.Save()
	if !result {
		c.EchoJsonErr("注册失败")
	} else {
		c.EchoJsonOk()
	}
}

func (c *UserController) Update() {
	c.Require("id")
	id := c.GetString("id")
	user := new(models.DhUser).Find(id)
	if user == nil {
		c.EchoJsonErr("用户不存在")
		c.StopRun()
	}
	user.Name = c.GetString("name")
	user.Corp = c.GetString("corp")
	user.Email = c.GetString("email")
	user.Mobile = c.GetString("mobile")
	if c.GetString("status") == "1" {
		user.Status = 1 
	} else {
		user.Status = 0
	}
	if c.GetString("is_dataI_user") == "1" {
		user.IsDataIUser = 1 
	} else {
		user.IsDataIUser = 0
	}
	if c.GetString("is_dataX_user") == "1" {
		user.IsDataXUser = 1 
	} else {
		user.IsDataXUser = 0
	}
	result := user.Save()
	if !result {
		c.EchoJsonErr("注册失败")
	} else {
		c.EchoJsonOk()
	}
}

func (c *UserController) Remove() {
	c.Require("id")
	id := c.GetString("id")
	user := new(models.DhUser).Find(id)
	if user == nil {
		c.EchoJsonErr("用户不存在")
		c.StopRun()
	}
	result := user.SoftDelete(id)
	if !result {
		c.EchoJsonErr("删除失败")
	} else {
		c.EchoJsonOk()
	}
}