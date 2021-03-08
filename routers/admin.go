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
				beego.NSRouter("/form", &admin.LoginController{},"get:Form"),
				beego.NSRouter("/category", &admin.LoginController{},"get:Category"),
				beego.NSRouter("/category_add", &admin.LoginController{},"get:Category_add"),
				beego.NSRouter("/config", &admin.LoginController{},"get:Config"),
				beego.NSRouter("/list", &admin.LoginController{},"get:List"),
				beego.NSRouter("/main", &admin.LoginController{},"get:Main"),
   			)
   beego.AddNamespace(ns)
}
