package controllers

import (
	"blogapi/models"
	"strconv"
)

type TagController struct {
	BaseController
}

//@router /:tagId [get]
func (t *TagController) Get() {
	var (
		page  int
		err   error
		tagID int
	)
	if tagID, err = strconv.Atoi(t.Ctx.Input.Param(":tagId")); err != nil {
		t.SendJSON(404, "", "invalid params")
	}
	if page, err = t.GetInt("page"); err != nil || page < 1 {
		page = 1
	}
	pageSize := 15
	res, err := models.TagAll(tagID, page, pageSize, "post_tag")
	t.CheckError(err)
	t.SendJSON(200, res, "successful")
}
