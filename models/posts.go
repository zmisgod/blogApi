package models

import (
	"fmt"
	"strconv"
	"strings"
)

type Tag struct {
	ID            int
	category_name string
	category_id   int
}

//GetArticleLists 文章列表
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
	if len(ids) == 0 {
		for _, value := range data {
			value["tag"] = make([]interface{}, 0)
		}
		return data, nil
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
			postIDr := post["id"].(int64)
			postID := strconv.FormatInt(postIDr, 10)
			post["tag"] = tagsBind[postID]
		}
		return postResult, nil
	}
	return lists, nil
}

//ArticleOne 文章详情
func ArticleOne(articleID int) (interface{}, error) {
	row, _ := dbConn.Query(
		"select p.ID as id,p.post_title,p.post_content,p.comment_count,p.post_date,p.post_intro,tm.name as category_name,u.user_nicename as author,t.term_taxonomy_id as category_id from wps_posts as p left join wps_term_relationships as rs on rs.object_id = p.ID left join wps_term_taxonomy as t on t.term_taxonomy_id = rs.term_taxonomy_id left join wps_terms as tm on tm.`term_id` = t.term_id left join wps_users as u on p.post_author = u.ID where p.ID = ? and p.post_status=?  and t.taxonomy = ?",
		articleID,
		"publish",
		"category",
	)
	return DBQueryRow(row)
}

//ArticleTags 单条文章的tag列表
func ArticleTags(articleID int) (interface{}, error) {
	rows, _ := dbConn.Query(
		"select tm.name as category_name,t.term_taxonomy_id as category_id from wps_posts as p left join wps_term_relationships as rs on rs.object_id = p.ID left join wps_term_taxonomy as t on t.term_taxonomy_id = rs.term_taxonomy_id left join wps_terms as tm on tm.`term_id` = t.term_id where p.ID = ? and t.taxonomy = ?",
		articleID,
		"post_tag",
	)
	return DBQueryRows(rows)
}

//ArticleTagsLists 文章tag列表
func ArticleTagsLists(articleIds string) (interface{}, error) {
	sqlSte := fmt.Sprintf("select tm.name as category_name,t.term_taxonomy_id as category_id,p.ID as id from wps_posts as p left join wps_term_relationships as rs on rs.object_id = p.ID left join wps_term_taxonomy as t on t.term_taxonomy_id = rs.term_taxonomy_id left join wps_terms as tm on tm.`term_id` = t.term_id where p.ID in (%s) and t.taxonomy ='%s'", articleIds, "post_tag")
	rows, _ := dbConn.Query(sqlSte)
	return DBQueryRows(rows)
}

//OnlyShowYMD 只显示年月日
func OnlyShowYMD(date string) string {
	res := strings.Split(date, " ")
	return res[0]
}

//TagAll 所有的category／tag获取
func TagAll(tagID, page, pageSize int, articleType string) (interface{}, error) {
	postIDs := GetTagListPostIDs(tagID, page, pageSize, articleType)
	if postIDs == "" {
		return "", nil
	}
	//然后查询上面id对应的id信息
	postSQL := fmt.Sprintf("select p.ID as id,p.post_title,p.comment_count,p.post_date,p.post_intro,tm.name as category_name,u.user_nicename as author from wps_posts as p left join wps_term_relationships as rs on rs.object_id = p.ID left join wps_term_taxonomy as t on t.term_taxonomy_id = rs.term_taxonomy_id left join wps_terms as tm on tm.`term_id` = t.term_id left join wps_users as u on p.post_author = u.ID where p.ID in (%s) and  p.post_status= '%s'  and t.taxonomy = '%s' order by ID DESC", postIDs, "publish", "category")
	rows, _ := dbConn.Query(postSQL)
	if err != nil {
		return "", err
	}
	lists, err := DBQueryRows(rows)
	data, _ := lists.([]map[string]interface{})
	ids := make([]int64, 0)
	for _, value := range data {
		temDate := value["post_date"].(string)
		value["post_date"] = OnlyShowYMD(temDate)
		temPostID := value["id"].(string)
		v, err := strconv.ParseInt(temPostID, 10, 64)
		if err == nil {
			value["id"] = v
			ids = append(ids, v)
		}
	}
	if len(ids) == 0 {
		for _, value := range data {
			value["tag"] = make([]interface{}, 0)
		}
		return data, nil
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
			postID := strconv.FormatInt(post["id"].(int64), 10)
			post["tag"] = tagsBind[postID]
		}
		return postResult, nil
	}
	return data, nil
}

//GetTagListPostIDs 根据tagid获取post id
func GetTagListPostIDs(tagID, page, pageSize int, articleType string) string {
	//先获取tag对应的id的数据
	sql := fmt.Sprintf("select p.ID as id from wps_posts as p left join wps_term_relationships as re on re.object_id = p.ID left join wps_term_taxonomy as ta on ta.term_taxonomy_id = re.term_taxonomy_id left join wps_users as u on p.post_author = u.ID where re.term_taxonomy_id = %d and p.post_status = 'publish' and p.post_type = 'post' and ta.taxonomy = '%s' order by p.ID desc limit %d offset %d", tagID, articleType, pageSize, (page-1)*pageSize)
	tagIDRows, _ := dbConn.Query(sql)
	tagRowLists, err := DBQueryRows(tagIDRows)
	if err != nil {
		return ""
	}
	tagRowIDMap, _ := tagRowLists.([]map[string]interface{})
	tagToPostID := make([]int64, 0)
	for _, value := range tagRowIDMap {
		temPostIDr, err := strconv.ParseInt(value["id"].(string), 10, 64)
		if err == nil {
			tagToPostID = append(tagToPostID, temPostIDr)
		}
	}
	if len(tagToPostID) == 0 {
		return ""
	}
	var postIDs string
	for _, v := range tagToPostID {
		postIDs += "," + strconv.FormatInt(v, 10)
	}
	searchID := []rune(postIDs)
	searchFinnalID := string(searchID[1:])
	return searchFinnalID
}

//GetRelatedPostByCategoryID 根据tagid获取post info
func GetRelatedPostByCategoryID(tagID, page, pageSize int, articleType string) string {
	//先获取tag对应的id的数据
	sql := fmt.Sprintf("select p.ID as id, p.post_title from wps_posts as p left join wps_term_relationships as re on re.object_id = p.ID left join wps_term_taxonomy as ta on ta.term_taxonomy_id = re.term_taxonomy_id left join wps_users as u on p.post_author = u.ID where re.term_taxonomy_id = %d and p.post_status = 'publish' and p.post_type = 'post' and ta.taxonomy = '%s' order by p.ID desc limit %d offset %d", tagID, articleType, pageSize, (page-1)*pageSize)
	tagIDRows, _ := dbConn.Query(sql)
	tagRowLists, err := DBQueryRows(tagIDRows)
	fmt.Println(tagRowLists)
	if err != nil {
		return ""
	}
	tagRowIDMap, _ := tagRowLists.([]map[string]interface{})
	tagToPostID := make([]int64, 0)
	for _, value := range tagRowIDMap {
		temPostIDr, err := strconv.ParseInt(value["id"].(string), 10, 64)
		if err == nil {
			tagToPostID = append(tagToPostID, temPostIDr)
		}
	}
	if len(tagToPostID) == 0 {
		return ""
	}
	var postIDs string
	for _, v := range tagToPostID {
		postIDs += "," + strconv.FormatInt(v, 10)
	}
	searchID := []rune(postIDs)
	searchFinnalID := string(searchID[1:])
	return searchFinnalID
}
