package models

import (
	"database/sql"
	"errors"
	"fmt"

	"github.com/astaxie/beego"
)

//UserInfo 用户信息
type UserInfo struct {
	id        int
	NickName  string `json:"nickname"`
	Sex       int    `json:"sex"`
	HeadURL   string `json:"head_url"`
	Introduce string `json:"introduce"`
	birthday  []uint8
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
func GetUserInfo(userID int) (UserInfo, error) {
	user := UserInfo{}
	err := dbConn.QueryRow(fmt.Sprintf("select i.user_id,i.nickname, i.head_url,i.birthday,i.introduce,i.sex from  wps_users as u left join wps_users_info as i on u.id = i.user_id where u.id = %d", userID)).
		Scan(&user.id,
			&user.NickName,
			&user.HeadURL,
			&user.birthday,
			&user.Introduce,
			&user.Sex,
		)
	if err != nil {
		return user, errors.New("未找到此人")
	}
	user.HeadURL = beego.AppConfig.String("StaticPrefix") + user.HeadURL
	userLink, _ := GetUserLink(user.id)
	user.Birthday = string(user.birthday)
	user.UserLink = userLink
	return user, nil
}

//GetUserHeadImages 获取用户头像
func GetUserHeadImages(userIDs string) (map[int]string, error) {
	var (
		rows *sql.Rows
		err  error
	)
	userHeadImages := map[int]string{}
	rows, err = dbConn.Query(fmt.Sprintf("select i.user_id,i.head_url,i.sex from  wps_users as u left join wps_users_info as i on u.id = i.user_id where u.id in (%s)", userIDs))
	defer rows.Close()
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
		userHeadImages[user.id] = beego.AppConfig.String("StaticPrefix") + user.HeadURL
	}
	return userHeadImages, nil
}

//UpdateUserInfo 修改user_info
func UpdateUserInfo(nickname, headURL, introduce, birthday string, sex int, userID int64) (bool, error) {
	var id int
	err := dbConn.QueryRow(fmt.Sprintf("select id from wps_users_info where user_id = %d", userID)).Scan(&id)
	if err != nil {
		stmt, err := dbConn.Prepare("insert into wps_users_info (user_id, nickname, sex, head_url, introduce, birthday) values (?,?,?,?,?,?)")
		if err != nil {
			return false, err
		}
		_, err = stmt.Exec(id, nickname, sex, headURL, introduce, birthday)
		if err != nil {
			return false, err
		}
		return true, nil
	}
	stmt, err := dbConn.Prepare("update wps_users_info set nickname = ? , sex = ?, head_url = ?, introduce = ?, birthday = ? where user_id = ?")
	if err != nil {
		return false, err
	}
	_, err = stmt.Exec(nickname, sex, headURL, introduce, birthday, userID)
	if err != nil {
		return false, err
	}
	return true, nil
}
