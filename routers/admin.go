package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
	"xzw-blog/controllers/admin"
	"xzw-blog/models"
)

func init() {
	ns :=
   		beego.NewNamespace("/admin",
   				beego.NSBefore(func(c *context.Context) {
   					//获取当前请求的路由
   					pathname := c.Request.URL.String()
   					//获取session，返回的是接口类型，必须进行"类型断言"
					userinfo, ok := c.Input.Session("userinfo").(models.User)
					//对于为进行验证的路由都返回登录页，注意防止死循环
					if !(ok && userinfo.Username != "") {
						if pathname != "/admin/login" && pathname != "/admin/login/dologin" {
							c.Redirect(302, "/admin/login")
						}
					}
				}),
				beego.NSRouter("/login", &admin.LoginController{},"get:Login"),
			beego.NSRouter("/login/dologin", &admin.LoginController{},"post:DoLogin"),
			beego.NSRouter("/form", &admin.ManageController{},"get:Form"),
				beego.NSRouter("/category", &admin.ManageController{},"get:Category"),
				beego.NSRouter("/category_add", &admin.ManageController{},"get:Category_add"),
				beego.NSRouter("/config", &admin.ManageController{},"get:Config"),
				beego.NSRouter("/list", &admin.ManageController{},"get:List"),
				beego.NSRouter("/main", &admin.ManageController{},"get:Main"),
   			)
   beego.AddNamespace(ns)
}
