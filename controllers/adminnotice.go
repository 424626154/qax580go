package controllers

/*
通知
*/
import (
	"github.com/astaxie/beego"
	"qax580go/models"
)

type AdminNoticeController struct {
	beego.Controller
}

func (c *AdminNoticeController) Get() {
	bool, username := chackAccount(c.Ctx)
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
	bool, username := chackAccount(c.Ctx)
	if bool {

	} else {
		c.Redirect("/admin", 302)
		return
	}
	c.Data["isUser"] = bool
	c.Data["User"] = username
	c.TplName = "adminnotice.html"
}
