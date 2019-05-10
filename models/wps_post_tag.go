package models

import (
	"database/sql"
	"fmt"
	"strconv"

	"github.com/zmisgod/blogApi/util"
)

//AdminTagListSearch 后台标签列表搜索
type AdminTagListSearch struct {
	TagID int    `json:"tag_id"`
	Name  string `json:"name"`
	AdminBaseListSearch
}

//WpsTag 文章标签
type WpsTag struct {
	TagID   int    `json:"tag_id"`
	TagName string `json:"tag_name"`
}

//AdminWpsTag 后台专用数据结构
type AdminWpsTag struct {
	TagID  int    `json:"tag_id"`
	Name   string `json:"name"`
	Counts int    `json:"counts"`
}

//GetPostTagLists 文章tag
func GetPostTagLists(postID int) ([]WpsTag, error) {
	var (
		rows *sql.Rows
		err  error
	)
	// var Post CommentLists
	tagList := []WpsTag{}

	rows, err = dbConn.Query(fmt.Sprintf("select t.tag_id,t.name as tag_name from wps_post_tags as pt left join wps_tags as t on pt.tag_id = t.tag_id where pt.post_id = %d and pt.disabled = 0", postID))
	defer rows.Close()
	if err != nil {
		return tagList, err
	}

	for rows.Next() {
		var aTag WpsTag
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
	offset := (page - 1) * pageSize
	rows, err = dbConn.Query(fmt.Sprintf("select p.id,p.post_title,u.name,c.c_name,p.post_title,p.post_intro,p.created_at from wps_post_tags as pt left join wps_posts as p on pt.post_id = p.id left join wps_users as u on p.user_id = u.id left join wps_post_categories as c on c.id = p.cat_id where p.post_status = 1 and pt.tag_id = %d order by p.created_at desc limit %d,%d", cateID, offset, pageSize))
	defer rows.Close()
	if err != nil {
		return postList, err
	}

	// for rows.Next() {
	// 	var aPost PostList
	// 	err = rows.Scan(
	// 		&aPost.ID,
	// 		&aPost.PostTitle,
	// 		&aPost.UserName,
	// 		&aPost.CategoryName,
	// 		&aPost.PostTitle,
	// 		&aPost.PostIntro,
	// 		&aPost.createdAt,
	// 	)
	// 	tm := time.Unix(int64(aPost.createdAt), 0)
	// 	aPost.CreatedAt = tm.Format("2006-01-02 15:04")
	// 	tags, _ := GetPostTagLists(aPost.ID)
	// 	aPost.Tags = tags
	// 	num, _ := GetArticleNumsByPost(aPost.ID)
	// 	aPost.NumInfo = num
	// 	postList = append(postList, aPost)
	// }
	return postList, nil
}

//AdminGetTagLists 获取后台tag列表
func AdminGetTagLists(search AdminTagListSearch) ([]AdminWpsTag, int, map[string]interface{}) {
	var tags []AdminWpsTag
	var count int
	searchParams := make(map[string]interface{}, 1)
	searchParams["orderNameList"] = util.TagOrderName()
	whereCondition := " where 1=1"
	if search.TagID != 0 {
		whereCondition += " and tag_id = " + strconv.Itoa(search.TagID)
	}
	if search.Name != "" {
		whereCondition += " and name ='" + search.Name + "'"
	}
	sql := fmt.Sprintf("select tag_id,name, counts from wps_tags %s order by %s %s limit %d , %d", whereCondition, search.OrderbyName, search.OrderType, (search.Page-1)*search.PageSize, search.PageSize)
	countSQL := fmt.Sprintf("select count(1) as counts from wps_tags %s", whereCondition)
	err := dbConn.QueryRow(countSQL).Scan(&count)
	if err != nil {
		return tags, count, searchParams
	}
	rows, err := dbConn.Query(sql)
	defer rows.Close()
	if err != nil {
		return tags, count, searchParams
	}
	for rows.Next() {
		var tag AdminWpsTag
		err = rows.Scan(&tag.TagID, &tag.Name, &tag.Counts)
		if err != nil {

		} else {
			tags = append(tags, tag)
		}
	}

	return tags, count, searchParams
}
