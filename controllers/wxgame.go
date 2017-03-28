package controllers

import (
	// "crypto/sha1"
	// "encoding/hex"
	"encoding/json"
	// "fmt"
	"github.com/astaxie/beego"
	"io/ioutil"
	"math/rand"
	"net/http"
	"strings"
	"time"
)

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

//分享JSON
type ShareJson struct {
	AppId     string `json:"appId"`
	Timestamp int64  `json:"timestamp"`
	NonceStr  string `json:"nonceStr"`
	Signature string `json:"signature"`
}

type JsApiTicketJson struct {
	Id        int64
	Ticket    string `json:"ticket"`
	ExpiresIn int64  `json:"expires_in"`
	ErrCode   int64  `json:"errcode"`
	ErrMsg    string `json:"errmsg"`
}

const (
	APPID  = "wx570bbcc8cf9fdd80"
	SECRET = "c4b26e95739bc7defcc42e556cc7ae42"
)

type WeixinGameController struct {
	beego.Controller
}

func (c *WeixinGameController) Get() {
	url := c.Input().Get("url")
	if len(url) == 0 {
		beego.Debug("required parameter error")
		body := []byte(`{"errcode":1000,"errmsg":"url erro"}`)
		c.Ctx.WriteString(string(body))
		return
	}
	responseShareJson(url, c)
}

//返回分享JSON
func responseShareJson(url string, c *WeixinGameController) {
	timestamp := time.Now().Unix()
	noncestr := getWxGameNonceStr(16, KC_RAND_KIND_ALL)
	signature := getWxGameSignature(noncestr, timestamp, url, c)
	shareJson := &ShareJson{AppId: APPID, Timestamp: timestamp, NonceStr: noncestr, Signature: signature}
	body, err := json.Marshal(shareJson)
	if err != nil {
		beego.Debug(err)
	}
	beego.Debug("body :", body)
	c.Ctx.WriteString(string(body))
	return
}

// 随机字符串
func getWxGameNonceStr(size int, kind int) string {
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

//获得签名
func getWxGameSignature(noncestr string, timestamp int64, url string, c *WeixinGameController) string {
	signature := ""
	wx_url := "[REALM]?appid=[APPID]&secret=[SECRET]&grant_type=client_credential"
	realm_name := "https://api.weixin.qq.com/cgi-bin/token"
	wx_url = strings.Replace(wx_url, "[REALM]", realm_name, -1)
	wx_url = strings.Replace(wx_url, "[APPID]", APPID, -1)
	wx_url = strings.Replace(wx_url, "[SECRET]", SECRET, -1)
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
			signature = getWxGameJsapiTicket(atj.AccessToken, noncestr, timestamp, url, c)
		}
	} else {
		beego.Debug(err)
	}

	return signature
}

//获得jsapi_ticket
func getWxGameJsapiTicket(access_toke string, noncestr string, timestamp int64, url string, c *WeixinGameController) string {
	signature := ""
	wx_url := "[REALM]?access_token=[ACCESS_TOKEN]&type=jsapi"
	realm_name := "https://api.weixin.qq.com/cgi-bin/ticket/getticket"
	wx_url = strings.Replace(wx_url, "[REALM]", realm_name, -1)
	wx_url = strings.Replace(wx_url, "[ACCESS_TOKEN]", access_toke, -1)
	beego.Debug("getJsapiTicketUrl", wx_url)
	resp, err := http.Get(wx_url)
	if err != nil {
		beego.Debug(err)
	}

	defer resp.Body.Close()
	beego.Debug("----------------getWxGameJsapiTicket body--------------------")
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		beego.Debug(err)
	} else {
		beego.Debug(string(body))
	}
	// body := []byte(`{"errcode":0,"errmsg":"ok","ticket":"bxLdikRXVbTPdHSM05e5u5sUoXNKd8-41ZO3MhKoyN5OfkWITDGgnr2fwJ0m9E8NYzWKVZvdVtaUgWvsdshFKA","expires_in":7200}`)
	var ticket JsApiTicketJson
	if err := json.Unmarshal(body, &ticket); err == nil {
		if ticket.ErrCode == 0 {
			signature = signatureWxJs(ticket.Ticket, noncestr, timestamp, url)
		}
	} else {

	}
	return signature
}

// //验证签名
// func signatureWxJs(jsapi_ticket string, noncestr string, timestamp int64, url string) string {
// 	timestamp_str := fmt.Sprintf("%d", timestamp)
// 	signature_str := "jsapi_ticket=[JSAPI_TICKET]&noncestr=[NONCESTR]&timestamp=[TIMESTAMP]&url=[URL]"
// 	signature_str = strings.Replace(signature_str, "[JSAPI_TICKET]", jsapi_ticket, -1)
// 	signature_str = strings.Replace(signature_str, "[NONCESTR]", noncestr, -1)
// 	signature_str = strings.Replace(signature_str, "[TIMESTAMP]", timestamp_str, -1)
// 	signature_str = strings.Replace(signature_str, "[URL]", url, -1)
// 	signature := goWxJsSha1(signature_str)
// 	beego.Debug("signature:", signature)
// 	return signature
// }

// //对字符串进行SHA1哈希
// func goWxJsSha1(str string) string {
// 	s := sha1.New()
// 	s.Write([]byte(str))
// 	return hex.EncodeToString(s.Sum(nil))
// }
