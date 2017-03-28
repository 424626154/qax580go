package controllers

/**
洗相
*/
import (
	"github.com/astaxie/beego"
)

type WeiZhanController struct {
	beego.Controller
}

func (c *WeiZhanController) Home() {
	if c.Ctx.Input.IsGet() {
		beego.Debug("Home Get")
	}
	if c.Ctx.Input.IsPost() {
		beego.Debug("Home Post")
	}
	c.TplName = "wzhome.html"
}
