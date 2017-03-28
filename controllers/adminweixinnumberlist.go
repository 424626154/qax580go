package controllers

/*
后台推荐微信号列表
*/
import (
	"fmt"
	"github.com/astaxie/beego"
	"qax580go/models"
)

type AdminWeixinNumberListController struct {
	beego.Controller
}

func (c *AdminWeixinNumberListController) Get() {
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
		err := models.DeleteWeixinNumber(id)
		if err != nil {
			beego.Error(err)
		}
		url := "/admin/weixinnumberlist"
		c.Redirect(url, 302)
		return
	case "state":
		id := c.Input().Get("id")
		if len(id) == 0 {
			break
		}
		id = c.Input().Get("id")
		err := models.UpdateWeixinNumber(id, 1)
		if err != nil {
			beego.Error(err)
		}
		c.Redirect("/admin/weixinnumberlist", 302)
		return
	case "state1":
		id := c.Input().Get("id")
		if len(id) == 0 {
			break
		}
		id = c.Input().Get("id")
		err := models.UpdateWeixinNumber(id, 0)
		if err != nil {
			beego.Error(err)
		}
		c.Redirect("/admin/weixinnumberlist", 302)
		return
	case "upinfo":
		id := c.Input().Get("id")
		if len(id) == 0 {
			break
		}
		url := fmt.Sprintf("/admin/upweixinnumberinfo?id=%s", id)
		beego.Debug("up_rul", url)
		c.Redirect(url, 302)
		return
	case "upimg":
		id := c.Input().Get("id")
		if len(id) == 0 {
			break
		}
		url := fmt.Sprintf("/admin/upweixinnumberimg?id=%s", id)
		beego.Debug("up_rul", url)
		c.Redirect(url, 302)
		return
	}

	beego.Debug("AdminWcListController")
	objs, err := models.GetAllWeixinNumbers()
	if err != nil {
		beego.Error(err)
	}
	c.Data["WeixinNumbers"] = objs
	// beego.Debug(objs)
	c.TplName = "adminweixinnumberlist.html"
}
