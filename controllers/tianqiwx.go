package controllers

/*
周边wifi
*/
import (
	"encoding/json"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/config"
	"io/ioutil"
	"net/http"
	"qax580go/models"
	"strings"
	"time"
)

type Tianqi struct {
	Resultcode string   `json:"resultcode"`
	Reason     string   `json:"reason"`
	ErrorCode  int64    `json:"error_code"`
	TQResult   TQResult `json:"result"`
}
type TQResult struct {
	TQSk     TQSk       `json:"sk"`
	TQToday  TQToday    `json:"today"`
	TQFuture []TQFuture `json:"future"`
}
type TQSk struct {
	Temp          string `json:"temp"`
	WindDirection string `json:"wind_direction"`
	WindStrength  string `json:"wind_strength"`
	Humidity      string `json:"humidity"`
	Time          string `json:"time"`
}
type TQToday struct {
	City           string `json:"city"`
	DateY          string `json:"date_y"`
	Week           string `json:"week"`
	Temperature    string `json:"temperature"`
	Weather        string `json:"weather"`
	Wind           string `json:"wind"`
	Dressingindex  string `json:"dressing_index"`
	DressingAdvice string `json:"dressing_advice"`
	UvIndex        string `json:"uv_index"`
	WashIndex      string `json:"wash_index"`
	TravelIndex    string `json:"travel_index"`
	ExerciseIndex  string `json:"exercise_index"`
}

type TQFuture struct {
	Temperature string `json:"temperature"`
	Weather     string `json:"weather"`
	Wind        string `json:"wind"`
	Week        string `json:"week"`
	Date        string `json:"date"`
}

type TianqiWXController struct {
	beego.Controller
}

func (c *TianqiWXController) Get() {
	op := c.Input().Get("op")
	switch op {
	case "location":
		latitude := c.Input().Get("latitude")
		longitude := c.Input().Get("longitude")
		c.Data["latitude"] = latitude
		c.Data["longitude"] = longitude
		beego.Debug("latitude:", latitude)
		beego.Debug("longitude:", longitude)
		getTianqi(longitude, latitude, c)
		c.TplName = "tianqi.html"
		return
	}
	appId := ""
	iniconf, err := config.NewConfig("json", "conf/myconfig.json")
	if err != nil {
		beego.Debug(err)
	} else {
		appId = iniconf.String("qax580::appid")
	}
	timestamp := time.Now().Unix()
	noncestr := getNonceStr(16, KC_RAND_KIND_ALL)
	c.Data["AppId"] = appId
	c.Data["TimesTamp"] = timestamp
	c.Data["NonceStr"] = noncestr
	getTianqiWxJsToken(noncestr, timestamp, c)
	c.TplName = "tianqiwx.html"
	// getTianqi("116.366324", "39.905859", c)
	// c.TplName = "tianqi.html"
}

func (c *TianqiWXController) Post() {
	c.TplName = "tianqiwx.html"
}

func getTianqiWxJsToken(noncestr string, timestamp int64, c *TianqiWXController) {
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
			getTianqiJsapiTicket(atj.AccessToken, noncestr, timestamp, c)
		}
	} else {
		beego.Debug("----------------get Token json error--------------------")
		beego.Debug(err)
	}
}

func getTianqiJsapiTicket(access_toke string, noncestr string, timestamp int64, c *TianqiWXController) {
	isdebug := "true"
	iniconf, err := config.NewConfig("json", "conf/myconfig.json")
	if err != nil {
		beego.Debug(err)
	} else {
		isdebug = iniconf.String("qax580::isdebug")
	}
	wx_url := "[REALM]?access_token=[ACCESS_TOKEN]&type=jsapi"
	// if beego.AppConfig.Bool("qax580::isdebug") {
	realm_name := ""
	if isdebug == "true" {
		realm_name = "http://localhost:9092"
	} else {
		realm_name = "https://api.weixin.qq.com/cgi-bin/ticket/getticket"
	}
	wx_url = strings.Replace(wx_url, "[REALM]", realm_name, -1)
	wx_url = strings.Replace(wx_url, "[ACCESS_TOKEN]", access_toke, -1)
	beego.Debug("getWifiJsapiTicket", wx_url)

	resp, err := http.Get(wx_url)
	if err != nil {
		beego.Debug(err)
	}

	defer resp.Body.Close()
	beego.Debug("----------------getJsapiTicket body--------------------")
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		beego.Debug(err)
	} else {
		beego.Debug(string(body))
	}
	var ticket models.JsApiTicketJson
	if err := json.Unmarshal(body, &ticket); err == nil {
		beego.Debug("----------------getJsapiTicket json--------------------")
		beego.Debug(ticket)
		if ticket.ErrCode == 0 {
			c.Data["Ticket"] = signatureWxJs(ticket.Ticket, noncestr, timestamp, "http://www.baoguangguang.cn/tianqiwx")
		}

		return
	} else {
		beego.Debug("----------------getJsapiTicket error--------------------")
		beego.Debug(err)
	}
}

func getTianqi(lon string, lat string, c *TianqiWXController) {
	//http://v.juhe.cn/weather/geo?format=2&key=您申请的KEY&lon=116.39277&lat=39.933748
	url := "[REALM]?format=2&key=[KEY]&lon=[LON]&lat=[LAT]&r=3000&type=1"
	url = strings.Replace(url, "[REALM]", "http://v.juhe.cn/weather/geo", -1)
	url = strings.Replace(url, "[KEY]", "f99ee0d9fe9af2a3ae272c33331a61dc", -1)
	url = strings.Replace(url, "[LON]", lon, -1)
	url = strings.Replace(url, "[LAT]", lat, -1)
	beego.Debug("url:", url)

	resp, err := http.Get(url)
	if err != nil {
		beego.Debug(err)
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	beego.Debug("----------------getTianqi body--------------------")
	if err != nil {
		beego.Debug(err)
	} else {
		beego.Debug(string(body))
	}
	var obj Tianqi
	if err := json.Unmarshal(body, &obj); err == nil {
		beego.Debug("================json str 转struct==")
		// beego.Debug(obj)
		beego.Debug("ErrorCode", obj.ErrorCode)
		c.Data["ErrorCode"] = obj.ErrorCode
		if obj.ErrorCode != 0 {
			c.Data["ErrorInfo"] = getTianqiError(obj.ErrorCode)
		} else {
			// beego.Debug(obj.TQResult.TQFuture)
			c.Data["Sk"] = obj.TQResult.TQSk
			c.Data["Today"] = obj.TQResult.TQToday
			c.Data["Future"] = obj.TQResult.TQFuture
		}
	} else {
		beego.Debug(err)
	}
}

func getTianqiError(errorcode int64) string {
	error_info := ""
	if errorcode == 203901 {
		error_info = "查询城市不能为空"
	}
	if errorcode == 203902 {
		error_info = "查询不到该城市的天气"
	}
	if errorcode == 203903 {
		error_info = "查询出错，请重试"
	}
	if errorcode == 203904 {
		error_info = "错误的GPS坐标"
	}
	if errorcode == 203905 {
		error_info = "GPS坐标解析出错，请确认提供的坐标正确（暂支持国内）"
	}
	if errorcode == 203906 {
		error_info = "IP地址错误"
	}
	if errorcode == 203907 {
		error_info = "查询不到该IP地址相关的天气信息"
	}
	return error_info
}
