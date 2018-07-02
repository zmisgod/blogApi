package models

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/astaxie/beego"
)

// Topics 主题
type Topics struct {
	ID        int `json:"id"`
	userID    int
	Title     string `json:"title"`
	ImgURL    string `json:"img_url"`
	Content   string `json:"content"`
	sort      int
	createdAt int
	CreatedAt string `json:"created_at"`
}

//GetTopicsArticleLists 获取文章主题的文章列表
func GetTopicsArticleLists(topicsID, page, pageSize int) ([]PostList, error) {
	var (
		rows *sql.Rows
		err  error
	)
	var postList []PostList
	offset := (page - 1) * pageSize
	rows, err = dbConn.Query(fmt.Sprintf("select p.id,p.user_id,p.post_title,u.name   as user_name,c.c_name as category_name,p.post_title,p.post_intro,p.created_at from wps_special_topic as st left join wps_special_topic_lists as stl on st.id = stl.top_id left join wps_posts as p on p.id = stl.post_id left join wps_users as u on p.user_id = u.id left join wps_post_categories as c on c.id = p.cat_id where st.id = %d and p.post_status = 1 and st.disabled = 0 and stl.disabled = 0 order by stl.sort asc limit %d, %d", topicsID, offset, pageSize))
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
			&aPost.createdAt,
		)
		if err != nil {
			continue
		}
		tm := time.Unix(int64(aPost.createdAt), 0)
		aPost.CreatedAt = tm.Format("2006-01-02 15:04")

		tags, _ := GetPostTagLists(aPost.ID)
		aPost.Tags = tags

		num, _ := GetArticleNumsByPost(aPost.ID)
		aPost.NumInfo = num

		postList = append(postList, aPost)
	}
	return postList, nil
}

//GetTopicLists 获取文章主题的列表
func GetTopicLists(page, pageSize int) ([]Topics, error) {
	var (
		rows *sql.Rows
		err  error
	)
	// var postList Post
	topicsList := []Topics{}
	offset := (page - 1) * pageSize
	rows, err = dbConn.Query(fmt.Sprintf("select id,title,img_url,content,created_at from wps_special_topic order by sort asc limit %d,%d", offset, pageSize))
	if err != nil {
		return topicsList, err
	}

	for rows.Next() {
		var aTopics Topics
		err = rows.Scan(
			&aTopics.ID,
			&aTopics.Title,
			&aTopics.ImgURL,
			&aTopics.Content,
			&aTopics.createdAt,
		)
		if aTopics.ImgURL != "" {
			aTopics.ImgURL = beego.AppConfig.String("StaticPrefix") + aTopics.ImgURL
		}
		tm := time.Unix(int64(aTopics.createdAt), 0)
		aTopics.CreatedAt = tm.Format("2006-01-02 15:04")
		topicsList = append(topicsList, aTopics)
	}
	return topicsList, nil
}
