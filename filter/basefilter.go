package filter

import (
	"common.dh.cn/utils"
	"github.com/astaxie/beego/context"
)

var BaseFilter = func(ctx *context.Context) {
	gooid := ctx.Input.Session("gooid")
	cachgooid := utils.S("gooid")
	if cachgooid != gooid {
		ctx.Redirect(302, "/login")
	}
}
