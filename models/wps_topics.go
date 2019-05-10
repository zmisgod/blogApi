package models

import (
	"fmt"
	"strconv"

	"github.com/zmisgod/blogApi/util"
)

//AdminTopicListSearch 后台专题列表搜索
type AdminTopicListSearch struct {
	ID     int `json:"id"`
	UserID int `json:"user_id"`
	AdminBaseListSearch
}

// AdminWpsTopics 主题
type AdminWpsTopics struct {
	ID        int    `json:"id"`
	UserID    int    `json:"user_id"`
	Title     string `json:"title"`
	ImgURL    string `json:"img_url"`
	Content   string `json:"content"`
	Sort      int    `json:"sort"`
	CreatedAt int    `json:"created_at"`
	UpdatedAt int    `json:"updated_at"`
}

//AdminGetTopicLists 后台搜索主题列表
func AdminGetTopicLists(search AdminTopicListSearch, userID int) ([]AdminWpsTopics, int, map[string]interface{}) {
	var topics []AdminWpsTopics
	var count int
	searchParams := make(map[string]interface{}, 1)
	searchParams["orderNameList"] = util.TopicOrderName()
	whereCondition := ""
	if search.ID != 0 {
		whereCondition += " and id = " + strconv.Itoa(search.ID)
	}
	if CheckUserIDAuth(userID, "admin_topic_list_search_user_id") {
		if search.UserID != 0 {
			whereCondition += " and user_id = " + strconv.Itoa(search.UserID)
		}
	} else {
		whereCondition += " and user_id = " + strconv.Itoa(userID)
	}
	sql := fmt.Sprintf("select id, user_id, title, img_url,content,sort,created_at, updated_at from wps_special_topic where 1=1 %s order by %s %s limit %d,%d", whereCondition, search.OrderbyName, search.OrderType, (search.Page-1)*search.PageSize, search.PageSize)
	countSQL := fmt.Sprintf("select count(1) from wps_special_topic where 1=1 %s order by %s %s limit %d,%d", whereCondition, search.OrderbyName, search.OrderType, (search.Page-1)*search.PageSize, search.PageSize)
	rows, err := dbConn.Query(sql)
	defer rows.Close()
	if err != nil {
		return topics, count, searchParams
	}
	for rows.Next() {
		var topic AdminWpsTopics
		err = rows.Scan(&topic.ID, &topic.UserID, &topic.Title, &topic.ImgURL, &topic.Content, &topic.Sort, &topic.CreatedAt, &topic.UpdatedAt)
		if err != nil {
			continue
		}
		topics = append(topics, topic)
	}
	dbConn.QueryRow(countSQL).Scan(&count)
	return topics, count, searchParams
}

//AdminTopicPost 后台主题文章
type AdminTopicPost struct {
	ID          int    `json:"id"`
	PostTitle   string `json:"post_title"`
	PostIntro   string `json:"post_intro"`
	PublishTime string `json:"publish_time"`
}

//AdminGetTopicByID topic列表
func AdminGetTopicByID(topicID, userID int) (map[string]interface{}, error) {
	var topic AdminWpsTopics
	var articleList []AdminTopicPost
	result := make(map[string]interface{})
	sql := fmt.Sprintf("select id, user_id, title, img_url,content,sort,created_at, updated_at from wps_special_topic where id = %d and user_id = %d", topicID, userID)
	err := dbConn.QueryRow(sql).Scan(&topic.ID, &topic.UserID, &topic.Title, &topic.ImgURL, &topic.Content, &topic.Sort, &topic.CreatedAt, &topic.UpdatedAt)
	if err != nil {
		return result, err
	}
	sql = fmt.Sprintf("select p.id,p.post_title,p.post_intro,p.publish_time from wps_special_topic_lists as l left join wps_posts as p on p.id = l.post_id where l.topic_id = %d", topicID)
	rows, err := dbConn.Query(sql)
	defer rows.Close()
	if err != nil {
		fmt.Println(err)
		return result, err
	}
	for rows.Next() {
		var post AdminTopicPost
		err = rows.Scan(&post.ID, &post.PostTitle, &post.PostIntro, &post.PublishTime)
		if err != nil {
			continue
		}
		articleList = append(articleList, post)
	}
	result["info"] = topic
	result["list"] = articleList
	return result, nil
}
