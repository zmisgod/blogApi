package models

import (
	"fmt"
	"time"
)

var wpsHistory = "wps_history"

//SaveUserVisiteHistory 保存用户的浏览历史
func SaveUserVisiteHistory(vType, vIP, vUserAgent, requestURI, refer string) int64 {
	stmt, err := dbConn.Prepare("insert into " + wpsHistory + " (type,user_agent,ip,uri,refer,visite_time) values (?,?,?,?,?,?)")
	defer stmt.Close()
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
