package admin

/*
后台主页
*/
import (
	"fmt"
	"qax580go/models"
	"qax580go/qutil"
	"strconv"

	"github.com/astaxie/beego"
)

type AdminHomeController struct {
	beego.Controller
}

func (c *AdminHomeController) Get() {
	bool, username := qutil.ChackAccount(c.Ctx)
	if bool {

	} else {
		c.Redirect("/admin", 302)
		return
	}
	c.Data["isUser"] = bool
	c.Data["User"] = username
	op := c.Input().Get("op")
	switch op {
	case "del":
		id := c.Input().Get("id")
		if len(id) == 0 {
			break
		}
		id = c.Input().Get("id")
		err := models.DeletePost(id)
		if err != nil {
			beego.Error(err)
		}
		beego.Debug("is admin del " + id)
		c.Redirect("/admin/home", 302)
		return
	case "examine":
		id := c.Input().Get("id")
		if len(id) == 0 {
			break
		}
		id = c.Input().Get("id")
		err := models.UpdatePost(id)
		if err != nil {
			beego.Error(err)
		}
		//添加审核金钱
		post, err := models.GetOnePost(id)
		if err != nil {
			beego.Error(err)
		} else {
			if post.Label == 1 {
				err = models.AddWxUserMoney(post.OpenId, qutil.MONEY_EXAMINE_SUM)
				if err != nil {
					beego.Error(err)
				} else {
					_, err = models.AddUserMoneyRecord(post.OpenId, qutil.MONEY_EXAMINE_SUM, qutil.MONEY_EXAMINE)
					if err != nil {
						beego.Error(err)
					}
				}

			}
		}
		if post.Label == 1 {
			msg := fmt.Sprintf("您发布的[%s]已通过审核", post.Title)
			beego.Debug("msg:", msg)
			sid := fmt.Sprintf("%d", post.Id)
			err = models.AddNotice(username, post.OpenId, true, msg, sid, qutil.NTYPE_1)
			if err != nil {
				beego.Error(err)
			}
		}
		beego.Debug("is admin examine " + id)
		c.Redirect("/admin/home", 302)
		return
	case "examine1":
		id := c.Input().Get("id")
		if len(id) == 0 {
			break
		}
		id = c.Input().Get("id")
		err := models.UpdatePostExamine(id, 0)
		if err != nil {
			beego.Error(err)
		}
		beego.Debug("is admin examine1" + id)
		c.Redirect("/admin/home", 302)
		return
	case "state1no":
		id := c.Input().Get("id")
		if len(id) == 0 {
			break
		}
		id = c.Input().Get("id")
		err := models.UpdatePostState1(id, 1)
		if err != nil {
			beego.Error(err)
		}
		post, err := models.GetOnePost(id)
		if err != nil {
			beego.Error(err)
		}
		if post.Label == 1 {
			msg := fmt.Sprintf("您发布的[%s]存在违规，请您重新发布", post.Title)
			sid := fmt.Sprintf("%d", post.Id)
			err = models.AddNotice(username, post.OpenId, true, msg, sid, qutil.NTYPE_2)
			if err != nil {
				beego.Error(err)
			}
		}
		beego.Debug("is admin state1no" + id)
		c.Redirect("/admin/home", 302)
		return
	case "state1ok":
		id := c.Input().Get("id")
		if len(id) == 0 {
			break
		}
		id = c.Input().Get("id")
		err := models.UpdatePostState1(id, 0)
		if err != nil {
			beego.Error(err)
		}
		beego.Debug("is admin state1ok" + id)
		c.Redirect("/admin/home", 302)
		return
	case "back":
		c.Ctx.SetCookie("username", "", -1, "/")
		c.Ctx.SetCookie("password", "", -1, "/")
		c.Redirect("/admin", 302)
		return
	}

	CurrentPage := int32(1)
	count, err := models.GetAdminPostCount() //后台用户数量
	NumberofPages := int32(30)
	temp := count / NumberofPages
	if (count % NumberofPages) != 0 {
		temp = temp + 1
	}
	CotalPages := temp
	pagetype := c.Input().Get("type")
	page := c.Input().Get("page")
	// beego.Debug("pagetype:", pagetype)

	if len(pagetype) != 0 && len(page) != 0 {
		switch pagetype {
		case "first": //首页
			CurrentPage = 1
		case "prev": //上一页
			pageint, error := strconv.Atoi(page)
			if error != nil {
				beego.Error(error)
			}
			CurrentPage = int32(pageint)
		case "next": //下一页
			pageint, error := strconv.Atoi(page)
			if error != nil {
				beego.Error(error)
			}
			CurrentPage = int32(pageint)
		case "last": //尾页
			CurrentPage = CotalPages
		case "page": //页码
			pageint, error := strconv.Atoi(page)
			if error != nil {
				beego.Error(error)
			}
			CurrentPage = int32(pageint)
		}
	}
	c.Data["CurrentPage"] = CurrentPage
	c.Data["CotalPages"] = CotalPages
	c.Data["NumberofPages"] = NumberofPages
	posts, err := models.QueryAdminPagePost(CurrentPage-1, NumberofPages)
	if err != nil {
		beego.Error(err)
	}
	c.Data["Posts"] = posts
	num, err := models.GetAllStateNum()
	if err != nil {
		beego.Error(err)
	}
	num1, err := models.GetAllState1Num()
	if err != nil {
		beego.Error(err)
	}
	c.Data["Num"] = num
	c.Data["Num1"] = num1
	c.Data["All"] = count
	c.TplName = "admin/adminhome.html"
}
