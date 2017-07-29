package admin

/*
通知
*/
import (
	"qax580go/models"
	"qax580go/qutil"

	"github.com/astaxie/beego"
)

type AdminNoticeController struct {
	beego.Controller
}

func (c *AdminNoticeController) Get() {
	bool, username := qutil.ChackAccount(c.Ctx)
	if bool {

	} else {
		c.Redirect("/admin", 302)
		return
	}
	c.Data["isUser"] = bool
	c.Data["User"] = username
	op := c.Input().Get("op")
	switch op {
	case "del":
		id := c.Input().Get("id")
		err := models.DeleteAdminUserNotice(id)
		if err != nil {
			beego.Error(err)
		}
		c.Redirect("/admin/notice", 302)
		return
	}
	objs, err := models.GetAllNotice()
	if err != nil {
		beego.Error(objs)
	}
	c.Data["Objs"] = objs
	c.TplName = "adminnotice.html"
}

func (c *AdminNoticeController) Post() {
	bool, username := qutil.ChackAccount(c.Ctx)
	if bool {

	} else {
		c.Redirect("/admin", 302)
		return
	}
	c.Data["isUser"] = bool
	c.Data["User"] = username
	c.TplName = "adminnotice.html"
}
