package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

	beego.GlobalControllerRouter["github.com/zmisgod/blogApi/controllers:CategoryController"] = append(beego.GlobalControllerRouter["github.com/zmisgod/blogApi/controllers:CategoryController"],
		beego.ControllerComments{
			Method: "Get",
			Router: `/:categoryId`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/zmisgod/blogApi/controllers:CommentController"] = append(beego.GlobalControllerRouter["github.com/zmisgod/blogApi/controllers:CommentController"],
		beego.ControllerComments{
			Method: "Get",
			Router: `/:articleId`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/zmisgod/blogApi/controllers:CommentController"] = append(beego.GlobalControllerRouter["github.com/zmisgod/blogApi/controllers:CommentController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/comment`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/zmisgod/blogApi/controllers:HomeController"] = append(beego.GlobalControllerRouter["github.com/zmisgod/blogApi/controllers:HomeController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/zmisgod/blogApi/controllers:HomeController"] = append(beego.GlobalControllerRouter["github.com/zmisgod/blogApi/controllers:HomeController"],
		beego.ControllerComments{
			Method: "Get",
			Router: `/:articleId`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/zmisgod/blogApi/controllers:SearchController"] = append(beego.GlobalControllerRouter["github.com/zmisgod/blogApi/controllers:SearchController"],
		beego.ControllerComments{
			Method: "Get",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/zmisgod/blogApi/controllers:TagController"] = append(beego.GlobalControllerRouter["github.com/zmisgod/blogApi/controllers:TagController"],
		beego.ControllerComments{
			Method: "Get",
			Router: `/:tagId`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

}
