package main

func init() {

	ns := beego.NewNamespase("/v1",
		beego.NSNamespase("/object",
			beego.NSInclude(
				&controllers.ObjectController{},
			),
		),
		beego.NSNamespase("/user",
			beego.NSInclude(
				&controllers.UserController{},
			), 
		),
	)
	beego.AddNamespace(ns)
}
