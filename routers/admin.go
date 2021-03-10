package routers

import (
	"fmt"
	"xzw-blog/controllers/admin"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
)

func init() {
	ns :=
   		beego.NewNamespace("/admin",
   				beego.NSBefore(func(c *context.Context) {
					fmt.Println("中间件，匹配路由之前执行，可以添加登录验证")
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
