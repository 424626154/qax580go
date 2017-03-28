package controllers

/*
后台登陆
*/
import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
	"qax580go/models"
	"strings"
)

type AdminLoginController struct {
	beego.Controller
}

func (c *AdminLoginController) Get() {
	c.TplName = "adminlogin.html"
	bool, _ := chackAccount(c.Ctx)
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

func chackAccount(ctx *context.Context) (bool, string) {
	ck, err := ctx.Request.Cookie("username")
	if err != nil {
		return false, ""
	}

	username := ck.Value

	ck, err = ctx.Request.Cookie("password")
	if err != nil {
		return false, ""
	}

	password := ck.Value

	admin, err := models.GetOneAdmin(username)
	if err != nil {
		return false, ""
	}
	if admin != nil && strings.EqualFold(username, admin.Username) && strings.EqualFold(password, admin.Password) {
		beego.Debug(" cookie username ", username)
		return true, username
	} else {
		return false, username
	}
}
