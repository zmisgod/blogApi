package main

import (
	"blogapi/models"
	_ "blogapi/routers"

	"github.com/astaxie/beego"
)

func init() {
	// cache, err := cache.NewCache("redis", `{"key":"collectionName","conn":":6039","dbNum":"0","password":""}`)
	models.Init()
}
func main() {
	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}
	beego.SetStaticPath("/static", "static")
	beego.Run()
}
