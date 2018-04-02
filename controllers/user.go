package controllers

import (
	"fmt"
	"math/rand"
	"strconv"

	"common.dh.cn/def"
	"common.dh.cn/models"
	"common.dh.cn/utils"
	"github.com/astaxie/beego/orm"

)

type UserController struct {
	AdminController
}

func (c *UserController) List() {
	var mpurl = "/admin/user/list?"
	c.init(1)
	var total, total_page int64
	var list []*models.DhUser
	c.TplName = "user/index.html"
	page, _ := c.GetInt64("page", 1)
	page_size, _ := c.GetInt64("page_size", 10)
	filters := map[string]interface{}{}
	search := c.GetString("search")
	status := c.GetString("status")
	businesstype := c.GetString("businesstype")

	if len(search) > 0 {
		cond := orm.NewCondition()
		if len(search) > 0 {
			c.Data["search"] = search
			mpurl = mpurl + "&search=" + search
			condor := cond.Or("name__icontains", search).Or("corp__icontains", search).
				Or("email__icontains", search).Or("mobile__icontains", search)
			if len(status) > 0 {
				c.Data["status"] = status
				int, _ := strconv.Atoi(status)
				mpurl = mpurl + "&status=" + status
				condor = cond.AndCond(condor).And("status", int)
			} else {
				c.Data["status"] = "nil"
			}
			if len(businesstype) > 0 {
				c.Data["businesstype"] = businesstype
				mpurl = mpurl + "&businesstype=" + businesstype
				condor = cond.AndCond(condor).And(businesstype, 1)
			} else {

				c.Data["businesstype"] = "nil"
			}
			number, _ := new(models.DhUser).Query().Offset((page - 1) * page_size).Limit(page_size).SetCond(condor).OrderBy("-create_time").All(&list)
			total, _ = new(models.DhUser).Query().SetCond(condor).Count()
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
		if len(businesstype) > 0 {
			c.Data["businesstype"] = businesstype
			filters[businesstype] = 1
			mpurl = mpurl + "&businesstype=" + businesstype
		} else {
			c.Data["businesstype"] = "nil"
		}

		total, total_page, list = new(models.DhUser).OrderPager(page, page_size, filters, "-create_time")
	}
	data := []utils.P{}
	if len(list) > 0 {
		for _, info := range list {
			_user := utils.P{}
			_user["ObjectId"] = info.ObjectId
			_user["Name"] = info.Name
			_user["Avatar"] = info.Avatar
			_user["Email"] = info.Email
			_user["Mobile"] = info.Mobile
			_user["Corp"] = info.Corp
			_user["Status"] = info.Status
			_user["IsDataIUser"] = info.IsDataIUser
			_user["IsDataXUser"] = info.IsDataXUser
			_user["CreateTime"] = info.CreateTime.Format("2006-01-02 15:04:05")
			data = append(data, _user)
		}
	}
	c.Data["List"] = data
	c.Data["Pagination"] = PagerHtml(int(total), int(total_page), int(page_size), int(page), mpurl)
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
func (c *UserController) Listremove() {
	c.Require("datas")
	datas := c.GetString("datas")
	plist := *utils.JsonDecodeArrays([]byte(datas))
	argerr := make([]string, 1)
	for _, v := range plist {
		dhUser := new(models.DhUser).Find(v["object_id"].(string))
		if dhUser == nil {
			argerr = append(argerr, v["object_id"].(string))
		} else {
			switch dhUser.Status {
			case -1, 0:
				dhUser.Status = 1
			case 1:
				dhUser.Status = -1
			}
			dhUser.Save()
		}
	}
	if len(argerr[0]) > 0 {
		c.EchoJsonErr("删除失败")
	}
	c.EchoJsonOk()

}
func (c *UserController) ChangeType() {
	c.TplName = "user/changeType.html"
}
func (c *UserController) UpdateStatusAva() {
	c.Require("id")
	id := c.GetString("id")
	status := c.GetString("status")
	user := new(models.DhUser).Find(id)
	if user == nil {
		c.EchoJsonErr("用户不存在")
		c.StopRun()
	}
	int, err := strconv.Atoi(status)
	if err == nil {
		user.Status = int
	}
	result := user.Save()
	if !result {
		c.EchoJsonErr("更新失败")
	} else {
		url := fmt.Sprintf("https://%v", "www.datahunter.cn")
		data := utils.P{"name": user.Name, "url": url,"corp":"北京数猎天下科技有限公司"}
		body := c.GetMailString("./views/templet/corp_invite.tpl", data)
		if body != "" {
			go utils.Mail(user.Email, utils.JoinStr("DataHunter注册验证 ", utils.DateTimeStr()), body)
		}

		c.EchoJsonOk()
	}
}
func (c *UserController) GetCorp() {
	c.Require("id")
	id := c.GetString("id")
	filtersUserCorp := map[string]interface{}{}
	filtersAllCorp := map[string]interface{}{}
	filtersAllCorp["status__gte"] = 0
	corpName := c.GetString("corpName")
	if corpName != "" && corpName != "undefined" {
		filtersAllCorp["name__contains"] = corpName
	}
	user := new(models.DhUser).Find(id)
	if user == nil {
		c.EchoJsonErr("用户不存在")
		c.StopRun()
	}
	filtersUserCorp["user_id"] = id
	userCorp := new(models.DhUserCorp).OrderList(filtersUserCorp, "-create_time")
	UserCorpData := []utils.P{}
	if len(userCorp) > 0 {
		for _, info := range userCorp {
			corp := new(models.DhCorp).Find(info.CorpId)
			if corp != nil {
				userCorp := utils.P{}
				userCorp["ObjectId"] = info.ObjectId
				userCorp["CorpId"] = info.CorpId
				userCorp["Role"] = info.Role
				userCorp["Userid"] = info.UserId
				userCorp["Name"] = corp.Name
				userCorp["Email"] = corp.Email

				UserCorpData = append(UserCorpData, userCorp)
			}
		}
	}
	c.Data["userCorpData"] = UserCorpData
	allCorp := new(models.DhCorp).OrderList(filtersAllCorp, "-create_time")
	allCorpData := []utils.P{}
	if len(allCorp) > 0 {
		for _, info := range allCorp {
			allCorp := utils.P{}
			allCorp["ObjectId"] = info.ObjectId
			allCorp["Name"] = info.Name
			allCorp["Email"] = info.Email
			allCorpData = append(allCorpData, allCorp)
		}
	}
	c.Data["allCorpData"] = allCorpData
	c.Data["name"] = user.Name
	c.Data["userid"] = user.ObjectId
	c.TplName = "user/manageCorp.html"

}
func (c *UserController) DelectAndAddCorp() {
	//title  1    为 删除用户
	//title  2    为添加用户
	//role   1,0    改变用户角色
	c.Require("id", "user_id")

	object_id := c.GetString("id")
	role := c.GetString("role")
	user_id := c.GetString("user_id")
	corp_id := c.GetString("corp_id")
	title := c.GetString("title")
	filtersUserCorp := map[string]interface{}{}
	filtersUserCorp["object_id"] = object_id
	filtersUserCorp["user_id"] = user_id
	if title == "1" {
		UserCorp := new(models.DhUserCorp).Find(filtersUserCorp)
		if UserCorp == nil {
			c.EchoJsonErr("团队不存在")
			c.StopRun()
		}

		if UserCorp.Role == "1" {
			UserCorpadminCount := map[string]interface{}{}
			UserCorpadminCount["corp_id"] = corp_id
			UserCorpadminCount["role"] = 1
			adminCount := new(models.DhUserCorp).Count(UserCorpadminCount)
			if adminCount < 2 {
				c.EchoJsonErr("管理员唯一不可删除")
				c.StopRun()
			}

		}
		result := UserCorp.Delete(object_id)
		if !result {
			c.EchoJsonErr("删除失败")
		} else {
			c.EchoJsonOk()

		}
	}
	if title == "2" {
		DhUserCorpfilter := map[string]interface{}{}
		Corp := new(models.DhCorp).Find(object_id)
		if Corp == nil {
			c.EchoJsonErr("团队不存在")
			c.StopRun()
		}

		DhUserCorp := new(models.DhUserCorp)
		DhUserCorp.Role = "0"
		DhUserCorp.CorpId = Corp.ObjectId
		DhUserCorp.UserId = user_id
		DhUserCorpfilter["corpid"] = Corp.ObjectId
		DhUserCorpfilter["userid"] = user_id
		DhUserCorpflag := new(models.DhUserCorp).Find(DhUserCorpfilter)
		if DhUserCorpflag != nil {
			c.EchoJsonErr("用户已经添加到团队中")
			c.StopRun()
		}
		result := DhUserCorp.Save()
		if !result {
			c.EchoJsonErr("添加用户失败")
		} else {
			c.EchoJsonOk()
		}

	}
	if role != "" {
		UserCorp := new(models.DhUserCorp).Find(filtersUserCorp)
		if UserCorp == nil {
			c.EchoJsonErr("团队不存在")
			c.StopRun()
		}
		userCorpfilterrole := map[string]interface{}{}
		userCorpfilterrole["role"] = "1"
		userCorpfilterrole["object_id"] = object_id
		if role == "0" {

			number := new(models.DhUserCorp).Count(userCorpfilterrole)
			if number < 2 {
				c.EchoJsonOk("管理员唯一不可改变")
				c.StopRun()

			}
		}
		UserCorp.Role = role
		result := UserCorp.Save()

		if !result {
			c.EchoJsonErr("注册失败")
		} else {
			c.EchoJsonOk()
		}

	}

}
func (c *UserController) ListChangeType() {
	c.Require("datas", "changType-X", "changType-I")
	datas := c.GetString("datas")
	changTypex, _ := c.GetInt("changType-X")
	changTypei, _ := c.GetInt("changType-I")
	plist := *utils.JsonDecodeArrays([]byte(datas))
	errmasage := make([]string, 1)
	for _, v := range plist {
		user := new(models.DhUser).Find(v["object_id"].(string))
		if user == nil {
			errmasage = append(errmasage, v["object_id"].(string))
		} else {
			user.IsDataIUser = changTypei
			user.IsDataXUser = changTypex
			reult := user.Save()
			fmt.Println(reult)
		}
	}
	if errmasage[0] != "" {
		c.EchoJsonErr("部分更新失败")
	}
	c.EchoJsonOk()
}
func (c *UserController) GetUserScreen() {
	c.Require("id")
	id := c.GetString("id")
	dhrelation := map[string]interface{}{}
	dhrelation["user_id"] = id
	dhrelation["relate_type"] = "dx_screen"

	user := new(models.DhUser).Find(id)
	if user == nil {
		c.EchoJsonErr("用户不存在")
		c.StopRun()
	}
	userScreen := new(models.DhRelation).OrderList(dhrelation, "-create_time")

	if userScreen == nil {
		c.EchoJsonErr("用户无大屏数据")
		c.StopRun()
	}
	userScreenData := []utils.P{}
	if len(userScreen) > 0 {
		for _, info := range userScreen {
			alluserscreen := utils.P{}
			alluserscreen["ObjectId"] = info.ObjectId
			alluserscreen["Name"] = info.Name

			userScreenData = append(userScreenData, alluserscreen)
		}
	}
	c.Data["userScreenData"] = userScreenData
	c.Data["name"] = user.Name
	c.Data["userid"] = user.ObjectId
	c.TplName = "user/manageScreen.html"

}
func (c *UserController) DelectUserScreen() {
	c.Require("id", "user_id")
	id := c.GetString("id")
	user_id := c.GetString("user_id")

	dhrelation := map[string]interface{}{}
	dhrelation["user_id"] = user_id
	dhrelation["relate_type"] = "dx_screen"
	dhrelation["object_id"] = id
	user := new(models.DhUser).Find(user_id)
	if user == nil {
		c.EchoJsonErr("用户不存在")
		c.StopRun()
	}
	userScreen := new(models.DhRelation).Delete(dhrelation)
	if userScreen != true {
		c.EchoJsonErr("大屏移除失败")
		c.StopRun()
	}

	c.EchoJsonOk()
}
func (c *UserController) GetUserData() {
	c.EchoJsonOk("测试")
}
