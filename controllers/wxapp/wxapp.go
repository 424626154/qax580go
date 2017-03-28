package controllers

/**
*微信小程序服务器
*
**/
import (
	"encoding/json"
	"github.com/astaxie/beego"
	"github.com/satori/go.uuid"
	"qax580go/models"
)

type WxAppController struct {
	beego.Controller
}

//获得UUID count=n [n0,n1,..]
func (c *WxAppController) Getuuid() {
	if c.Ctx.Input.IsGet() {
		beego.Debug("Getuuid Get")
	}
	if c.Ctx.Input.IsPost() {
		beego.Debug("Getuuid Post")
	}
	u1 := uuid.NewV4()
	beego.Debug(u1)
	responseJson := models.ResponseJson{}
	body, err := json.Marshal(responseJson)
	if err != nil {
		beego.Error(err)
	}
	response_json := string(body)
	c.Ctx.WriteString(response_json)
	return
}
