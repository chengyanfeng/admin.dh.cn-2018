package controllers

import (
	"common.dh.cn/utils"
	"common.dh.cn/models"

	"github.com/astaxie/beego/orm"
	"fmt"
	"strconv"
)

type SourceShareController struct {
	AdminController
}


func (c *SourceShareController) List() {
	var mpurl = "datasource_share/list?"
	c.init(5)
	var total, total_page int64
	var list []*models.DiDatasource
	c.TplName = "datasource_share/index.html"
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

			number, _ := new(models.DiDatasource).Query().Offset((page - 1) * page_size).Limit(page_size).SetCond(condor).OrderBy("-create_time").All(&list)
			total, _ = new(models.DiDatasource).Query().SetCond(condor).Count()
			if total % page_size != 0 {
				total_page = total / page_size + 1
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

		total, total_page, list = new(models.DiDatasource).OrderPager(page, page_size, filters, "-create_time")
		fmt.Println(list, "---------------------list--------------------")
	}
	data := []utils.P{}
	if len(list) > 0 {
		for _, info := range list {
			Screen := utils.P{}
			DiSourceShareFilter := utils.P{}
			DiSourceShareFilter["sourceid"] = info.ObjectId
			DiSourceShare := new(models.DiSourceShare).List(DiSourceShareFilter)
			if len(DiSourceShare) > 0 {
				DhCorpArray := make([]string, 1)
				for _, v := range DiSourceShare {
					DhCorpfilter := utils.P{}
					DhCorpfilter["object_id"] = v.Corpid
					DhCorp := new(models.DhCorp).Find(DhCorpfilter)
					DhCorpArray = append(DhCorpArray, DhCorp.Name)
				}
				Screen["CorpName"] = utils.ToString(DhCorpArray)
				Screen["Status"] = 1
			} else {
				Screen["CorpName"] = "---"
				Screen["Status"] = 0
			}
			Screen["ObjectId"] = info.ObjectId
			Screen["Name"] = info.Name
			Screen["format"] = info.Format
			Screen["CreateTime"] = info.CreateTime.Format("2006-01-02 15:04:05")
			Screen["UpdateTime"] = info.UpdateTime.Format("2006-01-02 15:04:05")
			data = append(data, Screen)
		}
	}
	c.Data["List"] = data
	fmt.Println(data, "---------------------data--------------------")
	c.Data["Pagination"] = PagerHtml(int(total), int(total_page), int(page_size), int(page), mpurl)
}

func (c *SourceShareController) Create(){

	c.TplName="datasource_share/create.html"


}