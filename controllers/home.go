package controllers

import (
	"fmt"

	"github.com/zmisgod/blogApi/models"
)

//HomeController homecontroller
type HomeController struct {
	BaseController
}

// @router / [get]
func (h *HomeController) Get() {
	var (
		err error
	)
	lists, err := models.GetArticleLists(h.page, h.pageSize)
	h.CheckError(err)
	h.SendJSON(200, lists, "successful")
}

//GetAll 获取所有
func (h *HomeController) GetAll() {
	fmt.Println("error")
}
