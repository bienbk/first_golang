package routers

import (
	"bien.com/hocgo/controllers"
	"github.com/beego/beego/v2/server/web"
)

func init() {
	ns :=
		web.NewNamespace("/v1",
			web.NSNamespace("/match",
				web.NSInclude(
					&controllers.TestController{},
				),
			),
		)
	web.AddNamespace(ns)
}