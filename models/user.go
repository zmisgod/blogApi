package models

import (
	"encoding/base64"
	"errors"
	"fmt"

	"github.com/astaxie/beego"

	"github.com/zmisgod/blogApi/util"
	"golang.org/x/crypto/bcrypt"
)

//CONST bcrypt need
const CONST = 10

//User 用户
type User struct {
	ID       int
	Name     string
	Email    string
	Password string
	Status   int
}

//CheckUserExists 检查用户是否存在
func CheckUserExists(username, password, oPassword string) (bool, error) {
	err = bcrypt.CompareHashAndPassword([]byte(oPassword), []byte(password))
	if err != nil {
		return false, err
	}
	return true, nil
}

//checkEmailIsExists 检查邮箱是否注册
func checkEmailIsExists(email string) bool {
	var id int
	err := dbConn.QueryRow(fmt.Sprintf("select id from wps_users where email = \"%s\"", email)).Scan(&id)
	if err != nil {
		return true
	}
	if id > 0 {
		return true
	}
	return false
}

//RegisterUser 注册用户
func RegisterUser(email, username, password string) (bool, error) {
	isExists := checkEmailIsExists(username)
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
func CheckUserAuth(authorition string) bool {
	authoritionDecode, err := base64.StdEncoding.DecodeString(authorition)
	if err != nil {
		return false
	}
	encryptKey := beego.AppConfig.String("UserAuthKey")
	userInfo, err := util.AesDecrypt(authoritionDecode, []byte(encryptKey))
	if err != nil {
		return false
	}
	fmt.Println(string(userInfo))
	return true
}

//GenerateUserAuth 生成用户认证key
func GenerateUserAuth(userID int) string {
	encryptKey := beego.AppConfig.String("UserAuthKey")
	userEncrypted, err := util.AesEncrypt([]byte("111212121212"), []byte(encryptKey))
	if err != nil {
		fmt.Println(err)
		return "error"
	}
	return base64.StdEncoding.EncodeToString(userEncrypted)
}
