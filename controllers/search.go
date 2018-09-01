package controllers

import "github.com/zmisgod/blogApi/models"

type SearchController struct {
	BaseController
}

//@router / [get]
func (t *SearchController) Get() {
	var (
		keyword string
		err     error
	)
	if keyword = t.GetString("keyword"); keyword == "" {
		t.SendError("empty search keyword")
	}
	res, err := models.SphinxSearch(keyword, t.page, t.pageSize)
	t.CheckError(err)
	t.SendData(res, "ok")
}
