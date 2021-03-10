package admin

import (
	"github.com/astaxie/beego"
)

type BaseController struct {
	beego.Controller
}

func (c *BaseController) Success(message string, rediect string) {
	c.Data["message"] = message
	c.Data["rediect"] = rediect
	c.TplName = "admin/public/success.html"
}

func (c *BaseController) Error(message string, rediect string) {
	c.Data["message"] = message
	c.Data["rediect"] = rediect
	c.TplName = "admin/public/error.html"
}
