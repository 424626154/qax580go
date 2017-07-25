package controllers

/**
* beermap
**/
import (
	"encoding/json"
	"qax580go/models"
	"qax580go/qutil"
	"strings"

	"github.com/astaxie/beego"
)

type BMResponseJson struct {
	Data    string `json:"data"`
	ErrCode int64  `json:"errcode"`
	ErrMsg  string `json:"errmsg"`
}

type BeerMapController struct {
	beego.Controller
}

//admin
func (c *BeerMapController) ALogin() {
	if c.Ctx.Input.IsGet() {
		beego.Debug("ALogin Get")
	}
	if c.Ctx.Input.IsPost() {
		beego.Debug("ALogin Post")
		username := c.Input().Get("user")
		password := c.Input().Get("password")
		autologin := c.Input().Get("autologin") == "on"
		beego.Debug("ALogin Post user:", username, "password:", password)
		if len(username) != 0 && len(password) != 0 {
			admin, err := models.GetOneBMAdmin(username)
			if err != nil {
				c.Redirect("/beermap/alogin", 302)
				return
			}
			if admin != nil && strings.EqualFold(username, admin.Username) && strings.EqualFold(password, admin.Password) {
				maxAge := 0
				if autologin {
					maxAge = 1<<31 - 1
				}
				c.Ctx.SetCookie(qutil.BM_PASSWORD, username, maxAge, "/")
				c.Ctx.SetCookie(qutil.BM_PASSWORD, password, maxAge, "/")
				beego.Debug("login ok------")
				c.Redirect("/beermap/admin", 302)
				return
			} else {
				c.Redirect("/beermap/alogin", 302)
				return
			}
		} else {
			c.Redirect("/beermap/alogin", 302)
			return
		}
	}
	c.TplName = "bmalogin.html"
}

func (c *BeerMapController) AMap() {
	if c.Ctx.Input.IsGet() {
		beego.Debug("AMap Get")
	}
	if c.Ctx.Input.IsPost() {
		beego.Debug("AMap Post")
	}
	c.TplName = "bmamap.html"
}

func (c *BeerMapController) Map() {
	if c.Ctx.Input.IsGet() {
		beego.Debug("Map Get")
	}
	if c.Ctx.Input.IsPost() {
		beego.Debug("Map Post")
	}
	c.TplName = "bmmap.html"
}

func (c *BeerMapController) Ajax() {
	if c.Ctx.Input.IsGet() {
		beego.Debug("Ajax Get")
	}
	response_obj := BMResponseJson{}
	response_obj.ErrCode = 1
	response_obj.ErrMsg = "未知错误"
	if c.Ctx.Input.IsPost() {
		beego.Debug("Ajax Post")
		op := c.Input().Get("op")
		switch op {
		case "addmaker":
			id := c.Input().Get("id")
			citycode := c.Input().Get("id")
			name := c.Input().Get("name")
			lng := c.Input().Get("lng")
			lat := c.Input().Get("lat")
			describe := c.Input().Get("describe")
			beego.Debug(id, name, lng, lat, describe)
			obj, err := models.AddBMMaker(id, citycode, name, lng, lat, describe)
			if err != nil {
				beego.Error(err)
				response_obj.ErrMsg = "存储失败"
			} else {
				response_obj.ErrCode = 0
				obj_body, err := json.Marshal(obj)
				if err != nil {
					beego.Error(err)
				}
				obj_json := string(obj_body)
				response_obj.Data = obj_json
			}
			break

		case "getmakers":
			objs, err := models.GetMakers()
			if err != nil {
				beego.Error(err)
				response_obj.ErrMsg = "获得数据失败"
			} else {
				response_obj.ErrCode = 0
				objs_body, err := json.Marshal(objs)
				if err != nil {
					beego.Error(err)
				}
				objs_json := string(objs_body)
				response_obj.Data = objs_json
				// beego.Debug(objs_json)
			}
			break
		}
	}

	body, err := json.Marshal(response_obj)
	if err != nil {
		beego.Error(err)
	}
	response_json := string(body)
	beego.Debug("RequestAjax response_json:", response_json)
	c.Ctx.WriteString(response_json)
}
