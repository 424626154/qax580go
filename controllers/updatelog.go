package controllers

/*
更新日志
*/
import (
	"qax580go/models"
	"qax580go/qutil"

	"github.com/astaxie/beego"
)

type UpdateLogController struct {
	beego.Controller
}

func (c *UpdateLogController) Get() {
	getULCookie(c)
	c.TplName = "updatelog.html"
}

func (c *UpdateLogController) Post() {
	c.TplName = "updatelog.html"
}

func getULCookie(c *UpdateLogController) string {
	isUser := false
	openid := c.Ctx.GetCookie(qutil.COOKIE_WX_OPENID)
	beego.Debug(openid)
	if len(openid) != 0 {
		wxuser, err := models.GetOneWxUserInfo(openid)
		if err != nil {
			beego.Error(err)
		} else {
			isUser = true
			beego.Debug(wxuser)
			c.Data["WxUser"] = wxuser
		}
	}
	c.Data["isUser"] = isUser
	return openid
}
