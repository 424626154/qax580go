package controllers

/*
后台聚合
*/
import (
	"github.com/astaxie/beego"
)

type AdminJuheController struct {
	beego.Controller
}

func (c *AdminJuheController) Get() {
	bool, username := chackAccount(c.Ctx)
	if bool {
		c.Data["isUser"] = bool
		c.Data["User"] = username
	} else {
		c.Redirect("/admin", 302)
		return
	}
	c.TplName = "adminjuhe.html"
}

func (c *AdminJuheController) Post() {
	c.TplName = "adminjuhe.html"
}
