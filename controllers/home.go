package controllers

import (
	"blogapi/models"
	"strconv"
)

type HomeController struct {
	BaseController
}

// @Get All Article
// @Description find Article by page
// @Param	page		query 	string	true		"the page you want to get"
// @Success 200 {object} ResponseData
// @Failure 403 :page is empty
// @router / [get]
func (h *HomeController) GetAll() {
	var (
		err  error
		page int
	)
	if page, err = h.GetInt("page"); err != nil || page < 1 {
		page = 1
	}
	pagesize := 12
	resultss, err := models.GetArticleLists(page, pagesize)
	h.CheckError(err)
	h.SendJSON(200, resultss, "successful")
}

//@router /:articleId [get]
func (h *HomeController) Get() {
	var (
		err       error
		articleID int
	)
	if articleID, err = strconv.Atoi(h.Ctx.Input.Param(":articleId")); err != nil {
		h.SendJSON(400, "", "invalid params")
	} else {
		res, err := models.ArticleOne(articleID)
		h.CheckError(err)
		data, ok := res.(map[string]interface{})
		if ok {
			tags, err := models.ArticleTags(articleID)
			if err != nil {
				tags = make([]interface{}, 0)
			}
			data["tag"] = tags
		}
		h.SendJSON(200, res, "successful")
	}
}
