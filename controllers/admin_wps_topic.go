package controllers

import (
	"strconv"

	"github.com/zmisgod/blogApi/models"
	"github.com/zmisgod/blogApi/util"
)

//AdminWpsTopicController 后台专栏管理
type AdminWpsTopicController struct {
	AdminController
}

//@router /list [get]
func (a *AdminWpsTopicController) List() {
	var search models.AdminTopicListSearch
	search.Page, _ = a.GetInt("page")
	if search.Page <= 0 {
		search.Page = 1
	}
	search.ID, _ = a.GetInt("id")
	search.PageSize, _ = a.GetInt("page_size")
	if search.PageSize <= 0 || search.PageSize >= 30 {
		search.PageSize = 15
	}
	//当前用户id
	userID, _ := strconv.Atoi(a.userInfo["userID"].(string))
	//搜索的用户id
	search.UserID, _ = a.GetInt("user_id")

	search.OrderType = a.GetString("order_type")
	if !util.InArraySting(search.OrderType, []string{"asc", "desc"}) {
		search.OrderType = "desc"
	}

	search.OrderbyName = a.GetString("order_by_name")
	if !util.InArraySting(search.OrderbyName, util.GetListMapValue(util.TopicOrderName(), "value")) {
		search.OrderbyName = "id"
	}

	lists, count, searchDatalist := models.AdminGetTopicLists(search, userID)
	if len(lists) > 0 {
		pagination := util.CombinePagination(search.Page, search.PageSize, count, search.OrderbyName, search.OrderType)
		a.SendDataLists(lists, pagination, searchDatalist, "ok")
	} else {
		a.SendDataError(searchDatalist, "没有相关数据")
	}
}

//@router /:id [get]
func (a *AdminWpsTopicController) One() {
	var (
		id  int
		err error
	)
	id, err = a.GetInt(":id")
	if err != nil {
		a.SendError(err.Error())
	}
	userID, _ := strconv.Atoi(a.userInfo["userID"].(string))
	articleInfo, err := models.AdminGetTopicByID(id, userID)
	if err != nil {
		a.SendError(err.Error())
	} else {
		a.SendData(articleInfo, "ok")
	}
}
