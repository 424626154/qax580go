package admin

/*
后台金钱详情
*/
import (
	"qax580go/models"

	"github.com/astaxie/beego"
)

type AdminMoneyInfoController struct {
	beego.Controller
}

func (c *AdminMoneyInfoController) Get() {
	beego.Debug("请求方式 get")

	c.TplName = "adminmoneyinfo.html"
}
func (c *AdminMoneyInfoController) Post() {

	openid := c.Input().Get("openid")
	beego.Debug("请求方式 post openid:", openid)
	if len(openid) != 0 {

		usermoneys, err := models.GetAllUserMoneyRecord(openid)
		beego.Debug("usermoneys :", usermoneys)
		if err != nil {
			beego.Error("err :", err)
		}
		c.Data["UserMoneys"] = usermoneys
	}
	c.TplName = "adminmoneyinfo.html"
}
