package controllers

import (
	"github.com/zmisgod/blogApi/util"

	"github.com/astaxie/beego"
)

type BaseController struct {
	beego.Controller
	moduleName     string
	controllerName string
	actionName     string
	options        map[string]string
	cache          *util.MyCache
	pageSize       int
	page           int
}

func (this *BaseController) Prepare() {
	this.pageSize = 12
	if page, err := this.GetInt("page"); err != nil || page < 1 {
		this.page = 1
	} else {
		this.page = page
	}
}

func (this *BaseController) SendJSON(code int, data interface{}, msg string) {
	out := make(map[string]interface{})
	out["code"] = code
	out["data"] = data
	out["msg"] = msg
	this.Data["json"] = out
	this.ServeJSON()
}

func (this *BaseController) CheckError(err error) {
	if err != nil {
		this.SendJSON(400, "", err.Error())
	}
}
