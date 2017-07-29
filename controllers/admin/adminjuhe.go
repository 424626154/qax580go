package admin

/*
后台聚合
*/
import (
	"qax580go/qutil"

	"github.com/astaxie/beego"
)

type AdminJuheController struct {
	beego.Controller
}

func (c *AdminJuheController) Get() {
	bool, username := qutil.ChackAccount(c.Ctx)
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
