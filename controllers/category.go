package controllers

import (
	"github.com/zmisgod/blogApi/models"
)

//CategoryController 分类的控制器
type CategoryController struct {
	BaseController
}

//@router /:categoryId [get]
func (h *CategoryController) Get() {
	cateID, err := h.GetInt(":categoryId")
	if err != nil {
		h.CheckError(err)
	}
	lists, err := models.GetArticleListsByCategoryID(cateID, h.page, h.pageSize)
	h.CheckError(err)
	h.SendData(lists, "successful")
}
