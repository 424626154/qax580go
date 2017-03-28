package controllers

/*
我的帮帮币
*/
import (
	"github.com/astaxie/beego"
	"qax580go/models"
)

type MymoneyController struct {
	beego.Controller
}

func (c *MymoneyController) Get() {
	user, isLogin := initMymoneyUser(c)
	if isLogin == false {
		c.Redirect("/login?from=mymoney", 302)
		return
	}
	openid := user.Uid
	beego.Debug(openid)
	c.TplName = "mymoney.html"

}

func initMymoneyUser(c *MymoneyController) (*models.User, bool) {
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
