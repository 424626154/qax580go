package admin

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"path"
	"qax580go/models"
	"strings"
	"time"

	"github.com/astaxie/beego"
)

type AdminPostController struct {
	beego.Controller
}

func (c *AdminPostController) AddPost() {
	beego.Debug(c.Input())
	request_json := `{"errcode":1,"errmsg":"addpost error"}`
	image_name := ""
	title := c.Input().Get("title")
	info := c.Input().Get("info")
	city := c.Input().Get("city")
	bfrom := c.Input().Get("bfrom")
	// beego.Debug("info:", info)
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

	}

	if len(title) != 0 && len(info) != 0 {
		b_from := false
		fromshow := ""
		fromurl := ""
		if bfrom == "true" {
			b_from = true
			fromshow = c.Input().Get("fromshow")
			fromurl = c.Input().Get("fromurl")
		}
		beego.Debug("b_from :", bfrom)
		id, err := models.AddPostLabel(title, info, 2, image_name, city, b_from, fromshow, fromurl)
		if err != nil {
			beego.Error(err)
		}
		beego.Debug("inset ok id:", id)
		request_json = `{"errcode":0,"errmsg":""}`
	}
	c.Ctx.WriteString(request_json)
}
