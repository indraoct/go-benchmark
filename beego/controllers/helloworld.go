package controllers

import "github.com/astaxie/beego"

type HelloWorldController struct {
	beego.Controller
}

// @router /hello [get]
func (this *HelloWorldController) Hello(){
	
	this.Data["json"] = "Hello World"
	this.ServeJSON()
}
