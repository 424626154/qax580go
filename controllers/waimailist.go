package controllers

/*
外卖列表
*/
import (
	"github.com/astaxie/beego"
	"qax580go/models"
)

type WaimaiListController struct {
	beego.Controller
}

func (c *WaimaiListController) Get() {
	getWaimaiListCookie(c)
	objs, err := models.GetAllCantingState1()
	if err != nil {
		beego.Error(err)
	}
	beego.Debug(objs)
	c.Data["Objs"] = objs
	c.TplName = "waimailist.html"

}

func (c *WaimaiListController) Post() {
	c.TplName = "waimailist.html"
}

func getWaimaiListCookie(c *WaimaiListController) string {
	isUser := false
	openid := c.Ctx.GetCookie(COOKIE_WX_OPENID)
	if len(openid) != 0 {
		wxuser, err := models.GetOneWxUserInfo(openid)
		if err != nil {
			beego.Error(err)
		} else {
			isUser = true
			c.Data["WxUser"] = wxuser
		}
	}
	c.Data["isUser"] = isUser
	return openid
}
