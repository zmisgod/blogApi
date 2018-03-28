package controllers

import (
	"github.com/zmisgod/blogApi/util"

	"github.com/astaxie/beego"
)

//BaseController 基础的控制器
type BaseController struct {
	beego.Controller
	moduleName     string
	controllerName string
	actionName     string
	options        map[string]string
	cache          *util.MyCache
	pageSize       int
	page           int
	token          string
	userID         int
	userName       string
	imgURL         string
}

//Prepare 准备数据
func (base *BaseController) Prepare() {
	base.pageSize = 12
	if page, err := base.GetInt("page"); err != nil || page < 1 {
		base.page = 1
	} else {
		base.page = page
	}
}

//SendJSON 返送json
func (base *BaseController) SendJSON(code int, data interface{}, msg string) {
	out := make(map[string]interface{})
	out["code"] = code
	out["data"] = data
	out["msg"] = msg
	base.Data["json"] = out
	base.ServeJSON()
}

//CheckError 检查错误
func (base *BaseController) CheckError(err error) {
	if err != nil {
		base.SendJSON(400, "", err.Error())
	}
}
