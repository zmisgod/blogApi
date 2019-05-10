package models

import (
	"errors"
	"fmt"
	"time"

	"github.com/zmisgod/blogApi/util"
	"golang.org/x/crypto/bcrypt"
)

//CONST bcrypt need
const CONST = 10

//UserLogin 用户登录信息
type UserLogin struct {
	util.WpsUsers
	LoginTime int64 `json:"login_time"`
}

//CheckUserExists 检查用户是否存在
func CheckUserExists(email, password string) (string, error) {
	userInfo, err := checkUserPassword(email, password)
	if err != nil {
		return "", err
	}
	auth, err := encodeUserAuthority(userInfo)
	return auth, err
}

//encodeUserAuthority 加密用户auth
func encodeUserAuthority(user UserLogin) (string, error) {
	jsonStr, err := util.JWTEncode(user.ID, user.Email, user.LoginTime)
	if err != nil {
		return "", err
	}
	return jsonStr, nil
}

//checkUserPassword 检查用户名密码是否正确，如果正确返回用户信息
func checkUserPassword(email, passPassword string) (UserLogin, error) {
	var user UserLogin
	err := dbConn.QueryRow(fmt.Sprintf("select id,name,email,password,status from wps_users where email = \"%s\"", email)).Scan(&user.ID, &user.Name, &user.Email, &user.Password, &user.Status)
	if err != nil {
		return user, err
	}
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(passPassword))
	if err != nil {
		var user UserLogin
		return user, err
	}
	user.LoginTime = time.Now().Unix()
	return user, nil
}

//checkEmailIsExists 检查邮箱是否注册
func checkEmailIsExists(email string) (bool, error) {
	var id int
	err := dbConn.QueryRow(fmt.Sprintf("select id from wps_users where email = \"%s\"", email)).Scan(&id)
	if err != nil {
		return false, err
	}
	if id > 0 {
		return false, nil
	}
	return false, nil
}

//RegisterUser 注册用户
func RegisterUser(email, username, password string) (bool, error) {
	isExists, err := checkEmailIsExists(email)
	if err != nil {
		return false, err
	}
	if isExists {
		return false, errors.New("email exists")
	}
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), CONST)
	if err != nil {
		return false, err
	}
	stmt, err := dbConn.Prepare("insert into wps_users (name, email, password, status) values (?,?,?,?,?,?)")
	if err != nil {
		return false, err
	}
	defer stmt.Close()
	_, err = stmt.Exec(username, email, hashedPassword, 0)
	if err != nil {
		return false, err
	}
	return true, nil
}

//CheckUserAuth 检查用户的身份认证
func CheckUserAuth(authorition string) (bool, map[string]interface{}) {
	if authorition == "" {
		return false, make(map[string]interface{})
	}
	res, err := util.JWTDecode(authorition)
	if err != nil {
		return false, make(map[string]interface{})
	}
	return true, res
}
