package home

import (
	"qax580go/models"
	"qax580go/qutil"

	"github.com/astaxie/beego"
)

type WeChat struct {
	beego.Controller
}

func (c *WeChat) WeChats() {
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
		objs, err := models.GetWeChats()
		if err != nil {
			beego.Error(err)
		}
		beego.Debug(objs)
		c.TplName = "home/wechats.html"
		c.Data["WeixinNumbers"] = objs
	}
}
