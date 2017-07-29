package qutil

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"qax580go/models"
	"strings"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/config"
	"github.com/astaxie/beego/context"
)

// const COOKIE_UID = "zz580_uid"

func GetCookeUid() string {
	return COOKIE_UID
}

func ChackAccount(ctx *context.Context) (bool, string) {
	ck, err := ctx.Request.Cookie("username")
	if err != nil {
		return false, ""
	}

	username := ck.Value

	ck, err = ctx.Request.Cookie("password")
	if err != nil {
		return false, ""
	}

	password := ck.Value

	admin, err := models.GetOneAdmin(username)
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

func GetImageUrl() string {
	url := ""
	isdebug := "flase"
	iniconf, err := config.NewConfig("json", "conf/myconfig.json")
	if err != nil {
		beego.Error(err)
	} else {
		isdebug = iniconf.String("qax580::isdebug")
		if isdebug == "true" {
			url = iniconf.String("qax580::imgurltest")
		} else {
			url = iniconf.String("qax580::imgurl")
		}

	}
	return url
}

type AccessTokenJson struct {
	AccessToken  string `json:"access_token"`
	ExpiresIn    int64  `json:"expires_in"`
	RefreshToken string `json:"refresh_token"`
	OpenID       string `json:"openid"`
	Scope        string `json:"scope"`
	Unionid      string `json:"unionid"`
	ErrCode      int64  `json:"errcode"`
	ErrMsg       string `json:"errmsg"`
}

func GetToken() (errcode int64, token string) {
	r_errcode := int64(0)
	r_token := ""
	// https: //api.weixin.qq.com/cgi-bin/token?grant_type=client_credential&appid=APPID&secret=APPSECRET
	appid := "wx570bbcc8cf9fdd80"
	secret := "c4b26e95739bc7defcc42e556cc7ae42"
	wx_url := "[REALM]?appid=[APPID]&secret=[SECRET]&grant_type=client_credential"
	realm_name := "https://api.weixin.qq.com/cgi-bin/token"
	wx_url = strings.Replace(wx_url, "[REALM]", realm_name, -1)
	wx_url = strings.Replace(wx_url, "[APPID]", appid, -1)
	wx_url = strings.Replace(wx_url, "[SECRET]", secret, -1)
	beego.Debug("token url:", wx_url)
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
	// body := []byte(`{"access_token":"ACCESS_TOKEN","expires_in":7200}`)
	var atj AccessTokenJson
	if err := json.Unmarshal(body, &atj); err == nil {
		beego.Debug(atj)
		if atj.ErrCode == 0 {
			r_token = atj.AccessToken
		} else {
			r_errcode = atj.ErrCode
		}
	} else {
		beego.Debug(err)
	}
	return r_errcode, r_token
}

func GetWxUser(openid string, access_token string) (models.Wxuserinfo, error) {
	user := models.Wxuserinfo{}
	response_json := `{"errcode":1,"errmsg":"getWxUser error"}`
	// ?access_token=ACCESS_TOKEN&openid=OPENID&lang=zh_CN
	query_url := "[REALM]?access_token=[ACCESS_TOKEN]&openid=[OPENID]&lang=zh_CN"
	isdebug := "true"
	iniconf, err := config.NewConfig("json", "conf/myconfig.json")
	if err != nil {
		beego.Debug(err)
	} else {
		isdebug = iniconf.String("qax580::isdebug")
	}
	realm_name := ""
	if isdebug == "true" {
		realm_name = "http://localhost:9091"
	} else {
		realm_name = "https://api.weixin.qq.com/cgi-bin/user/info"
	}
	realm_name = "https://api.weixin.qq.com/cgi-bin/user/info"
	query_url = strings.Replace(query_url, "[REALM]", realm_name, -1)
	query_url = strings.Replace(query_url, "[ACCESS_TOKEN]", access_token, -1)
	query_url = strings.Replace(query_url, "[OPENID]", openid, -1)
	beego.Debug("importUser url:", query_url)

	resp, err := http.Get(query_url)
	if err != nil {
		beego.Debug(err)
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		beego.Debug(err)
		body = []byte(response_json)
	} else {
		beego.Debug("wxqax get UserInfo body", string(body))
	}
	var uij models.Wxuserinfo
	if err := json.Unmarshal(body, &uij); err == nil {
		beego.Debug("wxqax get UserInfo json", uij)
		if uij.ErrCode == 0 {
			user = uij
		}

	} else {
		beego.Error(err)
	}
	return user, err
}

func GetError(errorcode int64) string {
	error_indo := ""
	if errorcode == 202201 {
		error_indo = "车次不能为空"
	}
	if errorcode == 202202 {
		error_indo = "查询不到车次的相关信息"
	}
	if errorcode == 202203 {
		error_indo = "出发站或终点站不能为空"
	}
	if errorcode == 202204 {
		error_indo = "查询不到结果"
	}
	if errorcode == 202205 {
		error_indo = "错误的出发站名称"
	}
	if errorcode == 202206 {
		error_indo = "错误的到达站名称"
	}
	if errorcode == 202207 {
		error_indo = "查询不到余票相关数据哦"
	}
	if errorcode == 202208 {
		error_indo = "错误的请求，请确认传递的参数正确"
	}
	if errorcode == 202209 {
		error_indo = "请求12306网络错误,请重试"
	}
	if errorcode == 202210 {
		error_indo = "12306账号密码错误"
	}
	if errorcode == 202211 {
		error_indo = "邮箱不存在"
	}
	if errorcode == 202212 {
		error_indo = "查询出错"
	}
	if errorcode == 202213 {
		error_indo = "提交订单超时，请重试"
	}
	if errorcode == 202214 {
		error_indo = "出票失败"
	}
	if errorcode == 202215 {
		error_indo = "排队失败"
	}
	if errorcode == 202216 {
		error_indo = "该车次无法预定"
	}
	if errorcode == 202217 {
		error_indo = "不合法的座位类型"
	}
	return error_indo
}
