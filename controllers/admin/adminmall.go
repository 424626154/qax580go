package admin

/*
后台商城
*/
import (
	"qax580go/models"
	"qax580go/qutil"

	"github.com/astaxie/beego"
)

type AdminMallController struct {
	beego.Controller
}

func (c *AdminMallController) Get() {
	bool, username := qutil.ChackAccount(c.Ctx)
	if bool {
		c.Data["isUser"] = bool
		c.Data["User"] = username
	} else {
		c.Redirect("/admin", 302)
		return
	}
	op := c.Input().Get("op")
	switch op {
	case "del":
		id := c.Input().Get("id")
		if len(id) == 0 {
			break
		}
		id = c.Input().Get("id")
		err := models.DeleteCommodity(id)
		if err != nil {
			beego.Error(err)
		}
		// beego.Debug("is admin del " + id)
		c.Redirect("/admin/mall", 302)
		return
	case "state":
		id := c.Input().Get("id")
		if len(id) == 0 {
			break
		}
		id = c.Input().Get("id")
		err := models.UpdateCommodityState(id, 1)
		if err != nil {
			beego.Error(err)
		}
		// beego.Debug("is admin state " + id)
		c.Redirect("/admin/mall", 302)
		return
	case "state1":
		id := c.Input().Get("id")
		if len(id) == 0 {
			break
		}
		id = c.Input().Get("id")
		err := models.UpdateCommodityState(id, 0)
		if err != nil {
			beego.Error(err)
		}
		// beego.Debug("is admin state1" + id)
		c.Redirect("/admin/mall", 302)
		return
	}

	commoditys, err := models.GetAllCommoditysAdmin()
	if err != nil {
		beego.Error(err)
	}
	c.TplName = "adminmall.html"
	c.Data["Commoditys"] = commoditys
}

func (c *AdminMallController) Post() {
	c.TplName = "adminmall.html"
}
