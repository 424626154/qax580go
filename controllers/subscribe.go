package controllers

/*
订阅与取消订阅
*/
import (
	"encoding/json"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/config"
	"io/ioutil"
	"net/http"
	"qax580go/models"
	"strings"
)

type SubsribeResponse struct {
	Errcode int64  `json:"errorcode"` //0 成功 1 失败 参数错误
	Errmsg  string `json:"errmsg"`
}

//http请求参数 subscribe_type from_openid to_user create_time
type SubsribeController struct {
	beego.Controller
}

func (c *SubsribeController) Get() {
	request(c)
}

func (c *SubsribeController) Post() {
	request(c)
}

func request(c *SubsribeController) {
	// response := &SubsribeResponse{}
	response_json := ""
	subscribe_type := c.Input().Get("subscribe_type")
	from_openid := c.Input().Get("from_openid")
	// to_user := c.Input().Get("to_user")
	// create_time := c.Input().Get("create_time")
	if len(subscribe_type) != 0 && len(from_openid) != 0 {
		// response.Errcode = 0
		// response.Errmsg = "ok"
		response_json = getSubscribeUserInfo(from_openid, c)
	} else {
		// response.Errcode = 1
		// response.Errmsg = "parameter error"
		response_json = `{"errcode":1,"errmsg":"parameter error"}`
	}

	// response_json, _ := json.Marshal(response)
	c.Ctx.WriteString(response_json)
}

func getSubscribeUserInfo(openid string, c *SubsribeController) string {
	response_json := ""
	isdebug := "true"
	iniconf, err := config.NewConfig("json", "conf/myconfig.json")
	if err != nil {
		beego.Debug(err)
	} else {
		isdebug = iniconf.String("qax580::isdebug")
	}
	wx_url := "[REALM]/wxqax/getuserinfo?openid=[OPENID]"
	// if beego.AppConfig.Bool("qax580::isdebug") {
	realm_name := ""
	if isdebug == "true" {
		realm_name = "http://localhost:9091"
	} else {
		realm_name = "http://182.92.167.29:8080"
	}
	wx_url = strings.Replace(wx_url, "[REALM]", realm_name, -1)
	wx_url = strings.Replace(wx_url, "[OPENID]", openid, -1)
	beego.Debug("----------------get UserInfo --------------------")
	beego.Debug(wx_url)

	resp, err := http.Get(wx_url)
	if err != nil {
		beego.Debug(err)
	}

	defer resp.Body.Close()
	beego.Debug("----------------get UserInfo body--------------------")
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		beego.Debug(err)
	} else {
		beego.Debug(string(body))
	}
	var uij models.Wxuserinfo
	if err := json.Unmarshal(body, &uij); err == nil {
		beego.Debug("----------------get UserInfo json--------------------")
		beego.Debug(uij)
		if uij.ErrCode == 0 {
			response_json = string(body)
		}
	} else {
		beego.Debug("----------------get UserInfo json error--------------------")
		beego.Debug(err)
	}
	return response_json
}
