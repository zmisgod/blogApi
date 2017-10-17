package controllers

import (
	"blogapi/models"
	"strconv"

	"github.com/astaxie/beego"
)

type HomeController struct {
	beego.Controller
}

// @Get Article
// @Description find Article by page
// @Param	page		query 	string	true		"the page you want to get"
// @Success 200 {object} models.Posts
// @Failure 403 :page is empty
// @router /:page [get]
func (h *HomeController) Get() {
	var (
		list     []*models.Posts
		pagesize int
		err      error
		page     int
	)

	if page, err = strconv.Atoi(h.Ctx.Input.Param(":page")); err != nil || page < 1 {
		page = 1
	}
	pagesize = 12
	query := new(models.Posts).Query().Filter("post_status", "publish")
	count, _ := query.Count()
	if count > 0 {
		query.Limit(pagesize, (page-1)*pagesize).All(&list)
	}
	h.Data["json"] = list
	h.ServeJSON()
}
