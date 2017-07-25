package admin

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

type AdminWxOfficial struct {
	beego.Controller
}

func (c *AdminWxOfficial) WxOfficials() {
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
			err := models.DeleteWxOfficial(id)
			if err != nil {
				beego.Error(err)
			}
			url := "/admin/wxofficials"
			c.Redirect(url, 302)
			return
		case "state":
			id := c.Input().Get("id")
			if len(id) == 0 {
				break
			}
			id = c.Input().Get("id")
			err := models.UpdateWxOfficial(id, 1)
			if err != nil {
				beego.Error(err)
			}
			c.Redirect("/admin/wxofficials", 302)
			return
		case "state1":
			id := c.Input().Get("id")
			if len(id) == 0 {
				break
			}
			id = c.Input().Get("id")
			err := models.UpdateWxOfficial(id, 0)
			if err != nil {
				beego.Error(err)
			}
			c.Redirect("/admin/wxofficials", 302)
			return
		case "upinfo":
			id := c.Input().Get("id")
			if len(id) == 0 {
				break
			}
			url := fmt.Sprintf("/admin/upwxofficialinfo?id=%s", id)
			beego.Debug("up_rul", url)
			c.Redirect(url, 302)
			return
		case "upimg":
			id := c.Input().Get("id")
			if len(id) == 0 {
				break
			}
			url := fmt.Sprintf("/admin/upwxofficialimg?id=%s", id)
			beego.Debug("up_rul", url)
			c.Redirect(url, 302)
			return
		case "ishome":
			id := c.Input().Get("id")
			if len(id) == 0 {
				break
			}
			show := c.Input().Get("show")
			isshow := false
			if show == "true" {
				isshow = true
			}
			err := models.UpdateWxOfficialIsHome(id, isshow)
			if err != nil {
				beego.Error(err)
			}
			c.Redirect("/admin/wxofficials", 302)
			return
		}

		wxnums, err := models.GetAdminWxOfficials()
		if err != nil {
			beego.Error(err)
		}
		beego.Debug(wxnums)
		c.TplName = "admin/wxofficials.html"
		c.Data["Wxnums"] = wxnums
	}
}

func (c *AdminWxOfficial) Add() {
	if c.Ctx.Input.IsGet() {
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

		c.TplName = "admin/addwxofficial.html"
	}
	if c.Ctx.Input.IsPost() {
		image_name := ""

		title := c.Input().Get("title")
		info := c.Input().Get("info")
		number := c.Input().Get("number")
		evaluate := c.Input().Get("evaluate")
		if len(title) != 0 && len(info) != 0 && len(number) != 0 {
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
			c.Redirect("/admin/wxofficials", 302)
		}

		if len(title) != 0 && len(info) != 0 && len(number) != 0 {
			err := models.AddWxOfficial(title, info, number, evaluate, image_name)
			if err != nil {
				beego.Error(err)
			}
			beego.Info(info)
			c.Redirect("/admin/wxofficials", 302)
		}
	}
}

func (c *AdminWxOfficial) UpInfo() {
	if c.Ctx.Input.IsGet() {
		id := c.Input().Get("id")
		obj, err := models.GetOneWxOfficial(id)
		if err != nil {
			beego.Error(err)
		}
		c.Data["Wxnum"] = obj
		c.TplName = "admin/upwxofficialinfo.html"
	}
	if c.Ctx.Input.IsPost() {
		id := c.Input().Get("id")
		title := c.Input().Get("title")
		info := c.Input().Get("info")
		number := c.Input().Get("number")
		evaluate := c.Input().Get("evaluate")
		if len(id) != 0 && len(title) != 0 && len(info) != 0 && len(number) != 0 {
			err := models.UpdateWxOfficialnfo(id, title, info, number, evaluate)
			if err != nil {
				beego.Error(err)
			}
			c.Redirect("/admin/wxofficials", 302)
			return
		}
		c.TplName = "admi/upwxofficialinfo.html"
	}
}

func (c *AdminWxOfficial) UpImg() {
	if c.Ctx.Input.IsGet() {
		id := c.Input().Get("id")
		obj, err := models.GetOneWxOfficial(id)
		if err != nil {
			beego.Error(err)
		}
		c.Data["Wxnum"] = obj
		c.TplName = "admin/upwxofficialimg.html"
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
				err := models.UpdateWxOfficialImg(id, image_name)
				if err != nil {
					beego.Error(err)
				} else {
					c.Redirect("/admin/wxofficials", 302)
					return
				}
			}
		}

		c.TplName = "admin/upwxofficialimg.html"
	}
}
