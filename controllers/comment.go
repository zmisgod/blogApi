package controllers

import (
	"encoding/json"
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

//CommentStruct 保存评论的数据结构
type CommentStruct struct {
	CommentID   int    `json:"comment_id"`
	Content     string `json:"content"`
	AuthorURL   string `json:"author_url"`
	AuthorEmail string `json:"author_email"`
	AuthorName  string `json:"author_name"`
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

	var commentRequest CommentStruct
	json.Unmarshal(com.Ctx.Input.RequestBody, &commentRequest)

	content = commentRequest.Content
	if len(content) < 15 {
		com.SendJSON(400, "error", "评论至少15字")
	}

	authorName := commentRequest.AuthorName
	if authorName == "" {
		com.SendJSON(400, "error", "empty authorName params")
	}

	authorAgent := com.userAgent
	if authorAgent == "" {
		com.SendJSON(400, "error", "do not try to post anything")
	}

	authorIP := com.Ctx.Input.IP()

	num := models.SaveArticleComment(articleID, commentRequest.CommentID, authorName, commentRequest.AuthorEmail, commentRequest.AuthorURL, content, authorIP, authorAgent)
	com.SendJSON(200, num, "ok")
}
