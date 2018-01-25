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
	Authname,_:=c.GetSecureCookie("2rdsfada3@#$%^&*","Authname")
	c.Data["Authname"]=Authname
}

func (c *SourceController) List() {
	c.init(2)
	var mpurl ="/source/list?"
	c.TplName = "source/index.html"
	var total,total_page int64
	typelist:=[]utils.P{}
	var list []*models.DiDatasourcePub
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





			number,_:=new(models.DiDatasourcePub).Query().Offset((page-1)*page_size).Limit(page_size).SetCond(condor).OrderBy("-create_time").All(&list)
			total,_=new(models.DiDatasourcePub).Query().SetCond(condor).Count()
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
			filters["datasource_type_id"]=sourceType
			mpurl=mpurl+"&sourceType="+sourceType

		}else {
			c.Data["sourceType"] ="nil"
		}

		total,total_page,list = new(models.DiDatasourcePub).OrderPager(page, page_size, filters,"-create_time")
	}
	data := []utils.P{}
	if len(list) > 0 {
		for _, info := range list {
			DiDatasourceTypefilter:=map[string]interface{}{}
			DiDatasourceTypefilter["object_id"]=info.DatasourceTypeId
			DiDatasourceType:=new(models.DiDatasourceType).Find(DiDatasourceTypefilter["object_id"])
			if DiDatasourceType ==nil{
				c.EchoJsonErr("没有当前类型")
			}
			dhdatasource := utils.P{}
			dhdatasource["ObjectId"] = info.ObjectId
			dhdatasource["Name"] = info.Name
			dhdatasource["TypeName"] = DiDatasourceType.Name
			dhdatasource["UpdateTime"] = info.CreateTime.Format("2006-01-02 15:04:05")
			dhdatasource["Status"] = info.Status
			data = append(data, dhdatasource)
		}
	}

	DiDatasourceTypes:=[] models.DiDatasourceType{}
	new(models.DiDatasourceType).Query().GroupBy("object_id").All(&DiDatasourceTypes,"name")
	if len(DiDatasourceTypes)>0{
		for _,v:=range DiDatasourceTypes{
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
	DiDatasourceData := new(models.DiDatasourcePub).Find(id)
	if DiDatasourceData == nil {
		c.EchoJsonErr("数据源不存在")
		c.StopRun()
	}
	if c.GetString("name")!=""{
		DiDatasourceData.Name = c.GetString("name")
	}
	if c.GetString("SourceType")!=""{
		DiDatasourceTypefilter:=map[string]interface{}{}
		DiDatasourceTypefilter["name"]=c.GetString("SourceType")
		DiDatasourceType:=new(models.DiDatasourceType).Find(DiDatasourceTypefilter["object_id"])
		if DiDatasourceType ==nil{
			c.EchoJsonErr("没有当前类型")
		}
		DiDatasourceData.DatasourceTypeId = DiDatasourceType.ObjectId
	}
	if c.GetString("status")!=""{
		int,err:=strconv.Atoi(c.GetString("status"))
		if err==nil{
			DiDatasourceData.Status=int
		}

	}
	result := DiDatasourceData.Save()
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
		DiDatasourceData := new(models.DiDatasourcePub).Find(v["object_id"].(string))
		if DiDatasourceData == nil {
			argerr=append(argerr,v["object_id"].(string))
		}else {
			result := DiDatasourceData.Delete(DiDatasourceData.ObjectId)
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
func (c *SourceController) ShowData() {

	c.TplName = "source/showData.html"
}
func (c *SourceController) Remove() {
	c.Require("id")
	id := c.GetString("id")
	DiDatasourceData := new(models.DiDatasourcePub).Find(id)
	if DiDatasourceData == nil {
		c.EchoJsonErr("大屏膜版不存在")
		c.StopRun()
	}
	result := DiDatasourceData.Delete(id)
	if !result {
		c.EchoJsonErr("删除失败")
	} else {
		c.EchoJsonOk()
	}
}

func (c *SourceController) Edit() {
	c.Require("id")
	id := c.GetString("id")
	DiDatasourceData := new(models.DiDatasourcePub).Find(id)
	if DiDatasourceData == nil {
		c.EchoJsonErr("数据不存在")
		c.StopRun()
	}
	typelist:=[]utils.P{}
	DiDatasourceTypes:=[] models.DiDatasourceType{}
	new(models.DiDatasourceType).Query().GroupBy("object_id").All(&DiDatasourceTypes,"name")
	if len(DiDatasourceTypes)>0{
		for _,v:=range DiDatasourceTypes{
			ty:=utils.P{}
			ty["sourceType"]=v.Name
			if v.ObjectId==DiDatasourceData.DatasourceTypeId {
			ty["Type"]=v.Name
			}
			typelist=append(typelist,ty)

		}
	}

	c.Data["sourceTypelist"]=typelist
	c.Data["object"] = &DiDatasourceData
	fmt.Println(DiDatasourceData,"abc=------------------------afaf")
	c.TplName = "source/edit.html"
}