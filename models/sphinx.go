package models

//Sphinx Search  已作废
//todo Elastic Search

import (
	"strings"
	"time"
)

//SphinxSearch demo
func SphinxSearch(keyword string, page, pageSize int) (interface{}, error) {
	sphinxClient.SetLimits((page-1)*pageSize, pageSize, 1000, 0)
	// 查询，第一个参数是我们要查询的关键字，第二个是索引名称test1，第三个是备注
	res, err := sphinxClient.Query(keyword, "main", "search article!")
	if err != nil {
		return nil, err
	}
	var articleMap []interface{}

	for _, match := range res.Matches {
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
		postAt, ok := match.AttrValues[2].(uint32)
		if ok {
			postDate := strings.Split(time.Unix(int64(postAt), 0).Format("2006-01-02"), " ")
			tempData["post_date"] = postDate[0]
		} else {
			tempData["post_date"] = ""
		}
		articleMap = append(articleMap, tempData)
	}
	return articleMap, nil
}
