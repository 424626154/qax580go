package controllers

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

type WptController struct {
	beego.Controller
}

/**
微平台主页
*/
func (c *WptController) Home() {
	if c.Ctx.Input.IsGet() {
		beego.Debug("WptController Home Get")
	}
	if c.Ctx.Input.IsPost() {
		beego.Debug("WptController Home Post")
	}
	initWptUser(c)
	objs, err := models.GetAllWptTJ(1)
	if err != nil {
		beego.Error(err)
	}
	c.Data["Objs"] = objs
	objs1, err := models.GetAllWptTJ(0)
	if err != nil {
		beego.Error(err)
	}
	c.Data["Obj1s"] = objs1
	// shaer_url := "http://www.baoguangguang.cn/wpt/home"
	// wxShare := getShare(tl_appid, tl_secret, shaer_url)
	// c.Data["WxShare"] = wxShare
	c.TplName = "wpthome.html"
}

func (c *WptController) Search() {
	if c.Ctx.Input.IsGet() {
		beego.Debug("WptController Search Get")
	}
	if c.Ctx.Input.IsPost() {
		beego.Debug("WptController Search Post")
	}
	initWptUser(c)
	search := c.Input().Get("search")
	if len(search) != 0 {
		objs, err := models.GetAllWptLike(search)
		if err != nil {
			beego.Error(err)
		}
		beego.Debug("Search :", objs)
		c.Data["Objs"] = objs
	} else {
		objs, err := models.GetAllWpt()
		if err != nil {
			beego.Error(err)
		}
		c.Data["Objs"] = objs

	}

	c.TplName = "wptsearch.html"
}

/**
用户微平台列表
*/

func (c *WptController) AdminHome() {
	bool, username := chackAccount(c.Ctx)
	if bool {

	} else {
		c.Redirect("/admin", 302)
		return
	}
	c.Data["isUser"] = bool
	c.Data["User"] = username

	if c.Ctx.Input.IsGet() {
		beego.Debug("WptController AdminHome Get")
	}
	if c.Ctx.Input.IsPost() {
		beego.Debug("WptController AdminHome Post")
	}
	op := c.Input().Get("op")
	switch op {
	case "del":
		id := c.Input().Get("id")
		err := models.DelWpt(id)
		if err != nil {
			beego.Error(err)
		}
		break
	case "state":
		id := c.Input().Get("id")
		err := models.UpWptState(id, 1)
		if err != nil {
			beego.Error(err)
		}
		break
	case "state1":
		id := c.Input().Get("id")
		err := models.UpWptState(id, 0)
		if err != nil {
			beego.Error(err)
		}
		break
	case "tuijian":
		id := c.Input().Get("id")
		err := models.UpWptTuijian(id, 1)
		if err != nil {
			beego.Error(err)
		}
		break
	case "tuijian1":
		id := c.Input().Get("id")
		err := models.UpWptTuijian(id, 0)
		if err != nil {
			beego.Error(err)
		}
		break
	}

	objs, err := models.GetAllWpts()
	if err != nil {
		beego.Error(err)
	}
	c.Data["Objs"] = objs
	c.TplName = "wptadminhome.html"
}

/**
添加平台
*/

func (c *WptController) AdminAdd() {
	bool, username := chackAccount(c.Ctx)
	if bool {

	} else {
		c.Redirect("/admin", 302)
		return
	}
	c.Data["isUser"] = bool
	c.Data["User"] = username
	if c.Ctx.Input.IsGet() {
		beego.Debug("WptController AdminAdd Get")
	}
	if c.Ctx.Input.IsPost() {
		beego.Debug("WptController AdminAdd Post")
	}

	title := c.Input().Get("title")
	wid := c.Input().Get("wid")
	info := c.Input().Get("info")
	wrange := c.Input().Get("wrange")
	qrcode_name := ""
	if len(title) != 0 && len(wid) != 0 && len(info) != 0 {
		// 上传微信二维码
		_, fh, err := c.GetFile("image1")
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
			qrcode_name = hex.EncodeToString(h.Sum(nil))
			beego.Debug(qrcode_name) // 输出加密结果
			err = c.SaveToFile("image1", path.Join("imagehosting", qrcode_name))
			if err != nil {
				beego.Error(err)
				qrcode_name = ""
			}
		}

		err = models.AddWpt(title, info, wid, wrange, qrcode_name)
		if err != nil {
			beego.Error(err)
		}

		url := "/wpt/adminhome"
		c.Redirect(url, 302)
		return
	}

	c.TplName = "wptadminadd.html"
}

/**
修改图片
*/
func (c *WptController) AdminUpImg() {
	bool, username := chackAccount(c.Ctx)
	if bool {
		c.Data["isUser"] = bool
		c.Data["User"] = username
	} else {
		c.Redirect("/admin", 302)
		return
	}
	c.Data["isUser"] = bool
	c.Data["User"] = username

	if c.Ctx.Input.IsGet() {
		beego.Debug("AdminUpImg Get")
	}
	if c.Ctx.Input.IsPost() {
		beego.Debug("AdminUpImg Post")

		id := c.Input().Get("id")
		originalqrcode := c.Input().Get("originalqrcode")
		qrcode_name := originalqrcode
		if len(id) != 0 {
			// 二维码
			_, fh, err := c.GetFile("image1")
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
				qrcode_name = hex.EncodeToString(h.Sum(nil))
				beego.Debug("qrcode_name:", qrcode_name) // 输出加密结果
				err = c.SaveToFile("image1", path.Join("imagehosting", qrcode_name))
				if err != nil {
					beego.Error(err)
					qrcode_name = originalqrcode
				}
			}

			beego.Debug("上传前图片originalqrcode", originalqrcode, "上传后图片qrcode_name", qrcode_name)
			if qrcode_name == originalqrcode {
				beego.Debug("未修改图片")
				c.Redirect("/wpt/adminhome", 302)
				return
			}

			if len(qrcode_name) != 0 {
				err := models.UpWptImg(id, qrcode_name)
				if err != nil {
					beego.Error(err)
				} else {
					c.Redirect("/wpt/adminhome", 302)
					return
				}
			}
		}
	}
	id := c.Input().Get("id")
	obj, err := models.GetOneWpt(id)
	if err != nil {
		beego.Error(err)
	}
	c.Data["Obj"] = obj
	c.TplName = "wptadminupimg.html"
}

/**
修改内容
*/

func (c *WptController) AdminUpInfo() {
	bool, username := chackAccount(c.Ctx)
	if bool {
		c.Data["isUser"] = bool
		c.Data["User"] = username
	} else {
		c.Redirect("/admin", 302)
		return
	}
	c.Data["isUser"] = bool
	c.Data["User"] = username
	// beego.Debug("c.Input() :", c.Input())
	if c.Ctx.Input.IsGet() {
		beego.Debug("AdminUpInfo Get")
	}
	if c.Ctx.Input.IsPost() {
		beego.Debug("AdminUpInfo Post")
		id := c.Input().Get("id")
		title := c.Input().Get("title")
		info := c.Input().Get("info")
		wid := c.Input().Get("wid")
		wrange := c.Input().Get("wrange")
		if len(title) != 0 && len(info) != 0 && len(wid) != 0 && len(wrange) != 0 {
			err := models.UpWptInfo(id, title, info, wid, wrange)
			if err != nil {
				beego.Error(err)
			}
			url := "/wpt/adminhome"
			c.Redirect(url, 302)
			return
		}
	}

	id := c.Input().Get("id")
	obj, err := models.GetOneWpt(id)
	if err != nil {
		beego.Error(err)
	}
	c.Data["Obj"] = obj
	c.TplName = "wptadminupinfo.html"
}

func initWptUser(c *WptController) (*models.User, bool) {
	isLogin := false
	openid := c.Ctx.GetCookie(COOKIE_UID)
	muser := &models.User{}
	beego.Debug(openid)
	if len(openid) != 0 {
		user, err := models.GetOneUserUid(openid)
		if err != nil {
			beego.Error(err)
		} else {
			muser = user
			// beego.Debug(user)
			isLogin = true
			c.Data["User"] = user
		}
	}
	c.Data["isLogin"] = isLogin
	return muser, isLogin
}
