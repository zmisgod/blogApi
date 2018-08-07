package models

import (
	"database/sql"
	"fmt"
)

//UserLink 用户的自定义链接
type UserLink struct {
	LinkType int    `json:"link_type"`
	Suffix   string `json:"suffix"`
}

//GetUserLink 用户第三方的链接
func GetUserLink(userID int) ([]UserLink, error) {
	var (
		rows *sql.Rows
		err  error
	)
	// var Post CommentLists
	userLink := []UserLink{}

	rows, err = dbConn.Query(fmt.Sprintf("select l.link_type,l.suffix from  wps_users as u left join wps_users_link as l on u.id = l.user_id where u.id = %d", userID))
	defer rows.Close()
	if err != nil {
		return userLink, err
	}

	for rows.Next() {
		var aUserLink UserLink
		err = rows.Scan(
			&aUserLink.LinkType,
			&aUserLink.Suffix,
		)
		if err != nil {
			continue
		}
		userLink = append(userLink, aUserLink)
	}
	return userLink, nil
}
