package models

type postDetail struct {
	ID         int    `json:"id"`
	post_title string `json:"post_title"`
}

func ArticleAll(page, pagesize int) (interface{}, error) {
	rows, _ := dbConn.Query(
		"select ID,post_title,post_author,post_status,comment_count,post_date,post_intro from wps_posts where post_status = 'publish' order by ID DESC limit ? offset ?",
		pagesize,
		(page-1)*pagesize,
	)
	return DBQueryRows(rows)
}

func ArticleOne(articleId int) (interface{}, error) {
	row := dbConn.QueryRow("select ID,post_title,post_author,post_status,comment_count,post_date,post_intro,post_content  from wps_posts where ID = %d AND post_status= '%s'", articleId, "publish")
	result := new(postDetail)
	row.Scan(&result.ID, &result.post_title)
	return result, nil
}
