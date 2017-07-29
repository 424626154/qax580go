package admin

/*
发布消息
*/
import (
	"qax580go/models"
	"qax580go/qutil"
	"strconv"

	"github.com/astaxie/beego"
)

type AdminUpUserMoneyController struct {
	beego.Controller
}

func (c *AdminUpUserMoneyController) Get() {
	id := c.Input().Get("id")
	c.Data["IsId"] = false
	if len(id) != 0 {
		user, err := models.GetOneUserUid(id)
		if err != nil {

		} else {
			c.Data["User"] = user
			c.Data["IsId"] = true
		}
	}
	c.TplName = "adminupusermoney.html"
}

func (c *AdminUpUserMoneyController) Post() {
	id := c.Input().Get("id")
	op := c.Input().Get("op")
	beego.Debug("op :", op)
	switch op {
	case "up":
		beego.Debug("op  ok")
		if len(id) != 0 {
			user, err := models.GetOneUserId(id)
			// beego.Debug("user:", user, id)
			if err != nil {
				beego.Error(err)
			} else {
				moneytype := c.Input().Get("moneytype")
				if len(moneytype) != 0 {
					addmoney := int64(0)
					if moneytype == "1" {
						addmoney = qutil.MONEY_SUBSCRIBE_SUM
					} else if moneytype == "2" {
						addmoney = qutil.MONEY_EXAMINE_SUM
					} else if moneytype == "3" {
						addmoney = qutil.MONEY_BELIKE_SUM
					}
					// beego.Debug("User:", user.Uid)
					err = models.AddUserMoney(user.Uid, addmoney)
					moneytype_i, err := strconv.ParseInt(moneytype, 10, 64)
					if err != nil {
						beego.Debug("err :", err)
					}
					_, err = models.AddUserMoneyRecord(user.Uid, addmoney, moneytype_i)
					if err != nil {
						beego.Error(err)
					}
				}
			}
		}
		c.Redirect("/admin/userlist", 302)
		return
	}
	c.TplName = "adminupusermoney.html"
}
