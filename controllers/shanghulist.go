package controllers

/*
信息详情
*/
import (
	"github.com/astaxie/beego"
	"qax580go/models"
)

type ShangHuListController struct {
	beego.Controller
}

func (c *ShangHuListController) Get() {
	beego.Debug("ShangHuListController Get")
	getShanghuListCookie(c)
	mytype := c.Input().Get("type")
	if len(mytype) != 0 {
		objs, err := models.GetAllTypeShangHus(mytype)
		if err != nil {
			beego.Error(err)
		} else {
			beego.Debug(objs)
		}
		c.Data["Objs"] = objs
	}

	c.TplName = "shanghulist.html"
}
func (c *ShangHuListController) Post() {
	getShanghuListCookie(c)
	c.TplName = "shanghulist.html"
}
func getShanghuListCookie(c *ShangHuListController) string {
	isUser := false
	openid := c.Ctx.GetCookie(COOKIE_WX_OPENID)
	beego.Debug("------------openid--------")
	beego.Debug(openid)
	if len(openid) != 0 {
		wxuser, err := models.GetOneWxUserInfo(openid)
		if err != nil {
			beego.Error(err)
		} else {
			isUser = true
			beego.Debug("--------------wxuser----------")
			beego.Debug(wxuser)
			c.Data["WxUser"] = wxuser
		}
	}
	c.Data["isUser"] = isUser
	return openid
}
