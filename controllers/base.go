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
	base.authority = base.Ctx.Input.Header("Authorization")
	if base.authority != "" {
		checkAuth := models.CheckUserAuth(base.authority)
		if !checkAuth {
			base.SendAuthRequire("请登录")
		}
	}
	base.requestURI = base.Ctx.Input.URI()
	base.refer = base.Ctx.Input.Refer()
	c, _ := base.GetControllerAndAction()
	controllerPrefix := strings.Replace(c, "Controller", "", 10)
	devMode := beego.AppConfig.String("runmode")
	if devMode != "dev" {
		saveLog := true
		if controllerPrefix != "Crh" {
			validIPs := beego.AppConfig.String("VaildIp")
			validIPLists := strings.Split(validIPs, ",")
			if base.refer == "" {
				count := 0
				for _, v := range validIPLists {
					if base.ip == v {
						count++
						saveLog = false
					}
				}
				if count == 0 {
					base.SendError("my api do not for you")
				}
			}
		}
		if saveLog {
			//用户请求日志
			models.SaveUserVisiteHistory(controllerPrefix, base.ip, base.userAgent, base.requestURI, base.refer)
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

//SendData 发送成功数据
func (base *BaseController) SendData(data interface{}, msg string) {
	base.sendJSON(200, data, msg)
}

//sendJSON 返送json
func (base *BaseController) sendJSON(code int, data interface{}, msg string) {
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
		base.SendError(err.Error())
	}
}

//SendError 参数验证失败
func (base *BaseController) SendError(msg string) {
	base.sendJSON(400, "", msg)
}

//SendAuthRequire 发送用户身份认证请求
func (base *BaseController) SendAuthRequire(msg string) {
	base.sendJSON(401, "", msg)
}

//SendInternalError 发送5xx错误，统称code 500
func (base *BaseController) SendInternalError(msg string) {
	base.sendJSON(500, "", msg)
}
