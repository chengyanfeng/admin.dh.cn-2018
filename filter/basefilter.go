package filter

import (
	"github.com/astaxie/beego/context"
)

var BaseFilter = func(ctx *context.Context) {
	auth := ctx.Input.Session("auth")
	if auth == nil {
		ctx.Redirect(302, "/login")

	}
}
