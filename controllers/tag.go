package controllers

import "github.com/zmisgod/blogApi/models"

type TagController struct {
	BaseController
}

//@router /:tagId [get]
func (h *TagController) Get() {
	var (
		err error
	)
	tagID, err := h.GetInt(":tagId")
	if err != nil {
		h.CheckError(err)
	}
	lists, err := models.GetArticleListsByTagID(tagID, h.page, h.pageSize)
	h.CheckError(err)
	h.SendJSON(200, lists, "successful")
}
