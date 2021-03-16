package admin

import (
	"github.com/astaxie/beego/cache"
	"github.com/astaxie/beego/utils/captcha"
	"xzw-blog/forms"
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
	BaseController
}

func (c *LoginController) Login() {
	c.TplName = "admin/login.html"
}

func (c *LoginController) DoLogin() {
	flag := cpt.VerifyReq(c.Ctx.Request)
	var form forms.UserInfo
	c.ParseForm(&form)
//调utils方法，认证 username 和 password，判断表单提交的用户信息的准确性
	if flag {
		c.Success("验证码正确", "/admin/main")
	} else {
		c.Error("验证码错误", "/admin/login")
	}
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
