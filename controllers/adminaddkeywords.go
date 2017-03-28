package controllers

/*
添加关键字
*/
import (
	"github.com/astaxie/beego"
	"qax580go/models"
)

type AdminaAddKeywordsController struct {
	beego.Controller
}

func (c *AdminaAddKeywordsController) Get() {
	bool, username := chackAccount(c.Ctx)
	if bool {
		c.Data["isUser"] = bool
		c.Data["User"] = username
	} else {
		c.Redirect("/admin", 302)
		return
	}
	c.Data["Image"] = ""
	c.TplName = "adminaddkeywords.html"

}
func (c *AdminaAddKeywordsController) Post() {
	title := c.Input().Get("title")
	if len(title) != 0 {
		beego.Debug("key", title)
		err := models.AddKeywords(title)
		if err != nil {
			beego.Error(err)
		}
		c.Redirect("/admin/keywords", 302)
		return
	}
	c.TplName = "adminaddkeywords.html"

}
