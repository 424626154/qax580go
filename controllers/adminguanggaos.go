package controllers

/*
后台广告列表
*/
import (
	"github.com/astaxie/beego"
	"qax580go/models"
)

type AdminGuanggaosController struct {
	beego.Controller
}

func (c *AdminGuanggaosController) Get() {
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
		err := models.DeleteGuanggao(id)
		if err != nil {
			beego.Error(err)
		}
		// beego.Debug("is admin del " + id)
		c.Redirect("/admin/guanggaos", 302)
		return
	case "state":
		id := c.Input().Get("id")
		if len(id) == 0 {
			break
		}
		id = c.Input().Get("id")
		err := models.UpdateGuanggao(id, 1)
		if err != nil {
			beego.Error(err)
		}
		// beego.Debug("is admin state " + id)
		c.Redirect("/admin/guanggaos", 302)
		return
	case "state1":
		id := c.Input().Get("id")
		if len(id) == 0 {
			break
		}
		id = c.Input().Get("id")
		err := models.UpdateGuanggao(id, 0)
		if err != nil {
			beego.Error(err)
		}
		// beego.Debug("is admin state1" + id)
		c.Redirect("/admin/guanggaos", 302)
		return
	}

	guanggaos, err := models.GetAllGuanggaos()
	if err != nil {
		beego.Error(err)
	}
	c.TplName = "adminguanggaos.html"
	c.Data["Guanggaos"] = guanggaos
	// beego.Debug(guanggaos)
}
