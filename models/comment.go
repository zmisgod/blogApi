package models

import "database/sql"

type Comment struct {
	CommentAuthor  string `json:"comment_author"`
	CommentDateGmt string `json:"comment_date_gmt"`
	CommentContent string `json:"comment_content"`
	CommentID      int    `json:"comment_ID"`
	CommentKarma   int    `json:"comment_karma"`
}

type CommentLists struct {
	Comments []Comment
}

func GetArticleCommentLists(articleId, page, pageSize int, orderby string, commentId int) (interface{}, error) {
	var (
		rows *sql.Rows
		err  error
	)

	rows, err = dbConn.Query(
		"select comment_author , comment_date_gmt, comment_content, comment_ID,comment_karma from wps_comments where comment_post_ID  = ? and comment_parent= ? order by ? limit ? offset ?",
		articleId,
		commentId,
		orderby,
		pageSize,
		(page-1)*pageSize,
	)

	if err != nil {
		return "", err
	}
	return DBQueryRows(rows)
}
