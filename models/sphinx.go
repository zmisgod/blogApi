package models

import (
	"fmt"
	"strings"
	"time"
)

func SphinxSearch(keyword string, page, pageSize int) (interface{}, error) {
	SphinxClient = SphinxConnect()
	SphinxClient.SetLimits((page-1)*pageSize, pageSize, 1000, 0)
	// 查询，第一个参数是我们要查询的关键字，第二个是索引名称test1，第三个是备注
	res, err := SphinxClient.Query(keyword, "main", "search article!")
	if err != nil {
		return nil, err
	}
	var articleMap []interface{}

	for _, match := range res.Matches {
		fmt.Println(match)
		tempData := make(map[string]interface{})
		tempData["id"] = match.DocId
		title, ok := match.AttrValues[0].(string)
		if ok {
			postTitle := strings.Replace(title, strings.ToUpper(keyword), "<b style='color:red'>"+strings.ToUpper(keyword)+"</b>", -1)
			tempData["post_intro"] = strings.Replace(postTitle, strings.ToLower(keyword), "<b style='color:red'>"+strings.ToLower(keyword)+"</b>", -1)
		}
		intro, ok := match.AttrValues[3].(string)
		if ok {
			postIntro := strings.Replace(intro, strings.ToUpper(keyword), "<b style='color:red'>"+strings.ToUpper(keyword)+"</b>", -1)
			tempData["post_title"] = strings.Replace(postIntro, strings.ToLower(keyword), "<b style='color:red'>"+strings.ToLower(keyword)+"</b>", -1)
		}
		postAt, ok := match.AttrValues[2].(int64)
		if ok {
			fmt.Println(postAt)
			postDate := strings.Split(time.Unix(postAt, 0).Format("2006-01-02"), " ")
			fmt.Println(postDate)
			tempData["post_date"] = postDate[0]
		}
		articleMap = append(articleMap, tempData)
	}
	return articleMap, nil
}
