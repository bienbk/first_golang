package routers

import (
	"github.com/beego/beego/v2/server/web"
	"github.com/first_golang/controllers"
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
