package controllers

/*
发布消息
*/
import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"path"
	"qax580go/models"
	"qax580go/qutil"
	"strings"
	"time"

	"github.com/astaxie/beego"
)

type UplodeController struct {
	beego.Controller
}

func (c *UplodeController) Get() {
	isLogin := isUplodeLogin(c)
	beego.Debug("isLogin:", isLogin)
	if isLogin == false {
		c.Redirect("/login?from=uplode", 302)
		return
	}
	c.Data["FromType"] = getUplodeFromType(c)
	c.TplName = "uplode.html"
}

func (c *UplodeController) Post() {
	image_name := ""
	title := c.Input().Get("title")
	info := c.Input().Get("info")
	city := c.Input().Get("city")
	beego.Debug("city", city)
	if len(title) != 0 && len(info) != 0 {
		// 获取附件
		_, fh, err := c.GetFile("image")
		if err != nil {
			beego.Error(err)
		}
		var attachment string
		if fh != nil {
			// 保存附件
			attachment = fh.Filename
			t := time.Now().Unix()
			str2 := fmt.Sprintf("%d", t)
			s := []string{attachment, str2}
			h := md5.New()
			h.Write([]byte(strings.Join(s, ""))) // 需要加密的字符串
			image_name = hex.EncodeToString(h.Sum(nil))
			beego.Info(image_name) // 输出加密结果
			err = c.SaveToFile("image", path.Join("imagehosting", image_name))
			if err != nil {
				beego.Error(err)
				image_name = ""
			}
		}
		if err != nil {
			beego.Error(err)
		}
	} else {
		c.Redirect("/uplode", 302)
	}
	if len(title) != 0 && len(info) != 0 {
		user := getUplodeUser(c)
		beego.Debug("----------AddPostLabelWx--------")
		err := models.AddPostLabelWx(title, info, 1, image_name, user.Uid, user.Name, user.Sex, user.Head, city)
		if err != nil {
			beego.Error(err)
		}
		beego.Debug("inset ok")
		c.Redirect("/mymessage", 302)
	}
}

func isUplodeLogin(c *UplodeController) bool {
	isLogin := false
	uid := c.Ctx.GetCookie(qutil.COOKIE_UID)
	if len(uid) != 0 {
		user, err := models.GetOneUserUid(uid)
		if err != nil {
			beego.Error(err)
		} else if len(user.Uid) != 0 {
			isLogin = true
			c.Data["User"] = user
		}
	}
	c.Data["isLogin"] = isLogin
	return isLogin
}

func getUplodeUser(c *UplodeController) *models.User {
	openid := c.Ctx.GetCookie(qutil.COOKIE_UID)
	muser := &models.User{}
	beego.Debug(openid)
	if len(openid) != 0 {
		user, err := models.GetOneUserUid(openid)
		if err != nil {
			beego.Error(err)
		} else {
			beego.Debug(user)
			muser = user
		}
	}
	return muser
}

/**
*来源类型
 */
func getUplodeFromType(c *UplodeController) string {
	from_type := c.Ctx.GetCookie(qutil.COOKIE_FROM_TYPE)
	if len(from_type) == 0 {
		from_type = qutil.COOKIE_FROM_ALL
	}
	return from_type
}
