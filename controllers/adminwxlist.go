package controllers

/*
后台推荐公众号列表
*/
import (
	"fmt"
	"github.com/astaxie/beego"
	"qax580go/models"
)

type AdminWcListController struct {
	beego.Controller
}

func (c *AdminWcListController) Get() {
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
	case "del":
		id := c.Input().Get("id")
		if len(id) == 0 {
			break
		}
		id = c.Input().Get("id")
		err := models.DeleteWxnum(id)
		if err != nil {
			beego.Error(err)
		}
		url := "/admin/wxlist"
		c.Redirect(url, 302)
		return
	case "state":
		id := c.Input().Get("id")
		if len(id) == 0 {
			break
		}
		id = c.Input().Get("id")
		err := models.UpdateWxnum(id, 1)
		if err != nil {
			beego.Error(err)
		}
		c.Redirect("/admin/wxlist", 302)
		return
	case "state1":
		id := c.Input().Get("id")
		if len(id) == 0 {
			break
		}
		id = c.Input().Get("id")
		err := models.UpdateWxnum(id, 0)
		if err != nil {
			beego.Error(err)
		}
		c.Redirect("/admin/wxlist", 302)
		return
	case "upinfo":
		id := c.Input().Get("id")
		if len(id) == 0 {
			break
		}
		url := fmt.Sprintf("/admin/upwxnuminfo?id=%s", id)
		beego.Debug("up_rul", url)
		c.Redirect(url, 302)
		return
	case "upimg":
		id := c.Input().Get("id")
		if len(id) == 0 {
			break
		}
		url := fmt.Sprintf("/admin/upwxnumimg?id=%s", id)
		beego.Debug("up_rul", url)
		c.Redirect(url, 302)
		return
	}

	beego.Debug("AdminWcListController")
	wxnums, err := models.GetAllWxnums()
	if err != nil {
		beego.Error(err)
	}
	beego.Debug(wxnums)
	c.TplName = "adminwxlist.html"
	c.Data["Wxnums"] = wxnums
}
