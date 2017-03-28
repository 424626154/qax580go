package controllers

/*
添加广告
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

type AdminaAddGuanggaoController struct {
	beego.Controller
}

func (c *AdminaAddGuanggaoController) Get() {
	bool, username := chackAccount(c.Ctx)
	if bool {
		c.Data["isUser"] = bool
		c.Data["User"] = username
	} else {
		c.Redirect("/admin", 302)
		return
	}
	c.Data["Image"] = ""
	c.TplName = "adminaddguanggao.html"

}
func (c *AdminaAddGuanggaoController) Post() {
	image_name := ""
	imageitem0 := ""
	imageitem1 := ""
	imageitem2 := ""
	title := c.Input().Get("title")
	info := c.Input().Get("info")
	blink := c.Input().Get("blink")
	link := c.Input().Get("link")
	bimg := c.Input().Get("bimg")
	if len(title) != 0 && len(info) != 0 {
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
		//上传imageitem0
		_, fh, err = c.GetFile("imageitem0")
		beego.Debug("上传imageitem0:", fh)
		if err != nil {
			beego.Error(err)
		}
		if fh != nil {
			// 保存附件
			attachment = fh.Filename
			t := time.Now().Unix()
			str2 := fmt.Sprintf("%d%s", t, "imageitem0")
			s := []string{attachment, str2}
			h := md5.New()
			h.Write([]byte(strings.Join(s, ""))) // 需要加密的字符串
			imageitem0 = hex.EncodeToString(h.Sum(nil))
			beego.Info(imageitem0) // 输出加密结果
			err = c.SaveToFile("imageitem0", path.Join("imagehosting", imageitem0))
			if err != nil {
				beego.Error(err)
				imageitem0 = ""
			}
		}
		//上传imageitem1
		_, fh, err = c.GetFile("imageitem1")
		beego.Debug("上传imageitem1:", fh)
		if err != nil {
			beego.Error(err)
		}
		if fh != nil {
			// 保存附件
			attachment = fh.Filename
			t := time.Now().Unix()
			str2 := fmt.Sprintf("%d%s", t, imageitem1)
			s := []string{attachment, str2}
			h := md5.New()
			h.Write([]byte(strings.Join(s, ""))) // 需要加密的字符串
			imageitem1 = hex.EncodeToString(h.Sum(nil))
			beego.Info(imageitem1) // 输出加密结果
			err = c.SaveToFile("imageitem1", path.Join("imagehosting", imageitem1))
			if err != nil {
				beego.Error(err)
				imageitem1 = ""
			}
		}
		//上传imageitem2
		_, fh, err = c.GetFile("imageitem2")
		beego.Debug("上传imageitem2:", fh)
		if err != nil {
			beego.Error(err)
		}
		if fh != nil {
			// 保存附件
			attachment = fh.Filename
			t := time.Now().Unix()
			str2 := fmt.Sprintf("%d%s", t, "imageitem2")
			s := []string{attachment, str2}
			h := md5.New()
			h.Write([]byte(strings.Join(s, ""))) // 需要加密的字符串
			imageitem2 = hex.EncodeToString(h.Sum(nil))
			beego.Info(imageitem2) // 输出加密结果
			err = c.SaveToFile("imageitem2", path.Join("imagehosting", imageitem2))
			if err != nil {
				beego.Error(err)
				imageitem2 = ""
			}
		}

		b_link := false
		s_link := ""
		if blink == "true" {
			b_link = true
			s_link = link
		}
		b_img := false
		if bimg == "true" {
			b_img = true
		}
		beego.Debug("info", info)
		_, err = models.AddGuanggao(title, info, image_name, b_link, s_link, b_img, imageitem0, imageitem1, imageitem2)
		if err != nil {
			beego.Error(err)
		}
		c.Redirect("/admin/guanggaos", 302)
		return
	}
	c.TplName = "adminaddguanggao.html"

}
