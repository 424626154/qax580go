package controllers

/*
后台修改推荐公众号图片
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

type AdminUpWxnumImgController struct {
	beego.Controller
}

func (c *AdminUpWxnumImgController) Get() {
	id := c.Input().Get("id")
	obj, err := models.GetOneWxnum(id)
	if err != nil {
		beego.Error(err)
	}
	c.Data["Wxnum"] = obj
	c.TplName = "adminupwxnumimg.html"
}
func (c *AdminUpWxnumImgController) Post() {
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
			err := models.UpdateWxnumImg(id, image_name)
			if err != nil {
				beego.Error(err)
			} else {
				c.Redirect("/admin/wxlist", 302)
				return
			}
		}
	}

	c.TplName = "adminupwxnumimg.html"
}
