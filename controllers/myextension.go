package controllers

/*
我的推广
*/
import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"qax580go/models"
	"qax580go/qutil"
	"strings"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/config"
)

type MyExtensionController struct {
	beego.Controller
}

func (c *MyExtensionController) Get() {
	openid := getMyExtensionCookie(c)
	beego.Debug("openid:", openid)
	if len(openid) != 0 {
		c.Data["isOpenid"] = true
		getMyExtensionToken(openid, c)
	} else {
		c.Data["isOpenid"] = false
	}
	c.TplName = "myextension.html"
}

func (c *MyExtensionController) Post() {
	c.TplName = "myextension.html"
}
func getMyExtensionCookie(c *MyExtensionController) string {
	isUser := false
	openid := c.Ctx.GetCookie(qutil.COOKIE_WX_OPENID)
	beego.Debug("------------openid--------")
	beego.Debug(openid)
	if len(openid) != 0 {
		wxuser, err := models.GetOneWxUserInfo(openid)
		if err != nil {
			beego.Error(err)
		} else {
			isUser = true
			beego.Debug("--------------wxuser----------")
			beego.Debug(wxuser)
			c.Data["WxUser"] = wxuser
		}
	}
	c.Data["isUser"] = isUser
	return openid
}
func getMyExtensionToken(openid string, c *MyExtensionController) {
	isdebug := "true"
	iniconf, err := config.NewConfig("json", "conf/myconfig.json")
	if err != nil {
		beego.Debug(err)
	} else {
		isdebug = iniconf.String("qax580::isdebug")
	}
	wx_url := "[REALM]?appid=[APPID]&secret=[SECRET]&grant_type=client_credential"
	// if beego.AppConfig.Bool("qax580::isdebug") {
	realm_name := ""
	if isdebug == "true" {
		realm_name = "http://localhost:9093"
	} else {
		realm_name = "https://api.weixin.qq.com/cgi-bin/token"
	}
	appid := "wx570bbcc8cf9fdd80"
	secret := "c4b26e95739bc7defcc42e556cc7ae42"
	wx_url = strings.Replace(wx_url, "[REALM]", realm_name, -1)
	wx_url = strings.Replace(wx_url, "[APPID]", appid, -1)
	wx_url = strings.Replace(wx_url, "[SECRET]", secret, -1)
	beego.Debug("----------------get Token --------------------")
	beego.Debug(wx_url)
	resp, err := http.Get(wx_url)
	if err != nil {
		beego.Debug(err)
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	beego.Debug("----------------get Token body--------------------")
	if err != nil {
		beego.Debug(err)
	} else {
		beego.Debug(string(body))
	}

	var atj models.AccessTokenJson
	if err := json.Unmarshal(body, &atj); err == nil {
		beego.Debug("----------------get Token json--------------------")
		beego.Debug(atj)
		if atj.ErrCode == 0 {
			getMyExtensionQRCode(atj.AccessToken, openid, c)
		}
	} else {
		beego.Debug("----------------get Token json error--------------------")
		beego.Debug(err)
	}
}

/*
获得二维码
http请求方式: POST
URL: https://api.weixin.qq.com/cgi-bin/qrcode/create?access_token=TOKEN
POST数据格式：json
POST数据例子：{"action_name": "QR_LIMIT_SCENE", "action_info": {"scene": {"scene_id": 123}}}
或者也可以使用以下POST数据创建字符串形式的二维码参数：
{"action_name": "QR_LIMIT_STR_SCENE", "action_info": {"scene": {"scene_str": "123"}}}
*/
func getMyExtensionQRCode(token string, openid string, c *MyExtensionController) {
	isdebug := "true"
	iniconf, err := config.NewConfig("json", "conf/myconfig.json")
	if err != nil {
		beego.Debug(err)
	} else {
		isdebug = iniconf.String("qax580::isdebug")
	}
	wx_url := "[REALM]?access_token=[TOKEN]"
	// if beego.AppConfig.Bool("qax580::isdebug") {
	realm_name := ""
	if isdebug == "true" {
		realm_name = "http://localhost:9095"
	} else {
		realm_name = "https://api.weixin.qq.com/cgi-bin/qrcode/create"
	}

	wx_url = strings.Replace(wx_url, "[REALM]", realm_name, -1)
	wx_url = strings.Replace(wx_url, "[TOKEN]", token, -1)
	beego.Debug("----------------get QRCode --------------------")
	beego.Debug(wx_url)
	// post_json := `{"action_name": "QR_LIMIT_SCENE", "action_info": {"scene": {"1": [SCENESTR]}}}`
	post_json := `{"action_name": "QR_LIMIT_SCENE", "action_info": {"scene": {"scene_id": 1}}}`
	post_json = strings.Replace(post_json, "[SCENESTR]", openid, -1)
	beego.Debug("post_json :", post_json)
	post_body := bytes.NewBuffer([]byte(post_json))
	resp, err := http.Post(wx_url, "application/json;charset=utf-8", post_body)
	if err != nil {
		beego.Debug(err)
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	beego.Debug("----------------get QRCode body--------------------")
	if err != nil {
		beego.Debug(err)
	} else {
		beego.Debug(string(body))
	}

	var atj models.QRCodeJson
	if err := json.Unmarshal(body, &atj); err == nil {
		beego.Debug("----------------get QRCode json--------------------")
		beego.Debug(atj)
		if atj.ErrCode == 0 {
			beego.Debug("qrcode url", atj.Url)
			c.Data["Url"] = "https://www.baidu.com"
		}
	} else {
		beego.Debug("----------------get QRCode json error--------------------")
		beego.Debug(err)
	}
}
