package controllers
import (
	"github.com/astaxie/beego/orm"
	"common.dh.cn/utils"
	"common.dh.cn/models"
	"common.dh.cn/controllers"
	"fmt"
	"strconv"

)
type SourcetypeController struct {
	controllers.BaseController
}
func (c *SourcetypeController) init(i int) {
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
	c.Data["Menu"]=Menu
}


func (c *SourcetypeController) List() {
	c.init(2)
	var mpurl ="/sourcetype/list?"

	c.TplName = "sourcetype/index.html"
	var total,total_page int64
	var list []*models.DhDatasource
	page,_ := c.GetInt64("page",1)
	page_size,_ := c.GetInt64("page_size",3)
	search := c.GetString("search")
	status:= c.GetString("status")
	filters := map[string]interface{}{}
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
			}

			number,_:=new(models.DhDatasource).Query().Offset((page-1)*page_size).Limit(page_size).SetCond(condor).All(&list)
			total,_=new(models.DhDatasource).Query().SetCond(condor).Count()
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
		}

		total,total_page,list = new(models.DhDatasource).OrderPager(page-1, page_size, filters,"-create_time")
	}
	data := []utils.P{}
	if len(list) > 0 {
		for _, info := range list {
			dhdatasource := utils.P{}
			dhdatasource["ObjectId"] = info.ObjectId
			dhdatasource["DataSourceTypeName"] = info.Name
			dhdatasource["DataSourceTypeTime"] = info.CreateTime.Format("2006-01-02 15:04:05")
			dhdatasource["DataSourceTypeStatus"] = info.Status
			data = append(data, dhdatasource)
		}
	}
	c.Data["List"] = data
	c.Data["Pagination"] = PagerHtml(int(total), int(total_page), int(page_size), int(page),mpurl)
}