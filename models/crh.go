package models

import (
	"strconv"

	"github.com/zmisgod/goTool/drawsvg"
)

//CHSR China High Speed Rewaiy
type CHSR struct {
	Resize int    `json:"resize"`
	List   []List `json:"list"`
	Circle Circle `json:"circle"`
}

//Circle info
type Circle struct {
	Radis       int    `json:"radis"`
	Color       string `json:"color"`
	BorderWidth int    `json:"border_width"`
}

//List 线路
type List struct {
	TrainID   int       `json:"train_id"`
	TrainName string    `json:"train_name"`
	Type      int       `json:"type"`
	MaxGroup  int       `json:"max_group"`
	Color     string    `json:"color"`
	Width     int       `json:"width"`
	Station   []Station `json:"station"`
}

//Station 地理位置信息
type Station struct {
	ID          int    `json:"id"`
	StationName string `json:"station_name"`
	Longtitude  string `json:"longtitude"`
	Latitude    string `json:"latitude"`
	Type        int    `json:"type"`
	Directive   int    `json:"directive"`
}

//CRHGenerate 生成
func CRHGenerate(trainLists CHSR) string {
	svg := drawsvg.Create()
	svg.SetResize(1)
	svg.SetCircle(trainLists.Circle.Radis, trainLists.Circle.BorderWidth, trainLists.Circle.Color, "#fff")
	for _, one := range trainLists.List {
		var dpath drawsvg.Path
		dpath.Aid = strconv.Itoa(one.TrainID)
		dpath.Alt = one.TrainName
		dpath.Fill = "transparent"
		dpath.Stroke = one.Color
		dpath.StrokeWidth = one.Width
		dpath.FillOpacity = "0.4"
		dpath.MaxGroup = one.MaxGroup + 1
		for i := 0; i < dpath.MaxGroup; i++ {
			var ipaths []drawsvg.IPath
			for _, value := range one.Station {
				if value.Type == i {
					var ipath drawsvg.IPath
					ipath.ID = value.ID
					ipath.Group = value.Type
					ipath.Long = value.Latitude
					ipath.Lat = value.Longtitude
					ipath.Directive = value.Directive
					ipaths = append(ipaths, ipath)
				}
			}
			if len(ipaths) > 0 {
				dpath.PathInfo = ipaths
				svg.SetPath(dpath)
			}
		}
	}
	content := svg.Draw()
	return content
}
