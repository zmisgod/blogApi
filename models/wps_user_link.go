package models

import (
	"database/sql"
	"fmt"

	"github.com/zmisgod/blogApi/util"
)

//GetUserLinkLists 用户第三方的链接
func GetUserLinkLists(userID int) []string {
	var (
		rows *sql.Rows
		err  error
	)
	userLinks := make([]string, 0)

	rows, err = dbConn.Query(fmt.Sprintf("select l.link_type,l.suffix from wps_users as u left join wps_users_link as l on u.id = l.user_id where u.id = %d", userID))
	defer rows.Close()
	if err != nil {
		return userLinks
	}

	for rows.Next() {
		var userLink util.WpsUsersLink
		err = rows.Scan(
			&userLink.LinkType,
			&userLink.Suffix,
		)
		if err != nil {
			continue
		}
		if value, ok := util.UserLinkType[userLink.LinkType]; ok {
			fullURL := value + "/" + userLink.Suffix
			userLinks = append(userLinks, fullURL)
		}
	}
	return userLinks
}
