package controllers

/*
添加后台用户
*/
import (
	"github.com/astaxie/beego"
	"qax580go/models"
)

type AdminAddUserController struct {
	beego.Controller
}

func (c *AdminAddUserController) Get() {
	bool, username := chackAccount(c.Ctx)
	if bool {

	} else {
		c.Redirect("/admin", 302)
		return
	}
	posts, err := models.GetAllPostsAdmin()
	if err != nil {
		beego.Error(err)
	}
	c.Data["Posts"] = posts
	c.Data["isUser"] = bool
	c.Data["User"] = username
	c.TplName = "adminadduser.html"
}
func (c *AdminAddUserController) Post() {
	username := c.Input().Get("user")
	password := c.Input().Get("password")
	if len(username) != 0 && len(password) != 0 {
		err := models.AddAdmin(username, password)
		if err != nil {
			beego.Error(err)
		}
		c.Redirect("/admin/userlist", 302)
	} else {
		c.Redirect("/admin/adduser", 302)
	}

}
