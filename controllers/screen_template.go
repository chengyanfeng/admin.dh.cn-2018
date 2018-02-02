package controllers

import (
	"fmt"
	"strconv"

	"common.dh.cn/models"
	"common.dh.cn/utils"
	"github.com/astaxie/beego/orm"
)

type ScreenTemplateController struct {
	AdminController
}

func (c *ScreenTemplateController) List() {
	var mpurl = "/screen_template/list?"
	c.init(3)
	var total, total_page int64
	var list []*models.DxScreenTemplate
	c.TplName = "screen_template/index.html"
	page, _ := c.GetInt64("page", 1)
	page_size, _ := c.GetInt64("page_size", 10)
	filters := map[string]interface{}{}
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

			number, _ := new(models.DxScreenTemplate).Query().Offset((page - 1) * page_size).Limit(page_size).SetCond(condor).OrderBy("-create_time").All(&list)
			total, _ = new(models.DxScreenTemplate).Query().SetCond(condor).Count()
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

		total, total_page, list = new(models.DxScreenTemplate).OrderPager(page, page_size, filters, "-create_time")
	}
	data := []utils.P{}
	if len(list) > 0 {
		for _, info := range list {
			Screen := utils.P{}
			Screen["ObjectId"] = info.ObjectId
			Screen["Screenid"] = info.ScreenId
			Screen["Name"] = info.Name
			Screen["Status"] = info.Status
			Screen["CreateTime"] = info.CreateTime.Format("2006-01-02 15:04:05")
			Screen["UpdateTime"] = info.UpdateTime.Format("2006-01-02 15:04:05")
			data = append(data, Screen)
		}
	}
	fmt.Println(data, "------------------------data-------------")
	c.Data["List"] = data
	c.Data["Pagination"] = PagerHtml(int(total), int(total_page), int(page_size), int(page), mpurl)
}
func (c *ScreenTemplateController) Update() {
	c.Require("id")
	id := c.GetString("id")
	dxScreenTemplate := new(models.DxScreenTemplate).Find(id)
	if dxScreenTemplate == nil {
		c.EchoJsonErr("用户不存在")
		c.StopRun()
	}
	if c.GetString("name") != "" {
		dxScreenTemplate.Name = c.GetString("name")
	}
	if c.GetString("logos") != "" {
		dxScreenTemplate.Name = c.GetString("name")
	}
	if c.GetString("status") != "" {
		int, err := strconv.Atoi(c.GetString("status"))
		if err == nil {
			dxScreenTemplate.Status = int
		}

	}
	result := dxScreenTemplate.Save()
	if !result {
		c.EchoJsonErr("更新失败")
	} else {
		c.EchoJsonOk()
	}
}
func (c *ScreenTemplateController) ListRemove() {
	c.Require("datas")
	datas := c.GetString("datas")
	plist := *utils.JsonDecodeArrays([]byte(datas))
	argerr := make([]string, 1)
	for _, v := range plist {
		DXScreenTemplate := new(models.DxScreenTemplate).Find(v["object_id"].(string))
		if DXScreenTemplate == nil {
			argerr = append(argerr, v["object_id"].(string))
		} else {
			switch DXScreenTemplate.Status {
			case -1, 0:
				DXScreenTemplate.Status = 1
			case 1:
				DXScreenTemplate.Status = -1
			}
			DXScreenTemplate.Save()
		}

	}
	if len(argerr[0]) > 0 {
		c.EchoJsonErr("部分更新失败")
	}
	c.EchoJsonOk()

}
func (c *ScreenTemplateController) Remove() {
	c.Require("id")
	id := c.GetString("id")
	dxScreenTemplate := new(models.DxScreenTemplate).Find(id)
	if dxScreenTemplate == nil {
		c.EchoJsonErr("大屏模版不存在")
		c.StopRun()
	}
	result := dxScreenTemplate.Delete(id)
	if !result {
		c.EchoJsonErr("删除失败")
	} else {
		c.EchoJsonOk()
	}
}
func (c *ScreenTemplateController) Create() {
	c.TplName = "screen_template/create.html"
}
func (c *ScreenTemplateController) Edit() {
	c.TplName = "screen_template/edit.html"
}
