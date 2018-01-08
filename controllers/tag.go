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
		err   error
		tagID int
	)
	if tagID, err = strconv.Atoi(t.Ctx.Input.Param(":tagId")); err != nil {
		t.SendJSON(404, "", "invalid params")
	}
	res, err := models.TagAll(tagID, t.page, t.pageSize, "post_tag")
	t.CheckError(err)
	t.SendJSON(200, res, "successful")
}
