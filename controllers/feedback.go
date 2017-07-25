package controllers

/*
意见反馈
*/
import (
	"qax580go/models"
	"qax580go/qutil"

	"github.com/astaxie/beego"
)

type FeedbackController struct {
	beego.Controller
}

func (c *FeedbackController) Get() {
	user, isLogin := initFeedbackUser(c)
	if isLogin == false {
		c.Redirect("/login?from=feedback", 302)
		return
	}
	openid := user.Uid
	info := c.Input().Get("info")
	if len(info) != 0 {
		beego.Debug("------------AddFeedback--------")
		beego.Debug(openid)
		err := models.AddFeedback(info, openid, user.Name, user.Sex, user.Head)
		if err != nil {
			beego.Error(err)
		}
		c.Redirect("/", 302)
	}
	c.TplName = "feedback.html"
}

func initFeedbackUser(c *FeedbackController) (*models.User, bool) {
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
	c.Data["isLogin"] = isLogin
	return muser, isLogin
}
