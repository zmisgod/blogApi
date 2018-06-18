package models

import (
	"database/sql"
	"fmt"
	"time"
)

// Comment wps_comments
type Comment struct {
	ID          int    `json:"id"`
	UserID      int    `json:"user_id"`
	Content     string `json:"content"`
	AuthorName  string `json:"author_name"`
	AuthorEmail string `json:"author_email"`
	AuthorURL   string `json:"author_url"`
	createdAt   int
	CreatedAt   string `json:"created_at"`
}

//GetArticleCommentLists 获取文章评论列表
func GetArticleCommentLists(postID, page, pageSize int, orderby string) ([]Comment, error) {
	var (
		rows *sql.Rows
		err  error
	)
	// var commentList CommentLists
	commentList := []Comment{}

	offset := (page - 1) * pageSize
	rows, err = dbConn.Query(fmt.Sprintf("select created_at,author_email,author_name,author_url,content,id,user_id from wps_comments where post_id = %d order by created_at desc limit %d,%d", postID, offset, pageSize))
	if err != nil {
		return commentList, err
	}
	for rows.Next() {
		var aComment Comment
		err = rows.Scan(
			&aComment.createdAt,
			&aComment.AuthorEmail,
			&aComment.AuthorName,
			&aComment.AuthorURL,
			&aComment.Content,
			&aComment.ID,
			&aComment.UserID,
		)
		if err != nil {
			continue
		}
		tm := time.Unix(int64(aComment.createdAt), 0)
		aComment.CreatedAt = tm.Format("2006-01-02 03:04")
		commentList = append(commentList, aComment)
	}
	return commentList, nil
}
