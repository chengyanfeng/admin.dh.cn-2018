package controllers

import (

	"common.dh.cn/controllers"


)

type LoginController struct {
	controllers.BaseController
}

func (c *LoginController)Get(){

	c.TplName="index/login.html"

}