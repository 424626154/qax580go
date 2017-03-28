package controllers

/*
登录
*/
import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
	"qax580go/models"
)

type LoginController struct {
	beego.Controller
}

func (this *LoginController) Get() {
	beego.Debug("login Get")
	this.Data["Alert"] = ""
	from := this.Input().Get("from")
	beego.Debug("from:", from)
	this.Data["From"] = from
	this.TplName = "login.html"
}

func (this *LoginController) Post() {
	beego.Debug("login Post")
	this.Data["Alert"] = ""
	// 获取表单信息
	email := this.Input().Get("email")
	pwd := this.Input().Get("pwd")
	from := this.Input().Get("from")
	beego.Debug("from:", from)
	autoLogin := this.Input().Get("autoLogin") == "on"
	if len(email) != 0 && len(pwd) != 0 {
		user, err := models.GetOneUser(email)
		if err != nil {
			beego.Error(err)
		}
		if len(user.Email) != 0 {
			if pwd == user.Password {
				maxAge := 0
				if autoLogin {
					maxAge = 1<<31 - 1
				}

				this.Ctx.SetCookie(COOKIE_UID, user.Uid, maxAge, "/")
				if len(from) != 0 {
					switch from {
					case "uplode":
						this.Redirect("/uplode", 302)
						break
					case "mymsg":
						this.Redirect("/mymessage", 302)
						break
					case "feedback":
						this.Redirect("/feedback", 302)
						break
					case "mynotice":
						this.Redirect("/mynotice", 302)
						break
					case "mymoney":
						this.Redirect("/mymoney", 302)
						break
					default:
						this.Redirect("/", 302)
						break
					}
				} else {
					this.Redirect("/", 302)
				}
				return
			} else {
				this.Data["Alert"] = "密码错误"
			}
		} else {
			this.Data["Alert"] = "邮箱未注册"
		}
	} else {

	}
	this.TplName = "login.html"
}

func checkAccount(ctx *context.Context) bool {
	ck, err := ctx.Request.Cookie("uname")
	if err != nil {
		return false
	}

	uname := ck.Value

	ck, err = ctx.Request.Cookie("pwd")
	if err != nil {
		return false
	}

	pwd := ck.Value
	return uname == beego.AppConfig.String("adminName") &&
		pwd == beego.AppConfig.String("adminPass")
}
