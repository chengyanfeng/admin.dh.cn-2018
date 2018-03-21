package controllers

import (
	"common.dh.cn/utils"
	"common.dh.cn/models"

	"fmt"
	"strconv"
	"strings"

)

type SourceShareController struct {
	AdminController
}

func (c *SourceShareController) List() {
	defer func(){ // 必须要先声明defer，否则不能捕获到panic异常
		fmt.Println("c")
		if err:=recover();err!=nil{
			fmt.Println(err) // 这里的err其实就是panic传入的内容，55
		}
		fmt.Println("d")
	}()
	var mpurl = "/admin/sourceshare/list?"
	c.init(5)
	var total, total_page int64
	var list []*models.DiDatasource
	var Dhlist []*models.DhRelation
	c.TplName = "datasource_share/index.html"
	page, _ := c.GetInt64("page", 1)
	page_size, _ := c.GetInt64("page_size", 10)
	filters := map[string]interface{}{}
	userid := utils.P{}
	userid["user_id"] = c.GetSession("Object_id")
	filters["user_id"]=c.GetSession("Object_id")
	filters["corp_id"]=c.GetSession("Object_id")
	search := c.GetString("search")
	shareflag := c.GetString("status")
	if len(search) > 0 { /*
		qs:=new(models.DiDatasource).Query()
		if len(search) > 0 {
			c.Data["search"] = search
			mpurl = mpurl + "&search=" + search
			qs=qs.Filter("name__icontains",search)
			if len(shareflag ) > 0 {
				c.Data["shareflag"] = shareflag
				int, _ := strconv.Atoi(shareflag )
				mpurl = mpurl + "&status=" + shareflag
				qs=qs.Filter("share_flag",int)
			} else {
				c.Data["shareflag"] = "nil"
			}

			number, _ := qs.Filter("user_id",corp_id_list).Offset((page - 1) * page_size).Limit(page_size).OrderBy("-create_time").All(&Dhlist)
			total, _ = qs.Filter("group_id__in",corp_id_list).Count()
			if total%page_size != 0 {
				total_page = total/page_size + 1
			} else {
				total_page = total / page_size
			}
			fmt.Println(number)
		}

	*/} else {
		if len(shareflag) > 0 {
			c.Data["shareflag"] = shareflag
			int, _ := strconv.Atoi(shareflag)
			filters["status"] = int

			mpurl = mpurl + "&status=" + shareflag
		} else {
			c.Data["shareflag"] = "nil"
		}
		total, total_page, Dhlist=	new(models.DhRelation).OrderPager(page, page_size, filters, "-create_time")
		fmt.Println(list, "---------------------list--------------------")
	}
	data := []utils.P{}
	if len(Dhlist) > 0 {
		for _, info := range Dhlist {
				Screen := utils.P{}
				didatasource:=new(models.DiDatasource).Find(info.RelateId)
				if didatasource!=nil{
					if didatasource.ShareFlag==1{
						var corpNameArry =[]string{}
						var dhsharelist []*models.DiDataSourceShare
						new(models.DiDataSourceShare).Query().Filter("datasource_id",didatasource.ObjectId).GroupBy("corp_id").All(&dhsharelist,"corp_id")
						for _,v:=range dhsharelist{
							dhcorp:=new(models.DhCorp).Find(v.CorpId)
							corpNameArry=append(corpNameArry, dhcorp.Name)
						}
						Screen["CorpName"] = utils.ToString(corpNameArry)
					}else {
						Screen["CorpName"] = utils.ToString("---")
					}
						Screen["Status"] = 0
						Screen["ObjectId"] = didatasource.ObjectId
						Screen["Url"] = didatasource.Url
						Screen["Name"] = didatasource.Name
						Screen["Format"] = didatasource.Format
						Screen["CreateTime"] = didatasource.CreateTime.Format("2006-01-02 15:04:05")
						Screen["UpdateTime"] = didatasource.UpdateTime.Format("2006-01-02 15:04:05")
						data = append(data, Screen)
						}
						}
	}
	c.Data["List"] = data
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
	filetype, name := filetype(nameandtype)
	url := c.GetString("upurl")
	DiDatasource := new(models.DiDatasource)
	DiDatasource.Type = "file"
	DiDatasource.Format = filetype
	DiDatasource.Name = name
	DiDatasource.Url = url
	result := DiDatasource.Save()
	if !result {
		c.EchoJsonErr("添加失败")
	} else {
		c.EchoJsonOk()
	}

}

func filetype(filename string) (filetype string, name string) {
	int := strings.LastIndex(filename, ".")
	name = filename[0:int]
	filetype = filename[int:]
	return filetype, name
}

func (c *SourceShareController) Remove() {
	c.Require("id")
	id := c.GetString("id")
	DiDatasource := new(models.DiDatasource).Find(id)
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
	id := c.GetString("id")
	c.Data["url"] = id
	c.TplName = "datasource_share/showData.html"
}

func (c *SourceShareController) ShareCorp() {
	id := c.GetString("id")
	corpidlist := c.GetString("corpIdlist")
	corplist := strings.Split(corpidlist, ",")
	fmt.Print(corpidlist)
	//数据源id
	c.Data["id"] = id
	//获取数据源name
	didatasource := new(models.DiDatasource).Find(id)
	c.Data["DatasourceName"] = didatasource.Name

	//获取所有的团队
	dhcorps := []models.DhCorp{}
	filtersAllUser := map[string]interface{}{}
	filtersAllUser["status__gte"] = 0
	list := new(models.DhCorp).OrderList(filtersAllUser, "-create_time")
	for _, v := range list {
		dhcorp := models.DhCorp{}
		dhcorp.ObjectId = v.ObjectId
		dhcorp.Email = v.Email
		dhcorp.Name = v.Name
		if len(corplist) > 0 {
			for _, corp := range corplist {
				if v.ObjectId == corp {
					diSourceSharefilter := utils.P{}
					diSourceSharefilter["datasource_id"] = id
					diSourceSharefilter["corpid"] = corp
					DiSourceShare := new(models.DiDataSourceShare).List(diSourceSharefilter)
					//由于查询出的成员是一样的分享状态，和字段控制，所以只取第一个就行了
					if DiSourceShare[0].Fields == "1" {
						dhcorp.Status = 2
						break
					} else {
						//由于创建dhcorp 时候没有创建多余字段，所以我暂时利用vcode 字段来存储东西
						dhcorp.Vcode = DiSourceShare[0].Fields
						dhcorp.Status = 0
						break
					}
				} else {
					dhcorp.Status = 1

				}
			}
			dhcorps = append(dhcorps, dhcorp)
		} else {
			dhcorp.Status = 1
			dhcorps = append(dhcorps, dhcorp)
		}

	}
	c.Data["dhcorps"] = dhcorps
	c.TplName = "datasource_share/sharecorp.html"
}

func (c *SourceShareController) DbConnect() {

	c.TplName = "datasource_share/dbconnect.html"
}

func (c *SourceShareController) SaveShareCorp() {

	data := c.GetString("data")

	a := *utils.JsonDecode([]byte(utils.ToString(data)))

	datasourceid := a["datasourceid"]
	datasourcename := a["datasourcename"]
	fmt.Print(datasourceid)
	args := a["args"]
	fmt.Print(args)
	m := args.([]interface{})
	for _, v := range m {
		deletefilter := utils.P{}
		corp := make(map[string]interface{})
		fiter := v.(map[string]interface{})
		corp["corp_id"] = fiter["corpid"]
		deletefilter["corp_id"] = fiter["corpid"]
		deletefilter["datasource_id"] = datasourceid
		delectlist := new(models.DiDataSourceShare).OrderList(deletefilter)
		fmt.Print(delectlist)
		for _, v := range delectlist {
			//先全部删除团队和数据源的信息
			flag := new(models.DiDataSourceShare).Delete(v.ObjectId)
			fmt.Print(flag)
		}

		parameter := fiter["filter"]
		//查询出团队的所有成员
		DhUserCorp := new(models.DhUserCorp).OrderList(corp, "-create_time")
		fmt.Print(DhUserCorp)
		for _, v := range DhUserCorp {
			disourceshare := new(models.DiDataSourceShare)
			disourceshare.DatasourceId = utils.ToString(datasourceid);
			disourceshare.UserId = v.ObjectId
			disourceshare.CorpId = v.CorpId
			disourceshare.Name = utils.ToString(datasourcename)
			disourceshare.Fields = utils.ToString(parameter)
			//如果字段不是为1 的话，说明是部分显示
			if parameter == "1" {
				disourceshare.IsFullShow = "1"
			} else {
				disourceshare.IsFullShow = "0"
			}
			flag := disourceshare.Save()
			if flag == true {
				//保存到关联表中
				//先把上面保存的数据，根据uerid,和corpid,和dataource_id 三个条件确定一个数据 查询出来
				dhrelationmap := map[string]interface{}{}
				dhrelationmap["user_id"] = v.UserId
				dhrelationmap["corp_id"] = v.CorpId
				dhrelationmap["datasource_id"] = utils.ToString(datasourceid)
				//再把这条数据保存到DhRelation
				//先查询出数据源的name
				didatasource:=new(models.DiDatasource).Find(utils.ToString(datasourceid))
				didatasource.ShareFlag=1
				//更改分享字段
				flag:= didatasource.Save()
				  if flag==true{
					  dhrelation := new(models.DhRelation)
					  dhrelation.Name = didatasource.Name
					  dhrelation.CorpId = v.CorpId
					  dhrelation.UserId = v.UserId
					  dhrelation.RelateId =utils.ToString(datasourceid)
					  dhrelation.Auth = "admin_share"
					  dhrelation.RelateType = "di_datasource"

					  shareflag := dhrelation.Save()
					  if shareflag == true {
						  //暂时不做处理
					  }
				  }

			}

		}

	}
	c.EchoJsonOk()
}
