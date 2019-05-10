package util

import (
	"fmt"
	"strconv"

	"github.com/astaxie/beego"
	jwt "github.com/dgrijalva/jwt-go"
)

//JWTEncode jwt加密
func JWTEncode(userID int, email string, loginTime int64) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userID":    strconv.Itoa(userID),
		"email":     email,
		"loginTime": strconv.Itoa(int(loginTime)),
	})
	return token.SignedString([]byte(beego.AppConfig.String("UserAuthKey")))
}

//JWTDecode jwt解密
func JWTDecode(tokenString string) (map[string]interface{}, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(beego.AppConfig.String("UserAuthKey")), nil
	})

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, nil
	}
	return make(map[string]interface{}), err
}
