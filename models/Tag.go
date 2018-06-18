package models

import (
	"database/sql"
	"fmt"
	"time"
)

//Tag 文章标签
type Tag struct {
	TagID   int    `json:"tag_id"`
	TagName string `json:"tag_name"`
}

//GetPostTagLists 文章tag
func GetPostTagLists(postID int) ([]Tag, error) {
	var (
		rows *sql.Rows
		err  error
	)
	// var Post CommentLists
	tagList := []Tag{}

	rows, err = dbConn.Query(fmt.Sprintf("select t.tag_id,t.name as tag_name from wps_post_tags as pt left join wps_tags as t on pt.tag_id = t.tag_id where pt.post_id = %d and pt.disabled = 0", postID))
	if err != nil {
		return tagList, err
	}

	for rows.Next() {
		var aTag Tag
		err = rows.Scan(
			&aTag.TagID,
			&aTag.TagName,
		)
		if err != nil {
			continue
		}
		tagList = append(tagList, aTag)
	}
	return tagList, nil
}

//GetArticleListsByTagID 根据分类id获取文章列表
func GetArticleListsByTagID(cateID, page, pageSize int) ([]PostList, error) {
	var (
		rows *sql.Rows
		err  error
	)
	// var postList Post
	postList := []PostList{}

	rows, err = dbConn.Query(fmt.Sprintf("select p.id,p.post_title,u.name,c.c_name,p.post_title,p.post_intro,p.comment_status,p.comment_count from wps_post_tags as pt left join wps_posts as p on pt.post_id = p.id left join wps_users as u on p.user_id = u.id left join wps_post_categories as c on c.id = p.cat_id where p.post_status = 1 and pt.tag_id = %d order by p.created_at desc limit %d,%d", cateID, page, pageSize))
	if err != nil {
		return postList, err
	}

	for rows.Next() {
		var aPost PostList
		err = rows.Scan(
			&aPost.ID,
			&aPost.PostTitle,
			&aPost.UserName,
			&aPost.CategoryName,
			&aPost.PostTitle,
			&aPost.PostIntro,
			&aPost.CommentCount,
			&aPost.createdAt,
		)
		tm := time.Unix(int64(aPost.createdAt), 0)
		aPost.CreatedAt = tm.Format("2006-01-02 03:04")
		tags, _ := GetPostTagLists(aPost.ID)
		aPost.Tags = tags
		postList = append(postList, aPost)
	}
	return postList, nil
}
