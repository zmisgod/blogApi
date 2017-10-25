package models

import (
	"time"
)

type postDetail struct {
	ID            int       `json:"id"`
	post_title    string    `json:"post_title"`
	post_author   string    `json:"post_author"`
	post_status   string    `json:"post_status"`
	comment_count int       `json:"comment_count"`
	post_date     time.Time `json:"post_date"`
	post_intro    string    `json:"post_intro"`
	post_content  string    `json:"post_content"`
}

func ArticleAll(page, pagesize int) (interface{}, error) {
	rows, _ := dbConn.Query(
		"select ID,post_title,post_author,post_status,comment_count,post_date,post_intro from wps_posts where post_status = 'publish' order by ID DESC limit ? offset ?",
		pagesize,
		(page-1)*pagesize,
	)
	return DBQueryRows(rows)
}

func ArticleOne(articleId int, args ...string) (interface{}, error) {
	row, _ := dbConn.Query("select ID,post_title,post_author,post_status,comment_count,post_date,post_intro,post_content  from wps_posts where ID = ? AND post_status= ?", articleId, "publish")
	return DBQueryRow(row)
}
