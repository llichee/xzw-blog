package admin

import "github.com/astaxie/beego"

type ManageController struct {
	beego.Controller
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

func (c *ManageController) Main() {
	c.TplName = "admin/main.tpl"
}
