package models

import (
	"database/sql"
	"fmt"
	"time"
)

//Link 友链
type Link struct {
	LinkURL         string `json:"link_url"`
	LinkName        string `json:"link_name"`
	LinkImage       string `json:"link_image"`
	LinkDescription string `json:"link_description"`
}

//GetLinks 友链列表
func GetLinks() ([]Link, error) {
	var (
		rows *sql.Rows
		err  error
	)
	// var commentList CommentLists
	linkList := []Link{}

	rows, err = dbConn.Query(fmt.Sprintf("select link_url,link_name,link_image,link_description from wps_links where start_time <= %d and end_time >= %d and link_status = 1", time.Now().Unix(), time.Now().Unix()))
	defer rows.Close()
	if err != nil {
		return linkList, err
	}
	for rows.Next() {
		var aLink Link
		err = rows.Scan(
			&aLink.LinkURL,
			&aLink.LinkName,
			&aLink.LinkImage,
			&aLink.LinkDescription,
		)
		if err != nil {
			continue
		}
		linkList = append(linkList, aLink)
	}
	return linkList, nil
}
