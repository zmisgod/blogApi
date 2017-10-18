package controllers

import "github.com/astaxie/beego"

type BaseController struct {
	beego.Controller
}

func (this *BaseController) SendJSON(code int, data interface{}, msg string) {
	out := make(map[string]interface{})
	out["code"] = code
	out["data"] = data
	out["msg"] = msg
	this.Data["json"] = out
	this.ServeJSON()
}
