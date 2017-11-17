package controllers

import "blogapi/models"

type SearchController struct {
	BaseController
}

//@router / [get]
func (t *SearchController) Get() {
	var (
		page     int
		pageSize int
		keyword  string
		err      error
	)
	if keyword = t.GetString("keyword"); keyword == "" {
		t.SendJSON(403, "", "empty search keyword")
	}
	if page, err = t.GetInt("page"); err != nil || page < 1 {
		page = 1
	}
	pageSize = 12
	res, err := models.SphinxSearch(keyword, page, pageSize)
	if err != nil {
		t.SendJSON(400, "", "ok")
	}
	t.SendJSON(200, res, "ok")
}
