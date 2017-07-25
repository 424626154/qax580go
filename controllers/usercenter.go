package controllers

/**
*个人中心
 */
import (
	"qax580go/models"
	"qax580go/qutil"

	"github.com/astaxie/beego"
)

type UserCenterController struct {
	beego.Controller
}

func (c *UserCenterController) Get() {
	beego.Debug("UserCenterController Get")
	c.Data["Alert"] = ""
	_, isLogin := initUserCenter(c)
	if isLogin == false {
		c.Redirect("/login", 302)
		return
	}
	c.TplName = "usercenter.html"
}

func (c *UserCenterController) Post() {
	beego.Debug("UserCenterController Post")
	c.Data["Alert"] = ""
	c.TplName = "usercenter.html"
}

func initUserCenter(c *UserCenterController) (*models.User, bool) {
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
			beego.Debug(user)
			isLogin = true
			c.Data["User"] = user
		}
	}
	c.Data["isLogin"] = isLogin
	return muser, isLogin
}
