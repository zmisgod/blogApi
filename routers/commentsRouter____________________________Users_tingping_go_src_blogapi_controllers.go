package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

	beego.GlobalControllerRouter["blogapi/controllers:CategoryController"] = append(beego.GlobalControllerRouter["blogapi/controllers:CategoryController"],
		beego.ControllerComments{
			Method: "Get",
			Router: `/:categoryId`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["blogapi/controllers:HomeController"] = append(beego.GlobalControllerRouter["blogapi/controllers:HomeController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["blogapi/controllers:HomeController"] = append(beego.GlobalControllerRouter["blogapi/controllers:HomeController"],
		beego.ControllerComments{
			Method: "Get",
			Router: `/:articleId`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["blogapi/controllers:SearchController"] = append(beego.GlobalControllerRouter["blogapi/controllers:SearchController"],
		beego.ControllerComments{
			Method: "Get",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["blogapi/controllers:TagController"] = append(beego.GlobalControllerRouter["blogapi/controllers:TagController"],
		beego.ControllerComments{
			Method: "Get",
			Router: `/:tagId`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

}
