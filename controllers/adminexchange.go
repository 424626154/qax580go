package controllers

/*
后台商城
*/
import (
	"github.com/astaxie/beego"
	"qax580go/models"
)

type AdminExchangeController struct {
	beego.Controller
}

func (c *AdminExchangeController) Get() {
	bool, username := chackAccount(c.Ctx)
	if bool {
		c.Data["isUser"] = bool
		c.Data["User"] = username
	} else {
		c.Redirect("/admin", 302)
		return
	}
	op := c.Input().Get("op")
	switch op {
	case "state":
		id := c.Input().Get("id")
		if len(id) == 0 {
			break
		}
		err := models.UpdateUorderState(id, 1)
		if err != nil {
			beego.Debug(err)
		} else {
			uoder, err := models.GetOneUorder(id)

			if err != nil {
				beego.Error(err)
			} else {
				com, err := models.GetOneCommodity1(uoder.CommodityId)
				if err != nil {
					beego.Error(err)
				} else {
					err = models.ConsumeWxUserMoney(uoder.OpenId, com.Money)
					if err != nil {
						beego.Error(err)
					} else {
						models.AddUserMoneyRecord(uoder.OpenId, com.Money, MONEY_EXCHANGE)
						if err != nil {
							beego.Error(err)
						} else {
						}
					}
				}
			}
		}
		c.Redirect("/admin/exchange", 302)
		return
	case "state1":
		id := c.Input().Get("id")
		if len(id) == 0 {
			break
		}
		err := models.UpdateUorderState(id, 0)
		if err != nil {
			beego.Debug(err)
		} else {

		}
		c.Redirect("/admin/exchange", 302)
		return
	}

	uorders, err := models.GetAllUorders()
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
			user, err := models.GetOneWxUserInfo(uorders[i].OpenId)
			if err != nil {
				beego.Debug(err)
			} else {
				obj.NickeName = user.NickeName
				obj.HeadImgurl = user.HeadImgurl
			}
			showOrders = append(showOrders, obj)
		}

	}
	// beego.Debug("ShowOrders :", showOrders)
	c.Data["ShowOrders"] = showOrders
	c.TplName = "adminexchange.html"
}

func (c *AdminExchangeController) Post() {
	c.TplName = "adminexchange.html"
}
