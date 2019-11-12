package main

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)


func init() {

	beego.GlobalControllerRouter
	["coursera/frameworks/user/controllers:ObjectController"] = append
	(beego.GlobalControllerRouter
	["coursera/frameworks/user/controllers:ObjectController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil })

	beego.GlobalControllerRouter
	["coursera/frammewowrks/user/controllers:ObjectController"] = append
}
