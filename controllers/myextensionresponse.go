package controllers

/*
我的推广
*/
import (
	"github.com/astaxie/beego"
)

type MyExtensionResponseController struct {
	beego.Controller
}

func (c *MyExtensionResponseController) Get() {
	openid := c.Input().Get("openid")
	c.Data["Openid"] = openid
	c.TplName = "myextensionresponse.html"
}
