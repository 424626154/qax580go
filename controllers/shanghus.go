package controllers

/*
信息详情
*/
import (
	"github.com/astaxie/beego"
	"qax580go/models"
)

type ShangHusController struct {
	beego.Controller
}

func (c *ShangHusController) Get() {
	beego.Debug("ShangHusController Get")
	getShanghusCookie(c)
	objs, err := models.GetAllShangHus()
	if err != nil {
		beego.Error(err)
	} else {
		beego.Debug(objs)
	}
	c.Data["Objs"] = objs
	c.TplName = "shanghus.html"
}
func (c *ShangHusController) Post() {
	getShanghusCookie(c)
	objs, err := models.GetAllShangHus()
	if err != nil {
		beego.Error(err)
	} else {
		beego.Debug(objs)
	}
	c.Data["Objs"] = objs
	c.TplName = "shanghus.html"
}
func getShanghusCookie(c *ShangHusController) string {
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
