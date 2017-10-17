package models

import (
	"time"

	"github.com/astaxie/beego/orm"
)

type Posts struct {
	ID           int       `json:"id"`
	PostAuthor   int       `json:"post_author"`
	PostStatus   string    `json:"post_status"`
	CommentCount int       `json:"comment_count"`
	PostDate     time.Time `json:"post_date"`
	PostContent  string    `json:"post_content"`
	PostIntro    string    `json:"post_intro"`
}

func (a *Posts) Insert() error {
	if _, err := orm.NewOrm().Insert(a); err != nil {
		return err
	}
	return nil
}

func (a *Posts) TableName() string {
	return TableName("posts")
}

func (a *Posts) Read(filds ...string) error {
	if err := orm.NewOrm().Read(a, filds...); err != nil {
		return err
	}
	return nil
}

func (a *Posts) Update(filds ...string) error {
	if _, err := orm.NewOrm().Update(a, filds...); err != nil {
		return err
	}
	return nil
}

func (a *Posts) Query() orm.QuerySeter {
	return orm.NewOrm().QueryTable(a)
}
