package controllers

import (
	"common.dh.cn/utils"
	"common.dh.cn/models"

	"github.com/astaxie/beego/orm"
	"fmt"
	"strconv"
	"strings"
)

type SourceShareController struct {
	AdminController
}

func (c *SourceShareController) List() {
	var mpurl = "/admin/sourceshare/list?"
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

		total, total_page, list = new(models.DiDatasource).OrderPager(page, page_size, filters, "-create_time")
		fmt.Println(list, "---------------------list--------------------")
	}
	data := []utils.P{}
	if len(list) > 0 {
		for _, info := range list {
			Screen := utils.P{}
			DiSourceShareFilter := utils.P{}
			DiSourceShareFilter["DiDatasourceId"] = info.ObjectId
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
			Screen["Url"]=info.Url
			Screen["Name"] = info.Name
			Screen["Format"] = info.Format
			Screen["CreateTime"] = info.CreateTime.Format("2006-01-02 15:04:05")
			Screen["UpdateTime"] = info.UpdateTime.Format("2006-01-02 15:04:05")
			data = append(data, Screen)
		}
	}
	c.Data["List"] = data
	fmt.Println(data, "---------------------data--------------------")
	c.Data["Pagination"] = PagerHtml(int(total), int(total_page), int(page_size), int(page), mpurl)
}
/**
暂时无用，因为已经全部调用前端接口！！！所以，暂时不用这个接口
 */
func (c *SourceShareController) Create() {

	c.TplName = "datasource_share/create.html"

}
/**
暂时无用，因为已经全部调用前端接口！！！所以，暂时不用这个接口
 */
func (c *SourceShareController) Add() {

	nameandtype := c.GetString("filename")
	filetype,name:=filetype(nameandtype)
	url := c.GetString("upurl")
	DiDatasource := new(models.DiDatasource)
	DiDatasource.Type="file"
	DiDatasource.Format=filetype
	DiDatasource.Name =name
	DiDatasource.Url = url
	result := DiDatasource.Save()
	if !result {
		c.EchoJsonErr("添加失败")
	} else {
		c.EchoJsonOk()
	}

}

func filetype(filename string)(filetype string,name string){
			int:=strings.LastIndex(filename,".")
			name=filename[0:int]
			filetype=filename[int:]
			return filetype,name
}

func (c *SourceShareController) Remove() {
	c.Require("id")
	id := c.GetString("id")
	DiDatasource:= new(models.DiDatasource).Find(id)
	if DiDatasource == nil {
		c.EchoJsonErr("数据源不存在")
		c.StopRun()
	}
	result := DiDatasource.Delete(id)
	if !result {
		c.EchoJsonErr("删除失败")
	} else {
		c.EchoJsonOk()
	}
}

func (c *SourceShareController) ListRemove() {
	c.Require("datas")
	datas := c.GetString("datas")
	plist := *utils.JsonDecodeArrays([]byte(datas))
	argerr := make([]string, 1)
	for _, v := range plist {
		DiDatasources := new(models.DiDatasource).Find(v["object_id"].(string))
		if DiDatasources == nil {
			argerr = append(argerr, v["object_id"].(string))
		} else {
			result := DiDatasources.Delete(v["object_id"].(string))
			if !result {
				argerr = append(argerr, v["object_id"].(string))
			}
		}
	}
	if len(argerr[0]) > 0 {
		c.EchoJsonErr("部分删除失败")
	}
	c.EchoJsonOk()

}
func (c *SourceShareController) ShowData() {
	id:=c.GetString("id")
	c.Data["url"]=id
	c.TplName = "datasource_share/showData.html"
}

func (c *SourceShareController) ShareCorp() {
	id:=c.GetString("id")
	c.Data["id"]=id
	//获取所有的团队
	 dhcorps:=[]models.DhCorp{}
	filtersAllUser := map[string]interface{}{}
	filtersAllUser["status__gte"] = 0
	list := new(models.DhCorp).OrderList(filtersAllUser, "-create_time")
	for _,v:=range list{
		dhcorp:=models.DhCorp{}
		dhcorp.ObjectId=v.ObjectId
		dhcorp.Email=v.Email
		dhcorp.Name=v.Name
		dhcorp.Status=1
		dhcorps=append(dhcorps, dhcorp)
	}
	c.Data["dhcorps"]=dhcorps
	c.TplName = "datasource_share/sharecorp.html"
}


func (c *SourceShareController) DbConnect() {

	c.TplName = "datasource_share/dbconnect.html"
}

func (c *SourceShareController) SaveShareCorp() {
	p := c.FormToP("datasourceid", "args")

	 print(utils.ToString((p["datasourceid"])),"---------------")
	c.Data["url"]=p
	c.TplName = "datasource_share/showData.html"
}


