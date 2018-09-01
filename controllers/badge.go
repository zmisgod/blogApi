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
	var width int = 225
	var height int = 20
	var fSize int = 11
	var rx float64 = 3

	var lText string = "https://gen.zmis.me"
	var lColor string = "#00FFED"
	var lFColor string = "#fff"
	var lLength float64 = 124
	var lX float64 = 62

	var rText string = "author: zmisgod"
	var rColor string = "#5A0988"
	var rFColor string = "#fff"
	var rLength float64 = 101
	var rX float64 = 175

	badge := models.SetBadge(width, height, fSize, rx, lText, lColor, lFColor, lLength, lX, rText, rColor, rFColor, rLength, rX)
	res, err := badge.Template()
	b.CheckError(err)
	b.SendData(res, "ok")
}
