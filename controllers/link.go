package controllers

import (
	"github.com/zmisgod/blogApi/models"
)

type LinkController struct {
	BaseController
}

//@router / [get]
func (h *LinkController) Get() {
	lists, err := models.GetLinks()
	h.CheckError(err)
	h.SendData(lists, "successful")
}
