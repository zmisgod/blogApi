package models

import (
	"fmt"
	"time"
)

//History 浏览记录
type History struct {
	ID         int    `json:"id"`
	Type       string `json:"type"`
	UserAgent  string `json:"user_agent"`
	IP         string `json:"ip"`
	visiteTime int
	VisiteTime string `json:"visite_time"`
	params     string
	Params     map[string]interface{}
}

//SaveUserVisiteHistory 保存用户的浏览历史
func SaveUserVisiteHistory(vType, vIP, vUserAgent, requestURI, refer string) int64 {
	var (
		err error
	)
	stmt, err := dbConn.Prepare("insert into wps_history (type,user_agent,ip,uri,refer,visite_time) values (?,?,?,?,?,?)")
	if err == nil {
		result, err := stmt.Exec(vType, vUserAgent, vIP, requestURI, refer, time.Now().Unix())
		if err == nil {
			lastID, err := result.LastInsertId()
			if err == nil {
				return lastID
			}
		} else {
			fmt.Println(err)
		}
	} else {
		fmt.Println(err)
	}
	return 0
}
