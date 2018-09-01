package controllers

import (
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
	h.SendData(lists, "successful")
}
