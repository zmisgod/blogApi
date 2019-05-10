package controllers

import (
	"fmt"
	"strconv"
	"time"

	"github.com/astaxie/beego"
	"github.com/zmisgod/blogApi/util"
)

//AdminUploadController 上传图片
type AdminUploadController struct {
	AdminController
}

//@router /image [post]
func (a *AdminUploadController) Image() {
	fileSize, _ := a.GetInt("size", 0)
	fmt.Println(fileSize)
	systemMaxUploadFileSize, err := beego.AppConfig.Int("MaxUploadFileSize")
	if fileSize <= 0 {
		a.SendError("图片尺寸为空")
	}
	if fileSize > systemMaxUploadFileSize {
		a.SendError("超过了最大上传大小的图片")
	}
	fileName := a.GetString("name")
	if fileName == "" {
		a.SendError("图片名称为空")
	}
	staticPath := beego.AppConfig.String("StaticPath")
	floderName := util.TodayFloderName(staticPath)
	fileName = util.Md5String(strconv.Itoa(int(time.Now().Unix()))) + fileName
	fullPath := floderName + fileName
	err = a.SaveToFile("file", fullPath)
	if err != nil {
		a.SendError(err.Error())
	} else {
		staticHost := beego.AppConfig.String("StaticPrefix")
		resultURL := staticHost + "/" + fullPath
		a.SendData(resultURL, "ok")
	}
}
