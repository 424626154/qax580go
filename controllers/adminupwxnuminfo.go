package controllers

/*
后台修改推荐公众号内容
*/
import (
	"github.com/astaxie/beego"
	"qax580go/models"
)

type AdminUpWxnumInfoController struct {
	beego.Controller
}

func (c *AdminUpWxnumInfoController) Get() {
	id := c.Input().Get("id")
	obj, err := models.GetOneWxnum(id)
	if err != nil {
		beego.Error(err)
	}
	c.Data["Wxnum"] = obj
	c.TplName = "adminupwxnuminfo.html"
}
func (c *AdminUpWxnumInfoController) Post() {
	id := c.Input().Get("id")
	title := c.Input().Get("title")
	info := c.Input().Get("info")
	number := c.Input().Get("number")
	evaluate := c.Input().Get("evaluate")
	if len(id) != 0 && len(title) != 0 && len(info) != 0 && len(number) != 0 {
		err := models.UpdateWxnumInfo(id, title, info, number, evaluate)
		if err != nil {
			beego.Error(err)
		}
		c.Redirect("/admin/wxlist", 302)
		return
	}
	c.TplName = "adminupwxnuminfo.html"
}
