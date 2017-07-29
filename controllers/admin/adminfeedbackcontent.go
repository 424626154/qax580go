package admin

/*
后台意见反馈详情
*/
import (
	"qax580go/models"
	"qax580go/qutil"

	"github.com/astaxie/beego"
)

type AdminFeedbackContentController struct {
	beego.Controller
}

func (c *AdminFeedbackContentController) Get() {
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

	op := c.Input().Get("op")
	switch op {
	case "con":
		id := c.Input().Get("id")
		if len(id) == 0 {
			break
		}
		id = c.Input().Get("id")
		feedback, err := models.GetOneFeedback(id)
		if err != nil {
			beego.Error(err)
		}
		c.Data["Feedback"] = feedback
		beego.Debug("is adminfeedbackcontent " + feedback.Info)
		c.TplName = "adminfeedbackcontent.html"
		return
	}
	c.TplName = "adminfeedbackcontent.html"

}
