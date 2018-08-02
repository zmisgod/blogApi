package controllers

import (
	"strings"

	"github.com/zmisgod/blogApi/models"

	"github.com/astaxie/beego"
)

//BaseController 基础的控制器
type BaseController struct {
	beego.Controller
	pageSize   int
	page       int
	ip         string
	userAgent  string
	authority  string
	requestURI string
	refer      string
}

//Prepare 准备数据
func (base *BaseController) Prepare() {
	base.pageSize = 12
	if page, err := base.GetInt("page"); err != nil || page < 1 {
		base.page = 1
	} else {
		base.page = page
	}
	base.ip = base.Ctx.Input.IP()
	base.userAgent = base.Ctx.Input.UserAgent()
	base.authority = base.Ctx.Input.Header("authorization")
	base.requestURI = base.Ctx.Input.URI()
	base.refer = base.Ctx.Input.Refer()
	c, _ := base.GetControllerAndAction()
	controllerPrefix := strings.Replace(c, "Controller", "", 10)
	devMode := beego.AppConfig.String("runmode")
	if devMode != "dev" {
		//用户请求日志
		models.SaveUserVisiteHistory(controllerPrefix, base.ip, base.userAgent, base.requestURI, base.refer)
		validIPs := beego.AppConfig.String("VaildIp")
		validIPLists := strings.Split(validIPs, ",")
		if base.refer == "" {
			count := 0
			for _, v := range validIPLists {
				if base.ip == v {
					count++
				}
			}
			if count == 0 {
				base.SendJSON(400, "", "my api do not for you")
			}
		}
	}
	//文章浏览记录
	if controllerPrefix == "Article" {
		postID, err := base.GetInt(":articleId")
		if err != nil {
			base.CheckError(err)
		}
		models.AutoSubPostView(postID)
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
