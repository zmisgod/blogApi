package controllers

import (
	"strconv"

	"github.com/zmisgod/blogApi/models"
	"github.com/zmisgod/blogApi/util"
)

//AdminController 后台用户验证
type AdminController struct {
	BaseController
}

//Prepare 用户验证
func (base *AdminController) Prepare() {
	//身份验证
	base.authority = base.Ctx.Input.Header("Authorization")
	if base.authority == "" || base.authority == "null" {
		base.SendAuthRequire("请登录")
		return
	}
	checkAuth, userInfo := models.CheckUserAuth(base.authority)
	if !checkAuth {
		base.SendAuthRequire("请登录")
		return
	}
	loginTimeStr := userInfo["loginTime"].(string)
	loginTime, err := strconv.Atoi(loginTimeStr)
	if err != nil {
		base.SendAuthRequire("请登录")
		return
	}
	notExpire := util.CheckAuthNotExpire(loginTime)
	if !notExpire {
		base.SendAuthRequire("请登录")
		return
	}
	base.userInfo = userInfo
	base.ip = base.Ctx.Input.IP()
	base.userAgent = base.Ctx.Input.UserAgent()
	base.requestURI = base.Ctx.Input.URI()
	base.refer = base.Ctx.Input.Refer()
}
