package controllers

/*
商城
*/
import (
	"github.com/astaxie/beego"
	"qax580go/models"
)

type ExchangeController struct {
	beego.Controller
}

func (c *ExchangeController) Get() {
	beego.Debug("ExchangeController Get")
	openid := getExchangeeCookie(c)
	uorders, err := models.GetAllUserUorders(openid)
	if err != nil {
		beego.Error(err)
	}
	var showOrders []models.ShowOrder
	for i := 0; i < len(uorders); i++ {
		com, err := models.GetOneCommodity1(uorders[i].CommodityId)
		if err != nil {
			beego.Error(err)
		} else {
			obj := models.ShowOrder{Id: uorders[i].Id, OpenId: uorders[i].OpenId, CommodityId: uorders[i].CommodityId,
				State: uorders[i].State, CreateTime: uorders[i].CreateTime, Time: uorders[i].Time,
				ExchangeTime: uorders[i].ExchangeTime, Time1: uorders[i].Time1, Commodity: com}
			// beego.Debug("com", com)
			// beego.Debug("obj", obj)
			// showOrders[i] = obj
			showOrders = append(showOrders, obj)
		}

	}
	c.Data["ShowOrders"] = showOrders
	beego.Debug(showOrders)
	c.TplName = "exchange.html"
}
func (c *ExchangeController) Post() {
	beego.Debug("ExchangeController Post")
	openid := getExchangeeCookie(c)
	uorders, err := models.GetAllUserUorders(openid)
	if err != nil {
		beego.Error(err)
	}
	var showOrders []models.ShowOrder
	for i := 0; i < len(uorders); i++ {
		com, err := models.GetOneCommodity1(uorders[i].CommodityId)
		if err != nil {
			beego.Error(err)
		} else {
			obj := models.ShowOrder{Id: uorders[i].Id, OpenId: uorders[i].OpenId, CommodityId: uorders[i].CommodityId,
				State: uorders[i].State, CreateTime: uorders[i].CreateTime, Time: uorders[i].Time,
				ExchangeTime: uorders[i].ExchangeTime, Time1: uorders[i].Time1, Commodity: com}
			// beego.Debug("com", com)
			// beego.Debug("obj", obj)
			// showOrders[i] = obj
			showOrders = append(showOrders, obj)
		}

	}
	c.Data["ShowOrders"] = showOrders
	beego.Debug(showOrders)
	c.TplName = "exchange.html"
}
func getExchangeeCookie(c *ExchangeController) string {
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
