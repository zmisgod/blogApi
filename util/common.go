package util

import (
	"crypto/md5"
	"encoding/hex"
	"os"
	"reflect"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/astaxie/beego"
)

//CheckEmpty 检查interface是否为空
func CheckEmpty(data interface{}) bool {
	v := reflect.ValueOf(data)
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}
	return v.Interface() == reflect.Zero(v.Type()).Interface()
}

//Md5String 将string加密成 md5
func Md5String(cacheKey string) string {
	md5Ctx := md5.New()
	md5Ctx.Write([]byte(cacheKey))
	cipherStr := md5Ctx.Sum(nil)
	return hex.EncodeToString(cipherStr)
}

//CheckEmail 检查邮箱是否正确
func CheckEmail(email string) (b bool) {
	if m, _ := regexp.MatchString("^([a-zA-Z0-9_-])+@([a-zA-Z0-9_-])+(.[a-zA-Z0-9_-])+", email); !m {
		return false
	}
	return true
}

//ImplodeInt like php function implode() 将int数据的字符转成字符串，并用 implodeStr 字符分割
func ImplodeInt(implodeStr string, array []int) string {
	str := ""
	for k, v := range array {
		prefix := ""
		if k != 0 {
			prefix = ","
		}
		str += prefix + strconv.Itoa(v)
	}
	return str
}

//InArraySting like php function : in_array()
func InArraySting(value string, values []string) bool {
	isExist := false
	for _, v := range values {
		if v == value {
			isExist = true
			break
		}
	}
	return isExist
}

//GetListMapValue 根据数组中map获取其中的value的数据
func GetListMapValue(list []map[string]string, value string) []string {
	res := make([]string, 0)
	for _, v := range list {
		if _, ok := v[value]; ok {
			res = append(res, v[value])
		}
	}
	return res
}

//CombinePagination 组装pagination数据
func CombinePagination(page, pageSize, total int, sortBy, descending string) map[string]interface{} {
	pagination := make(map[string]interface{})
	pagination["page"] = page
	pagination["pageSize"] = pageSize
	pagination["sortBy"] = sortBy
	pagination["descending"] = descending
	pagination["total"] = total
	return pagination
}

//CheckAuthNotExpire 检查auth中的登录时间没有过期
func CheckAuthNotExpire(loginTime int) bool {
	expireTime, err := beego.AppConfig.Int64("UserTokenExpireTime")
	if err != nil {
		expireTime = 86400
	}
	if time.Now().Unix()-int64(loginTime) <= expireTime {
		return true
	}
	return false
}

//ConvertUtf8ToTimeTime []uint8 to time.Time
func ConvertUtf8ToTimeTime(atime []uint8) time.Time {
	pt, _ := strconv.ParseInt(string(atime), 10, 64)
	return time.Unix(pt, 0)
}

//ArrayIntToString 数字数组转换成string
func ArrayIntToString(list []int, spliteStr string) string {
	if len(list) > 0 {
		strList := []string{}
		for _, v := range list {
			strList = append(strList, strconv.Itoa(v))
		}
		return strings.Join(strList, spliteStr)
	}
	return ""
}

//CreateFloder 创建今日的文件夹
func CreateFloder(path string) bool {
	_, err := os.Stat(path)
	if err != nil {
		os.MkdirAll(path, 0766)
	}
	return true
}

//TodayFloderName 今日的文件名
func TodayFloderName(staticPath string) string {
	nowTime := time.Now().Format("2006-01-02")
	res := strings.Split(nowTime, "-")
	path := staticPath + res[0] + "/" + res[1] + "/" + res[2] + "/"
	CreateFloder(path)
	return path
}

//CreateTodayFloder 创建今天的文件夹
func CreateTodayFloder(staticPath string) bool {
	todayStr := TodayFloderName(staticPath)
	return CreateFloder(todayStr)
}
