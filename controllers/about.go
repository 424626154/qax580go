package controllers

/*
关于
*/
import (
	"github.com/astaxie/beego"
	"qax580go/models"
)

type AboutController struct {
	beego.Controller
}

func (c *AboutController) Get() {
	initAboutUser(c)
	c.TplName = "about.html"
}

func (c *AboutController) Post() {
	c.TplName = "about.html"
}

func initAboutUser(c *AboutController) (*models.User, bool) {
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
