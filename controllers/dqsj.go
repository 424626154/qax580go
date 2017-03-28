package controllers

/**
大签世界
*/
import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/config"
	"github.com/astaxie/beego/context"
	"io/ioutil"
	"math/rand"
	"net/http"
	"path"
	"qax580go/models"
	"strconv"
	"strings"
	"time"
)

type DqsjController struct {
	beego.Controller
}

//后台
func (c *DqsjController) Admin() {
	if c.Ctx.Input.IsGet() {
		beego.Debug("Admin Get")
	}
	if c.Ctx.Input.IsPost() {
		beego.Debug("Admin Post")
	}
	bool, username := chackDqsjAccount(c.Ctx)
	if bool {

	} else {
		c.Redirect("/dqsj/adminlogin", 302)
		return
	}
	c.Data["User"] = username
	c.Data["isUser"] = bool
	beego.Debug("username:", username)
	op := c.Input().Get("op")
	switch op {
	case "back":
		c.Ctx.SetCookie(DQSJ_USERNAME, "", -1, "/")
		c.Ctx.SetCookie(DQSJ_PASSWORD, "", -1, "/")
		c.Redirect("/dqsj/admin", 302)
		return
	}

	c.TplName = "dqsjadmin.html"
}

//后台登录
func (c *DqsjController) AdminLogin() {
	if c.Ctx.Input.IsGet() {
		beego.Debug("AdminLogin Get")
	}
	if c.Ctx.Input.IsPost() {
		beego.Debug("AdminLogin Post")
		username := c.Input().Get("user")
		password := c.Input().Get("password")
		autologin := c.Input().Get("autologin") == "on"
		beego.Debug("AdminLogin Post user:", username, "password:", password)
		if len(username) != 0 && len(password) != 0 {
			admin, err := models.GetOneDqsjAdmin(username)
			if err != nil {
				c.Redirect("/dqsj/adminlogin", 302)
				return
			}
			if admin != nil && strings.EqualFold(username, admin.Username) && strings.EqualFold(password, admin.Password) {
				maxAge := 0
				if autologin {
					maxAge = 1<<31 - 1
				}
				c.Ctx.SetCookie(DQSJ_USERNAME, username, maxAge, "/")
				c.Ctx.SetCookie(DQSJ_PASSWORD, password, maxAge, "/")
				beego.Debug("login ok------")
				c.Redirect("/dqsj/admin", 302)
				return
			} else {
				c.Redirect("/dqsj/adminlogin", 302)
				return
			}
		} else {
			c.Redirect("/dqsj/adminlogin", 302)
			return
		}
	}
	c.TplName = "dqsjadminlogin.html"
}
func (c *DqsjController) AdminCai() {
	if c.Ctx.Input.IsGet() {
		beego.Debug("AdminCai Get")
	}
	if c.Ctx.Input.IsPost() {
		beego.Debug("AdminCai Post")
	}
	bool, _ := chackDqsjAccount(c.Ctx)
	if bool {

	} else {
		c.Redirect("/dqsj/adminlogin", 302)
		return
	}

	op := c.Input().Get("op")
	beego.Debug("op :", op)
	switch op {
	case "del":
		id := c.Input().Get("id")
		err := models.DeleteAllCaiItem(id)
		if err != nil {
			beego.Error(err)
		}
		err = models.DeleteCaiGroup(id)
		if err != nil {
			beego.Error(err)
		}
		c.Redirect("/dqsj/admincai", 302)
		return
	case "state0":
		id := c.Input().Get("id")
		err := models.UpdateCaiGroup(id, 1)
		if err != nil {
			beego.Error(err)
		}
		c.Redirect("/dqsj/admincai", 302)
		return
	case "state1":
		id := c.Input().Get("id")
		err := models.UpdateCaiGroup(id, 0)
		if err != nil {
			beego.Error(err)
		}
		c.Redirect("/dqsj/admincai", 302)
		return
	case "itemdel":
		id := c.Input().Get("id")
		err := models.DeleteCaiGrItem(id)
		if err != nil {
			beego.Error(err)
		}
		c.Redirect("/dqsj/admincai", 302)
		return
	case "itemstate0":
		id := c.Input().Get("id")
		err := models.UpdateCaiItem(id, 1)
		if err != nil {
			beego.Error(err)
		}
		c.Redirect("/dqsj/admincai", 302)
		return
	case "itemstate1":
		id := c.Input().Get("id")
		err := models.UpdateCaiItem(id, 0)
		if err != nil {
			beego.Error(err)
		}
		c.Redirect("/dqsj/admincai", 302)
		return
	case "tipsdel":
		id := c.Input().Get("id")
		err := models.DeleteCaiTips(id)
		if err != nil {
			beego.Error(err)
		}
		c.Redirect("/dqsj/admincai", 302)
		return
	case "tipsstate0":
		id := c.Input().Get("id")
		err := models.UpdateCaiTips(id, 1)
		if err != nil {
			beego.Error(err)
		}
		c.Redirect("/dqsj/admincai", 302)
		return
	case "tipsstate1":
		id := c.Input().Get("id")
		err := models.UpdateCaiTips(id, 0)
		if err != nil {
			beego.Error(err)
		}
		c.Redirect("/dqsj/admincai", 302)
		return
	}

	objs, err := models.GetAllCaiGroup()
	if err != nil {
		beego.Error(err)
	}
	var showCaiGroup []models.DqsjShowCaiGroup
	for i := 0; i < len(objs); i++ {
		objitem, err := models.GetAllCaiItem(objs[i].Id)
		if err != nil {
			beego.Error(err)
		} else {
			obgshow := models.DqsjShowCaiGroup{Id: objs[i].Id, Name: objs[i].Name,
				OrderId: objs[i].OrderId, State: objs[i].State, Time: objs[i].Time, CaiItems: objitem}
			showCaiGroup = append(showCaiGroup, obgshow)
		}

	}

	c.Data["ShowCaiGroup"] = showCaiGroup

	tips, err := models.GetAllCaiTips()
	if err != nil {
		beego.Error(err)
	}
	c.Data["CaiTips"] = tips
	c.TplName = "dqsjadmincai.html"
}
func (c *DqsjController) AdminAddCaiGroup() {
	bool, _ := chackDqsjAccount(c.Ctx)
	if bool {

	} else {
		c.Redirect("/dqsj/adminlogin", 302)
		return
	}
	if c.Ctx.Input.IsGet() {
		beego.Debug("AdminAddCaiGroup Get")
	}
	if c.Ctx.Input.IsPost() {
		beego.Debug("AdminAddCaiGroup Post")
		name := c.Input().Get("name")
		orderid := c.Input().Get("orderid")
		if len(name) != 0 && len(orderid) != 0 {
			err := models.AddCaiGroup(name, orderid)
			if err != nil {
				beego.Error(err)
			}
			c.Redirect("/dqsj/admincai", 302)
		} else {
			c.Redirect("/dqsj/adminaddcaigroup", 302)
		}

	}
	c.TplName = "dqsjadminaddcaigroup.html"
}
func (c *DqsjController) AdminAddCaiItem() {
	bool, _ := chackDqsjAccount(c.Ctx)
	if bool {

	} else {
		c.Redirect("/dqsj/adminlogin", 302)
		return
	}
	if c.Ctx.Input.IsGet() {
		beego.Debug("AdminAddCaiItem Get")
		groupid := c.Input().Get("groupid")
		c.Data["GroupId"] = groupid
	}
	if c.Ctx.Input.IsPost() {
		beego.Debug("AdminAddCaiItem Post")
		image_name := ""
		name := c.Input().Get("name")
		price := c.Input().Get("price")
		pricedesc := c.Input().Get("pricedesc")
		groupid := c.Input().Get("groupid")
		if len(name) != 0 && len(price) != 0 {
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
				// beego.Info(image_name) // 输出加密结果
				err = c.SaveToFile("image", path.Join("imagehosting", image_name))
				if err != nil {
					beego.Error(err)
					image_name = ""
				}
			}
			err = models.AddCaiItem(name, image_name, groupid, price, pricedesc)
			if err != nil {
				beego.Error(err)
			}
			c.Redirect("/dqsj/admincai", 302)
			return
		}
	}
	c.TplName = "dqsjadminaddcaiitem.html"
}
func (c *DqsjController) AdminCaiUpCon() {
	bool, _ := chackDqsjAccount(c.Ctx)
	if bool {

	} else {
		c.Redirect("/dqsj/adminlogin", 302)
		return
	}
	if c.Ctx.Input.IsGet() {
		beego.Debug("AdminCaiUpCon Get")
		id := c.Input().Get("id")
		beego.Debug("AdminCaiUpCon Get id:", id)
		if len(id) == 0 {
			c.Redirect("/dqsj/admincai", 302)
		}
		obj, err := models.GetOneCaiItem(id)
		if err != nil {
			beego.Error(err)
		}
		c.Data["CaiItem"] = obj
	}
	if c.Ctx.Input.IsPost() {
		beego.Debug("AdminCaiUpCon Post")
		id := c.Input().Get("id")
		name := c.Input().Get("name")
		price := c.Input().Get("price")
		pricedesc := c.Input().Get("pricedesc")
		if len(id) > 0 && len(name) > 0 && len(price) > 0 && len(pricedesc) > 0 {
			err := models.UpdateCaiItemCon(id, name, price, pricedesc)
			if err != nil {
				beego.Error(err)
			}
			c.Redirect("/dqsj/admincai", 302)
			return
		}

	}
	c.TplName = "dqsjadmincaiupcon.html"
}
func (c *DqsjController) AdminPan() {
	if c.Ctx.Input.IsGet() {
		beego.Debug("AdminPan Get")
	}
	if c.Ctx.Input.IsPost() {
		beego.Debug("AdminPan Post")
	}
	bool, _ := chackDqsjAccount(c.Ctx)
	if bool {

	} else {
		c.Redirect("/dqsj/adminlogin", 302)
		return
	}

	op := c.Input().Get("op")
	beego.Debug("op :", op)
	switch op {
	case "del":
		id := c.Input().Get("id")
		err := models.DeletePanItem(id)
		if err != nil {
			beego.Error(err)
		}
		c.Redirect("/dqsj/adminpan", 302)
		return
	case "state0":
		id := c.Input().Get("id")
		err := models.UpdatePanItem(id, 1)
		if err != nil {
			beego.Error(err)
		}
		c.Redirect("/dqsj/adminpan", 302)
		return
	case "state1":
		id := c.Input().Get("id")
		err := models.UpdatePanItem(id, 0)
		if err != nil {
			beego.Error(err)
		}
		c.Redirect("/dqsj/adminpan", 302)
		return
	}
	panitem, err := models.GetAllPanItem()
	if err != nil {
		beego.Error(err)
	}
	//计算总概率
	allProbability := int64(0)
	for i := 0; i < len(panitem); i++ {
		if panitem[i].State == 1 {
			allProbability += panitem[i].Probability
		}
	}
	for i := 0; i < len(panitem); i++ {
		panitem[i].AllProbability = allProbability
	}
	c.Data["PanItem"] = panitem

	config, err := models.GetConfig()
	bpan := false
	if config != nil {
		bpan = config.Bpan
	}
	c.Data["Bpan"] = bpan
	c.TplName = "dqsjadminpan.html"
}
func (c *DqsjController) AdminAddPan() {
	bool, _ := chackDqsjAccount(c.Ctx)
	if bool {

	} else {
		c.Redirect("/dqsj/adminlogin", 302)
		return
	}
	if c.Ctx.Input.IsGet() {
		beego.Debug("AdminAddPan Get")
	}
	if c.Ctx.Input.IsPost() {
		beego.Debug("AdminAddPan Post")
		name := c.Input().Get("name")
		info := c.Input().Get("info")
		probability := c.Input().Get("probability")
		if len(name) != 0 && len(info) != 0 {
			err := models.AddPanItem(name, info, probability)
			if err != nil {
				beego.Error(err)
			}
			c.Redirect("/dqsj/adminpan", 302)
		} else {
			c.Redirect("/dqsj/adminaddpan", 302)
		}

	}
	c.TplName = "dqsjadminaddpan.html"
}
func (c *DqsjController) AdminUpPan() {
	bool, _ := chackDqsjAccount(c.Ctx)
	if bool {

	} else {
		c.Redirect("/dqsj/adminlogin", 302)
		return
	}
	if c.Ctx.Input.IsGet() {
		beego.Debug("AdminUpPan Get")
	}
	if c.Ctx.Input.IsPost() {
		beego.Debug("AdminUpPan Post")
		pid := c.Input().Get("pid")
		name := c.Input().Get("name")
		info := c.Input().Get("info")
		probability := c.Input().Get("probability")
		if len(name) != 0 && len(info) != 0 {
			err := models.UpPanItem(name, info, probability, pid)
			if err != nil {
				beego.Error(err)
			}
			c.Redirect("/dqsj/adminpan", 302)
		} else {
			c.Redirect("/dqsj/adminuppan", 302)
		}

	}
	pid := c.Input().Get("pid")
	obj, err := models.GetOnePanItem(pid)
	if err != nil {
		beego.Error(err)
	}
	c.Data["DqsjPanItem"] = obj
	c.TplName = "dqsjadminuppan.html"
}
func (c *DqsjController) AdminAddCaiTips() {
	bool, _ := chackDqsjAccount(c.Ctx)
	if bool {

	} else {
		c.Redirect("/dqsj/adminlogin", 302)
		return
	}
	if c.Ctx.Input.IsGet() {
		beego.Debug("AdminAddCaiTips Get")
	}
	if c.Ctx.Input.IsPost() {
		beego.Debug("AdminAddCaiTips Post")
		info := c.Input().Get("info")
		if len(info) != 0 {
			err := models.AddCaiTips(info)
			if err != nil {
				beego.Error(err)
			}
			c.Redirect("/dqsj/admincai", 302)
		} else {
			c.Redirect("/dqsj/adminaddcaitips", 302)
		}

	}
	c.TplName = "dqsjadminaddcaitips.html"
}

func (c *DqsjController) AdminGg() {
	bool, _ := chackDqsjAccount(c.Ctx)
	if bool {

	} else {
		c.Redirect("/dqsj/adminlogin", 302)
		return
	}
	if c.Ctx.Input.IsGet() {
		beego.Debug("AdminGuangGao Get")
	}
	if c.Ctx.Input.IsPost() {
		beego.Debug("AdminGuangGao Post")
	}
	op := c.Input().Get("op")
	switch op {
	case "del":
		id := c.Input().Get("id")
		if len(id) == 0 {
			break
		}
		id = c.Input().Get("id")
		err := models.DeleteDqsjGuanggao(id)
		if err != nil {
			beego.Error(err)
		}
		// beego.Debug("is admin del " + id)
		c.Redirect("/dqsj/admingg", 302)
		return
	case "state":
		id := c.Input().Get("id")
		if len(id) == 0 {
			break
		}
		id = c.Input().Get("id")
		err := models.UpdateDqsjGuanggao(id, 1)
		if err != nil {
			beego.Error(err)
		}
		// beego.Debug("is admin state " + id)
		c.Redirect("/dqsj/admingg", 302)
		return
	case "state1":
		id := c.Input().Get("id")
		if len(id) == 0 {
			break
		}
		id = c.Input().Get("id")
		err := models.UpdateDqsjGuanggao(id, 0)
		if err != nil {
			beego.Error(err)
		}
		// beego.Debug("is admin state1" + id)
		c.Redirect("/dqsj/admingg", 302)
		return
	}
	guanggaos, err := models.GetAllDqsjGuanggaos()
	if err != nil {
		beego.Error(err)
	}
	c.Data["Guanggaos"] = guanggaos
	c.TplName = "dqsjadminguanggao.html"
}

func (c *DqsjController) AdminAddGg() {
	bool, _ := chackDqsjAccount(c.Ctx)
	if bool {

	} else {
		c.Redirect("/dqsj/adminlogin", 302)
		return
	}
	if c.Ctx.Input.IsGet() {
		beego.Debug("AdminAddGg Get")
	}
	if c.Ctx.Input.IsPost() {
		beego.Debug("AdminAddGg Post")
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
			_, err = models.AddDqsjGuanggao(title, info, image_name, b_link, s_link, b_img, imageitem0, imageitem1, imageitem2)
			if err != nil {
				beego.Error(err)
			}
			c.Redirect("/dqsj/admingg", 302)
			return
		}
	}
	c.TplName = "dqsjadminaddgg.html"
}
func (c *DqsjController) AdminUpGgCon() {
	bool, _ := chackDqsjAccount(c.Ctx)
	if bool {

	} else {
		c.Redirect("/dqsj/adminlogin", 302)
		return
	}
	if c.Ctx.Input.IsGet() {
		beego.Debug("AdminUpGgCon Get")
		id := c.Input().Get("id")
		if len(id) != 0 {
			gg, err := models.GetOneDqsjGuanggao(id)
			if err != nil {
				beego.Error(err)
			}
			c.Data["Guanggao"] = gg
			beego.Debug("gg :", gg)
		}

	}
	if c.Ctx.Input.IsPost() {
		beego.Debug("AdminUpGgCon Post")
		title := c.Input().Get("title")
		info := c.Input().Get("info")
		blink := c.Input().Get("blink")
		link := c.Input().Get("link")
		id := c.Input().Get("id")
		if len(id) != 0 && len(title) != 0 && len(info) != 0 {
			b_link := false
			s_link := ""
			if blink == "true" {
				b_link = true
				s_link = link
			}

			err := models.UpdateDqsjGuanggaoInfo(id, title, info, b_link, s_link)
			if err != nil {
				beego.Error(err)
			}
			c.Redirect("/dqsj/admingg", 302)
		}
	}
	c.TplName = "dqsjadminupggcon.html"
}

func (c *DqsjController) AdminUpGgImg() {
	bool, _ := chackDqsjAccount(c.Ctx)
	if bool {

	} else {
		c.Redirect("/dqsj/adminlogin", 302)
		return
	}
	if c.Ctx.Input.IsGet() {
		beego.Debug("AdminUpGgImg Get")
		id := c.Input().Get("id")
		if len(id) != 0 {
			gg, err := models.GetOneDqsjGuanggao(id)
			if err != nil {
				beego.Error(err)
			}
			c.Data["Guanggao"] = gg
			beego.Debug("gg :", gg)
		}

	}
	if c.Ctx.Input.IsPost() {
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
				err := models.UpdateDqsjGuanggaoImg(id, image_name, b_img, item0_name, item1_name, item2_name)
				if err != nil {
					beego.Error(err)
				} else {
					c.Redirect("/dqsj/admingg", 302)
					return
				}
			}
		}
	}
	c.TplName = "dqsjadminupggimg.html"
}

func (c *DqsjController) AdminHuoDong() {
	bool, _ := chackDqsjAccount(c.Ctx)
	if bool {

	} else {
		c.Redirect("/dqsj/adminlogin", 302)
		return
	}
	if c.Ctx.Input.IsGet() {
		beego.Debug("AdminHuoDong Get")
	}
	if c.Ctx.Input.IsPost() {
		beego.Debug("AdminHuoDong Post")
	}
	op := c.Input().Get("op")
	switch op {
	case "uphuodong":
		huodong := c.Input().Get("huodong")
		beego.Debug("huodong:", huodong)
		if len(huodong) == 0 {
			break
		}
		err := models.ModifyDqsjHomeHD(huodong)
		if err != nil {
			beego.Error(err)
		}
		// beego.Debug("is admin del " + id)
		c.Redirect("/dqsj/adminhuodong", 302)
		return
	case "del":
		id := c.Input().Get("id")
		if len(id) == 0 {
			break
		}
		id = c.Input().Get("id")
		err := models.DeleteDqsjHD(id)
		if err != nil {
			beego.Error(err)
		}
		// beego.Debug("is admin del " + id)
		c.Redirect("/dqsj/adminhuodong", 302)
		return
	case "state":
		id := c.Input().Get("id")
		if len(id) == 0 {
			break
		}
		id = c.Input().Get("id")
		err := models.UpdateDqsjHD(id, 1)
		if err != nil {
			beego.Error(err)
		}
		// beego.Debug("is admin state " + id)
		c.Redirect("/dqsj/adminhuodong", 302)
		return
	case "state1":
		id := c.Input().Get("id")
		if len(id) == 0 {
			break
		}
		id = c.Input().Get("id")
		err := models.UpdateDqsjHD(id, 0)
		if err != nil {
			beego.Error(err)
		}
		// beego.Debug("is admin state1" + id)
		c.Redirect("/dqsj/adminhuodong", 302)
		return
	case "additem":
		content := c.Input().Get("content")
		if len(content) == 0 {
			break
		}
		err := models.AddDqsjHD(content)
		if err != nil {
			beego.Error(err)
		}
		// beego.Debug("is admin state1" + id)
		c.Redirect("/dqsj/adminhuodong", 302)
		return
	}

	obj, err := models.GetOneDqsjHome()
	if err != nil {
		beego.Debug(err)
	}
	c.Data["HuoDong"] = ""
	if obj != nil {
		c.Data["HuoDong"] = obj.HuoDong
	}
	obj1, err := models.GetAllDqsjHD()
	if err != nil {
		beego.Debug(err)
	}
	c.Data["DqsjHuoDong"] = obj1
	c.TplName = "dqsjadminhuodong.html"
}
func (c *DqsjController) AdminGua() {
	bool, _ := chackDqsjAccount(c.Ctx)
	if bool {

	} else {
		c.Redirect("/dqsj/adminlogin", 302)
		return
	}
	if c.Ctx.Input.IsGet() {
		beego.Debug("AdminGua Get")
	}
	if c.Ctx.Input.IsPost() {
		beego.Debug("AdminGua Post")
	}
	op := c.Input().Get("op")
	beego.Debug("op :", op)
	switch op {
	case "del":
		id := c.Input().Get("id")
		err := models.DeleteGuaItem(id)
		if err != nil {
			beego.Error(err)
		}
		c.Redirect("/dqsj/admingua", 302)
		return
	case "state0":
		id := c.Input().Get("id")
		err := models.UpdateGuaItem(id, 1)
		if err != nil {
			beego.Error(err)
		}
		c.Redirect("/dqsj/admingua", 302)
		return
	case "state1":
		id := c.Input().Get("id")
		err := models.UpdateGuaItem(id, 0)
		if err != nil {
			beego.Error(err)
		}
		c.Redirect("/dqsj/admingua", 302)
		return
	}
	panitem, err := models.GetAllGuaItem()
	if err != nil {
		beego.Error(err)
	}
	//计算总概率
	allProbability := int64(0)
	for i := 0; i < len(panitem); i++ {
		if panitem[i].State == 1 {
			allProbability += panitem[i].Probability
		}
	}
	for i := 0; i < len(panitem); i++ {
		panitem[i].AllProbability = allProbability
	}
	c.Data["GuaItem"] = panitem

	c.TplName = "dqsjadmingua.html"
}
func (c *DqsjController) AdminAddGua() {
	bool, _ := chackDqsjAccount(c.Ctx)
	if bool {

	} else {
		c.Redirect("/dqsj/adminlogin", 302)
		return
	}
	if c.Ctx.Input.IsGet() {
		beego.Debug("AdminAddGua Get")
	}
	if c.Ctx.Input.IsPost() {
		beego.Debug("AdminAddGua Post")
		name := c.Input().Get("name")
		info := c.Input().Get("info")
		probability := c.Input().Get("probability")
		if len(name) != 0 && len(info) != 0 {
			err := models.AddGuaItem(name, info, probability)
			if err != nil {
				beego.Error(err)
			}
			c.Redirect("/dqsj/admingua", 302)
		} else {
			c.Redirect("/dqsj/adminaddgua", 302)
		}

	}
	c.TplName = "dqsjadminaddgua.html"
}
func (c *DqsjController) AdminUpGua() {
	bool, _ := chackDqsjAccount(c.Ctx)
	if bool {

	} else {
		c.Redirect("/dqsj/adminlogin", 302)
		return
	}
	if c.Ctx.Input.IsGet() {
		beego.Debug("AdminUpGua Get")
	}
	if c.Ctx.Input.IsPost() {
		beego.Debug("AdminUpGua Post")
		pid := c.Input().Get("pid")
		name := c.Input().Get("name")
		info := c.Input().Get("info")
		probability := c.Input().Get("probability")
		if len(name) != 0 && len(info) != 0 {
			err := models.UpGuaItem(name, info, probability, pid)
			if err != nil {
				beego.Error(err)
			}
			c.Redirect("/dqsj/admingua", 302)
		} else {
			c.Redirect("/dqsj/adminupgua", 302)
		}

	}
	pid := c.Input().Get("pid")
	obj, err := models.GetOneGuaItem(pid)
	if err != nil {
		beego.Error(err)
	}
	c.Data["GuaItem"] = obj
	c.TplName = "dqsjadminupgua.html"
}

func (c *DqsjController) AdminShare() {
	bool, _ := chackDqsjAccount(c.Ctx)
	if bool {

	} else {
		c.Redirect("/dqsj/adminlogin", 302)
		return
	}
	if c.Ctx.Input.IsGet() {
		beego.Debug("AdminShare Get")
	}
	if c.Ctx.Input.IsPost() {
		beego.Debug("AdminShare Post")
		name := c.Input().Get("name")
		if len(name) != 0 {
			_, err := models.UpConfigShare(name)
			if err != nil {
				beego.Debug(err)
			}
		}
	}
	config, err := models.GetConfig()
	if err != nil {
		beego.Debug(err)
	}
	c.Data["Title"] = config.ShareTitle
	c.TplName = "dqsjadminshare.html"
}

/*
*大签世界会员
 */
func (c *DqsjController) AdminMember() {
	bool, _ := chackDqsjAccount(c.Ctx)
	if bool {

	} else {
		c.Redirect("/dqsj/adminlogin", 302)
		return
	}
	op := c.Input().Get("op")
	beego.Debug("op :", op)
	switch op {
	case "del":
		id := c.Input().Get("id")
		err := models.DeleteMember(id)
		if err != nil {
			beego.Error(err)
		}
		c.Redirect("/dqsj/adminmember", 302)
		return
	}
	objs, err := models.GetAllMember()
	if err != nil {
		beego.Error(err)
	}
	like := c.Input().Get("like")
	if len(like) > 0 {
		objs, err = models.GetLikeMember(like)
		if err != nil {
			beego.Error(err)
		}
	}
	c.Data["Objs"] = objs
	c.TplName = "dqsjadminmember.html"
}

func (c *DqsjController) AdminAddMember() {
	bool, _ := chackDqsjAccount(c.Ctx)
	if bool {

	} else {
		c.Redirect("/dqsj/adminlogin", 302)
		return
	}
	if c.Ctx.Input.IsGet() {
		beego.Debug("AdminAddMember Get")
	}
	if c.Ctx.Input.IsPost() {
		beego.Debug("AdminAddMember Post")
		account := c.Input().Get("account")
		name := c.Input().Get("name")
		phone := c.Input().Get("phone")
		beernum := c.Input().Get("beernum")
		if len(account) != 0 {
			err := models.AddMember(account, name, phone, beernum)
			if err != nil {
				beego.Error(err)
			}
			c.Redirect("/dqsj/adminmember", 302)
		} else {
			c.Redirect("/dqsj/adminaddmember", 302)
		}

	}
	c.TplName = "dqsjadminaddmember.html"
}

func (c *DqsjController) AdminUpMember() {
	bool, _ := chackDqsjAccount(c.Ctx)
	if bool {

	} else {
		c.Redirect("/dqsj/adminlogin", 302)
		return
	}
	if c.Ctx.Input.IsGet() {
		beego.Debug("AdminUpMember Get")
	}
	if c.Ctx.Input.IsPost() {
		beego.Debug("AdminUpMember Post")
		pid := c.Input().Get("pid")
		account := c.Input().Get("account")
		name := c.Input().Get("name")
		phone := c.Input().Get("phone")
		beernum := c.Input().Get("beernum")
		if len(account) != 0 && len(pid) != 0 {
			err := models.UpMember(account, name, phone, beernum, pid)
			if err != nil {
				beego.Error(err)
			}
			c.Redirect("/dqsj/adminmember", 302)
		} else {
			c.Redirect("/dqsj/adminupmember", 302)
		}

	}
	pid := c.Input().Get("pid")
	obj, err := models.GetOneMember(pid)
	if err != nil {
		beego.Error(err)
	}
	c.Data["Obj"] = obj
	beego.Debug("pid", pid, obj)
	c.TplName = "dqsjadminupmember.html"
}
func (c *DqsjController) AdminMemberSet() {
	bool, _ := chackDqsjAccount(c.Ctx)
	if bool {

	} else {
		c.Redirect("/dqsj/adminlogin", 302)
		return
	}
	if c.Ctx.Input.IsGet() {
		beego.Debug("AdminMember Get")
	}
	if c.Ctx.Input.IsPost() {
		beego.Debug("AdminMember Post")
	}
	obj, err := models.GetMemberSet()
	if err != nil {
		beego.Error(err)
	}
	beego.Debug("obj:", obj)
	pass := ""
	if obj != nil {
		pass = obj.DelPass
	}
	c.Data["DelPass"] = pass
	c.TplName = "dqsjadminmemberset.html"
}

//http post
func (c *DqsjController) Post() {
	bool, _ := chackDqsjAccount(c.Ctx)
	if bool {

	} else {
		c.Redirect("/dqsj/adminlogin", 302)
		return
	}
	if c.Ctx.Input.IsGet() {
		beego.Debug("Post Get")
	}
	if c.Ctx.Input.IsPost() {
		beego.Debug("Post Post")
	}
	op := c.Input().Get("op")
	request_json := `{"errcode":1,"errmsg":"request_json error"}`
	switch op {
	case "wx":
		err := models.UpWxAttributeTime(int64(0), int64(0))
		if err != nil {
			beego.Debug(err)
		} else {
			request_json = `{"errcode":0,"errmsg":""}`
		}
		break
	case "pan":
		bpan := c.Input().Get("bpan")
		beego.Debug("bpan", bpan)
		if len(bpan) > 0 {
			vbpan := false
			if bpan == "true" {
				vbpan = true
			}
			_, err := models.UpConfigPan(vbpan)
			if err != nil {
				beego.Error(err)
			} else {
				request_json = fmt.Sprintf(`{"errcode":0,"errmsg":"","data":"%s"}`, bpan)
			}
		}
		break
	case "beer_ded":
		num := c.Input().Get("num")
		id := c.Input().Get("id")
		if len(num) > 0 && len(id) > 0 {
			beernumi, err := strconv.ParseInt(num, 10, 64)
			if err != nil {
				beego.Error(err)
			} else {
				obj, err := models.GetOneMember(id)
				if err != nil {
					beego.Error(err)
				} else {
					if obj != nil {
						beernum := obj.BeerNum
						if obj.BeerNum-beernumi > 0 {
							beernum = obj.BeerNum - beernumi
						} else {
							beernum = 0
						}
						err = models.UpMemberBeer(id, beernum)
						if err != nil {
							beego.Error(err)
						} else {
							request_json = fmt.Sprintf(`{"errcode":0,"errmsg":"","data":"%s"}`, beernum)
						}
					}
				}
			}
		}
		break
	case "up_set_delpass": //设置解锁密码
		pass := c.Input().Get("pass")
		if len(pass) > 0 {
			_, err := models.UpMemberSetDelPass(pass)
			if err != nil {
				beego.Debug(err)
			} else {
				request_json = fmt.Sprintf(`{"errcode":0,"errmsg":"","data":"%s"}`, "修改密码成功")
			}
		}
		break
	}
	c.Ctx.WriteString(request_json)
}

//主页
func (c *DqsjController) Home() {
	if c.Ctx.Input.IsGet() {
		beego.Debug("Home Get")
	}
	if c.Ctx.Input.IsPost() {
		beego.Debug("Home Post")
	}
	//微信分享
	token := getDqsjToken()
	if len(token) > 0 {
		beego.Debug("http_dqsj_token :", token)
	}
	appId := ""
	isdebug := "flase"
	iniconf, err := config.NewConfig("json", "conf/myconfig.json")
	if err != nil {
		beego.Debug(err)
	} else {
		appId = iniconf.String("qax580::appid")
		isdebug = iniconf.String("qax580::isdebug")
	}
	url := "http://www.baoguangguang.cn/dqsj/home"
	if isdebug == "true" {
		url = "http://localhost:8080/dqsj/home"
	}

	timestamp := time.Now().Unix()
	noncestr := getNonceStr(16, KC_RAND_KIND_ALL)
	ticket := getDqsjTicket(token)
	c.Data["AppId"] = appId
	c.Data["TimesTamp"] = timestamp
	c.Data["NonceStr"] = noncestr
	c.Data["Ticket"] = signatureWxJs(ticket, noncestr, timestamp, url)
	wxShareCon := models.WxShareCon{}
	wxShareCon.Title = getShareTitle()
	wxShareCon.Link = url
	wxShareCon.ImgUrl = "http://182.92.167.29:8080/static/img/dqsjicon.jpg"
	c.Data["WxShareCon"] = wxShareCon
	// beego.Debug(wxShareCon)

	//广告栏
	c.Data["ImgUrlPath"] = getImageUrl()
	guanggaos, err := models.GetAllDqsjGuanggaosState1()
	if err != nil {
		beego.Error(err)
	}
	c.Data["Guanggaos"] = guanggaos

	obj, err := models.GetOneDqsjHome()
	if err != nil {
		beego.Debug(err)
	}
	c.Data["HuoDong"] = ""
	if obj != nil {
		c.Data["HuoDong"] = obj.HuoDong
		beego.Debug("HuoDong:", obj.HuoDong)
	}
	obj1, err := models.GetAllDqsjHDState1()
	if err != nil {
		beego.Debug(err)
	}
	if obj1 != nil {
		for i := 0; i < len(obj1); i++ {
			obj1[i].ShowId = int64(i + 1)
		}
	}
	c.Data["DqsjHuoDong"] = obj1

	config, err := models.GetConfig()
	bpan := false
	if config != nil {
		bpan = config.Bpan
	}
	c.Data["Bpan"] = bpan

	c.TplName = "dqsjhome.html"
}

//菜单
func (c *DqsjController) Cai() {
	if c.Ctx.Input.IsGet() {
		beego.Debug("Cai Get")
	}
	if c.Ctx.Input.IsPost() {
		beego.Debug("Cai Post")
	}
	//微信分享
	token := getDqsjToken()
	if len(token) > 0 {
		beego.Debug("http_dqsj_token :", token)
	}
	appId := ""
	isdebug := "flase"
	iniconf, err := config.NewConfig("json", "conf/myconfig.json")
	if err != nil {
		beego.Debug(err)
	} else {
		appId = iniconf.String("qax580::appid")
		isdebug = iniconf.String("qax580::isdebug")
	}
	url := "http://www.baoguangguang.cn/dqsj/cai"
	if isdebug == "true" {
		url = "http://localhost:8080/dqsj/cai"
	}
	timestamp := time.Now().Unix()
	noncestr := getNonceStr(16, KC_RAND_KIND_ALL)
	ticket := getDqsjTicket(token)
	c.Data["AppId"] = appId
	c.Data["TimesTamp"] = timestamp
	c.Data["NonceStr"] = noncestr
	c.Data["Ticket"] = signatureWxJs(ticket, noncestr, timestamp, url)
	wxShareCon := models.WxShareCon{}
	wxShareCon.Title = getShareTitle()
	wxShareCon.Link = url
	wxShareCon.ImgUrl = "http://182.92.167.29:8080/static/img/dqsjicon.jpg"
	c.Data["WxShareCon"] = wxShareCon
	// beego.Debug(wxShareCon)

	objs, err := models.GetAllCaiGroupState1()
	if err != nil {
		beego.Error(err)
	}
	var showCaiGroup []models.DqsjShowCaiGroup
	for i := 0; i < len(objs); i++ {
		objitem, err := models.GetAllCaiItemState1(objs[i].Id)
		if err != nil {
			beego.Error(err)
		} else {
			obgshow := models.DqsjShowCaiGroup{Id: objs[i].Id, Name: objs[i].Name,
				OrderId: objs[i].OrderId, State: objs[i].State, Time: objs[i].Time, CaiItems: objitem}
			showCaiGroup = append(showCaiGroup, obgshow)
		}

	}

	c.Data["ShowCaiGroup"] = showCaiGroup
	tips, err := models.GetAllCaiTipsState1()
	if err != nil {
		beego.Error(err)
	}
	c.Data["CaiTips"] = tips

	config, err := models.GetConfig()
	bpan := false
	if config != nil {
		bpan = config.Bpan
	}
	c.Data["Bpan"] = bpan

	c.TplName = "dqsjcai.html"
}

//幸运盘
func (c *DqsjController) Pan() {
	if c.Ctx.Input.IsGet() {
		beego.Debug("Pan Get")
	}
	if c.Ctx.Input.IsPost() {
		beego.Debug("Pan Post")
	}
	//微信分享
	token := getDqsjToken()
	if len(token) > 0 {
		beego.Debug("http_dqsj_token :", token)
	}
	appId := ""
	isdebug := "flase"
	iniconf, err := config.NewConfig("json", "conf/myconfig.json")
	if err != nil {
		beego.Debug(err)
	} else {
		appId = iniconf.String("qax580::appid")
		isdebug = iniconf.String("qax580::isdebug")
	}
	url := "http://www.baoguangguang.cn/dqsj/pan"
	if isdebug == "true" {
		url = "http://localhost:8080/dqsj/pan"
	}
	timestamp := time.Now().Unix()
	noncestr := getNonceStr(16, KC_RAND_KIND_ALL)
	ticket := getDqsjTicket(token)
	c.Data["AppId"] = appId
	c.Data["TimesTamp"] = timestamp
	c.Data["NonceStr"] = noncestr
	c.Data["Ticket"] = signatureWxJs(ticket, noncestr, timestamp, url)
	wxShareCon := models.WxShareCon{}
	wxShareCon.Title = getShareTitle()
	wxShareCon.Link = url
	wxShareCon.ImgUrl = "http://182.92.167.29:8080/static/img/dqsjicon.jpg"
	c.Data["WxShareCon"] = wxShareCon
	// beego.Debug(wxShareCon)
	panitem, err := models.GetAllPanItemState1()
	if err != nil {
		beego.Error(err)
	}
	c.Data["PanItem"] = panitem
	beego.Debug("panitem :", panitem)
	c.TplName = "dqsjpan.html"
}

func (c *DqsjController) Gua() {
	if c.Ctx.Input.IsGet() {
		beego.Debug("Gua Get")
	}
	if c.Ctx.Input.IsPost() {
		beego.Debug("Gua Post")
	}
	//微信分享
	token := getDqsjToken()
	if len(token) > 0 {
		beego.Debug("http_dqsj_token :", token)
	}
	appId := ""
	isdebug := "flase"
	iniconf, err := config.NewConfig("json", "conf/myconfig.json")
	if err != nil {
		beego.Debug(err)
	} else {
		appId = iniconf.String("qax580::appid")
		isdebug = iniconf.String("qax580::isdebug")
	}
	url := "http://www.baoguangguang.cn/dqsj/gua"
	if isdebug == "true" {
		url = "http://localhost:8080/dqsj/gua"
	}
	timestamp := time.Now().Unix()
	noncestr := getNonceStr(16, KC_RAND_KIND_ALL)
	ticket := getDqsjTicket(token)
	c.Data["AppId"] = appId
	c.Data["TimesTamp"] = timestamp
	c.Data["NonceStr"] = noncestr
	c.Data["Ticket"] = signatureWxJs(ticket, noncestr, timestamp, url)
	wxShareCon := models.WxShareCon{}
	wxShareCon.Title = getShareTitle()
	wxShareCon.Link = url
	wxShareCon.ImgUrl = "http://182.92.167.29:8080/static/img/dqsjicon.jpg"
	c.Data["WxShareCon"] = wxShareCon

	guaitem, err := models.GetAllGuaItemState1()
	if err != nil {
		beego.Error(err)
	}
	var guas []models.DqsjGuaItem
	for i := 0; i < len(guaitem); i++ {
		for j := 0; j < int(guaitem[i].Probability); j++ {
			guas = append(guas, guaitem[i])
		}
	}
	rand.Seed(time.Now().UnixNano())
	ri := rand.Intn(len(guas))
	rguaitem := guas[ri]
	beego.Debug("rguaitem :", rguaitem)
	beego.Debug("guas len", len(guas), "ri :", ri)
	c.Data["GuaItem"] = rguaitem
	c.TplName = "dqsjgua1.html"
}
func (c *DqsjController) GuangGao() {
	if c.Ctx.Input.IsGet() {
		beego.Debug("GuangGao Get")
	}
	if c.Ctx.Input.IsPost() {
		beego.Debug("GuangGao Post")
	}
	//微信分享
	token := getDqsjToken()
	if len(token) > 0 {
		beego.Debug("http_dqsj_token :", token)
	}
	appId := ""
	isdebug := "flase"
	iniconf, err := config.NewConfig("json", "conf/myconfig.json")
	if err != nil {
		beego.Debug(err)
	} else {
		appId = iniconf.String("qax580::appid")
		isdebug = iniconf.String("qax580::isdebug")
	}
	url := "http://www.baoguangguang.cn/dqsj/guanggao"
	if isdebug == "true" {
		url = "http://localhost:8080/dqsj/guanggao"
	}
	id := c.Input().Get("id")
	url = fmt.Sprintf("%s?op=con&id=%s", url, id)
	beego.Debug("wx url :", url)
	timestamp := time.Now().Unix()
	noncestr := getNonceStr(16, KC_RAND_KIND_ALL)
	ticket := getDqsjTicket(token)
	c.Data["AppId"] = appId
	c.Data["TimesTamp"] = timestamp
	c.Data["NonceStr"] = noncestr
	c.Data["Ticket"] = signatureWxJs(ticket, noncestr, timestamp, url)
	wxShareCon := models.WxShareCon{}
	wxShareCon.Title = getShareTitle()
	wxShareCon.Link = url
	wxShareCon.ImgUrl = "http://182.92.167.29:8080/static/img/dqsjicon.jpg"
	c.Data["WxShareCon"] = wxShareCon
	// beego.Debug(wxShareCon)

	op := c.Input().Get("op")
	switch op {
	case "con":
		id := c.Input().Get("id")
		if len(id) == 0 {
			break
		}
		guangao, err := models.GetOneDqsjGuanggao(id)
		if err != nil {
			beego.Error(err)
		}
		c.Data["Guanggao"] = guangao
		beego.Debug("guangao :", guangao)
		c.TplName = "dqsjguanggao.html"
		return
	}

	c.TplName = "dqsjguanggao.html"
}

func (c *DqsjController) Client() {
	clientJson := ClientJson{}
	clientJson.ErrCode = 1
	clientJson.ErrMsg = "未知错误"
	if c.Ctx.Input.IsGet() {
		beego.Debug("Client Get")
	}
	// if c.Ctx.Input.IsPost() {
	beego.Debug("Client Post")
	op := c.Input().Get("op")
	beego.Debug("Client op:", op)
	switch op {
	case "login":
		username := c.Input().Get("username")
		password := c.Input().Get("password")
		beego.Debug("username :", username, "password :", password)
		if len(username) != 0 && len(password) != 0 {
			admin, err := models.GetOneDqsjAdmin(username)
			if err != nil {
				clientJson.ErrMsg = "查询用户名失败"
			} else {
				if admin != nil && strings.EqualFold(username, admin.Username) && strings.EqualFold(password, admin.Password) {
					clientJson.ErrCode = 0
					clientJson.Data = "登录成功"
				} else {
					clientJson.ErrMsg = "登录失败"
				}
			}
		} else {
			clientJson.ErrMsg = "用户名或密码为空"
		}
		break
	case "member":
		like := c.Input().Get("like")
		if len(like) > 0 {
			objs, err := models.GetLikeUserMember(like)
			if err != nil {
				beego.Error(err)
			} else {
				objs_json, err := json.Marshal(objs)
				if err != nil {
					beego.Error(err)
				} else {
					clientJson.ErrCode = 0
					clientJson.Data = string(objs_json)
				}
			}
		} else {
			objs, err := models.GetAllUserMember()
			if err != nil {
				beego.Error(err)
			} else {
				objs_json, err := json.Marshal(objs)
				if err != nil {
					beego.Error(err)
				} else {
					clientJson.ErrCode = 0
					clientJson.Data = string(objs_json)
				}
			}
		}
		break
	case "getone":
		mid := c.Input().Get("mid")
		obj, err := models.GetOneMember(mid)
		if err != nil {
			beego.Error(err)
			clientJson.ErrMsg = "参数错误"
		} else {
			obj_json, err := json.Marshal(obj)
			if err != nil {
				beego.Error(err)
			} else {
				clientJson.ErrCode = 0
				clientJson.Data = string(obj_json)
			}
		}
		break
	case "up":
		mid := c.Input().Get("mid")
		account := c.Input().Get("account")
		name := c.Input().Get("name")
		phone := c.Input().Get("phone")
		beernum := c.Input().Get("beernum")
		if len(account) != 0 && len(mid) != 0 {
			err := models.UpMember(account, name, phone, beernum, mid)
			if err != nil {
				beego.Error(err)
				clientJson.ErrMsg = "参数错误"
			} else {
				obj, err := models.GetOneMember(mid)
				if err != nil {
					beego.Error(err)
					clientJson.ErrMsg = "参数错误"
				} else {
					obj_json, err := json.Marshal(obj)
					if err != nil {
						beego.Error(err)
					} else {
						clientJson.ErrCode = 0
						clientJson.Data = string(obj_json)
					}
				}
			}

		} else {

		}
		break
	case "deduction":
		mid := c.Input().Get("id")
		num := c.Input().Get("num")
		if len(num) > 0 && len(mid) > 0 {
			beernumi, err := strconv.ParseInt(num, 10, 64)
			if err != nil {
				beego.Error(err)
			} else {
				obj, err := models.GetOneMember(mid)
				if err != nil {
					beego.Error(err)
				} else {
					if obj != nil {
						beernum := obj.BeerNum
						if obj.BeerNum-beernumi > 0 {
							beernum = obj.BeerNum - beernumi
						} else {
							beernum = 0
						}
						err = models.UpMemberBeer(mid, beernum)
						if err != nil {
							beego.Error(err)
						} else {
							clientJson.ErrCode = 0
							clientJson.Data = fmt.Sprintf(`{"Id":%s,"BeerNum":%d}"`, mid, beernum)
						}
					}
				}
			}
		}
		break
	case "add":
		account := c.Input().Get("account")
		name := c.Input().Get("name")
		phone := c.Input().Get("phone")
		beernum := c.Input().Get("beernum")
		if len(account) != 0 {
			err := models.AddMember(account, name, phone, beernum)
			if err != nil {
				beego.Error(err)
			} else {
				clientJson.ErrCode = 0
			}

		} else {
			clientJson.ErrMsg = "参数错误"
		}
		break
	case "memberdel":
		id := c.Input().Get("id")
		if len(id) != 0 {
			err := models.DeleteUserMember(id)
			if err != nil {
				beego.Error(err)
				clientJson.ErrMsg = "删除错误"
			} else {
				clientJson.ErrCode = 0
				clientJson.Data = id
			}
		} else {
			clientJson.ErrMsg = "参数错误"
		}
		break
	case "lock":
		pass := c.Input().Get("pass")
		beego.Debug("pass :", pass)
		if len(pass) > 0 {
			obj, err := models.GetMemberSet()
			if err != nil {
				beego.Debug(err)
			} else {
				if obj != nil {
					beego.Debug("obj.DelPass  :", obj.DelPass)
					if strings.EqualFold(obj.DelPass, pass) {
						clientJson.ErrCode = 0
						clientJson.Data = fmt.Sprintf(`{"errcode":0,"errmsg":"","data":"%s"}`, "解锁成功")
					} else {

					}
				}
			}
		}
		break
	default:
		clientJson.ErrMsg = "参数错误"
		break
	}
	// }
	//struct 到json str
	body, err := json.Marshal(clientJson)
	if err != nil {
		beego.Error(err)
	}
	response_json := string(body)
	beego.Debug("response_json :", response_json)
	c.Ctx.WriteString(response_json)
}

func getDqsjToken() string {
	//https://api.weixin.qq.com/cgi-bin/token?&appid=APPID&secret=APPSECRET
	wxAttribute, err := models.GetWxAttribute()
	if err != nil {
		beego.Debug(err)
	}
	if wxAttribute != nil {
		if len(wxAttribute.AccessToken) != 0 {
			current_time := time.Now().Unix()
			beego.Debug("current_time:", current_time, "wxAttribute.AccessTokenTime :", wxAttribute.AccessTokenTime, "current_time-wxAttribute.AccessTokenTime:", current_time-wxAttribute.AccessTokenTime)
			if current_time-wxAttribute.AccessTokenTime < 6000 {
				return wxAttribute.AccessToken
			}
		}
	}
	wx_url := "[REALM]?grant_type=client_credential&appid=[APPID]&secret=[SECRET]"
	realm_name := "https://api.weixin.qq.com/cgi-bin/token"
	appid := "wx570bbcc8cf9fdd80"
	secret := "c4b26e95739bc7defcc42e556cc7ae42"
	wx_url = strings.Replace(wx_url, "[REALM]", realm_name, -1)
	wx_url = strings.Replace(wx_url, "[APPID]", appid, -1)
	wx_url = strings.Replace(wx_url, "[SECRET]", secret, -1)
	beego.Debug("http_wx_token_url :", wx_url)
	resp, err := http.Get(wx_url)
	if err != nil {
		beego.Debug(err)
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		beego.Debug("http_wx_token_err :", err)
	} else {
		beego.Debug("http_wx_token_body :", string(body))
	}

	var atj models.AccessTokenJson
	if err := json.Unmarshal(body, &atj); err == nil {
		beego.Debug("http_wx_token_json :", atj)
		if atj.ErrCode == 0 {
			_, err = models.AddWxAttributeToken(atj.AccessToken)
			if err != nil {
				beego.Debug(err)
			}
			return atj.AccessToken
		} else {
			return ""
		}
	} else {
		beego.Debug("http_wx_token_err :", err)
		return ""
	}
}

func getDqsjTicket(access_toke string) string {
	wxAttribute, err := models.GetWxAttribute()
	if err != nil {
		beego.Debug(err)
	}
	if wxAttribute != nil {
		if len(wxAttribute.Ticket) != 0 {
			current_time := time.Now().Unix()
			beego.Debug("current_time:", current_time, "wxAttribute.TicketTime :", wxAttribute.TicketTime, "current_time-wxAttribute.TicketTime:", current_time-wxAttribute.TicketTime)
			if current_time-wxAttribute.TicketTime < 6000 {
				return wxAttribute.Ticket
			}
		}
	}

	wx_url := "[REALM]?access_token=[ACCESS_TOKEN]&type=jsapi"
	realm_name := "https://api.weixin.qq.com/cgi-bin/ticket/getticket"
	wx_url = strings.Replace(wx_url, "[REALM]", realm_name, -1)
	wx_url = strings.Replace(wx_url, "[ACCESS_TOKEN]", access_toke, -1)
	beego.Debug("http_wx_ticket_url :", wx_url)
	resp, err := http.Get(wx_url)
	if err != nil {
		beego.Debug(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		beego.Debug(err)
	} else {
		beego.Debug(string(body))
	}
	var ticket models.JsApiTicketJson
	if err := json.Unmarshal(body, &ticket); err == nil {
		beego.Debug("http_wx_ticket_ticketobj :", ticket)
		if ticket.ErrCode == 0 {
			_, err = models.AddWxAttributeTicket(ticket.Ticket)
			if err != nil {
				beego.Debug(err)
			}
			return ticket.Ticket
		}

		return ""
	} else {
		beego.Debug("http_wx_ticket_ticke :", err)
		return ""
	}
}

func chackDqsjAccount(ctx *context.Context) (bool, string) {
	ck, err := ctx.Request.Cookie(DQSJ_USERNAME)
	if err != nil {
		return false, ""
	}

	username := ck.Value

	ck, err = ctx.Request.Cookie(DQSJ_PASSWORD)
	if err != nil {
		return false, ""
	}

	password := ck.Value

	admin, err := models.GetOneDqsjAdmin(username)
	beego.Debug("GetOneDqsjAdmin admin:", admin)
	if err != nil {
		return false, ""
	}
	if admin != nil && strings.EqualFold(username, admin.Username) && strings.EqualFold(password, admin.Password) {
		beego.Debug(" cookie username ", username)
		return true, username
	} else {
		return false, username
	}

}

func getShareTitle() string {
	title := DQSJ_SHARE_TITLE
	config, err := models.GetConfig()
	if err != nil {
		beego.Error(err)
	}
	if config != nil && len(config.ShareTitle) != 0 {
		title = config.ShareTitle
	}
	beego.Debug("getShareTitle :", title)
	return title
}

type ClientJson struct {
	Rtype   string `json:"rtype"`
	Data    string `json:"data"`
	ErrCode int64  `json:"errcode"`
	ErrMsg  string `json:"errmsg"`
}
