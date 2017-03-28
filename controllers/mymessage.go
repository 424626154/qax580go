package controllers

/*
我的消息
*/
import (
	"github.com/astaxie/beego"
	"qax580go/models"
)

type MyMessageController struct {
	beego.Controller
}

func (c *MyMessageController) Get() {
	user, islogin := initMyMessageUser(c)
	if islogin == false {
		c.Redirect("/login?from=mymsg", 302)
		return
	}
	beego.Debug("user:", user)
	openid := user.Uid
	op := c.Input().Get("op")
	switch op {
	case "del":
		if len(openid) != 0 {
			id := c.Input().Get("id")
			if len(id) == 0 {
				break
			}
			id = c.Input().Get("id")
			err := models.DeletePostOpenid(id, openid)
			if err != nil {
				beego.Error(err)
			}
			beego.Debug("MyMessageController delete id = %s openid = %s ", id, openid)
			c.Redirect("/mymessage", 302)
		}
		return
	}

	if len(openid) != 0 {
		posts, err := models.GetAllPostsOpenid(openid)
		if err != nil {
			beego.Error(err)
		}
		c.Data["Posts"] = posts
	}
	c.TplName = "mymessage.html"
}

func (c *MyMessageController) Post() {
	c.TplName = "mymessage.html"
}
func initMyMessageUser(c *MyMessageController) (*models.User, bool) {
	isLogin := false
	openid := c.Ctx.GetCookie(COOKIE_UID)
	user := &models.User{}
	beego.Debug(openid)
	if len(openid) != 0 {
		user, err := models.GetOneUserUid(openid)
		if err != nil {
			beego.Error(err)
		} else {
			beego.Debug(user)
			isLogin = true
			c.Data["User"] = user
		}
	}
	c.Data["isLogin"] = isLogin
	return user, isLogin
}
