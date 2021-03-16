package main

import (
	"github.com/astaxie/beego"
	"xzw-blog/models"
	_ "xzw-blog/routers"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {
	models.DB.AutoMigrate(models.User{})
	beego.Run()
	defer models.DB.Close()
}

