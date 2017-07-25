package controllers

/**
*个人中心
 */
import (
	"qax580go/models"
	"qax580go/qutil"

	"github.com/astaxie/beego"
)

type UpUserController struct {
	beego.Controller
}

func (c *UpUserController) Get() {
	beego.Debug("UpUserController Get")
	c.Data["Alert"] = ""
	_, isLogin := initUpUser(c)
	if isLogin == false {
		c.Redirect("/login", 302)
		return
	}
	uid := c.Input().Get("email")
	beego.Debug("uid:", uid)
	c.TplName = "upuser.html"
}

func (c *UpUserController) Post() {
	beego.Debug("UpUserController Post")
	c.Data["Alert"] = ""
	name := c.Input().Get("name")
	id := c.Input().Get("id")
	if len(name) != 0 && len(id) != 0 {
		user, err := models.GetOneUserId(id)
		if err != nil {
			beego.Error(err)
		}
		_, err = models.UpdataUserInfo(user.Id, name)
		if err != nil {
			beego.Error(err)
		}
		c.Redirect("/usercenter", 302)
		return
	}
	c.TplName = "upuser.html"
}

func initUpUser(c *UpUserController) (*models.User, bool) {
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
