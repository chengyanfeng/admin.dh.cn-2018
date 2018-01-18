package controllers
import (
	"github.com/astaxie/beego/orm"
	"common.dh.cn/utils"
	"common.dh.cn/models"
	"common.dh.cn/controllers"
	"fmt"
	"strconv"
)


type SourceController struct {
	controllers.BaseController
}
func (c *SourceController) init(i int) {
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

func (c *SourceController) List() {
	c.init(2)
	var mpurl ="/source/list?"
	c.TplName = "source/index.html"
	var total,total_page int64
	typelist:=[]utils.P{}
	var list []*models.DiDatasource
	page,_ := c.GetInt64("page",1)
	page_size,_ := c.GetInt64("page_size",10)
	search := c.GetString("search")
	status:= c.GetString("status")
	sourceType:= c.GetString("sourceType")
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
			}else {

				c.Data["status"] = "nil"
			}
			if len(sourceType)>0{
				c.Data["sourceType"] = sourceType
				mpurl=mpurl+"&sourceType="+sourceType
				condor=cond.AndCond(condor).And("type",sourceType)
			}else {

				c.Data["sourceType"] ="nil"
			}





			number,_:=new(models.DiDatasource).Query().Offset((page-1)*page_size).Limit(page_size).SetCond(condor).OrderBy("-create_time").All(&list)
			total,_=new(models.DiDatasource).Query().SetCond(condor).Count()
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
		}else {

			c.Data["status"] = "nil"
		}
		if len(sourceType)>0{
			c.Data["sourceType"] =sourceType
			filters["type"]=sourceType
			mpurl=mpurl+"&sourceType="+sourceType

		}else {
			c.Data["sourceType"] ="nil"
		}

		total,total_page,list = new(models.DiDatasource).OrderPager(page, page_size, filters,"-create_time")
	}
	data := []utils.P{}
	if len(list) > 0 {
		for _, info := range list {
			dhdatasource := utils.P{}
			dhdatasource["ObjectId"] = info.ObjectId
			dhdatasource["Name"] = info.Name
			dhdatasource["SourceType"] = info.Type
			dhdatasource["UpdateTime"] = info.CreateTime.Format("2006-01-02 15:04:05")
			dhdatasource["Status"] = info.Status
			data = append(data, dhdatasource)
		}
	}

	DiDatasourceGroups:=[] models.DiDatasourceGroup{}
	new(models.DiDatasourceGroup).Query().GroupBy("object_id").All(&DiDatasourceGroups,"name")
	if len(DiDatasourceGroups)>0{
		for _,v:=range DiDatasourceGroups{
			ty:=utils.P{}
			ty["sourceType"]=v.Name
			ty["Type"]=sourceType
			typelist=append(typelist,ty)

		}
	}

	c.Data["sourceTypelist"]=typelist
	c.Data["List"] = data
	c.Data["Pagination"] = PagerHtml(int(total), int(total_page), int(page_size), int(page),mpurl)
}
func (c *SourceController) Update() {
	c.Require("id")
	id := c.GetString("id")
	DiDatasource := new(models.DiDatasource).Find(id)
	if DiDatasource == nil {
		c.EchoJsonErr("数据源不存在")
		c.StopRun()
	}
	if c.GetString("name")!=""{
		DiDatasource.Name = c.GetString("name")
	}
	if c.GetString("SourceType")!=""{
		DiDatasource.Type = c.GetString("SourceType")
	}
	if c.GetString("status")!=""{
		int,err:=strconv.Atoi(c.GetString("status"))
		if err==nil{
			DiDatasource.Status=int
		}

	}
	result := DiDatasource.Save()
	if !result {
		c.EchoJsonErr("更新失败")
	} else {
		c.EchoJsonOk()
	}
}

func (c *SourceController) Listremove() {
	c.Require("datas")
	datas := c.GetString("datas")
	plist:=*utils.JsonDecodeArrays([]byte(datas))
	argerr :=make([]string,1)
	for _,v :=range plist{
		DiDatasource := new(models.DiDatasource).Find(v["object_id"].(string))
		if DiDatasource == nil {
			argerr=append(argerr,v["object_id"].(string))
		}else {
			result := DiDatasource.Delete(DiDatasource.ObjectId)
			if !result{

				argerr=append(argerr,v["object_id"].(string))
			}
		}

	}
	if    len(argerr[0]) > 0 {
		c.EchoJsonErr("部分更新失败")
	}
	c.EchoJsonOk()


}

func (c *SourceController) Remove() {
	c.Require("id")
	id := c.GetString("id")
	DiDatasource := new(models.DiDatasource).Find(id)
	if DiDatasource == nil {
		c.EchoJsonErr("大屏膜版不存在")
		c.StopRun()
	}
	result := DiDatasource.Delete(id)
	if !result {
		c.EchoJsonErr("删除失败")
	} else {
		c.EchoJsonOk()
	}
}

func (c *SourceController) Edit() {
	c.Require("id")
	id := c.GetString("id")
	DiDatasource := new(models.DiDatasource).Find(id)
	if DiDatasource == nil {
		c.EchoJsonErr("数据不存在")
		c.StopRun()
	}
	typelist:=[]utils.P{}
	DiDatasourceGroups:=[] models.DiDatasourceGroup{}
	new(models.DiDatasourceGroup).Query().GroupBy("object_id").All(&DiDatasourceGroups,"name")
	if len(DiDatasourceGroups)>0{
		for _,v:=range DiDatasourceGroups{
			ty:=utils.P{}
			ty["sourceType"]=v.Name
			ty["Type"]=DiDatasource.Type
			typelist=append(typelist,ty)

		}
	}
	c.Data["sourceTypelist"]=typelist
	c.Data["object"] = &DiDatasource
	c.TplName = "source/edit.html"
}