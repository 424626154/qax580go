package admin

/*
后台添加菜单
*/
import (
	"fmt"
	"qax580go/models"
	"qax580go/qutil"

	"github.com/astaxie/beego"
)

type AdminCaidansController struct {
	beego.Controller
}

func (c *AdminCaidansController) Get() {
	bool, username := qutil.ChackAccount(c.Ctx)
	if bool {
		c.Data["isUser"] = bool
		c.Data["User"] = username
	} else {
		c.Redirect("/admin", 302)
		return
	}
	id := c.Input().Get("id")
	if len(id) == 0 {
		c.Redirect("/admin/waimailist", 302)
		return
	}
	id = c.Input().Get("id")
	obj, err := models.GetOneCanting(id)
	if err != nil {
		beego.Error(err)
	}
	obj1, err := models.GetAllCaidan(id)
	if err != nil {
		beego.Error(err)
	}
	c.Data["Canting"] = obj
	c.Data["Caidans"] = obj1
	// beego.Debug("Canting:", obj)
	beego.Debug("Caidans:", obj1)
	op := c.Input().Get("op")
	switch op {
	case "del":
		fid := c.Input().Get("fid")
		id := c.Input().Get("id")
		if len(id) == 0 {
			break
		}
		id = c.Input().Get("id")
		err := models.DeleteCaidan(id)
		if err != nil {
			beego.Error(err)
		}
		url := fmt.Sprintf("/admin/caidans?id=%s", fid)
		c.Redirect(url, 302)
		return
	}
	c.TplName = "admincaidans.html"
}

func (c *AdminCaidansController) Post() {
	c.TplName = "admincaidans.html"
}
