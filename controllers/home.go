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
		list    []*models.Posts
		article models.Posts
		err     error
		page    int
	)

	if page, err = h.GetInt("page"); err != nil || page < 1 {
		page = 1
	}
	pagesize := 12
	query := article.Query().Filter("post_status", "publish")
	count, _ := query.Count()
	if count > 0 {
		query.OrderBy("-id").Limit(pagesize, (page-1)*pagesize).All(&list)
	}
	result := make([]map[string]interface{}, len(list))

	for k, v := range list {
		row := make(map[string]interface{})
		row["id"] = v.Id
		row["post_author"] = v.PostAuthor
		row["post_status"] = v.PostStatus
		row["comment_count"] = v.CommentCount
		row["post_date"] = v.PostDate
		row["post_intro"] = v.PostIntro
		result[k] = row
	}
	h.SendJSON(200, result, "successful")
}

//@router /:articleId [get]
func (h *HomeController) Get() {
	var (
		err       error
		articleId int
		article   models.Posts
	)
	if articleId, err = strconv.Atoi(h.Ctx.Input.Param(":articleId")); err != nil {
		h.SendJSON(400, "", "invalid params")
	} else {
		article, err = models.OneArticle(articleId)
		if err != nil {
			h.SendJSON(400, "", "empty data")
		} else {
			h.SendJSON(200, article, "successful")
		}
	}
}
