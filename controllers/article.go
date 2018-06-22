package controllers

import "github.com/zmisgod/blogApi/models"

//ArticleController articlecontroller
type ArticleController struct {
	BaseController
}

//@router /:articleId [get]
func (h *ArticleController) Get() {
	var (
		err error
	)
	postID, err := h.GetInt(":articleId")
	if err != nil {
		h.CheckError(err)
	}
	lists, err := models.GetArticleDetail(postID)
	h.CheckError(err)
	models.AutoSubPostView(postID)
	h.SendJSON(200, lists, "successful")
}
