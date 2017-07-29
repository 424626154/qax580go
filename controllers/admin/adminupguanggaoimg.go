package admin

/*
后台修改广告图片
*/
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

type AdminUpGuanggaoImgController struct {
	beego.Controller
}

func (c *AdminUpGuanggaoImgController) Get() {
	id := c.Input().Get("id")
	if len(id) == 0 {
		c.Redirect("/admin/guanggaos", 302)
		return
	}
	// beego.Debug(id)
	guangao, err := models.GetOneCommodity(id)
	if err != nil {
		beego.Error(err)
	}
	c.Data["Guanggao"] = guangao
	c.TplName = "adminupguanggaoimg.html"
}
func (c *AdminUpGuanggaoImgController) Post() {

	id := c.Input().Get("id")
	originalimg := c.Input().Get("originalimg")
	originalitem0 := c.Input().Get("originalitem0")
	originalitem1 := c.Input().Get("originalitem1")
	originalitem2 := c.Input().Get("originalitem2")
	bimg := c.Input().Get("bimg")
	image_name := originalimg
	item0_name := originalitem0
	item1_name := originalitem1
	item2_name := originalitem2
	if len(id) != 0 {
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
				image_name = originalimg
			}
		}
		// 上传图片0
		_, fh, err = c.GetFile("imageitem0")
		beego.Debug("上传图片imageitem0:", fh)
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
			item0_name = hex.EncodeToString(h.Sum(nil))
			beego.Info(item0_name) // 输出加密结果
			err = c.SaveToFile("imageitem0", path.Join("imagehosting", item0_name))
			if err != nil {
				beego.Error(err)
				item0_name = originalitem0
			}
		}
		// 上传图片1
		_, fh, err = c.GetFile("imageitem1")
		beego.Debug("上传图片imageitem1:", fh)
		if err != nil {
			beego.Error(err)
		}
		if fh != nil {
			// 保存附件
			attachment = fh.Filename
			t := time.Now().Unix()
			str2 := fmt.Sprintf("%d%s", t, "imageitem1")
			s := []string{attachment, str2}
			h := md5.New()
			h.Write([]byte(strings.Join(s, ""))) // 需要加密的字符串
			item1_name = hex.EncodeToString(h.Sum(nil))
			beego.Info(item1_name) // 输出加密结果
			err = c.SaveToFile("imageitem1", path.Join("imagehosting", item1_name))
			if err != nil {
				beego.Error(err)
				item1_name = originalitem1
			}
		}
		// 上传图片2
		_, fh, err = c.GetFile("imageitem2")
		beego.Debug("上传图片imageitem2:", fh)
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
			item2_name = hex.EncodeToString(h.Sum(nil))
			beego.Info(item2_name) // 输出加密结果
			err = c.SaveToFile("imageitem2", path.Join("imagehosting", item2_name))
			if err != nil {
				beego.Error(err)
				item2_name = originalitem2
			}
		}
		b_img := false
		if bimg == "true" {
			b_img = true
		}
		beego.Debug("上传前图片", originalitem0, "上传后图片", item0_name)
		if len(image_name) != 0 || len(item0_name) != 0 || len(item1_name) != 0 || len(item2_name) != 0 {
			err := models.UpdateGuanggaoImg(id, image_name, b_img, item0_name, item1_name, item2_name)
			if err != nil {
				beego.Error(err)
			} else {
				c.Redirect("/admin/guanggaos", 302)
				return
			}
		}
	}
	c.TplName = "adminupguanggaoimg.html"
}
