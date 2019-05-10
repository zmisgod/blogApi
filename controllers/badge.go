package controllers

import (
	"github.com/zmisgod/blogApi/models"
)

//BadgeController create baege
type BadgeController struct {
	BaseController
}

//@router / [get]
func (b *BadgeController) Get() {
	var width = 225
	var height = 20
	var fSize = 11
	var rx float64 = 3

	var lText = "https://gen.zmis.me"
	var lColor = "#00FFED"
	var lFColor = "#fff"
	var lLength float64 = 124
	var lX float64 = 62

	var rText = "author: zmisgod"
	var rColor = "#5A0988"
	var rFColor = "#fff"
	var rLength float64 = 101
	var rX float64 = 175

	badge := models.SetBadge(width, height, fSize, rx, lText, lColor, lFColor, lLength, lX, rText, rColor, rFColor, rLength, rX)
	res, err := badge.Template()
	b.CheckError(err)
	b.SendData(res, "ok")
}
