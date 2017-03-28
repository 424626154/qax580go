package controllers

import (
	"github.com/astaxie/beego"
	"qax580go/models"
)

type WcListController struct {
	beego.Controller
}

func (c *WcListController) Get() {
	initWxlistUser(c)
	wxnums, err := models.GetAllWxnumsState1()
	if err != nil {
		beego.Error(err)
	}
	beego.Debug(wxnums)
	c.TplName = "wxlist.html"
	c.Data["Wxnums"] = wxnums
}
func initWxlistUser(c *WcListController) (*models.User, bool) {
	isLogin := false
	openid := c.Ctx.GetCookie(COOKIE_UID)
	muser := &models.User{}
	beego.Debug(openid)
	if len(openid) != 0 {
		user, err := models.GetOneUserUid(openid)
		if err != nil {
			beego.Error(err)
		} else {
			muser = user
			// beego.Debug(user)
			isLogin = true
			c.Data["User"] = user
		}
	}
	c.Data["isLogin"] = isLogin
	return muser, isLogin
}
