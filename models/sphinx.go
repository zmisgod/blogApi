package models

import (
	"github.com/yunge/sphinx"
)

func SphinxSearch(keyword string, page, pageSize int) (interface{}, error) {
	sphinxOptions := &sphinx.Options{
		Host:      "localhost",
		Timeout:   5000,
		Limit:     pageSize,
		MatchMode: sphinx.SPH_MATCH_ANY,
	}
	SphinxClient := sphinx.NewClient(sphinxOptions)
	if err := SphinxClient.Error(); err != nil {
		return nil, err
	}
	defer SphinxClient.Close()

	SphinxClient.SetLimits((page-1)*pageSize, pageSize, 10000000, 0)
	// 查询，第一个参数是我们要查询的关键字，第二个是索引名称test1，第三个是备注
	res, err := SphinxClient.Query(keyword, "main", "search article!")
	if err != nil {
		return nil, err
	}
	var articleMap []interface{}
	for _, match := range res.Matches {
		articleMap = append(articleMap, match)
	}
	return articleMap, nil
}
