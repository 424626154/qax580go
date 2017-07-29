package admin

/*
后台广告详情
*/
import (
	"qax580go/models"

	"github.com/astaxie/beego"
)

type AdminGuanggaoContentController struct {
	beego.Controller
}

func (c *AdminGuanggaoContentController) Get() {
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
		c.TplName = "adminguangaocontent.html"
		return
	}
	c.TplName = "adminguangaocontent.html"

}
