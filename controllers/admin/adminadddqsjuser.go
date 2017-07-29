package admin

/*
添加后台用户
*/
import (
	"qax580go/models"
	"qax580go/qutil"

	"github.com/astaxie/beego"
)

type AdminAddDqsjUserController struct {
	beego.Controller
}

func (c *AdminAddDqsjUserController) Get() {
	bool, username := qutil.ChackAccount(c.Ctx)
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
	c.TplName = "admin/adminadddqsjuser.html"
}
func (c *AdminAddDqsjUserController) Post() {
	username := c.Input().Get("user")
	password := c.Input().Get("password")
	if len(username) != 0 && len(password) != 0 {
		err := models.AddDqsjAdmin(username, password)
		if err != nil {
			beego.Error(err)
		}
		c.Redirect("/admin/dqsj", 302)
	} else {
		c.Redirect("/admin/adddqsjuser", 302)
	}

}
