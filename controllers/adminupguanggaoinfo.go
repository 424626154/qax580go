package controllers

/*
后台修改广告内容
*/
import (
	"github.com/astaxie/beego"
	"qax580go/models"
)

type AdminUpGuanggaoInfoController struct {
	beego.Controller
}

func (c *AdminUpGuanggaoInfoController) Get() {
	id := c.Input().Get("id")
	// beego.Debug(id)
	guangao, err := models.GetOneGuanggao(id)
	if err != nil {
		beego.Error(err)
	}
	c.Data["Guanggao"] = guangao
	c.TplName = "adminupguanggaoinfo.html"
}
func (c *AdminUpGuanggaoInfoController) Post() {
	title := c.Input().Get("title")
	info := c.Input().Get("info")
	blink := c.Input().Get("blink")
	link := c.Input().Get("link")
	id := c.Input().Get("id")
	if len(id) != 0 && len(title) != 0 && len(info) != 0 {
		b_link := false
		s_link := ""
		if blink == "true" {
			b_link = true
			s_link = link
		}

		err := models.UpdateGuanggaoInfo(id, title, info, b_link, s_link)
		if err != nil {
			beego.Error(err)
		}
		c.Redirect("/admin/guanggaos", 302)
	}
	c.TplName = "adminupguanggaoinfo.html"
}
