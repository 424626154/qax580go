package admin //添加外面入口
/*
添加菜单
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

type AdminAddCaidanController struct {
	beego.Controller
}

func (c *AdminAddCaidanController) Get() {
	bool, username := qutil.ChackAccount(c.Ctx)
	if bool {
		c.Data["isUser"] = bool
		c.Data["User"] = username
	} else {
		c.Redirect("/admin", 302)
		return
	}
	fid := c.Input().Get("fid")
	c.Data["Fid"] = fid
	c.TplName = "admin/adminaddcaidan.html"
}
func (c *AdminAddCaidanController) Post() {
	image_name := ""
	name := c.Input().Get("name")
	info := c.Input().Get("info")
	fid := c.Input().Get("fid")
	mtype := c.Input().Get("mtype")
	price := c.Input().Get("price")
	if len(name) != 0 && len(info) != 0 && len(fid) != 0 && len(mtype) != 0 && len(price) != 0 {
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
	if len(name) != 0 && len(info) != 0 && len(fid) != 0 && len(mtype) != 0 {
		_, err := models.AddCaidan(fid, name, info, image_name, int8(0), mtype, price)
		if err != nil {
			beego.Error(err)
		}
		url := fmt.Sprintf("/admin/caidans?id=%s", fid)
		c.Redirect(url, 302)
		return
	}
	c.TplName = "admin/adminaddcaidan.html"
}
