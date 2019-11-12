package main

import (
	"coursera/frameworks/web_bp/controllers"
	"github.com/astaxie/beego"
)


func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/custom", &controllers.MainController{}, "*:Custom")
}
