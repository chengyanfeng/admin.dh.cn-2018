package filter
import (
	"common.dh.cn/utils"
	"github.com/astaxie/beego/context"
)
var BaseFilter = func(ctx *context.Context) {
	if ctx.Request.RequestURI  !="/"  {
		if ctx.Request.RequestURI  !="/login" {
			if ctx.Request.RequestURI !="/login/quit" {
				gooid:= ctx.Input.Session("gooid")
			if len(utils.ToStrings(gooid))>0 {
					cachgooid:=utils.S("gooid")
					if cachgooid!=gooid {
						ctx.Redirect(302,"/")
					}

				}else {
					ctx.Redirect(302,"/")
				}
			}
		}

	}
}















