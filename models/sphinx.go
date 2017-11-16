package models

func SphinxSearch(keyword string, page, pageSize int) (interface{}, error) {
	SphinxClient.SetLimits((page-1)*pageSize, pageSize, 1000, 0)
	// 查询，第一个参数是我们要查询的关键字，第二个是索引名称test1，第三个是备注
	res, err := SphinxClient.Query(keyword, "main", "search article!")
	if err != nil {
		return nil, err
	}
	var articleMap []interface{}

	for _, match := range res.Matches {
		var tempData map[string]interface{}
		tempData["id"] = match.DocId
		tempData["post_title"] = match.AttrValues[0]
		tempData["post_intro"] = match.AttrValues[3]
		articleMap = append(articleMap, tempData)
	}
	return articleMap, nil
}
