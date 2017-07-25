package controllers

/**
*注册验证
 */
import (
	"qax580go/models"
	"qax580go/qutil"

	"github.com/astaxie/beego"
)

type RegisterVerifyController struct {
	beego.Controller
}

func (c *RegisterVerifyController) Get() {
	verify := c.Input().Get("verify")
	c.Data["Alert"] = ""
	if len(verify) != 0 {
		user, err := models.GetOneUserVerify(verify)
		if err != nil {
			beego.Error(err)
		}
		if len(user.Verify) != 0 {
			_, err := models.UpdataUserVerify(user.Id, true)
			if err != nil {
				beego.Error(err)
			}
			maxAge := 1<<31 - 1
			c.Ctx.SetCookie(qutil.COOKIE_UID, user.Uid, maxAge, "/")
			c.Redirect("/", 302)
			return
		} else {
			c.Data["Alert"] = "验证参数错误"
		}
	} else {
		c.Data["Alert"] = "验证连接失效"
	}
	c.TplName = "registerverify.html"
}
func (c *RegisterVerifyController) Post() {
	beego.Debug("post RegisterVerify")
	c.TplName = "registerverify.html"
}
