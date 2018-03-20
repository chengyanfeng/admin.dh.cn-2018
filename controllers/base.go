package controllers

import (
	"strconv"
	"strings"
	"common.dh.cn/models"
	"common.dh.cn/controllers"
	"common.dh.cn/utils"


)

var accountSub = []utils.P{
	utils.P{
		"On":   0,
		"Path": "/admin/user/list",
		"Name": "用户管理",
	},
	utils.P{
		"On":   0,
		"Path": "/admin/corp/list",
		"Name": "团队管理",
	},
}
var datasourceSub = []utils.P{
	utils.P{
		"On":   0,
		"Path": "/admin/datasource_type/list",
		"Name": "分类管理",
	},
	utils.P{
		"On":   0,
		"Path": "/admin/datasource_pub/list",
		"Name": "数据源管理",
	},
}
var screenSub = []utils.P{
	utils.P{
		"On":   0,
		"Path": "/admin/screen_template/list",
		"Name": "大屏模版管理",
	},
	utils.P{
		"On":   0,
		"Path": "/admin/screen/list",
		"Name": "用户大屏管理",
	},
}
var sourceshare = []utils.P{
	utils.P{
		"On": 0,
		"Path": "/admin/sourceshare/list",
		"Name": "工作表共享管理",
	},
}
var invicodeSub = []utils.P{
	utils.P{
		"On":   0,
		"Path": "/admin/invitation_code/list",
		"Name": "邀请码管理",
	},
}
var Menu = []utils.P{
	utils.P{
		"On":   0,
		"Path": "/admin/",
		"Name": "首页",

	},
	utils.P{
		"On":   0,
		"Path": "/",
		"Name": "账号管理",
		"Sub":  accountSub,
	},
	utils.P{
		"On":   0,
		"Path": "/",
		"Name": "公共数据源管理",
		"Sub":  datasourceSub,
	},
	utils.P{
		"On":   0,
		"Path": "/",
		"Name": "大屏管理",
		"Sub":  screenSub,
	},
	utils.P{
		"On":   0,
		"Path": "/",
		"Name": "邀请码管理",
		"Sub":  invicodeSub,
	},
	utils.P{
		"On": 0,
		"Path": "/",
		"Name": "数据权限管理",
		"Sub": sourceshare,
	},
}

type AdminController struct {
	controllers.BaseController
}
func (c *AdminController)getList()([]utils.P){
	filter:=utils.P{}
	FListMap:=[]utils.P{}
	list:=new(models.DhList).List(filter)
	for _,fv :=range list{
		if fv.FObjectId==""{
			childmapList:=[]utils.P{}
			listmap:=utils.P{}
			listmap["On"]=0
			listmap["Path"]=fv.Path
			listmap["Name"]=fv.Name
			for _,child:=range list{

				if child.FObjectId==fv.ObjectId{
					childmap:=utils.P{}
					childmap["On"]=0
					childmap["Path"]=child.Path
					childmap["Name"]=child.Name
					childmapList=append(childmapList,childmap)
				}
			}
			if len(childmapList)>0{
				listmap["Sub"]=childmapList
			}
			FListMap=append(FListMap,listmap)
		}
	}
	return FListMap
}



func (c *AdminController) init(i int) {
	c.Layout = "common/layout.html"
	c.LayoutSections = make(map[string]string)
	c.LayoutSections["HtmlHead"] = "common/header.html"
	c.LayoutSections["HtmlFooter"] = "common/footer.html"
	/*Menu:=c.getList()*/
	Menu:=Menu
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
	Authname := c.GetSession("Authname")
	c.Data["Authname"] = Authname
	c.Data["auth_id"]=c.GetSession("auth")
	c.Data["Object_id"]=c.GetSession("Object_id")
	c.Data["Menu"] = Menu
}

func PagerHtml(num int, totalpage int, perpage int, curpage int, mpurl string) string {
	if num == 0 {
		return ""
	}
	var max_page, begin, end int
	html := ""
	if num > perpage {
		html = "<ul>"
		html += "<li><a href=\"" + mpurl + "\">首页</a></li>"
		if curpage-1 > 0 {
			if mpurl == "javascript:;" {
				html += "<li><a href=\"" + mpurl + "\">&#8249;</a></li>"
			} else if strings.Contains(mpurl, "?") {
				html += "<li><a href=\"" + mpurl + "&page=" + string(curpage-1) + "\">&#8249;</a></li>"
			} else {
				html += "<li><a href=\"" + mpurl + "?page=" + string(curpage-1) + "\">&#8249;</a></li>"
			}
		} else {
			html += "<li><a href=\"" + mpurl + "\">&#8249;</a></li>"
		}
		//最大页数
		if totalpage <= 9 {
			max_page = totalpage
		} else {
			max_page = 9
		}
		rank := 4
		if curpage >= max_page {
			if (curpage - rank) > 0 {
				begin = curpage - rank
			} else {
				begin = 1
			}
		} else {
			begin = 1
		}
		if curpage >= max_page {
			if (curpage + rank) <= totalpage {
				end = curpage + rank
			} else {
				end = totalpage
			}
		} else {
			end = max_page
		}
		for i := begin; i <= end; i++ {
			var link string
			if mpurl == "javascript:;" {
				link = "javascript:;"
			} else if strings.Contains(mpurl, "?") {
				link = mpurl + "&page=" + strconv.Itoa(i)
			} else {
				link = mpurl + "?page=" + strconv.Itoa(i)
			}
			class := ""
			if i == curpage {
				link = "javascript:;"
				class = "active"
			}
			html += "<li class=\"" + class + "\"><a href=\"" + link + "\">" + strconv.Itoa(i) + "</a></li>"
		}
		if (curpage + 1) < totalpage {
			if mpurl == "javascript:;" {
				html += "<li><a href=\"" + mpurl + "\">&#8250;</a></li>"
			} else if strings.Contains(mpurl, "?") {
				html += "<li><a href=\"" + mpurl + "&page=" + strconv.Itoa(curpage+1) + "\">&#8250;</a></li>"
			} else {
				html += "<li><a href=\"" + mpurl + "?page=" + strconv.Itoa(curpage+1) + "\">&#8250;</a></li>"
			}
		} else {
			if mpurl == "javascript:;" {
				html += "<li><a href=\"" + mpurl + "\">&#8250;</a></li>"
			} else if strings.Contains(mpurl, "?") {
				html += "<li><a href=\"" + mpurl + "&page=" + strconv.Itoa(totalpage) + "\">&#8250;</a></li>"
			} else {
				html += "<li><a href=\"" + mpurl + "?page=" + strconv.Itoa(totalpage) + "\">&#8250;</a></li>"
			}
		}
		if mpurl == "javascript:;" {
			html += "<li><a href=\"" + mpurl + "\">尾页</a></li>"
		} else if strings.Contains(mpurl, "?") {
			html += "<li><a href=\"" + mpurl + "&page=" + strconv.Itoa(totalpage) + "\">尾页</a></li>"
		} else {
			html += "<li><a href=\"" + mpurl + "?page=" + strconv.Itoa(totalpage) + "\">尾页</a></li>"
		}
		html += "<li style=\"color:#516372;\">&nbsp;&nbsp;<span style=\"border:0\">总记录数：<b>" + strconv.Itoa(num) + "</b>&nbsp;&nbsp;&nbsp;&nbsp;页数：<b>" + strconv.Itoa(totalpage) + "</b></span>&nbsp;&nbsp;&nbsp;&nbsp;</li>"
		html += "<ul>"
	}
	return html
}
