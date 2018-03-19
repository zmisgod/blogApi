package controllers

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/zmisgod/blogApi/models"
)

type CommentController struct {
	BaseController
}

//@router /:articleId [get]
func (com *CommentController) Get() {
	var (
		err       error
		articleID int
		commentID int
		orderby   string
	)
	if articleID, err = strconv.Atoi(com.Ctx.Input.Param(":articleId")); err != nil {
		com.SendJSON(400, "", "invalid params")
	}
	commentID, err = com.GetInt("comment_id")

	orderby = "comment_ID desc"

	res, err := models.GetArticleCommentLists(articleID, com.page, com.pageSize, orderby, commentID)
	com.CheckError(err)
	com.SendJSON(200, res, "ok")
}

//@router /comment [post]
func (com *CommentController) Post() {
	var (
		commentID      int
		commentContent string = strings.TrimSpace(com.GetString("comment_content"))
		articleID      int
		err            error
	)
	commentID, err = com.GetInt("comment_id")
	if err != nil {
		com.SendJSON(400, "error", "invalid comment_id params")
	}
	articleID, err = com.GetInt("article_id")
	if err != nil {
		com.SendJSON(400, "error", "invalid article_id params")
	}
	fmt.Println(commentID)
	fmt.Println(commentContent)
	fmt.Println(articleID)
	com.SendJSON(200, "comment_id", "ok")
}
