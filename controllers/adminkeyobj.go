package controllers

/*
后台关键字内容
*/
import (
	"fmt"
	"github.com/astaxie/beego"
	"qax580go/models"
)

type AdminKeyobjController struct {
	beego.Controller
}

func (c *AdminKeyobjController) Get() {
	beego.Debug("AdminKeyobjController Get")
	bool, username := chackAccount(c.Ctx)
	if bool {
		c.Data["isUser"] = bool
		c.Data["User"] = username
	} else {
		c.Redirect("/admin", 302)
		return
	}
	op := c.Input().Get("op")
	keyid := c.Input().Get("keyid")
	switch op {
	case "del":
		id := c.Input().Get("id")
		if len(id) == 0 {
			break
		}
		id = c.Input().Get("id")
		err := models.DeleteKeyobj(id)
		if err != nil {
			beego.Error(err)
		}
		// beego.Debug("is admin del " + id)
		url := fmt.Sprintf("/admin/keyobj?keyid=%s", keyid)
		c.Redirect(url, 302)
		return
	case "state":
		id := c.Input().Get("id")
		if len(id) == 0 {
			break
		}
		id = c.Input().Get("id")
		err := models.UpdateKeyobjState(id, 1)
		if err != nil {
			beego.Error(err)
		}
		// beego.Debug("is admin state " + id)
		url := fmt.Sprintf("/admin/keyobj?keyid=%s", keyid)
		c.Redirect(url, 302)
		return
	case "state1":
		id := c.Input().Get("id")
		if len(id) == 0 {
			break
		}
		id = c.Input().Get("id")
		err := models.UpdateKeyobjState(id, 0)
		if err != nil {
			beego.Error(err)
		}
		// beego.Debug("is admin state1" + id)
		url := fmt.Sprintf("/admin/keyobj?keyid=%s", keyid)
		c.Redirect(url, 302)
		return
	}
	commoditys, err := models.GetAllKeyobj(keyid)
	if err != nil {
		beego.Error(err)
	}
	c.TplName = "adminkeyobj.html"
	beego.Debug("keyid :", keyid)
	c.Data["KeyId"] = keyid
	c.Data["Objs"] = commoditys
}

func (c *AdminKeyobjController) Post() {
	beego.Debug("AdminKeyobjController Post")
	c.TplName = "adminkeyobj.html"
}
