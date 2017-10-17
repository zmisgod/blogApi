package main

import (
	"blogapi/models"
	_ "blogapi/routers"

	"github.com/astaxie/beego"
)

func init() {
	models.Init()
}
func main() {
	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}
	beego.Run()
}
