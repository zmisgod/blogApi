package models

import (
	"strconv"
	"strings"
)

//SvgBadge 自定义的徽章
type SvgBadge struct {
	width  int
	height int
	fSize  int
	rx     float64

	lText   string
	lColor  string
	lFColor string
	lLength float64
	lX      float64

	rText   string
	rColor  string
	rFColor string
	rLength float64
	rX      float64
}

//SetBadge 设置badge信息
func SetBadge(width, height, fSize int, rx float64, lText, lColor, lFColor string, lLength, lX float64, rText, rColor, rFColor string, rLength, rX float64) *SvgBadge {
	var badge SvgBadge
	badge.width = width
	badge.height = height
	badge.fSize = fSize
	badge.rx = rx

	badge.lText = lText
	badge.lColor = lColor
	badge.lFColor = lFColor
	badge.lLength = lLength
	badge.lX = lX

	badge.rText = rText
	badge.rColor = rColor
	badge.rFColor = rFColor
	badge.rLength = rLength
	badge.rX = rX
	return &badge
}

//SaveBadge 保存徽章
func (badge *SvgBadge) SaveBadge() (int, error) {
	return 1, nil
}

//Template 模板
func (badge *SvgBadge) Template() (string, error) {
	result := `<svg xmlns="http://www.w3.org/2000/svg" width="$width" height="$height">
	<metadata>Created by zmisgod. https://gen.zmis.me</metadata>
	<linearGradient id="b" x2="0" y2="100%">
	<stop offset="0" stop-color="#bbb" stop-opacity=".1"></stop>
	<stop offset="1" stop-opacity=".1"></stop>
	</linearGradient> 
	<mask id="a"><rect width="$width" height="$height" rx="$rx" fill="#fff"></rect></mask>
	<g mask="url(#a)">
	<rect width="$l_length" height="$height" fill="$l_color"></rect> 
	<rect x="$l_length" width="$r_length" height="$height" fill="$r_color"></rect> 
	<rect width="$width" height="$height" fill="url(#b)"></rect></g>
	<g text-anchor="middle" font-family="DejaVu Sans,Verdana,Geneva,sans-serif" font-size="$f_size">
	<text x="$l_x" y="15" fill="#010101" fill-opacity=".3">$l_text</text> 
	<text x="$l_x" y="14" fill="$l_f_color" id="left_text">$l_text</text> 
	<text x="$r_x" y="15" fill="#010101" fill-opacity=".3">$r_text</text> 
	<text x="$r_x" y="14" fill="$r_f_color" id="right_text">$r_text</text>
	</g></svg>`
	result = strings.Replace(result, "$width", strconv.Itoa(badge.width), 100)
	result = strings.Replace(result, "$height", strconv.Itoa(badge.height), 100)
	result = strings.Replace(result, "$rx", strconv.FormatFloat(badge.rx, 'E', -1, 64), 100)
	result = strings.Replace(result, "$f_size", strconv.Itoa(badge.fSize), 100)
	result = strings.Replace(result, "$l_text", badge.lText, 100)
	result = strings.Replace(result, "$l_color", badge.lColor, 100)
	result = strings.Replace(result, "$l_f_color", badge.lFColor, 100)
	result = strings.Replace(result, "$l_length", strconv.FormatFloat(badge.lLength, 'E', -1, 64), 100)
	result = strings.Replace(result, "$l_x", strconv.FormatFloat(badge.lX, 'E', -1, 64), 100)
	result = strings.Replace(result, "$r_text", badge.rText, 100)
	result = strings.Replace(result, "$r_color", badge.rColor, 100)
	result = strings.Replace(result, "$r_f_color", badge.rFColor, 100)
	result = strings.Replace(result, "$r_length", strconv.FormatFloat(badge.rLength, 'E', -1, 64), 100)
	result = strings.Replace(result, "$r_x", strconv.FormatFloat(badge.rX, 'E', -1, 64), 100)
	return result, nil
}
