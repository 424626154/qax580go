package controllers

/**
*忘记密码验证
 */
import (
	"github.com/astaxie/beego"
	"qax580go/models"
)

type ForgetpwdVerifyController struct {
	beego.Controller
}

func (c *ForgetpwdVerifyController) Get() {
	beego.Debug("get forgetpwdverify")
	c.Data["Alert"] = ""
	verify := c.Input().Get("verify")
	beego.Debug("verify:", verify)
	c.Data["Verify"] = verify
	if len(verify) != 0 {

	} else {
		c.Data["Alert"] = "验证连接失效"
	}
	c.TplName = "forgetpwdverify.html"
}
func (c *ForgetpwdVerifyController) Post() {
	beego.Debug("post forgetpwdverify")
	c.Data["Alert"] = ""
	c.Data["Verify"] = ""
	verify := c.Input().Get("verify")
	pwd := c.Input().Get("pwd")
	beego.Debug("verify:", verify)
	if len(verify) != 0 && len(pwd) != 0 {
		user, err := models.GetOneUserVerify(verify)
		if err != nil {
			beego.Error(err)
		}
		if len(user.Verify) != 0 {

			_, err := models.UpdataUserPassword(user.Id, pwd)
			if err != nil {
				beego.Error(err)
			}
			maxAge := 1<<31 - 1
			c.Ctx.SetCookie(COOKIE_UID, user.Uid, maxAge, "/")
			c.Redirect("/", 302)
			return
		} else {
			c.Data["Alert"] = "验证参数错误"
		}
	} else {
		c.Data["Alert"] = "验证参数错误"
	}
	c.TplName = "forgetpwdverify.html"
}
