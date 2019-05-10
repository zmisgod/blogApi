package controllers

import (
	"encoding/json"
	"io/ioutil"

	"github.com/zmisgod/blogApi/models"
)

type UserController struct {
	BaseController
}

//LoginParams 登录传递的参数
type LoginParams struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

//@router / [post]
func (h *UserController) Login() {
	ioReader := h.Ctx.Request.Body
	bytes, err := ioutil.ReadAll(ioReader)
	if err != nil {
		h.SendError("empty data")
	}
	var loginParams LoginParams
	err = json.Unmarshal(bytes, &loginParams)
	if err != nil {
		h.SendError("邮箱或密码为空，请重试")
	}
	res, err := models.CheckUserExists(loginParams.Email, loginParams.Password)
	if err != nil {
		h.SendError("error")
	}
	h.SendData(res, "ok")
}
