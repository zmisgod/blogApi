package controllers

import "blogapi/models"

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
		t.SendJSON(403, "", "empty search keyword")
	}
	res, err := models.SphinxSearch(keyword, t.page, t.pageSize)
	t.CheckError(err)
	t.SendJSON(200, res, "ok")
}
