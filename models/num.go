package models

import (
	"database/sql"
	"fmt"
)

//Num wps_post_nums
type Num struct {
	postID     int
	ViewNum    int `json:"view_num"`
	LikeNum    int `json:"like_num"`
	DislikeNum int `json:"dislike_num"`
	CommentNum int `json:"comment_num"`
}

//GetArticleNumsByPost 获取文章的数量详情
func GetArticleNumsByPost(postID int) (Num, error) {
	var (
		rows *sql.Rows
		err  error
	)
	// var commentList CommentLists
	postNum := Num{}

	rows, err = dbConn.Query(fmt.Sprintf("select post_id,view_num, like_num, dislike_num, comment_num from wps_post_nums where post_id = %d", postID))
	defer rows.Close()
	if err != nil {
		return postNum, err
	}
	for rows.Next() {
		err = rows.Scan(
			&postNum.postID,
			&postNum.ViewNum,
			&postNum.LikeNum,
			&postNum.DislikeNum,
			&postNum.CommentNum,
		)
		if err != nil {
			return postNum, err
		}
	}
	return postNum, nil
}
