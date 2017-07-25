package home

import (
	"qax580go/qutil"

	"qax580go/models"

	"github.com/astaxie/beego"
)

type WxOfficial struct {
	beego.Controller
}

func (c *WxOfficial) WxOfficials() {
	if c.Ctx.Input.IsGet() {
		isLogin := false
		openid := c.Ctx.GetCookie(qutil.COOKIE_UID)
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
		beego.Debug(muser)
		c.Data["isLogin"] = isLogin
		wxnums, err := models.GetWxOfficials()
		if err != nil {
			beego.Error(err)
		}
		beego.Debug(wxnums)
		c.TplName = "home/wxofficials.html"
		c.Data["Wxnums"] = wxnums
	}
}
