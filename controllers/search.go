package controllers

import "blogapi/models"
import "strconv"
import "fmt"

type SearchController struct {
	BaseController
}

//@router / [get]
func (t *SearchController) Get() {
	var (
		page     int
		pageSize int
		keyword  string
	)
	fmt.Println("_+_+_+")
	if keyword = t.Ctx.Input.Param(":keyword"); keyword == "" {
		t.SendJSON(403, "", "empty search keyword")
	}
	if page, err := strconv.Atoi(t.Ctx.Input.Param(":page")); err != nil || page < 1 {
		page = 1
	}
	if pageSize, err := strconv.Atoi(t.Ctx.Input.Param(":pageSize")); err != nil || pageSize < 1 {
		pageSize = 15
	}
	res, err := models.SphinxSearch(keyword, page, pageSize)
	if err != nil {
		t.SendJSON(400, "", "ok")
	}
	t.SendJSON(200, res, "ok")
}
