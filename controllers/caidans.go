package controllers

/*
菜单
*/

import (
	"github.com/astaxie/beego"
	"qax580go/models"
)

type CaidansController struct {
	beego.Controller
}

func (c *CaidansController) Get() {
	getCaidanCookie(c)
	id := c.Input().Get("id")
	if len(id) == 0 {
		c.Redirect("/waimailist", 302)
		return
	}
	id = c.Input().Get("id")
	obj, err := models.GetOneCanting(id)
	if err != nil {
		beego.Error(err)
	}
	obj1, err := models.GetAllCaidan(id)
	if err != nil {
		beego.Error(err)
	}
	c.Data["Canting"] = obj
	c.Data["Caidans"] = obj1
	// beego.Debug("Canting:", obj)
	beego.Debug("Caidans:", obj1)
	c.TplName = "caidans.html"
}

func (c *CaidansController) Post() {
	c.TplName = "caidans.html"
}

func getCaidanCookie(c *CaidansController) string {
	isUser := false
	openid := c.Ctx.GetCookie(COOKIE_WX_OPENID)
	if len(openid) != 0 {
		wxuser, err := models.GetOneWxUserInfo(openid)
		if err != nil {
			beego.Error(err)
		} else {
			isUser = true
			c.Data["WxUser"] = wxuser
		}
	}
	c.Data["isUser"] = isUser
	return openid
}
