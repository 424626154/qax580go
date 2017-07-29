package admin

/*
后台商户
*/
import (
	"qax580go/models"
	"qax580go/qutil"

	"github.com/astaxie/beego"
)

type AdminShanghusController struct {
	beego.Controller
}

func (c *AdminShanghusController) Get() {
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
		err := models.DeleteShangHu(id)
		if err != nil {
			beego.Error(err)
		}
		beego.Debug("is admin shanghu del " + id)
		c.Redirect("/admin/shanghus", 302)
		return
	case "state":
		id := c.Input().Get("id")
		if len(id) == 0 {
			break
		}
		id = c.Input().Get("id")
		err := models.UpdateShangHuState(id, 1)
		if err != nil {
			beego.Error(err)
		}
		// beego.Debug("is admin state " + id)
		c.Redirect("/admin/shanghus", 302)
		return
	case "state1":
		id := c.Input().Get("id")
		if len(id) == 0 {
			break
		}
		id = c.Input().Get("id")
		err := models.UpdateShangHuState(id, 0)
		if err != nil {
			beego.Error(err)
		}
		// beego.Debug("is admin state1" + id)
		c.Redirect("/admin/shanghus", 302)
		return
	case "recommend":
		id := c.Input().Get("id")
		if len(id) == 0 {
			break
		}
		id = c.Input().Get("id")
		err := models.UpdateShangHuRecommend(id, 1)
		if err != nil {
			beego.Error(err)
		}
		beego.Debug("is admin recommend " + id)
		c.Redirect("/admin/shanghus", 302)
		return
	case "recommend1":
		id := c.Input().Get("id")
		if len(id) == 0 {
			break
		}
		id = c.Input().Get("id")
		err := models.UpdateShangHuRecommend(id, 0)
		if err != nil {
			beego.Error(err)
		}
		beego.Debug("is admin recommend1 " + id)
		c.Redirect("/admin/shanghus", 302)
		return
	}

	objs, err := models.GetAdminAllShangHus()
	if err != nil {
		beego.Error(err)
	}
	c.TplName = "adminshanghus.html"
	c.Data["Objs"] = objs
}

func (c *AdminShanghusController) Post() {
	c.TplName = "adminshanghus.html"
}
