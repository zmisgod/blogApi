package models

import (
	"database/sql"
	"fmt"
)

//User 用户信息
type User struct {
	id        int
	NickName  string     `json:"nickname"`
	Sex       int        `json:"sex"`
	HeadURL   string     `json:"head_url"`
	Introduce string     `json:"introduce"`
	Birthday  string     `json:"birthday"`
	UserLink  []UserLink `json:"user_link"`
}

//GetUserInfo 用户信息
func GetUserInfo(userID int) (User, error) {
	var (
		rows *sql.Rows
		err  error
	)
	// var Post CommentLists
	user := User{}

	rows, err = dbConn.Query(fmt.Sprintf("select i.user_id,i.nickname, i.head_url,i.birthday,i.introduce,i.sex from  wps_users as u left join wps_users_info as i on u.id = i.user_id where u.id = %d", userID))
	if err != nil {
		return user, err
	}

	for rows.Next() {
		err = rows.Scan(
			&user.id,
			&user.NickName,
			&user.HeadURL,
			&user.Birthday,
			&user.Introduce,
			&user.Sex,
		)
		if err != nil {
			continue
		}
		userLink, _ := GetUserLink(user.id)
		user.UserLink = userLink
		break
	}
	return user, nil
}
