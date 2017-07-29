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

type AdminUpCommodityImgController struct {
	beego.Controller
}

func (c *AdminUpCommodityImgController) Get() {
	id := c.Input().Get("id")
	if len(id) == 0 {
		c.Redirect("/admin/guanggaos", 302)
		return
	}
	// beego.Debug(id)
	obj, err := models.GetOneCommodity(id)
	if err != nil {
		beego.Error(err)
	}
	c.Data["Commodity"] = obj
	c.TplName = "adminupcommodityimg.html"
}
func (c *AdminUpCommodityImgController) Post() {

	id := c.Input().Get("id")
	originalimg := c.Input().Get("originalimg")
	image_name := originalimg
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

		beego.Debug("上传前图片", originalimg, "上传后图片", image_name)
		if len(image_name) != 0 {
			err := models.UpdateCommodityimg(id, image_name)
			if err != nil {
				beego.Error(err)
			} else {
				c.Redirect("/admin/mall", 302)
				return
			}
		}
	}
	c.TplName = "adminupcommodityimg.html"
}
