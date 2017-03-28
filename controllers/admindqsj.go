package controllers

/*
后台用户列表
*/
import (
	"github.com/astaxie/beego"
	"qax580go/models"
)

type AdminDqsjUserListController struct {
	beego.Controller
}

func (c *AdminDqsjUserListController) Get() {
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

	admins, err := models.GetAllDqsjAdmins()
	if err != nil {
		beego.Error(err)
	}
	c.TplName = "admindqsj.html"
	beego.Debug(admins)
	c.Data["Admins"] = admins
	beego.Error(admins)
	op := c.Input().Get("op")
	switch op {
	case "del":
		id := c.Input().Get("id")
		if len(id) == 0 {
			break
		}
		id = c.Input().Get("id")
		err := models.DeleteDqsjAdmin(id)
		if err != nil {
			beego.Error(err)
		}
		beego.Debug("is admin del " + id)
		c.Redirect("/admin/dqsj", 302)
		return
	}
}
