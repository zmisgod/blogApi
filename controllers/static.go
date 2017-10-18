package controllers

import (
	"github.com/astaxie/beego"
)

type StaticController struct {
	beego.Controller
}

func (c *StaticController) Static() {
	c.TplName = "index.tpl"
}
