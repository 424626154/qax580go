package admin

/*
添加外面入口
*/
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

type AdminAddWaimaiController struct {
	beego.Controller
}

func (c *AdminAddWaimaiController) Get() {
	bool, username := qutil.ChackAccount(c.Ctx)
	if bool {
		c.Data["isUser"] = bool
		c.Data["User"] = username
	} else {
		c.Redirect("/admin", 302)
		return
	}
	c.TplName = "admin/adminaddwaimai.html"
}
func (c *AdminAddWaimaiController) Post() {
	image_name := ""
	name := c.Input().Get("name")
	address := c.Input().Get("address")
	phone := c.Input().Get("phone")
	starth := c.Input().Get("starth")
	startm := c.Input().Get("startm")
	endh := c.Input().Get("endh")
	endm := c.Input().Get("endm")
	beego.Debug(phone, starth, startm, endh, endm)
	if len(name) != 0 && len(address) != 0 {
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
		c.Redirect("/admin/addwaimai", 302)
	}
	if len(name) != 0 && len(address) != 0 {
		_, err := models.AddCanting(name, address, image_name, int8(0), phone, starth, startm, endh, endm)
		if err != nil {
			beego.Error(err)
		}
		c.Redirect("/admin/waimailist", 302)
		return
	}
	c.TplName = "admin/aadminaddwaimai.html"
}
