package controllers

/*
后台外卖列表
*/
import (
	"github.com/astaxie/beego"
	"qax580go/models"
)

type AdminWaimaiListController struct {
	beego.Controller
}

func (c *AdminWaimaiListController) Get() {
	bool, username := chackAccount(c.Ctx)
	if bool {
		c.Data["isUser"] = bool
		c.Data["User"] = username
	} else {
		c.Redirect("/admin", 302)
		return
	}
	objs, err := models.GetAllCanting()
	if err != nil {
		beego.Error(err)
	}
	c.Data["Objs"] = objs
	beego.Debug("Objs :", objs)
	op := c.Input().Get("op")
	switch op {
	case "del":
		id := c.Input().Get("id")
		if len(id) == 0 {
			break
		}
		id = c.Input().Get("id")
		err := models.DeleteCanting(id)
		if err != nil {
			beego.Error(err)
		}
		c.Redirect("/admin/waimailist", 302)
		return
	case "state":
		id := c.Input().Get("id")
		if len(id) == 0 {
			break
		}
		id = c.Input().Get("id")
		err := models.UpdateCanting(id, 1)
		if err != nil {
			beego.Error(err)
		}
		beego.Debug("op = state")
		c.Redirect("/admin/waimailist", 302)
		return
	case "state1":
		id := c.Input().Get("id")
		if len(id) == 0 {
			break
		}
		id = c.Input().Get("id")
		err := models.UpdateCanting(id, 0)
		if err != nil {
			beego.Error(err)
		}
		c.Redirect("/admin/waimailist", 302)
		return
	}
	c.TplName = "adminwaimailist.html"
}
