package routers

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
	"xzw-blog/controllers/admin"
	"xzw-blog/models"
	_ "xzw-blog/models"
)

func init() {
	ns :=
   		beego.NewNamespace("/"+beego.AppConfig.String("adminPath"),
   				beego.NSBefore(func(c *context.Context) {

   					//获取当前请求的路由
   					pathname := c.Request.URL.String()
   					//获取session，返回的是接口类型，必须进行"类型断言"
					//userinfo, ok := c.Input.Session("userinfo").(models.User)
					userinfo := models.User{
						ID:        1,
						Username:  "hdadmin",
						Password:  "efd76d0b8427d4234655cb23bee8f507",
						CreatedAt: nil,
						UpdatedAt: nil,
						DeletedAt: nil,
					}
					fmt.Println(userinfo)
					//对于为进行验证的路由都返回登录页，注意防止死循环
					//if !(ok && userinfo.Username != "") {
					if userinfo.Username == "" {
							if pathname != "/"+beego.AppConfig.String("adminPath")+"/login" && pathname != "/"+beego.AppConfig.String("adminPath")+"/login/dologin" {
								c.Redirect(302, "/"+beego.AppConfig.String("adminPath")+"/login")

						} else {
							fmt.Println("登陆成功不跳转登录页了")
						}
					}

				}),
				beego.NSRouter("/", &admin.ManageController{}),
				beego.NSRouter("/login", &admin.LoginController{},"get:Login"),
				beego.NSRouter("/login/dologin", &admin.LoginController{},"post:DoLogin"),

				beego.NSRouter("/form", &admin.ManageController{},"get:Form"),
				beego.NSRouter("/category", &admin.ManageController{},"get:Category"),
				beego.NSRouter("/category_add", &admin.ManageController{},"get:Category_add"),
				beego.NSRouter("/config", &admin.ManageController{},"get:Config"),
				beego.NSRouter("/list", &admin.ManageController{},"get:List"),
   			)
   beego.AddNamespace(ns)
}
