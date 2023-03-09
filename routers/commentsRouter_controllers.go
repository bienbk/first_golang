package routers

import (
	beego "github.com/beego/beego/v2/server/web"
	"github.com/beego/beego/v2/server/web/context/param"
)

func init() {

    beego.GlobalControllerRouter["bien.com/hocgo/controllers:TestController"] = append(beego.GlobalControllerRouter["bien.com/hocgo/controllers:TestController"],
        beego.ControllerComments{
            Method: "Xinchao",
            Router: "/add",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["bien.com/hocgo/controllers:TestController"] = append(beego.GlobalControllerRouter["bien.com/hocgo/controllers:TestController"],
        beego.ControllerComments{
            Method: "DivFunction",
            Router: "/div",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["bien.com/hocgo/controllers:TestController"] = append(beego.GlobalControllerRouter["bien.com/hocgo/controllers:TestController"],
        beego.ControllerComments{
            Method: "Multipli",
            Router: "/multiplication",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["bien.com/hocgo/controllers:TestController"] = append(beego.GlobalControllerRouter["bien.com/hocgo/controllers:TestController"],
        beego.ControllerComments{
            Method: "SubFunction",
            Router: "/sub",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

}
