package controllers

/*
推荐微信公众号
*/
import (
	"github.com/astaxie/beego"
	"qax580go/models"
)

type WeixinNumberListController struct {
	beego.Controller
}

func (c *WeixinNumberListController) Get() {
	initWeixinNumberListUser(c)
	objs, err := models.GetAllWeixinNumbersState1()
	if err != nil {
		beego.Error(err)
	}
	beego.Debug(objs)
	c.TplName = "weixinnumberlist.html"
	c.Data["WeixinNumbers"] = objs
}
func initWeixinNumberListUser(c *WeixinNumberListController) (*models.User, bool) {
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
