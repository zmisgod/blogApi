package main

import (
	"github.com/zmisgod/blogApi/models"
	_ "github.com/zmisgod/blogApi/routers"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/plugins/cors"
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
	if beego.BConfig.RunMode != "dev" {
		beego.InsertFilter("*", beego.BeforeRouter, cors.Allow(&cors.Options{
			AllowMethods:     []string{"*"},
			AllowHeaders:     []string{"Origin", "Authorization", "Access-Control-Allow-Origin"},
			ExposeHeaders:    []string{"Content-Length", "Access-Control-Allow-Origin"},
			AllowCredentials: false,
			AllowOrigins:     []string{"https://*.zmis.me"},
		}))
	}
	beego.Run()
}
