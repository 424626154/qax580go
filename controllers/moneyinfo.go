package controllers

/*
金钱详情
*/
import (
	"github.com/astaxie/beego"
	"qax580go/models"
)

type MoneyInfoController struct {
	beego.Controller
}

func (c *MoneyInfoController) Get() {
	c.TplName = "moneyinfo.html"
}
func (c *MoneyInfoController) Post() {
	openid := getMoneyInfoCookie(c)
	usermoneys, err := models.GetAllUserMoneyRecord(openid)
	beego.Debug("usermoneys :", usermoneys)
	if err != nil {
		beego.Error("err :", err)
	}
	c.Data["UserMoneys"] = usermoneys
	c.TplName = "moneyinfo.html"
}
func getMoneyInfoCookie(c *MoneyInfoController) string {
	isUser := false
	openid := c.Ctx.GetCookie(COOKIE_WX_OPENID)
	beego.Debug("------------openid--------")
	beego.Debug(openid)
	if len(openid) != 0 {
		wxuser, err := models.GetOneWxUserInfo(openid)
		if err != nil {
			beego.Error(err)
		} else {
			isUser = true
			beego.Debug("--------------wxuser----------")
			beego.Debug(wxuser)
			c.Data["WxUser"] = wxuser
		}
	}
	c.Data["isUser"] = isUser
	return openid
}
