package models

//CustomBadge 自定义的徽章
type CustomBadge struct {
	width   int
	height  int
	rx      float64
	fSize   int
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

//SaveBadge 保存徽章
func SaveBadge(customBadge CustomBadge) (int, error) {
	return 1, nil
}

func template(CustomBadge CustomBadge) string {
	return `<svg xmlns="http://www.w3.org/2000/svg" width="$width" height="$height">
	<metadata>
	Created by zmisgod. https://gen.zmis.me
	</metadata>
	<linearGradient id="b" x2="0" y2="100%">
	<stop offset="0" stop-color="#bbb" stop-opacity=".1"></stop>
	<stop offset="1" stop-opacity=".1"></stop>
	</linearGradient> 
	<mask id="a">
	<rect width="$width" height="$height" rx="$rx" fill="#fff"></rect>
	</mask>
	<g mask="url(#a)">
	<rect width="$l_length" height="$height" fill="$l_color"></rect> 
	<rect x="$l_length" width="$r_length" height="$height" fill="$r_color"></rect> 
	<rect width="$width" height="$height" fill="url(#b)"></rect>
	</g>
	<g text-anchor="middle" font-family="DejaVu Sans,Verdana,Geneva,sans-serif" font-size="$f_size">
	<text x="$l_x" y="15" fill="#010101" fill-opacity=".3">$l_text</text> 
	<text x="$l_x" y="14" fill="$l_f_color" id="left_text">$l_text</text> 
	<text x="$r_x" y="15" fill="#010101" fill-opacity=".3">$r_text</text> 
	<text x="$r_x" y="14" fill="$r_f_color" id="right_text">$r_text</text>
	</g>
	</svg>`
}
