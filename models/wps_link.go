package models

import (
	"fmt"
	"time"

	"github.com/zmisgod/blogApi/util"
)

var wpsLinks = "wps_links"

//GetLinks 友链列表
func GetLinks() ([]util.WpsLink, error) {
	linkList := []util.WpsLink{}

	rows, err := dbConn.Query(fmt.Sprintf("select link_url,link_name,link_image,link_description from "+wpsLinks+" where start_time <= %d and end_time >= %d and link_status = 1", time.Now().Unix(), time.Now().Unix()))
	defer rows.Close()
	if err != nil {
		return linkList, err
	}
	for rows.Next() {
		var aLink util.WpsLink
		err = rows.Scan(
			&aLink.LinkURL,
			&aLink.LinkName,
			&aLink.LinkImage,
			&aLink.LinkDescription,
		)
		if err != nil {
			continue
		}
		linkList = append(linkList, aLink)
	}
	return linkList, nil
}
