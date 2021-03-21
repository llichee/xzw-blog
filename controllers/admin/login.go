package admin

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/cache"
	"github.com/astaxie/beego/utils/captcha"
	"xzw-blog/forms"
	"xzw-blog/utils"
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
	num, userinfo := utils.Auth(form.Username, form.Password)
	fmt.Println("===================")
	fmt.Println(num, userinfo)
	fmt.Println("===================")
	if flag {
		if num > 0 {
			c.SetSession("userinfo", userinfo)
			c.Success("登陆成功", "/"+beego.AppConfig.String("adminPath")+"/")
		} else {
			c.Error("用户名或密码错误", "/"+beego.AppConfig.String("adminPath")+"/login")
		}
	} else {
		c.Error("验证码错误", "/"+beego.AppConfig.String("adminPath")+"/login")
	}

}

