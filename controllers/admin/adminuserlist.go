package admin

/*
后台微信用户列表
*/
import (
	"qax580go/models"
	"qax580go/qutil"

	"github.com/astaxie/beego"
)

type AdminUserListController struct {
	beego.Controller
}

func (c *AdminUserListController) Get() {
	bool, username := qutil.ChackAccount(c.Ctx)
	if bool {

	} else {
		c.Redirect("/admin", 302)
		return
	}
	c.Data["isUser"] = bool
	c.Data["User"] = username
	c.TplName = "adminuserlist.html"
	admins, err := models.GetAllUsers()
	if err != nil {
		beego.Error(err)
	}
	// beego.Debug(admins)
	c.Data["Users"] = admins
}
func (c *AdminUserListController) Post() {
	op := c.Input().Get("op")
	if op == "del" {
		openid := c.Input().Get("openid")
		if len(openid) != 0 {
			user, err := models.GetOneUserUid(openid)
			if err == nil {
				err = models.DeleteUser(user.Id)
				if err != nil {
					beego.Error(err)
				} else {
					err = models.DeleteUserMoneyRecord(user.Uid)
					if err != nil {
						beego.Error(err)
					}
				}
			} else {
				beego.Error(err)
			}

		}
		url := "/admin/userlist"
		c.Redirect(url, 302)
		return
	}
	c.TplName = "adminuserlist.html"
}
