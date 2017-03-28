package controllers

/*
推荐测试
*/
import (
	"github.com/astaxie/beego"
	"qax580go/models"
)

type RecommendController struct {
	beego.Controller
}

func (c *RecommendController) Get() {
	c.TplName = "recommend.html"
}

func (c *RecommendController) Post() {
	op := c.Input().Get("op")
	switch op {
	case "rec":
		index, err := models.GetQueryIndex("123456")
		if err != nil {
			beego.Error(err)
		}
		posts, err := models.QueryPagePost(index, 5)
		if err != nil {
			beego.Error(err)
		}
		if len(posts) != 0 {

		} else {
			beego.Debug("no info")
		}
		c.Data["Posts"] = posts
		beego.Debug(posts)
	}
	c.TplName = "recommend.html"
}
