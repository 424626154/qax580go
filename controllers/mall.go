package controllers

/*
商城
*/
import (
	"fmt"
	"qax580go/models"
	"qax580go/qutil"

	"github.com/astaxie/beego"
)

type MallController struct {
	beego.Controller
}

func (c *MallController) Get() {
	beego.Debug("MallController Get")
	initMallUser(c)
	c.TplName = "mall.html"
}
func (c *MallController) Post() {
	beego.Debug("MallController Post")
	initMallUser(c)
	commoditys, err := models.GetAllCommoditys()
	if err != nil {
		beego.Error(err)
	}
	// beego.Debug("commoditys :", commoditys)
	c.Data["Commoditys"] = commoditys
	op := c.Input().Get("op")
	beego.Debug("op", op)
	if op == "exchange" {
		id := c.Input().Get("id")
		beego.Debug("id", id)
		openid := c.Input().Get("openid")
		beego.Debug("openid", openid)
		if len(id) != 0 && len(openid) != 0 {
			err := models.AddUorder(openid, id)
			if err != nil {
				beego.Error(err)
			}
		}
		if len(openid) != 0 {
			url := fmt.Sprintf("/exchange?openid=%s", openid)
			c.Redirect(url, 302)
			return
		}
	}
	c.TplName = "mall.html"
}
func initMallUser(c *MallController) (*models.User, bool) {
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
