package models

import (
	"strings"
)

func SphinxSearch(keyword string, page, pageSize int) (interface{}, error) {
	SphinxClient.SetLimits((page-1)*pageSize, pageSize, 1000, 0)
	// 查询，第一个参数是我们要查询的关键字，第二个是索引名称test1，第三个是备注
	res, err := SphinxClient.Query(keyword, "main", "search article!")
	if err != nil {
		return nil, err
	}
	var articleMap []interface{}

	for _, match := range res.Matches {
		tempData := make(map[string]interface{})
		tempData["id"] = match.DocId
		title, ok := match.AttrValues[0].(string)
		if ok {
			tempData["post_title"] = strings.Replace(title, keyword, "<b style='color:red'>"+keyword+"</b>", -1)
		}
		intro, ok := match.AttrValues[3].(string)
		if ok {
			tempData["post_intro"] = strings.Replace(intro, keyword, "<b style='color:red'>"+keyword+"</b>", -1)
		}
		articleMap = append(articleMap, tempData)
	}
	return articleMap, nil
}
