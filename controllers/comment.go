package controllers

import (
	"strconv"

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
		orderby   string
	)
	if articleID, err = strconv.Atoi(com.Ctx.Input.Param(":articleId")); err != nil {
		com.SendJSON(400, "", "invalid params")
	}

	orderby = "comment_ID desc"

	res, err := models.GetArticleCommentLists(articleID, com.page, com.pageSize, orderby)
	com.CheckError(err)
	com.SendJSON(200, res, "ok")
}

//@router /:articleId [post]
func (com *CommentController) Post() {
	var (
		content   string
		err       error
		articleID int
	)
	if articleID, err = strconv.Atoi(com.Ctx.Input.Param(":articleId")); err != nil {
		com.SendJSON(400, "", "invalid params")
	}

	commentID, err := com.GetInt("comment_id")
	if err != nil {
		com.SendJSON(400, "error", "invalid comment_id params")
	}

	content = com.GetString("conetnt")
	if len(content) == 0 {
		com.SendJSON(400, "", "empty content")
	}

	authorURL := com.GetString("auithor_url")

	authorEmail := com.GetString("author_email")

	authorName := com.GetString("author_name")
	if authorName == "" {
		com.SendJSON(400, "error", "empty authorName params")
	}

	authorAgent := com.Ctx.Input.Cookie("User-Agent")
	if authorAgent == "" {
		com.SendJSON(400, "error", "do not try to post anything")
	}

	authorIP := com.Ctx.Input.IP()

	num := models.SaveArticleComment(articleID, commentID, authorName, authorEmail, authorURL, content, authorIP, authorAgent)
	com.SendJSON(200, num, "ok")
}
