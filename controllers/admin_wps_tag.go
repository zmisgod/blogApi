package controllers

import (
	"github.com/zmisgod/blogApi/models"
	"github.com/zmisgod/blogApi/util"
)

//AdminWpsTagController 后台分类管理
type AdminWpsTagController struct {
	AdminController
}

//@router /list [get]
func (a *AdminWpsTagController) List() {
	var search models.AdminTagListSearch
	search.Page, _ = a.GetInt("page")
	if search.Page <= 0 {
		search.Page = 1
	}
	search.PageSize, _ = a.GetInt("page_size")
	if search.PageSize <= 0 || search.PageSize >= 30 {
		search.PageSize = 15
	}

	search.OrderType = a.GetString("order_type")
	if !util.InArraySting(search.OrderType, []string{"asc", "desc"}) {
		search.OrderType = "desc"
	}

	search.OrderbyName = a.GetString("order_by_name")
	if !util.InArraySting(search.OrderbyName, util.GetListMapValue(util.TagOrderName(), "value")) {
		search.OrderbyName = "tag_id"
	}

	lists, count, searchDatalist := models.AdminGetTagLists(search)
	if len(lists) > 0 {
		pagination := util.CombinePagination(search.Page, search.PageSize, count, search.OrderbyName, search.OrderType)
		a.SendDataLists(lists, pagination, searchDatalist, "ok")
	} else {
		a.SendDataError(searchDatalist, "没有相关数据")
	}
}
