package controllers

/*
我的系统消息
*/
import (
	"qax580go/models"
	"qax580go/qutil"

	"github.com/astaxie/beego"
)

type MynoticeController struct {
	beego.Controller
}

func (c *MynoticeController) Get() {
	user, isLogin := initMynoticeUser(c)
	if isLogin == false {
		c.Redirect("/login?from=mynotice", 302)
		return
	}
	openid := user.Uid

	op := c.Input().Get("op")
	beego.Debug("op", op)
	switch op {
	case "del":
		id := c.Input().Get("id")
		if len(id) != 0 {
			err := models.DeleteUserNotice(id)
			if err != nil {
				beego.Error(err)
			}
		}
		beego.Debug("del id", id)
		c.Redirect("/mynotice", 302)
		return
	}

	objs, err := models.GetUeserAllNotice(openid)
	if err != nil {
		beego.Error(err)
	}
	for i := 0; i < len(objs); i++ {
		err := models.UpUeserNoticeRead(objs[i].Id, 1)
		if err != nil {
			beego.Error(err)
		} else {
			objs[i].ToRead = 1
		}
	}
	beego.Debug(objs)
	c.Data["Objs"] = objs
	c.TplName = "mynotice.html"

}

func initMynoticeUser(c *MynoticeController) (*models.User, bool) {
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
