package models

import (
	"database/sql"
	"fmt"
	"time"
)

//Post 文章基础结构
type Post struct {
	ID           int    `json:"id"`
	UserID       int    `json:"user_id"`
	UserName     string `json:"user_name"`
	CategoryName string `json:"category_name"`
	PostTitle    string `json:"post_title"`
	PostIntro    string `json:"post_intro"`
	CommentCount int    `json:"comment_count"`
	CoverURL     string `json:"cover_url"`
	createdAt    int
	CreatedAt    string `json:"created_at"`
	Contents     string `json:"contents"`
	Tags         []Tag  `json:"tags"`
}

//PostDetail 文章详情
type PostDetail struct {
	Post
	UserInfo User `json:"user_info"`
}

//PostList 文章列表
type PostList struct {
	Post
}

//GetArticleLists 获取文章列表
func GetArticleLists(page, pageSize int) ([]PostList, error) {
	var (
		rows *sql.Rows
		err  error
	)
	// var Post CommentLists
	postList := []PostList{}
	offset := (page - 1) * pageSize
	rows, err = dbConn.Query(fmt.Sprintf("select p.id,p.user_id, p.post_title,u.name as user_name,c.c_name as category_name,p.post_title,p.post_intro,p.comment_count,p.created_at from wps_posts as p left join wps_users as u on p.user_id = u.id left join wps_post_categories as c on c.id = p.cat_id where p.post_status = 1 order by p.created_at desc limit %d,%d", page, offset))
	if err != nil {
		return postList, err
	}

	for rows.Next() {
		var aPost PostList
		err = rows.Scan(
			&aPost.ID,
			&aPost.UserID,
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

//GetArticleDetail 获取文章详情
func GetArticleDetail(postID int) (PostDetail, error) {
	var (
		rows *sql.Rows
		err  error
	)
	var post PostDetail
	rows, err = dbConn.Query(fmt.Sprintf("select p.id,p.user_id, p.post_title,u.name as user_name,c.c_name as category_name,p.post_title,p.post_intro,p.comment_count,p.created_at,pc.contents from wps_posts as p left join wps_users as u on p.user_id = u.id left join wps_post_categories as c on c.id = p.cat_id left join wps_post_contents as pc on pc.id = p.id where p.post_status = 1 and p.id  = %d", postID))
	if err != nil {
		return post, err
	}

	for rows.Next() {
		err = rows.Scan(
			&post.ID,
			&post.UserID,
			&post.PostTitle,
			&post.UserName,
			&post.CategoryName,
			&post.PostTitle,
			&post.PostIntro,
			&post.CommentCount,
			&post.createdAt,
			&post.Contents,
		)
		if err != nil {
			continue
		}
		tm := time.Unix(int64(post.createdAt), 0)
		post.CreatedAt = tm.Format("2006-01-02 03:04")
		tags, _ := GetPostTagLists(post.ID)
		post.Tags = tags
		userInfo, _ := GetUserInfo(post.UserID)
		post.UserInfo = userInfo
		break
	}
	return post, nil
}
