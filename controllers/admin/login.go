package admin

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/cache"
	"github.com/astaxie/beego/utils/captcha"
)

var cpt *captcha.Captcha

func init() {
	store := cache.NewMemoryCache()
	cpt = captcha.NewWithFilter("/captcha/", store)
	cpt.ChallengeNums = 4
	cpt.StdWidth = 100
	cpt.StdHeight = 40
}

type LoginController struct {
	beego.Controller
}

func (c *LoginController) Login() {
	c.TplName = "admin/login.html"
}

func (c *LoginController) Form() {
	c.TplName = "admin/_form.tpl"
}

func (c *LoginController) Category() {
	c.TplName = "admin/category.tpl"
}

func (c *LoginController) Category_add() {
	c.TplName = "admin/category_add.tpl"
}

func (c *LoginController) Config() {
	c.TplName = "admin/config.html"
}

func (c *LoginController) List() {
	c.TplName = "admin/list.tpl"
}

func (c *LoginController) Main() {
	c.TplName = "admin/main.tpl"
}
