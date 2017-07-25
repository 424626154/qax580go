package admin

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"path"
	"strings"
	"time"

	"qax580go/qutil"

	"qax580go/models"

	"github.com/astaxie/beego"
)

type AdminWeChat struct {
	beego.Controller
}

func (c *AdminWeChat) WeChats() {
	if c.Ctx.Input.IsGet() {
		bool, username := qutil.ChackAccount(c.Ctx)
		if bool {
			c.Data["isUser"] = bool
			c.Data["User"] = username
		} else {
			c.Redirect("/admin", 302)
			return
		}
		op := c.Input().Get("op")
		switch op {
		case "del":
			id := c.Input().Get("id")
			if len(id) == 0 {
				break
			}
			id = c.Input().Get("id")
			err := models.DeleteWeChat(id)
			if err != nil {
				beego.Error(err)
			}
			url := "/admin/wechats"
			c.Redirect(url, 302)
			return
		case "state":
			id := c.Input().Get("id")
			if len(id) == 0 {
				break
			}
			id = c.Input().Get("id")
			err := models.UpdateWeChatState(id, 1)
			if err != nil {
				beego.Error(err)
			}
			c.Redirect("/admin/wechats", 302)
			return
		case "state1":
			id := c.Input().Get("id")
			if len(id) == 0 {
				break
			}
			id = c.Input().Get("id")
			err := models.UpdateWeChatState(id, 0)
			if err != nil {
				beego.Error(err)
			}
			c.Redirect("/admin/wechats", 302)
			return
		case "upinfo":
			id := c.Input().Get("id")
			if len(id) == 0 {
				break
			}
			url := fmt.Sprintf("/admin/upwechatinfo?id=%s", id)
			beego.Debug("up_rul", url)
			c.Redirect(url, 302)
			return
		case "upimg":
			id := c.Input().Get("id")
			if len(id) == 0 {
				break
			}
			url := fmt.Sprintf("/admin/upwechatimg?id=%s", id)
			beego.Debug("up_rul", url)
			c.Redirect(url, 302)
			return
		case "ishome":
			id := c.Input().Get("id")
			show := c.Input().Get("show")
			ishome := false
			if show == "true" {
				ishome = true
			}
			err := models.UpdateWeChatIsHome(id, ishome)
			if err != nil {
				beego.Error(err)
			}
			c.Redirect("/admin/wechats", 302)
			return
		}

		objs, err := models.GetAdminWeChats()
		if err != nil {
			beego.Error(err)
		}
		c.Data["WeixinNumbers"] = objs
		// beego.Debug(objs)
		c.TplName = "admin/wechats.html"
	}
}

func (c *AdminWeChat) Add() {
	if c.Ctx.Input.IsGet() {
		bool, username := qutil.ChackAccount(c.Ctx)
		if bool {
			c.Data["isUser"] = bool
			c.Data["User"] = username
		} else {
			c.Redirect("/admin", 302)
			return
		}
		c.TplName = "admin/addwechat.html"
	}
	if c.Ctx.Input.IsPost() {
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
			c.Redirect("/admin/wechats", 302)
		}

		if len(name) != 0 && len(info) != 0 && len(number) != 0 {
			err := models.AddWeChat(name, info, number, evaluate, image_name)
			if err != nil {
				beego.Error(err)
			}
			beego.Info(info)
			c.Redirect("/admin/wechats", 302)
		}

		c.TplName = "admin/addwechat.html"
	}
}

func (c *AdminWeChat) UpInfo() {
	if c.Ctx.Input.IsGet() {
		id := c.Input().Get("id")
		obj, err := models.GetOneWeChat(id)
		if err != nil {
			beego.Error(err)
		}
		c.Data["WeixinNumber"] = obj
		c.TplName = "admin/upwechatinfo.html"
	}
	if c.Ctx.Input.IsPost() {
		id := c.Input().Get("id")
		name := c.Input().Get("name")
		info := c.Input().Get("info")
		number := c.Input().Get("number")
		evaluate := c.Input().Get("evaluate")
		if len(id) != 0 && len(name) != 0 && len(info) != 0 && len(number) != 0 {
			err := models.UpdateWeChatInfo(id, name, info, number, evaluate)
			if err != nil {
				beego.Error(err)
			}
			c.Redirect("/admin/wechats", 302)
			return
		}
		c.TplName = "admin/upwechatinfo.html"
	}
}

func (c *AdminWeChat) UpImg() {
	if c.Ctx.Input.IsGet() {
		id := c.Input().Get("id")
		obj, err := models.GetOneWeChat(id)
		if err != nil {
			beego.Error(err)
		}
		c.Data["WeixinNumber"] = obj
		c.TplName = "admin/upwechatimg.html"
	}

	if c.Ctx.Input.IsPost() {
		id := c.Input().Get("id")
		image_name := ""
		if len(id) != 0 {
			// 获取附件
			_, fh, err := c.GetFile("image")
			beego.Debug("上传图片:", fh)
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
			if image_name != "" {
				err := models.UpdateWeChatImg(id, image_name)
				if err != nil {
					beego.Error(err)
				} else {
					c.Redirect("/admin/wechats", 302)
					return
				}
			}
		}

		c.TplName = "admin/upwechatimg.html"
	}
}
