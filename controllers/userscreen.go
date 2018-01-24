package controllers

import (

	"common.dh.cn/utils"
	"common.dh.cn/models"
	"common.dh.cn/controllers"
	"github.com/astaxie/beego/orm"
	"fmt"
	"strconv"

)

type UserscreenController struct {
	controllers.BaseController
}

func (c *UserscreenController) init(i int) {
	c.Layout = "common/layout.html"
	c.LayoutSections = make(map[string]string)
	c.LayoutSections["HtmlHead"] = "common/header.html"
	c.LayoutSections["HtmlFooter"] = "common/footer.html"
	for   k,v :=range Menu{
		if k!=i{
			v["On"]=0
		}else {
			Menu[i]["On"]=1
			if Menu[i]["Sub"]!=nil{
				a:= Menu[i]["Sub"].(interface{})
				b:= a.([]utils.P)
				for _,v:=range b {
					v["On"]=1
				}
			}
		}
	}
	Authname,_:=c.GetSecureCookie("2rdsfada3@#$%^&*","Authname")
	c.Data["Authname"]=Authname
	c.Data["Menu"]=Menu
}


func (c *UserscreenController) List() {
	defer func(){
		if err:=recover();err!=nil{
			c.EchoJsonErr("出现异常")
		}

	}()

	var mpurl ="/userscreen/list?"
	c.init(3)
	var total,total_page int64
	var list []*models.DxScreen
	c.TplName = "userscreen/index.html"
	page,_ := c.GetInt64("page",1)
	page_size,_ := c.GetInt64("page_size",10)
	filters := map[string]interface{}{}
	filters["status__gte"]=0
	search := c.GetString("search")
	status:= c.GetString("status")

	if len(search)>0{
		cond := orm.NewCondition()
		if len(search)>0{
			c.Data["search"] = search
			mpurl=mpurl+"&search="+search
			condor := cond.Or("name__icontains", search)

			if len(status)>0{
				c.Data["status"] = status
				int,_:=strconv.Atoi(status)
				mpurl=mpurl+"&status="+status
				condor=cond.AndCond(condor).And("status",int)
			}else{
				c.Data["status"] = "nil"
			}

			number,_:=new(models.DxScreen).Query().Offset((page-1)*page_size).Limit(page_size).SetCond(condor).OrderBy("-create_time").All(&list)
			total,_=new(models.DxScreen).Query().SetCond(condor).Count()
			if total%page_size!=0{
				total_page=total/page_size+1
			}else {
				total_page=total/page_size
			}
			fmt.Println(number)
		}

	}else {
		if len(status)>0{
			c.Data["status"] = status
			int,_:=strconv.Atoi(status)
			filters["status"]=int
			mpurl=mpurl+"&status="+status
		}else{
			c.Data["status"] = "nil"
		}

		total,total_page,list = new(models.DxScreen).OrderPager(page, page_size, filters,"-create_time")
	}
	data := []utils.P{}
	if len(list) > 0 {
		for _, info := range list {
			filtersdxScreen := map[string]interface{}{}
			filtersdxScreen["relate_id"]=info.ObjectId
			filtersdxScreen["relate_type"]="dx_screen"
			dxScreen:= new(models.DhRelation).Find(filtersdxScreen)
			user:=new(models.DhUser).Find(dxScreen.UserId)
			if user==nil{
				c.EchoJsonErr("用户查询失败")
			}
			Screen := utils.P{}
			Screen["ObjectId"] = info.ObjectId
			Screen["Name"] = info.Name
			Screen["CreateUserId"] = user.ObjectId
			Screen["CreateUser"] = user.Name
			Screen["Status"] = info.Status
			 p:=*utils.JsonDecode([]byte(info.Config))
			Screen["Config"] =utils.ToString(p["width"])+utils.ToString("*")+utils.ToString(p["height"])
			Screen["CreateTime"] = info.CreateTime.Format("2006-01-02 15:04:05")
			Screen["UpdateTime"] = info.UpdateTime.Format("2006-01-02 15:04:05")
			data = append(data, Screen)
		}
	}
	c.Data["List"] = data
	c.Data["Pagination"] = PagerHtml(int(total), int(total_page), int(page_size), int(page),mpurl)
}
func (c *UserscreenController) Update() {
	c.Require("id")
	id := c.GetString("id")
	dxScreenTemplate := new(models.DxScreenTemplate).Find(id)
	if dxScreenTemplate == nil {
		c.EchoJsonErr("用户不存在")
		c.StopRun()
	}
	if c.GetString("name")!=""{
		dxScreenTemplate.Name = c.GetString("name")
	}
	if c.GetString("status")!=""{
		int,err:=strconv.Atoi(c.GetString("status"))
		if err==nil{
			dxScreenTemplate.Status=int
		}

	}
	result := dxScreenTemplate.Save()
	if !result {
		c.EchoJsonErr("更新失败")
	} else {
		c.EchoJsonOk()
	}
}
func (c *UserscreenController) Listremove() {
	c.Require("datas")
	datas := c.GetString("datas")
	plist:=*utils.JsonDecodeArrays([]byte(datas))
	argerr :=make([]string,1)
	for _,v :=range plist{
		DxScreen := new(models.DxScreen).Find(v["object_id"].(string))
		if DxScreen == nil {
			argerr=append(argerr,v["object_id"].(string))
		}else {
			 DxScreen.Status =-1

			DxScreen.Save()
		}
	}
	if    len(argerr[0]) > 0 {
		c.EchoJsonErr("部分更新失败")
	}
	c.EchoJsonOk()


}

func (c *UserscreenController) Remove() {
	c.Require("id")
	id := c.GetString("id")
	DxScreen := new(models.DxScreen).Find(id)
	if DxScreen == nil {
		c.EchoJsonErr("大屏膜版不存在")
		c.StopRun()
	}
	result := DxScreen.SoftDelete(id)
	if !result {
		c.EchoJsonErr("删除失败")
	} else {
		c.EchoJsonOk()
	}
}
func (c *UserscreenController) Add() {
	DxScreenTemplate := new(models.DxScreenTemplate)
	DxScreenTemplate.Name = c.GetString("name")
	DxScreenTemplate.Status = 0
	result :=DxScreenTemplate.Save()
	if !result {
		c.EchoJsonErr("创建失败")
	} else {
		c.EchoJsonOk()
	}
}
func (c *UserscreenController) Create(){

	c.TplName="userscreen/create.html"
}
