package controllers

import (
	"fmt"

	"github.com/zmisgod/blogApi/models"
)

type UserController struct {
	BaseController
}

//@router / [post]
func (h *UserController) Login() {
	res, err := models.CheckUserExists("zmisgod", "111111", "111")
	if err != nil {
		h.SendError("error")
	}
	h.SendData(res, "ok")
}

//@router /login [get]
func (h *UserController) Register() {
	authority := h.GetString("test")
	res := models.GenerateUserAuth(1)
	if authority != "" {
		models.CheckUserAuth(authority)
	}
	fmt.Println(res)
	h.SendData("ok", res)
}
