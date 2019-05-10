package controllers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strconv"

	"github.com/zmisgod/blogApi/models"
	"github.com/zmisgod/blogApi/util"
)

//WpsArticleController 文章
type AdminWpsArticleController struct {
	AdminController
}

//@router /list [get]
func (a *AdminWpsArticleController) List() {
	var search models.PostListsSearch
	search.Page, _ = a.GetInt("page")
	if search.Page <= 0 {
		search.Page = 1
	}
	search.UserID, _ = a.GetInt("user_id")
	userID, _ := strconv.Atoi(a.userInfo["userID"].(string))

	search.PostStatus, _ = a.GetInt("post_status")
	search.PostType, _ = a.GetInt("post_type")
	search.CommentStatus, _ = a.GetInt("comment_status")

	search.PageSize, _ = a.GetInt("page_size")
	if search.PageSize <= 0 || search.PageSize >= 30 {
		search.PageSize = 15
	}
	search.TagID, _ = a.GetInt("tag_id")

	search.OrderType = a.GetString("order_type")
	if !util.InArraySting(search.OrderType, []string{"asc", "desc"}) {
		search.OrderType = "desc"
	}

	search.OrderbyName = a.GetString("order_by_name")
	if !util.InArraySting(search.OrderbyName, util.GetListMapValue(util.PostOrderName(), "value")) {
		search.OrderbyName = "id"
	}
	lists, count, searchDatalist := models.AdminGetArticleLists(search, userID)
	if len(lists) > 0 {
		pagination := util.CombinePagination(search.Page, search.PageSize, count, search.OrderbyName, search.OrderType)
		a.SendDataLists(lists, pagination, searchDatalist, "ok")
	} else {
		a.SendDataError(searchDatalist, "没有相关数据")
	}
}

//@router /save [post]
func (a *AdminWpsArticleController) Save() {
	ioReader := a.Ctx.Request.Body
	bytes, err := ioutil.ReadAll(ioReader)
	if err != nil {
		a.SendError("empty data")
	}
	var pass models.PassPost
	err = json.Unmarshal(bytes, &pass)
	if err != nil {
		a.SendError("参数验证失败")
	}
	userIDStr := a.userInfo["userID"].(string)
	userID, _ := strconv.Atoi(userIDStr)
	postID, err := models.SaveArticle(pass, userID)
	if err != nil {
		fmt.Println(err)
		a.SendError(err.Error())
	} else {
		a.SendData(postID, "ok")
	}
}

//@router /:id [get]
func (a *AdminWpsArticleController) One() {
	var (
		id  int
		err error
	)
	id, err = a.GetInt(":id")
	if err != nil {
		a.CheckError(err)
	}
	articleInfo, searchDatalist := models.AdminGetArticleByID(id)
	if articleInfo.ID > 0 {
		a.SendDataLists(articleInfo, nil, searchDatalist, "ok")
	} else {
		a.SendDataError(searchDatalist, "error")
	}
}

//@router /:id [post]
func (a *AdminWpsArticleController) AOne() {
	a.SendError("ok")
}
