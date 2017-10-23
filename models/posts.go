package models

import (
	"fmt"
	"time"

	"github.com/astaxie/beego/orm"
)

type Posts struct {
	Id           int       `json:"id"`
	PostTitle    string    `json:"post_title"`
	PostAuthor   string    `json:"post_author"`
	PostStatus   string    `json:"post_status"`
	CommentCount int       `json:"comment_count"`
	PostDate     time.Time `json:"post_date"`
	PostIntro    string    `json:"post_intro"`
}
type PostInfo struct {
	PostContent  string    `json:"post_content"`
	Id           int       `json:"id"`
	PostTitle    string    `json:"post_title"`
	PostAuthor   string    `json:"post_author"`
	PostStatus   string    `json:"post_status"`
	CommentCount int       `json:"comment_count"`
	PostDate     time.Time `json:"post_date"`
	PostIntro    string    `json:"post_intro"`
}

func ArticleAll(page, pagesize int) interface{} {
	o := orm.NewOrm()
	var lists []Posts
	o.Raw(fmt.Sprintf("select * from wps_posts where post_status = 'publish' order by ID DESC limit %d", pagesize)).QueryRows(&lists)
	return lists
}

func Tests(page int) interface{} {
	stmt, err := db_conn.Query("select ID,post_content from wps_posts where post_status = 'publish' and ID = ?", page)
	if err != nil {
		fmt.Println(err)
	}
	for stmt.Next() {
		var ID int
		var post_content string
		err = stmt.Scan(&ID, &post_content)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(post_content)
		fmt.Println(ID)
	}
	return stmt
}

func ArticleOne(articleId int) (PostInfo, string) {
	o := orm.NewOrm()
	var articleDetail PostInfo
	sql := fmt.Sprintf("select * from wps_posts where ID = %d AND post_status= '%s'", articleId, "publish")
	o.Raw(sql).QueryRow(&articleDetail)
	if articleDetail.Id == 0 {
		return articleDetail, "error"
	}
	return articleDetail, ""
}
