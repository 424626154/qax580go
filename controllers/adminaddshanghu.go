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

type SHType struct {
	Name string
	Type int16
}
type AdminAddShanghuController struct {
	beego.Controller
}

func (c *AdminAddShanghuController) Get() {
	bool, username := chackAccount(c.Ctx)
	if bool {
		c.Data["isUser"] = bool
		c.Data["User"] = username
	} else {
		c.Redirect("/admin", 302)
		return
	}
	c.Data["Image"] = ""
	c.TplName = "adminaddshanghu.html"

}
func (c *AdminAddShanghuController) Post() {
	image_name := ""
	title := c.Input().Get("title")
	info := c.Input().Get("info")
	mtype := c.Input().Get("type")
	// city := c.Input().Get("city")
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

		mtype_i := int16(0)
		typemap := make(map[string]int16)
		typemap[SH_CANYIN] = CANYIN_TYPE
		typemap[SH_QICHE] = QICHE_TYPE
		typemap[SH_WEIXIU] = WEIXIU_TYPE
		typemap[SH_PEIXUN] = PEIXUN_TYPE
		if v, ok := typemap[mtype]; ok {
			mtype_i = v
		} else {

		}
		beego.Debug("mtype_i", mtype_i)
		err = models.AddShangHu(title, info, image_name, mtype_i)
		if err != nil {
			beego.Error(err)
		}
		c.Redirect("/admin/shanghus", 302)
		return
	}
	c.TplName = "adminaddshanghu.html"

}
