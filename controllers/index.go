package controllers

import (
	"fmt"
	_ "github.com/astaxie/beego/httplib"
	"common.dh.cn/def"
	"common.dh.cn/utils"
	"io/ioutil"
	"net/http"
	"bytes"
	"mime/multipart"
	"io"

)

type IndexController struct {
	AdminController
}

func (c *IndexController) Get() {
	c.init(0)
	c.TplName = "index/index.html"
}

func (c *IndexController) FileUploader() {
	data := c.GetString("data")
	bin := utils.Base64Decode(data)
	name := c.GetString("name")
	index, _ := c.GetInt64("index")
	size, _ := c.GetInt64("size")
	total, _ := c.GetInt64("total")
	ext := utils.ToString(utils.Pathinfo(name)["extension"])
	if size > def.MAX_UPLOAD_SIZE {
		c.EchoJsonErr("文件尺寸不得大于", def.MAX_UPLOAD_SIZE)
	}
	md5 := utils.Md5(name)
	file_path := fmt.Sprintf("upload/%v.%v", md5, ext)
	utils.WriteFile(file_path, bin)
	real_url := c.Hostname() + "/" + file_path
	result := utils.P{
		"result":     1,
		"index":      index,
		"total":      total,
		"name":       "name",
		"real_url":   real_url,
		"file_index": c.GetString("file_index"),
	}
	c.EchoJson(result)
}
//从接口中获取信息
func (c *IndexController) Filetest() {
		data,filename,_:= c.GetFile("data")
		var b bytes.Buffer
		w:=multipart.NewWriter(&b)
		file1,_:=w.CreateFormFile("bin",filename.Filename)
		io.Copy(file1,data)
		fw,err:=w.CreateFormField("auth")
		fw.Write([]byte("2f6f0ce5c7ca6ea09a2818f72ce4851d"))
		w.Close()
		req, err := http.NewRequest("POST", "https://dev.datahunter.cn/v2/api/upload", &b)
		req.Header.Set("Content-Type", w.FormDataContentType())
		req.Body=ioutil.NopCloser(&b)
	if err != nil {
		// handle error
	}
	client := &http.Client{}
	resp, err := client.Do(req)
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		// handle error
	}
	utils.ToP(string(body))
	fmt.Println(utils.ToP(string(body)))
	defer resp.Body.Close()
	c.EchoJson(utils.ToP(string(body)))
}