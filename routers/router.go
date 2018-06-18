// @APIVersion 1.0.0
// @Title beego Test API
// @Description beego has a very cool tools to autogenerate documents for your API
// @Contact astaxie@gmail.com
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"github.com/zmisgod/blogApi/controllers"

	"github.com/astaxie/beego"
)

func init() {
	// beego.Router("/", &controllers.StaticController{}, "get:Static")
	ns := beego.NewNamespace("/v1",
		beego.NSNamespace("/tag",
			beego.NSInclude(
				&controllers.TagController{},
			),
		),
		beego.NSNamespace("/article",
			beego.NSInclude(
				&controllers.ArticleController{},
			),
		),
		beego.NSNamespace("/category",
			beego.NSInclude(
				&controllers.CategoryController{},
			),
		),
		beego.NSNamespace("/home",
			beego.NSInclude(
				&controllers.HomeController{},
			),
		),
		beego.NSNamespace("/search",
			beego.NSInclude(
				&controllers.SearchController{},
			),
		),
		beego.NSNamespace("/comment",
			beego.NSInclude(
				&controllers.CommentController{},
			),
		),
		beego.NSNamespace("/link",
			beego.NSInclude(
				&controllers.LinkController{},
			),
		),
	)
	beego.AddNamespace(ns)
}
