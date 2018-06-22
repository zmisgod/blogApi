package models

import (
	"database/sql"
	"fmt"
	"time"
)

//GetArticleListsByCategoryID 根据分类id获取文章列表
func GetArticleListsByCategoryID(cateID, page, pageSize int) ([]PostList, error) {
	var (
		rows *sql.Rows
		err  error
	)
	// var postList PostList
	postList := []PostList{}

	offset := (page - 1) * pageSize
	rows, err = dbConn.Query(fmt.Sprintf("select p.id,p.post_title,u.name as user_name,c.c_name as category_name,p.post_title,p.post_intro,p.created_at from wps_posts as p left join wps_users as u on p.user_id = u.id left join wps_post_categories as c on c.id = p.cat_id where p.post_status = 1 and p.cat_id = %d order by p.created_at desc limit %d,%d", cateID, offset, pageSize))
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
			&aPost.createdAt,
		)
		tm := time.Unix(int64(aPost.createdAt), 0)
		aPost.CreatedAt = tm.Format("2006-01-02 03:04")
		tags, _ := GetPostTagLists(aPost.ID)
		aPost.Tags = tags
		num, _ := GetArticleNumsByPost(aPost.ID)
		aPost.NumInfo = num
		postList = append(postList, aPost)
	}
	return postList, nil
}
