package admin

import (
	"fmt"
)

type ManageController struct {
	BaseController
}

func (c *ManageController) Get() {
	userinfo:=c.GetSession("userinfo")
	fmt.Printf("11111111111111")
	fmt.Printf("%#v",userinfo)

	c.TplName = "admin/main.html"
}

func (c *ManageController) Form() {
	c.TplName = "admin/_form.tpl"
}

func (c *ManageController) Category() {
	c.TplName = "admin/category.tpl"
}

func (c *ManageController) Category_add() {
	c.TplName = "admin/category_add.tpl"
}

func (c *ManageController) Config() {
	c.TplName = "admin/config.html"
}

func (c *ManageController) List() {
	c.TplName = "admin/list.tpl"
}


