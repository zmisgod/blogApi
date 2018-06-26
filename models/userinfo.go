package models

import (
	"database/sql"
	"fmt"

	"github.com/astaxie/beego"
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

//userHeadImages 用户头像
type userImages struct {
	id      int
	sex     int
	HeadURL string `json:"head_url"`
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
		user.HeadURL = beego.AppConfig.String("StaticPrefix") + user.HeadURL
		userLink, _ := GetUserLink(user.id)
		user.UserLink = userLink
		break
	}
	return user, nil
}

//GetUserHeadImages 获取用户头像
func GetUserHeadImages(userIDs string) (map[int]string, error) {
	var (
		rows *sql.Rows
		err  error
	)
	// var Post CommentLists
	userHeadImages := map[int]string{}
	rows, err = dbConn.Query(fmt.Sprintf("select i.user_id,i.head_url,i.sex from  wps_users as u left join wps_users_info as i on u.id = i.user_id where u.id in (%s)", userIDs))
	if err != nil {
		return userHeadImages, err
	}

	for rows.Next() {
		var user userImages
		err = rows.Scan(
			&user.id,
			&user.HeadURL,
			&user.sex,
		)
		if err != nil {
			continue
		}
		userHeadImages[user.id] = user.HeadURL
	}
	return userHeadImages, nil
}
