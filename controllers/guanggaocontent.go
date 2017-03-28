package controllers

/*
广告详情
*/
import (
	"github.com/astaxie/beego"
	"qax580go/models"
)

type GuanggaoContentController struct {
	beego.Controller
}

func (c *GuanggaoContentController) Get() {
	op := c.Input().Get("op")
	switch op {
	case "con":
		id := c.Input().Get("id")
		if len(id) == 0 {
			break
		}
		guangao, err := models.GetOneGuanggao(id)
		if err != nil {
			beego.Error(err)
		}
		c.Data["Guanggao"] = guangao
		beego.Debug("guangao :", guangao)
		c.TplName = "guanggaocontent.html"
		return
	}
	c.TplName = "guanggaocontent.html"

}
