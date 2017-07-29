package admin

/*
发布消息
*/
import (
	"qax580go/models"

	"github.com/astaxie/beego"
)

type AdminUpWxuserInfoController struct {
	beego.Controller
}

func (c *AdminUpWxuserInfoController) Get() {

	c.TplName = "adminupwxuserinfo.html"
}

func (c *AdminUpWxuserInfoController) Post() {
	openid := c.Input().Get("openid")
	op := c.Input().Get("op")
	if op == "up" {
		subscribe := c.Input().Get("subscribe")
		if len(subscribe) != 0 {
			err := models.UpWxUserSubscribe(openid, subscribe)
			if err != nil {
				beego.Error(err)
			}
		}
		url := "/admin/wxuserlist"
		c.Redirect(url, 302)
		return
	}
	c.Data["IsOpenid"] = false
	if len(openid) != 0 {
		user, err := models.GetOneWxUserInfo(openid)
		if err == nil {
			c.Data["IsOpenid"] = true
			c.Data["User"] = user
		}

	}
	c.TplName = "adminupwxuserinfo.html"
}
