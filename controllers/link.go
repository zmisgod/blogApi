package controllers

import (
	"github.com/zmisgod/blogApi/models"
)

type LinkController struct {
	BaseController
}

//@router / [get]
func (h *LinkController) Get() {
	var (
		err error
	)
	lists, err := models.GetLinks()
	h.CheckError(err)
	h.SendJSON(200, lists, "successful")
}
