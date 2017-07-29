package admin

/*
后台登陆
*/
import (
	"qax580go/models"
	"qax580go/qutil"
	"strings"

	"github.com/astaxie/beego"
)

type AdminLoginController struct {
	beego.Controller
}

func (c *AdminLoginController) Get() {
	c.TplName = "admin/adminlogin.html"
	bool, _ := qutil.ChackAccount(c.Ctx)
	if bool {
		c.Redirect("/admin/home", 302)
		return
	} else {

	}

}

func (c *AdminLoginController) Post() {
	username := c.Input().Get("user")
	password := c.Input().Get("password")
	autologin := c.Input().Get("autologin") == "on"
	if len(username) != 0 && len(password) != 0 {
		admin, err := models.GetOneAdmin(username)
		if err != nil {
			c.Redirect("/admin", 302)
			return
		}
		if admin != nil && strings.EqualFold(username, admin.Username) && strings.EqualFold(password, admin.Password) {
			maxAge := 0
			if autologin {
				maxAge = 1<<31 - 1
			}
			c.Ctx.SetCookie("username", username, maxAge, "/")
			c.Ctx.SetCookie("password", password, maxAge, "/")
			beego.Debug("login ok")
			c.Redirect("/admin/home", 302)
			return
		} else {
			c.Redirect("/admin", 302)
			return
		}
	} else {
		c.Redirect("/admin", 302)
		return
	}
}
