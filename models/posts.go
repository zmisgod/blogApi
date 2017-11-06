package models

import (
	"fmt"
	"strconv"
	"strings"
)

func ArticleAll(page, pagesize int) (interface{}, error) {
	rows, _ := dbConn.Query(
		"select p.ID as id,p.post_title,p.comment_count,p.post_date,p.post_intro,tm.name as category_name,u.user_nicename as author from wps_posts as p left join wps_term_relationships as rs on rs.object_id = p.ID left join wps_term_taxonomy as t on t.term_taxonomy_id = rs.term_taxonomy_id left join wps_terms as tm on tm.`term_id` = t.term_id left join wps_users as u on p.post_author = u.ID where p.post_status=?  and t.taxonomy = ? order by ID DESC limit ? offset ?",
		"publish",
		"category",
		pagesize,
		(page-1)*pagesize,
	)
	lists, err := DBQueryRows(rows)
	if err != nil {
		return "", err
	}
	return lists, nil
}

type Tag struct {
	ID            int
	category_name string
	category_id   int
}

func GetArticleLists(page, pagesize int) (interface{}, error) {
	rows, _ := dbConn.Query(
		"select p.ID as id,p.post_title,p.comment_count,p.post_date,p.post_intro,tm.name as category_name,u.user_nicename as author from wps_posts as p left join wps_term_relationships as rs on rs.object_id = p.ID left join wps_term_taxonomy as t on t.term_taxonomy_id = rs.term_taxonomy_id left join wps_terms as tm on tm.`term_id` = t.term_id left join wps_users as u on p.post_author = u.ID where p.post_status=?  and t.taxonomy = ? order by ID DESC limit ? offset ?",
		"publish",
		"category",
		pagesize,
		(page-1)*pagesize,
	)
	lists, err := DBQueryRows(rows)
	if err != nil {
		return "", err
	}
	data, _ := lists.([]map[string]interface{})
	ids := make([]int64, 0)
	for _, value := range data {
		temDate := value["post_date"].(string)
		value["post_date"] = OnlyShowYMD(temDate)
		v, ok := value["id"].(int64)
		if ok {
			ids = append(ids, v)
		}
	}
	var tagIds string
	for _, v := range ids {
		tagIds += "," + strconv.FormatInt(v, 10)
	}
	tagRune := []rune(tagIds)
	finalTag := string(tagRune[1:])
	tags, err := ArticleTagsLists(finalTag)
	if err == nil {
		mTags, _ := tags.([]map[string]interface{})
		tagsBind := make(map[string]interface{}, 0)
		for _, vs := range ids {
			var tagDetail []interface{}
			tagID := strconv.FormatInt(vs, 10)
			for _, de := range mTags {
				tagDID, _ := de["id"].(string)
				if tagID == tagDID && len(de) != 0 {
					tagDetail = append(tagDetail, de)
				}
			}
			tagsBind[tagID] = tagDetail
		}
		postResult, _ := lists.([]map[string]interface{})
		for _, post := range postResult {
			postId := post["id"].(int64)
			postID := strconv.FormatInt(postId, 10)
			post["tag"] = tagsBind[postID]
		}
		return postResult, nil
	}
	return lists, nil
}

func ArticleOne(articleId int) (interface{}, error) {
	row, _ := dbConn.Query(
		"select p.ID as id,p.post_title,p.post_content,p.comment_count,p.post_date,p.post_intro,tm.name as category_name,u.user_nicename as author,t.term_taxonomy_id as category_id from wps_posts as p left join wps_term_relationships as rs on rs.object_id = p.ID left join wps_term_taxonomy as t on t.term_taxonomy_id = rs.term_taxonomy_id left join wps_terms as tm on tm.`term_id` = t.term_id left join wps_users as u on p.post_author = u.ID where p.ID = ? and p.post_status=?  and t.taxonomy = ?",
		articleId,
		"publish",
		"category",
	)
	return DBQueryRow(row)
}

func ArticleTags(articleId int) (interface{}, error) {
	rows, _ := dbConn.Query(
		"select tm.name as category_name,t.term_taxonomy_id as category_id from wps_posts as p left join wps_term_relationships as rs on rs.object_id = p.ID left join wps_term_taxonomy as t on t.term_taxonomy_id = rs.term_taxonomy_id left join wps_terms as tm on tm.`term_id` = t.term_id where p.ID = ? and t.taxonomy = ?",
		articleId,
		"post_tag",
	)
	return DBQueryRows(rows)
}

func ArticleTagsLists(articleIds string) (interface{}, error) {
	sqlSte := fmt.Sprintf("select tm.name as category_name,t.term_taxonomy_id as category_id,p.ID as id from wps_posts as p left join wps_term_relationships as rs on rs.object_id = p.ID left join wps_term_taxonomy as t on t.term_taxonomy_id = rs.term_taxonomy_id left join wps_terms as tm on tm.`term_id` = t.term_id where p.ID in (%s) and t.taxonomy ='%s'", articleIds, "post_tag")
	rows, _ := dbConn.Query(sqlSte)
	return DBQueryRows(rows)
}

func OnlyShowYMD(date string) string {
	res := strings.Split(date, " ")
	return res[0]
}
