package controllers

/*
添加关键字对象
*/
import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"github.com/astaxie/beego"
	"path"
	"qax580go/models"
	"strings"
	"time"
)

type AdminaAddKeyobjController struct {
	beego.Controller
}

func (c *AdminaAddKeyobjController) Get() {
	beego.Debug("AdminaAddKeyobjController Get")
	bool, username := chackAccount(c.Ctx)
	if bool {
		c.Data["isUser"] = bool
		c.Data["User"] = username
	} else {
		c.Redirect("/admin", 302)
		return
	}
	keyid := c.Input().Get("keyid")
	beego.Debug("keyid :", keyid)
	c.Data["KeyId"] = keyid
	c.Data["Image"] = ""
	c.TplName = "adminaddkeyobj.html"

}
func (c *AdminaAddKeyobjController) Post() {
	beego.Debug("AdminaAddKeyobjController Post")
	image_name := ""
	title := c.Input().Get("title")
	info := c.Input().Get("info")
	murl := c.Input().Get("url")
	keyid := c.Input().Get("keyid")
	if len(title) != 0 && len(info) != 0 {
		// 获取附件
		_, fh, err := c.GetFile("image")
		// beego.Debug("上传图片:", fh)
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
		beego.Debug("keyid", keyid)
		err = models.AddKeyobj(keyid, title, info, image_name, murl)
		if err != nil {
			beego.Error(err)
		}
		url := fmt.Sprintf("/admin/keyobj?keyid=%s", keyid)
		c.Redirect(url, 302)
		return
	}
	c.TplName = "adminaddkeyobj.html"

}
