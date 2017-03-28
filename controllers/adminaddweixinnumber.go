package controllers

/*
后台添加微信号
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

type AdminAddWeixinNumberController struct {
	beego.Controller
}

func (c *AdminAddWeixinNumberController) Get() {
	bool, username := chackAccount(c.Ctx)
	if bool {
		c.Data["isUser"] = bool
		c.Data["User"] = username
	} else {
		c.Redirect("/admin", 302)
		return
	}
	c.TplName = "adminaddweixinnumber.html"

}
func (c *AdminAddWeixinNumberController) Post() {
	beego.Debug("AdminAddWeixinNumberController Post")
	image_name := ""
	name := c.Input().Get("name")
	info := c.Input().Get("info")
	number := c.Input().Get("number")
	evaluate := c.Input().Get("evaluate")
	if len(name) != 0 && len(info) != 0 && len(number) != 0 {
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
		c.Redirect("/admin/weixinnumberlist", 302)
	}

	if len(name) != 0 && len(info) != 0 && len(number) != 0 {
		err := models.AddWeixinNumber(name, info, number, evaluate, image_name)
		if err != nil {
			beego.Error(err)
		}
		beego.Info(info)
		c.Redirect("/admin/weixinnumberlist", 302)
	}

	c.TplName = "adminaddweixinnumber.html"
}
