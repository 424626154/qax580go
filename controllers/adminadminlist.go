package controllers

/*
后台用户列表
*/
import (
	"github.com/astaxie/beego"
	"qax580go/models"
)

type AdminAdminListController struct {
	beego.Controller
}

func (c *AdminAdminListController) Get() {
	bool, username := chackAccount(c.Ctx)
	if bool {

	} else {
		c.Redirect("/admin", 302)
		return
	}
	// posts, err := models.GetAllPostsAdmin()
	// if err != nil {
	// 	beego.Error(err)
	// }
	// c.Data["Posts"] = posts
	c.Data["isUser"] = bool
	c.Data["User"] = username

	admins, err := models.GetAllAdmins()
	if err != nil {
		beego.Error(err)
	}
	c.TplName = "adminadminlist.html"
	c.Data["Admins"] = admins
	beego.Debug(admins)
	op := c.Input().Get("op")
	switch op {
	case "del":
		id := c.Input().Get("id")
		if len(id) == 0 {
			break
		}
		id = c.Input().Get("id")
		err := models.DeleteAdmin(id)
		if err != nil {
			beego.Error(err)
		}
		beego.Debug("is admin del " + id)
		c.Redirect("/admin/adminlist", 302)
		return
	}
}
