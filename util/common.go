package util

import (
	"crypto/md5"
	"encoding/hex"
	"reflect"
	"regexp"
)

func CheckEmpty(data interface{}) bool {
	v := reflect.ValueOf(data)
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}
	return v.Interface() == reflect.Zero(v.Type()).Interface()
}

func Md5String(cacheKey string) string {
	md5Ctx := md5.New()
	md5Ctx.Write([]byte(cacheKey))
	cipherStr := md5Ctx.Sum(nil)
	return hex.EncodeToString(cipherStr)
}

func CheckEmail(email string) (b bool) {
	if m, _ := regexp.MatchString("^([a-zA-Z0-9_-])+@([a-zA-Z0-9_-])+(.[a-zA-Z0-9_-])+", email); !m {
		return false
	}
	return true
}
