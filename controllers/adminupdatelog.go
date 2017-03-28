package controllers

/*
更新日志
*/
import (
	"github.com/astaxie/beego"
)

type AdminUpdateLogController struct {
	beego.Controller
}

func (c *AdminUpdateLogController) Get() {
	bool, username := chackAccount(c.Ctx)
	if bool {

	} else {
		c.Redirect("/admin", 302)
		return
	}
	c.Data["isUser"] = bool
	c.Data["User"] = username
	c.TplName = "adminupdatelog.html"
}

func (c *AdminUpdateLogController) Post() {
	bool, username := chackAccount(c.Ctx)
	if bool {

	} else {
		c.Redirect("/admin", 302)
		return
	}
	c.Data["isUser"] = bool
	c.Data["User"] = username
	c.TplName = "adminupdatelog.html"
}
