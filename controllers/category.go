package controllers

import (
	"blogapi/models"
	"strconv"
)

type CategoryController struct {
	BaseController
}

//@router /:categoryId [get]
func (t *CategoryController) Get() {
	var (
		page       int
		err        error
		categoryID int
	)
	if categoryID, err = strconv.Atoi(t.Ctx.Input.Param(":categoryId")); err != nil {
		t.SendJSON(404, "", "invalid params")
	}
	if page, err = t.GetInt("page"); err != nil || page < 1 {
		page = 1
	}
	pageSize := 15
	res, err := models.TagAll(categoryID, page, pageSize, "category")
	t.CheckError(err)
	t.SendJSON(200, res, "successful")
}
