package controllers

/*
天气预报
*/
import (
	"crypto/sha1"
	"encoding/hex"
	"encoding/json"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/config"
	"io/ioutil"
	"math/rand"
	"net/http"
	"qax580go/models"
	// "strconv"
	"fmt"
	"strings"
	"time"
)

const (
	KC_RAND_KIND_NUM   = 0 // 纯数字
	KC_RAND_KIND_LOWER = 1 // 小写字母
	KC_RAND_KIND_UPPER = 2 // 大写字母
	KC_RAND_KIND_ALL   = 3 // 数字、大小写字母
)

type WeatherController struct {
	beego.Controller
}

func (c *WeatherController) Get() {

	op := c.Input().Get("op")
	switch op {
	case "location":
		latitude := c.Input().Get("latitude")
		longitude := c.Input().Get("longitude")
		c.Data["latitude"] = latitude
		c.Data["longitude"] = longitude
		beego.Debug("latitude:", latitude)
		beego.Debug("longitude:", longitude)
		getWeather(longitude, latitude, c)
		c.TplName = "location.html"
		return
	}
	c.TplName = "weather.html"
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
	getWeatherToken(noncestr, timestamp, c)
	// getWeather(c)
	c.TplName = "weather.html"
}

func (c *WeatherController) Post() {
	c.TplName = "weather.html"
}

// 随机字符串
func getNonceStr(size int, kind int) string {
	ikind, kinds, result := kind, [][]int{[]int{10, 48}, []int{26, 97}, []int{26, 65}}, make([]byte, size)
	is_all := kind > 2 || kind < 0
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < size; i++ {
		if is_all { // random ikind
			ikind = rand.Intn(3)
		}
		scope, base := kinds[ikind][0], kinds[ikind][1]
		result[i] = uint8(base + rand.Intn(scope))
	}
	return string(result)
}

func getWeatherToken(noncestr string, timestamp int64, c *WeatherController) {
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
			getJsapiTicket(atj.AccessToken, noncestr, timestamp, c)
		}
	} else {
		beego.Debug("----------------get Token json error--------------------")
		beego.Debug(err)
	}
}

func getJsapiTicket(access_toke string, noncestr string, timestamp int64, c *WeatherController) {
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
	beego.Debug("getJsapiTicketUrl", wx_url)

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
			c.Data["Ticket"] = signatureWxJs(ticket.Ticket, noncestr, timestamp, "http://www.baoguangguang.cn/weather")
		}

		return
	} else {
		beego.Debug("----------------getJsapiTicket error--------------------")
		beego.Debug(err)
	}
}

//验证签名
func signatureWxJs(jsapi_ticket string, noncestr string, timestamp int64, url string) string {
	// timestamp_str := strconv.Itoa(timestamp)
	timestamp_str := fmt.Sprintf("%d", timestamp)
	signature_str := "jsapi_ticket=[JSAPI_TICKET]&noncestr=[NONCESTR]&timestamp=[TIMESTAMP]&url=[URL]"
	signature_str = strings.Replace(signature_str, "[JSAPI_TICKET]", jsapi_ticket, -1)
	signature_str = strings.Replace(signature_str, "[NONCESTR]", noncestr, -1)
	signature_str = strings.Replace(signature_str, "[TIMESTAMP]", timestamp_str, -1)
	signature_str = strings.Replace(signature_str, "[URL]", url, -1)
	beego.Debug("signature_str:", signature_str)
	signature := goWxJsSha1(signature_str)
	beego.Debug("signature:", signature)
	return signature
}

//对字符串进行SHA1哈希
func goWxJsSha1(str string) string {
	s := sha1.New()
	s.Write([]byte(str))
	return hex.EncodeToString(s.Sum(nil))
}

func getWeather(lon string, lat string, c *WeatherController) {
	url := "[REALM]?format=[FORMAT]&key=[KEY]&lon=[LON]&lat=[LAT]"
	url = strings.Replace(url, "[REALM]", "http://v.juhe.cn/weather/geo", -1)
	url = strings.Replace(url, "[FORMAT]", "2", -1)
	url = strings.Replace(url, "[KEY]", "f99ee0d9fe9af2a3ae272c33331a61dc", -1)
	url = strings.Replace(url, "[LON]", lon, -1)
	url = strings.Replace(url, "[LAT]", lat, -1)
	// url = "http://v.juhe.cn/weather/geo?format=2&key=f99ee0d9fe9af2a3ae272c33331a61dc&lon=116.39277&lat=39.933748"
	beego.Debug("url:", url)

	resp, err := http.Get(url)
	if err != nil {
		beego.Debug(err)
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	beego.Debug("----------------getWeather body--------------------")
	if err != nil {
		beego.Debug(err)
	} else {
		beego.Debug(string(body))
	}
	var uij models.WeatherJson
	if err := json.Unmarshal(body, &uij); err == nil {
		fmt.Println("================json str 转struct==")
		today := uij.Result.Today
		beego.Debug(today)
		c.Data["ErrorCode"] = uij.ErrorCode
		c.Data["Today"] = today
	} else {
		beego.Debug(err)
	}
}
