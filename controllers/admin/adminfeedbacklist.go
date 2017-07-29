package admin

/*
后台意见反馈列表
*/
import (
	"qax580go/models"
	"qax580go/qutil"

	"github.com/astaxie/beego"
)

type AdminFeedbackListController struct {
	beego.Controller
}

func (c *AdminFeedbackListController) Get() {
	bool, username := qutil.ChackAccount(c.Ctx)
	if bool {

	} else {
		c.Redirect("/admin", 302)
		return
	}
	posts, err := models.GetAllPostsAdmin()
	if err != nil {
		beego.Error(err)
	}
	c.Data["Posts"] = posts
	c.Data["isUser"] = bool
	c.Data["User"] = username

	feedbacks, err := models.GetAllFeedbacks()
	if err != nil {
		beego.Error(err)
	}
	c.TplName = "adminfeedbacklist.html"
	c.Data["Feedbacks"] = feedbacks
	beego.Debug(feedbacks)
}
