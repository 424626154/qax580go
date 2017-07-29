package admin

/*
更新日志
*/
import (
	"qax580go/qutil"

	"github.com/astaxie/beego"
)

type AdminUpdateLogController struct {
	beego.Controller
}

func (c *AdminUpdateLogController) Get() {
	bool, username := qutil.ChackAccount(c.Ctx)
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
	bool, username := qutil.ChackAccount(c.Ctx)
	if bool {

	} else {
		c.Redirect("/admin", 302)
		return
	}
	c.Data["isUser"] = bool
	c.Data["User"] = username
	c.TplName = "adminupdatelog.html"
}
