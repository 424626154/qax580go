package admin

/*
后台微信用户列表
*/
import (
	"qax580go/models"
	"qax580go/qutil"

	"github.com/astaxie/beego"
)

type WxUserListController struct {
	beego.Controller
}

func (c *WxUserListController) Get() {
	bool, username := qutil.ChackAccount(c.Ctx)
	if bool {

	} else {
		c.Redirect("/admin", 302)
		return
	}
	c.Data["isUser"] = bool
	c.Data["User"] = username
	c.TplName = "adminwxuserlist.html"
	admins, err := models.GetAllWxUsers()
	if err != nil {
		beego.Error(err)
	}
	beego.Debug(admins)
	c.Data["WxUsers"] = admins
}
func (c *WxUserListController) Post() {
	op := c.Input().Get("op")
	if op == "del" {
		openid := c.Input().Get("openid")
		if len(openid) != 0 {
			user, err := models.GetOneWxUserInfo(openid)
			if err == nil {
				err = models.DeleteWxUser(user.Id)
				if err != nil {
					beego.Error(err)
				} else {
					err = models.DeleteUserMoneyRecord(user.OpenId)
					if err != nil {
						beego.Error(err)
					}
				}
			} else {
				beego.Error(err)
			}

		}
		url := "/admin/wxuserlist"
		c.Redirect(url, 302)
		return
	}
	c.TplName = "adminwxuserlist.html"
}
