package admin

/*
信息详情
*/
import (
	"qax580go/models"

	"github.com/astaxie/beego"
)

type AdminContentController struct {
	beego.Controller
}

func (c *AdminContentController) Get() {
	op := c.Input().Get("op")
	switch op {
	case "con":
		id := c.Input().Get("id")
		if len(id) == 0 {
			break
		}
		id = c.Input().Get("id")
		post, err := models.GetOnePost(id)
		if err != nil {
			beego.Error(err)
		}
		c.Data["Post"] = post
		beego.Debug("is con " + post.Title)
		c.TplName = "admincontent.html"
		return
	}
	c.TplName = "admincontent.html"

}
