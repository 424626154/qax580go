package admin

/*
后台修改广告内容
*/
import (
	"qax580go/models"

	"github.com/astaxie/beego"
)

type AdminUpCommodityInfoController struct {
	beego.Controller
}

func (c *AdminUpCommodityInfoController) Get() {
	id := c.Input().Get("id")
	// beego.Debug("AdminUpCommodityInfoController id:", id)
	obj, err := models.GetOneCommodity(id)
	if err != nil {
		beego.Error(err)
	}
	// beego.Debug("AdminUpCommodityInfoController Commodity:", obj)
	c.Data["Commodity"] = obj
	c.TplName = "adminupcommodityinfo.html"
}
func (c *AdminUpCommodityInfoController) Post() {
	title := c.Input().Get("title")
	info := c.Input().Get("info")
	money := c.Input().Get("money")
	id := c.Input().Get("id")
	if len(id) != 0 && len(title) != 0 && len(info) != 0 && len(money) != 0 {
		err := models.UpdateCommodityInfo(id, title, info, money)
		if err != nil {
			beego.Error(err)
		}
		c.Redirect("/admin/mall", 302)
	}
	c.TplName = "adminupcommodityinfo.html"
}
