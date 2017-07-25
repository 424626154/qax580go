package controllers

/*
信息详情
*/
import (
	"qax580go/models"
	"qax580go/qutil"
	"strings"

	"github.com/astaxie/beego"
)

type ContentController struct {
	beego.Controller
}

func (c *ContentController) Get() {
	openid := getContentCookie(c)
	splat := c.Ctx.Input.Param(":splat")
	if len(splat) > 0 {
		result := strings.Index(splat, ".html")
		// beego.Debug("result:", result)
		if result > 0 {
			id := string([]byte(splat)[:result])
			beego.Debug("result:", result, "id:", id)
			// id = c.Input().Get("id")
			post, err := models.GetOnePost(id)
			if err != nil {
				beego.Error(err)
			}
			beego.Debug("id :", id)
			c.Data["Id"] = id
			c.Data["Post"] = post
			beego.Debug("is con " + post.Title)
			help_num, err := models.GatPostHelpNum(id)
			c.Data["HelpNum"] = help_num
			state, err := models.GatPaseHelpState(id, openid)
			c.Data["HelpState"] = state
			c.TplName = "content.html"
			return
		}
	}
	beego.Debug("get op splat:---------------", splat)
	op := c.Input().Get("op")
	beego.Debug("get op :---------------", op)
	switch op {
	case "con":
		id := c.Input().Get("id")
		if len(id) == 0 {
			break
		}
		// id = c.Input().Get("id")
		post, err := models.GetOnePost(id)
		if err != nil {
			beego.Error(err)
		}
		beego.Debug("id :", id)
		c.Data["Id"] = id
		c.Data["Post"] = post
		beego.Debug("is con " + post.Title)
		help_num, err := models.GatPostHelpNum(id)
		c.Data["HelpNum"] = help_num
		state, err := models.GatPaseHelpState(id, openid)
		c.Data["HelpState"] = state
		c.TplName = "content.html"
		return
	}
	c.TplName = "content.html"
}
func (c *ContentController) Post() {
	openid := getContentCookie(c)
	op := c.Input().Get("op")
	beego.Debug("post op :---------------", op)
	switch op {
	case "help":
		id_s := c.Input().Get("id")
		if len(openid) != 0 && len(id_s) != 0 {
			post, err := models.GetOnePost(id_s)
			if err != nil {
				beego.Error(err)
			}
			_, err = models.AddPosthelp(post.Id, openid, 1)
			if err != nil {
				beego.Error(err)
			} else {
				if post.Label == 1 {
					err = models.AddWxUserMoney(post.OpenId, 1)
					if err != nil {
						beego.Error(err)
					} else {
						_, err = models.AddUserMoneyRecord(post.OpenId, qutil.MONEY_BELIKE_SUM, qutil.MONEY_BELIKE)
					}
				}
				url := "/content?op=con&id=" + id_s
				c.Redirect(url, 302)
				return
			}
		}
		c.TplName = "content.html"
		break
	}
	c.TplName = "content.html"
}
func getContentCookie(c *ContentController) string {
	isUser := false
	openid := c.Ctx.GetCookie(qutil.COOKIE_WX_OPENID)
	beego.Debug("------------openid--------")
	beego.Debug(openid)
	if len(openid) != 0 {
		wxuser, err := models.GetOneWxUserInfo(openid)
		if err != nil {
			beego.Error(err)
		} else {
			isUser = true
			beego.Debug("--------------wxuser----------")
			beego.Debug(wxuser)
			c.Data["WxUser"] = wxuser
		}
	}
	c.Data["isUser"] = isUser
	return openid
}
