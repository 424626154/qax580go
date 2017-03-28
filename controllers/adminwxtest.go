package controllers

/*
后台推荐公众号列表
*/
import (
	"fmt"
	"github.com/astaxie/beego"
	"qax580go/models"
)

type AdminWxTestController struct {
	beego.Controller
}

func (c *AdminWxTestController) Get() {
	bool, username := chackAccount(c.Ctx)
	if bool {
		c.Data["isUser"] = bool
		c.Data["User"] = username
	} else {
		c.Redirect("/admin", 302)
		return
	}
	key0 := c.Input().Get("key0")
	c.Data["Key0"] = key0
	c.TplName = "adminwxtest.html"
}
func (c *AdminWxTestController) Post() {
	beego.Debug("AdminWxTestController Post")
	op := c.Input().Get("op")
	switch op {
	case "key0":
		key0 := c.Input().Get("key0")
		beego.Debug("key0 :", key0)
		//是否存在关键字
		key_count := int32(0)
		count, err := models.GetKeywordsCount(key0)
		if err != nil {
			beego.Error(err)
		} else {
			key_count = count
		}
		beego.Debug("count :", count)
		beego.Debug("key_count :", key_count)
		if key_count > 0 {
			obj, err := models.GetOneKeywords(key0)
			if err != nil {
				beego.Error(err)
			} else {
				objs, err := models.QueryFuzzyLimitKeyobj(obj.Id, 5)
				if err != nil {
					beego.Error(err)
				}
				if len(objs) > 0 {
					beego.Debug("Keyobjs :", objs)
				} else {

				}
			}
		} else {
			//信息查询
			posts, err := models.QueryFuzzyLimitPost(key0, 5)
			if err != nil {
				beego.Error(err)
			}
			// beego.Debug(requestBody.FromUserName)
			// beego.Debug(requestBody.ToUserName)
			if len(posts) > 0 {

			} else {
			}
		}
		url := fmt.Sprintf("/admin/wxtest?key0=%s", key0)
		c.Redirect(url, 302)
		return
	}
	c.TplName = "adminwxtest.html"
}
