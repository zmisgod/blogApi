package models

import (
	"fmt"
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
	data, ok := lists.([]map[string]interface{})
	if ok {
		var ids []int64
		for _, value := range data {
			v, ok := value["id"].(int64)
			if ok {
				ids = append(ids, v)
			}
		}
		fmt.Println(ids)
		// res := strings.Join(ids, ",")
		// fmt.Println(res)
		// tags, err := ArticleTagsLists(res)
		// if err != nil {
		// 	return "", err
		// }
		// fmt.Println(tags)
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
	rows, _ := dbConn.Query(
		"select tm.name as category_name,t.term_taxonomy_id as category_id from wps_posts as p left join wps_term_relationships as rs on rs.object_id = p.ID left join wps_term_taxonomy as t on t.term_taxonomy_id = rs.term_taxonomy_id left join wps_terms as tm on tm.`term_id` = t.term_id where p.ID in (?) and t.taxonomy = ?",
		articleIds,
		"post_tag",
	)
	return DBQueryRows(rows)
}
