package models

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"time"
)

var DB *gorm.DB
var err error

type User struct {
	ID int64 `gorm:"column(id);pk;auto"`
	Username string `gorm:"size(32)"`
	Password string `gorm:"size(1024)"`
	CreatedAt *time.Time `gorm:"auto_now_add;"`
	UpdatedAt *time.Time `gorm:"auto_now;"`
	DeletedAt *time.Time `gorm:"null;"`
}

func init() {
	DB, err = gorm.Open("mysql", "root:hengda888@/db_blog?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	} else {
		fmt.Println("数据库连接成功!")
	}
}

func (User) TableName() string {
	return "userinfo"
}
