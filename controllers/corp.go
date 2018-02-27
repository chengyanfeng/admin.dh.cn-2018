package controllers

import (
	"fmt"
	"strconv"

	"common.dh.cn/models"
	"common.dh.cn/utils"
	"github.com/astaxie/beego/orm"
)

type CorpController struct {
	AdminController
}

func (c *CorpController) List() {
	utils.S("name", "chenyanfeng")
	var mpurl = "/admin/corp/list?"
	c.init(1)
	c.TplName = "corp/index.html"
	var total, total_page int64
	var list []*models.DhCorp
	page, _ := c.GetInt64("page", 1)
	page_size, _ := c.GetInt64("page_size", 10)
	search := c.GetString("search")
	status := c.GetString("status")
	filters := map[string]interface{}{}
	if len(search) > 0 {
		cond := orm.NewCondition()
		if len(search) > 0 {
			c.Data["search"] = search
			mpurl = mpurl + "&search=" + search
			condor := cond.Or("name__icontains", search).
				Or("email__icontains", search).Or("mobile__icontains", search)
			if len(status) > 0 {
				c.Data["status"] = status
				int, _ := strconv.Atoi(status)
				mpurl = mpurl + "&status=" + status
				condor = cond.AndCond(condor).And("status", int)
			} else {
				c.Data["status"] = "nil"
			}

			number, _ := new(models.DhCorp).Query().Offset((page - 1) * page_size).Limit(page_size).SetCond(condor).OrderBy("-create_time").All(&list)
			total, _ = new(models.DhCorp).Query().SetCond(condor).Count()
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

		total, total_page, list = new(models.DhCorp).OrderPager(page, page_size, filters, "-create_time")
	}
	data := []utils.P{}
	if len(list) > 0 {
		for _, info := range list {
			countfilter := utils.P{}
			_crop := utils.P{}
			countfilter["corp_id"] = info.ObjectId
			CropCount := new(models.DhUserCorp).Count(countfilter)
			_crop["ObjectId"] = info.ObjectId
			_crop["CropName"] = info.Name
			_crop["CropEmail"] = info.Email
			_crop["CropMobile"] = info.Mobile
			_crop["CropCount"] = CropCount
			_crop["CreateTime"] = info.CreateTime.Format("2006-01-02 15:04:05")
			_crop["CropStatus"] = info.Status
			data = append(data, _crop)
		}
	}
	c.Data["List"] = data
	c.Data["Pagination"] = PagerHtml(int(total), int(total_page), int(page_size), int(page), mpurl)
}
func (c *CorpController) Create() {
	c.TplName = "corp/create.html"
}
func (c *CorpController) Update() {
	c.Require("id")
	id := c.GetString("id")
	corp := new(models.DhCorp).Find(id)
	if corp == nil {
		c.EchoJsonErr("用户不存在")
		c.StopRun()
	}
	if c.GetString("name") != "" {
		corp.Name = c.GetString("name")
	}
	if c.GetString("email") != "" {
		corp.Email = c.GetString("email")
	}
	if c.GetString("mobile") != "" {
		corp.Mobile = c.GetString("mobile")
	}
	if c.GetString("status") != "" {
		int, err := strconv.Atoi(c.GetString("status"))
		if err == nil {
			corp.Status = int
		}

	}
	result := corp.Save()
	if !result {
		c.EchoJsonErr("更新失败")
	} else {
		c.EchoJsonOk()
	}
}
func (c *CorpController) Edit() {
	c.Require("id")
	id := c.GetString("id")
	corp := new(models.DhCorp).Find(id)
	if corp == nil {
		c.EchoJsonErr("用户不存在")
		c.StopRun()
	}

	c.Data["object"] = &corp
	c.TplName = "corp/edit.html"
}
func (c *CorpController) Add() {
	Corp := new(models.DhCorp)
	Corp.Name = c.GetString("name")
	Corp.Email = c.GetString("email")
	Corp.Mobile = c.GetString("mobile")

	if c.GetString("status") == "1" {
		Corp.Status = 1
	} else {
		Corp.Status = 0
	}

	result := Corp.Save()
	if !result {
		c.EchoJsonErr("注册失败")
	} else {
		c.EchoJsonOk()
	}
}

func (c *CorpController) Remove() {
	c.Require("id")
	id := c.GetString("id")
	corp := new(models.DhCorp).Find(id)
	if corp == nil {
		c.EchoJsonErr("团队不存在")
		c.StopRun()
	}
	result := corp.SoftDelete(id)
	if !result {
		c.EchoJsonErr("删除失败")
	} else {
		c.EchoJsonOk()
	}
}
func (c *CorpController) GetUserCorp() {
	c.Require("id")
	id := c.GetString("id")
	filtersUser := map[string]interface{}{}
	filtersAllUser := map[string]interface{}{}
	filtersAllUser["status__gte"] = 0
	corpName := c.GetString("corpName")
	if corpName != "" && corpName != "undefined" {
		filtersAllUser["name__contains"] = corpName
	}
	dhcorp := new(models.DhCorp).Find(id)
	if dhcorp == nil {
		c.EchoJsonErr("团队不存在")
		c.StopRun()
	}

	filtersUser["corp_id"] = id
	DhUserCorp := new(models.DhUserCorp).OrderList(filtersUser, "-create_time")
	CropUserData := []utils.P{}
	if len(DhUserCorp) > 0 {
		for _, info := range DhUserCorp {
			users := utils.P{}
			user := new(models.DhUser).Find(info.UserId)
			if user == nil {

				c.EchoJsonErr("用户不存在")
				c.StopRun()
			}
			users["objectId"] = info.ObjectId
			users["userid"] = user.ObjectId
			users["name"] = user.Name
			users["email"] = user.Email
			users["role"] = info.Role
			users["corpid"] = info.CorpId
			CropUserData = append(CropUserData, users)
		}
	}

	DhUsers := new(models.DhUser).OrderList(filtersAllUser, "-create_time")
	AllUserData := []utils.P{}

	if len(DhUsers) > 0 {
		for _, info := range DhUsers {
			user := utils.P{}
			user["name"] = info.Name
			user["email"] = info.Email
			user["userid"] = info.ObjectId
			AllUserData = append(AllUserData, user)
		}

	}

	c.Data["CorpId"] = id
	c.Data["corpName"] = dhcorp.Name
	c.Data["CorpUserData"] = CropUserData
	c.Data["AllUserData"] = AllUserData
	c.TplName = "corp/manageUser.html"

}
func (c *CorpController) RemoveAndUser() {
	c.Require("id")
	id := c.GetString("id")
	user_id := c.GetString("user_id")
	removed := c.GetString("removed")
	userCorp := make(map[string]interface{})
	userCorp["user_id"] = user_id
	userCorp["corp_id"] = id
	corp := new(models.DhCorp).Find(id)
	if corp == nil {
		c.EchoJsonErr("团队不存在")
		c.StopRun()
	}
	if removed == "0" {
		dhusercorp := new(models.DhUserCorp).Find(userCorp)
		if dhusercorp == nil {

		} else {
			DhUserCorpCount := map[string]interface{}{}
			DhUserCorpCount["corp_id"] = id
			DhUserCorpCount["role"] = 3
			adminCount := new(models.DhUserCorp).Count(DhUserCorpCount)
			if adminCount < 2 {
				c.EchoJsonErr("管理员唯一不可删除")
				c.StopRun()
			}
			result := dhusercorp.Delete(userCorp)
			if result != true {
				c.EchoJsonErr("删除失败")
			}
			c.EchoJsonOk()
		}
	} else if removed == "1" {
		usercorpfilter := map[string]interface{}{}
		usercorp := new(models.DhUserCorp)
		usercorp.UserId = user_id
		usercorp.CorpId = id
		usercorp.Role = "1"
		usercorpfilter["userid"] = user_id
		usercorpfilter["corpid"] = id
		usercorpflag := new(models.DhUserCorp).Find(usercorpfilter)
		if usercorpflag != nil {
			c.EchoJsonErr("用户已经存在团队中")
			c.StopRun()
		}
		usercorp.Save()
		c.EchoJsonOk()
	} else {
		c.EchoJsonOk()
	}
}
func (c *CorpController) ChangeUserRole() {
	c.Require("id")
	id := c.GetString("id")
	role := c.GetString("role")
	corp := new(models.DhUserCorp).Find(id)
	if corp == nil {

		c.EchoJsonErr("团队不存在")
		c.StopRun()
	}
	userCorpfilterrole := map[string]interface{}{}
	userCorpfilterrole["role"] = "3"
	userCorpfilterrole["object_id"] = id
	if role == "0" {

		number := new(models.DhUserCorp).Count(userCorpfilterrole)
		if number < 2 {
			c.EchoJsonOk("管理员唯一不可改变")
			c.StopRun()

		}
	}

	corp.Role = role
	result := corp.Save()
	if !result {

		c.EchoJsonErr("修改失败")
	}
	c.EchoJsonOk()

}
func (c *CorpController) ListRemove() {
	c.Require("datas")
	datas := c.GetString("datas")
	plist := *utils.JsonDecodeArrays([]byte(datas))
	argerr := make([]string, 1)
	for _, v := range plist {
		dhcorp := new(models.DhCorp).Find(v["object_id"].(string))
		if dhcorp == nil {
			argerr = append(argerr, v["object_id"].(string))
		} else {
			switch dhcorp.Status {
			case -1, 0:
				dhcorp.Status = 1
			case 1:
				dhcorp.Status = -1
			}
			dhcorp.Save()
		}

	}
	if len(argerr[0]) > 0 {
		c.EchoJsonErr("部分更新失败")
	}
	c.EchoJsonOk()

}
