package filter
import (
	"common.dh.cn/utils"
	"github.com/astaxie/beego/context"
	"fmt"
)

var BaseFilter = func(ctx *context.Context) {
	fmt.Println(ctx.Request.RequestURI,"------------url---------")

	if ctx.Request.RequestURI  !="/"  {
		 if ctx.Request.RequestURI  !="/login" {
			 if ctx.Request.RequestURI !="/login/quit" {
	gooid:= ctx.Input.Session("gooid")


	if len(utils.ToString(gooid))>0 {
		cachgooid:=utils.S("gooid")
		fmt.Println(gooid,"---------------gooid-----------------")
		fmt.Println(cachgooid,"---------------gooid-----------------")
		if cachgooid!=gooid {
			fmt.Println(cachgooid,"---------------跳回去了-----------------")
			ctx.Redirect(301,"/")
		}

	}else {
		ctx.Redirect(301,"/")
	}
	    }
	 }

		 fmt.Println("------------------------login----------------")
	 }
}















