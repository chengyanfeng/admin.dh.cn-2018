package controllers

import (
	"fmt"
	"strconv"

	"common.dh.cn/models"
	"common.dh.cn/utils"
	"github.com/astaxie/beego/orm"
)

type InvitationCodeController struct {
	AdminController
}

func (c *InvitationCodeController) List() {
	defer func() {
		if err := recover(); err != nil {
			c.EchoJsonErr("出现异常")
		}

	}()

	var mpurl = "/admin/invitation_code/list?"
	c.init(4)
	var total, total_page int64
	var list []*models.DhIcode
	c.TplName = "invitation_code/index.html"
	page, _ := c.GetInt64("page", 1)
	page_size, _ := c.GetInt64("page_size", 10)
	filters := map[string]interface{}{}
	filters["status__gte"] = 0
	search := c.GetString("search")
	status := c.GetString("status")

	if len(search) > 0 {
		cond := orm.NewCondition()
		if len(search) > 0 {
			c.Data["search"] = search
			mpurl = mpurl + "&search=" + search
			condor := cond.Or("name__icontains", search)

			if len(status) > 0 {
				c.Data["status"] = status
				int, _ := strconv.Atoi(status)
				mpurl = mpurl + "&status=" + status
				condor = cond.AndCond(condor).And("status", int)
			} else {
				c.Data["status"] = "nil"
			}

			number, _ := new(models.DhIcode).Query().Offset((page - 1) * page_size).Limit(page_size).SetCond(condor).OrderBy("-create_time").All(&list)
			total, _ = new(models.DhIcode).Query().SetCond(condor).Count()
			if total%page_size != 0 {
				total_page = total/page_size + 1
			} else {
				total_page = total / page_size
			}
			fmt.Println(number)
		}

	} else {
		if len(status) > 0 {
			c.Data["status"] = status
			int, _ := strconv.Atoi(status)
			filters["status"] = int
			mpurl = mpurl + "&status=" + status
		} else {
			c.Data["status"] = "nil"
		}

		total, total_page, list = new(models.DhIcode).OrderPager(page, page_size, filters, "-create_time")
	}
	data := []utils.P{}
	if len(list) > 0 {
		for _, info := range list {
			filters := utils.P{}
			filters["icode"] = info.Code
			users := new(models.DhUser).List(filters)
			if len(users) == 0 {
				CodeList := utils.P{}
				CodeList["ObjectId"] = info.ObjectId
				CodeList["Code"] = info.Code
				CodeList["Name"] = ""
				CodeList["Email"] = ""
				CodeList["Status"] = info.Status
				CodeList["CreateTime"] = info.CreateTime.Format("2006-01-02 15:04:05")
				CodeList["UseTime"] = ""
				data = append(data, CodeList)
			} else {
				for _, v := range users {
					CodeList := utils.P{}
					CodeList["ObjectId"] = info.ObjectId
					CodeList["Code"] = info.Code
					CodeList["Name"] = v.Name
					CodeList["Email"] = v.Email
					CodeList["Status"] = info.Status
					CodeList["CreateTime"] = info.CreateTime.Format("2006-01-02 15:04:05")
					CodeList["UseTime"] = v.UpdateTime.Format("2006-01-02 15:04:05")
					data = append(data, CodeList)
				}
			}
		}
	}

	c.Data["List"] = data
	c.Data["Pagination"] = PagerHtml(int(total), int(total_page), int(page_size), int(page), mpurl)
}

func (c *InvitationCodeController) ListRemove() {
	c.Require("datas")
	datas := c.GetString("datas")
	plist := *utils.JsonDecodeArrays([]byte(datas))
	argerr := make([]string, 1)
	for _, v := range plist {
		dhIcode := new(models.DhIcode).Find(v["object_id"].(string))
		if dhIcode == nil {
			argerr = append(argerr, v["object_id"].(string))
		} else {
			dhIcode.Status = -1

			dhIcode.Save()
		}
	}
	if len(argerr[0]) > 0 {
		c.EchoJsonErr("部分更新失败")
	}
	c.EchoJsonOk()

}
func (c *InvitationCodeController) Add() {

	amont := c.GetString("id")
	argerr := make([]string, 1)
	if len(amont) > 0 && amont != "0" {

		j, _ := strconv.Atoi(amont)
		for i := 0; i < j; i++ {
			dhicode := models.DhIcode{}
			code := utils.GetRandomString(5)
			dhicode.Code = code
			result := dhicode.Save()
			if !result {
				argerr = append(argerr, "1")
			}
		}

	}
	if len(argerr[0]) > 0 {
		c.EchoJsonErr("部分创建失败")
	}
	c.EchoJsonOk()

}
func (c *InvitationCodeController) Remove() {
	c.Require("id")
	id := c.GetString("id")
	DhIcode := new(models.DhIcode).Find(id)
	if DhIcode == nil {
		c.EchoJsonErr("大屏模版不存在")
		c.StopRun()
	}
	result := DhIcode.SoftDelete(id)
	if !result {
		c.EchoJsonErr("删除失败")
	} else {
		c.EchoJsonOk()
	}
}
