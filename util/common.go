package util

import (
	"crypto/md5"
	"encoding/hex"
	"reflect"
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
