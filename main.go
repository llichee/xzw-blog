package main

import (
	"github.com/astaxie/beego"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"xzw-blog/models"
	_ "xzw-blog/routers"
	"xzw-blog/utils"
)

func main() {
	models.DB.AutoMigrate(models.User{})
	if u:=utils.GetUserByname("hdadmin"); u.Username == "" {
		user := models.User{
			ID:        0,
			Username:  "hdadmin",
			Password:  "hengda888",
			CreatedAt: nil,
			UpdatedAt: nil,
			DeletedAt: nil,
		}
		models.DB.Create(&user)
	}
	beego.Run()
	defer models.DB.Close()
}

