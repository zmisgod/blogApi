package controllers

import (
	"blogapi/models"
	"strconv"

	"github.com/astaxie/beego"
)

type HomeController struct {
	beego.Controller
	ResponseData
}

type ResponseData struct {
	code     int
	page     int
	pagesize int
	data     interface{}
	msg      string
}

// @Get All Article
// @Description find Article by page
// @Param	page		query 	string	true		"the page you want to get"
// @Success 200 {object} ResponseData
// @Failure 403 :page is empty
// @router / [get]
func (h *HomeController) GetAll() {
	var (
		list []*models.Posts
		err  error
		page int
	)

	if page, err = h.GetInt("page"); err != nil || page < 1 {
		page = 1
	}
	pagesize := 12
	query := new(models.Posts).Query().Filter("post_status", "publish")
	count, _ := query.Count()
	if count > 0 {
		query.OrderBy("-id").Limit(pagesize, (page-1)*pagesize).All(&list)
	}
	h.Data["json"] = list
	h.ServeJSON()
}

//@router /:articleId [get]
func (h *HomeController) Get() {
	var (
		err       error
		articleId int
		article   models.Posts
	)
	if articleId, err = strconv.Atoi(h.Ctx.Input.Param(":articleId")); err != nil {
		h.Abort("403")
	} else {
		article, err = models.OneArticle(articleId)
		if err != nil {
			h.Abort("403")
		} else {
			h.Data["json"] = article
		}
		h.ServeJSON()
	}
}
