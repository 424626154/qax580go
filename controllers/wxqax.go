package controllers

/*
微信http服务器
*/
import (
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/config"
	"io/ioutil"
	"net/http"
	"qax580go/models"
	"strings"
	"time"
)

type WxqaxController struct {
	beego.Controller
}

/**************请求接口*******************/
/*
关注与取消关注请求
／wxqax/sunscribe
//http请求参数 subscribe_type from_openid to_user create_time
*/
func (c *WxqaxController) Sunscribe() {
	response_json := `{"errcode":1,"errmsg":"Sunscribe error"}`
	subscribe_type := c.Input().Get("subscribe_type")
	from_openid := c.Input().Get("from_openid")
	appid := c.Input().Get("appid")
	secret := c.Input().Get("secret")
	if len(subscribe_type) != 0 && len(from_openid) != 0 {
		response_json = getWxAccessToken(c, from_openid, appid, secret)

		var user models.Wxuserinfo
		if err := json.Unmarshal([]byte(response_json), &user); err == nil {
			beego.Debug("----------------get Wxuserinfo json--------------------")
			beego.Debug(user)
			if user.ErrCode == 0 {
				state, err := models.SunscribeWxUserInfo(user)
				if err != nil {
					beego.Error(err)
					response_json = `{"errcode":1,"errmsg":"AddWxUserInfo error"}`
				} else {
					beego.Debug("SunscribeWxUserInfo state", state)
					if subscribe_type == "subscribe" && state == 0 {
						err = models.AddWxUserMoney(user.OpenId, 4)
						if err != nil {
							beego.Error(err)
							response_json = `{"errcode":1,"errmsg":"AddWxUserMoney error"}`
						} else {
							_, err = models.AddUserMoneyRecord(user.OpenId, MONEY_SUBSCRIBE_SUM, MONEY_SUBSCRIBE)
						}
					}
				}
			}
		} else {
			beego.Debug("----------------get Token json error--------------------")
			beego.Debug(err)
		}

	}
	c.Ctx.WriteString(response_json)

}

/**************请求接口*******************/
/*
 openid=
{
    "subscribe": 1,
    "openid": "o6_bmjrPTlm6_2sgVt7hMZOPfL2M",
    "nickname": "Band",
    "sex": 1,
    "language": "zh_CN",
    "city": "广州",
    "province": "广东",
    "country": "中国",
    "headimgurl":    "http://wx.qlogo.cn/mmopen/g3MonUZtNHkdmzicIlibx6iaFqAc56vxLSUfpb6n5WKSYVY0ChQKkiaJSgQ1dZuTOgvLLrhJbERQQ4eMsv84eavHiaiceqxibJxCfHe/0",
   "subscribe_time": 1382694957,
   "unionid": " o6_bmasdasdsad6_2sgVt7hMZOPfL"
   "remark": "",
   "groupid": 0
}

{"errcode":40013,"errmsg":"invalid appid"}
*/
func (c *WxqaxController) Getuserinfo() {
	response_json := `{"errcode":1,"errmsg":"Getuserinfo error"}`
	openid := c.Input().Get("openid")
	if len(openid) != 0 {
		response_json = getWxAccessToken(c, openid, qa_appid, qa_secret)
	} else {
		response_json = `{"errcode":1,"errmsg":"parameter error"}`
	}
	c.Ctx.WriteString(response_json)
}

func getWxAccessToken(c *WxqaxController, openid string, appid string, secret string) string {
	// https: //api.weixin.qq.com/cgi-bin/token?grant_type=client_credential&appid=APPID&secret=APPSECRET
	response_json := `{"errcode":1,"errmsg":"getWxAccessToken error"}`
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
		realm_name = "http://localhost:9090"
	} else {
		realm_name = "https://api.weixin.qq.com/cgi-bin/token"
	}
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
			response_json = getWxUserInfo(atj.AccessToken, openid, c)
		}
	} else {
		beego.Debug("----------------get Token json error--------------------")
		beego.Debug(err)
	}
	return response_json
}

func getWxUserInfo(access_toke, openid string, c *WxqaxController) string {
	response_json := `{"errcode":1,"errmsg":"getWxUserInfo error"}`
	isdebug := "true"
	iniconf, err := config.NewConfig("json", "conf/myconfig.json")
	if err != nil {
		beego.Debug(err)
	} else {
		isdebug = iniconf.String("qax580::isdebug")
	}
	wx_url := "[REALM]?access_token=[ACCESS_TOKEN]&openid=[OPENID]&lang=zh_CN"
	// if beego.AppConfig.Bool("qax580::isdebug") {
	realm_name := ""
	if isdebug == "true" {
		realm_name = "http://localhost:9091"
	} else {
		realm_name = "https://api.weixin.qq.com/cgi-bin/user/info"
	}
	wx_url = strings.Replace(wx_url, "[REALM]", realm_name, -1)
	wx_url = strings.Replace(wx_url, "[ACCESS_TOKEN]", access_toke, -1)
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

func getToken() (errcode int64, token string) {
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

func getWxUser(openid string, access_token string) (models.Wxuserinfo, error) {
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

/**
投票校验json
0 成功 用户已注册 1失败 openid参数错误 2 失败 pollsid参数错误 3 参与投票数据错误 100 失败 未知错误
*/
type PollCheck struct {
	ErrCode int32  `json:"errcode"`
	ErrMsg  string `json:"errmsg"`
}

/**
投票资格验证
*/
func (c *WxqaxController) Qualification() {
	response_json := `{"errcode":100,"errmsg":"未知错误"}`
	pollChecks := [7]PollCheck{PollCheck{int32(0), "用户已经关注，可以操作"}, PollCheck{int32(1), "openid参数错误"},
		PollCheck{int32(2), "pollsid参数错误"}, PollCheck{int32(3), "参与投票数据错误"}, PollCheck{int32(4), "获取token错误"},
		PollCheck{int32(5), "获取userinfo错误"}, PollCheck{int32(6), "用户未关注"}}
	var pollCheck PollCheck
	pollCheck.ErrCode = int32(100)
	pollCheck.ErrMsg = "未知错误"
	openid := c.Input().Get("openid")
	pollsid := c.Input().Get("pollsid")
	beego.Debug("/wxqax/qualification openid:", openid)
	beego.Debug("/wxqax/qualification pollsid:", pollsid)
	//验证openid
	if len(openid) == 0 {
		pollCheck.ErrCode = pollChecks[1].ErrCode
		pollCheck.ErrMsg = pollChecks[1].ErrMsg
		res_json, err := json.Marshal(pollCheck)
		if err != nil {
			beego.Error(err)
		} else {
			response_json = string(res_json)
		}
		c.Ctx.WriteString(response_json)
		return
	}
	//验证pollsid
	if len(pollsid) == 0 {
		pollCheck.ErrCode = pollChecks[2].ErrCode
		pollCheck.ErrMsg = pollChecks[2].ErrMsg
		res_json, err := json.Marshal(pollCheck)
		if err != nil {
			beego.Error(err)
		} else {
			response_json = string(res_json)
		}
		c.Ctx.WriteString(response_json)
		return
	}
	obj, err := models.GetOnePolls(pollsid)
	if err != nil {
		beego.Error(err)
		pollCheck.ErrCode = pollChecks[3].ErrCode
		pollCheck.ErrMsg = pollChecks[3].ErrMsg
		res_json, err := json.Marshal(pollCheck)
		if err != nil {
			beego.Error(err)
		} else {
			response_json = string(res_json)
		}
		c.Ctx.WriteString(response_json)
		return
	} else {
		appid := obj.Appid
		secret := obj.Secret
		beego.Debug("/wxqax/qualification Appid:", appid)
		beego.Debug("/wxqax/qualification Secret:", secret)
		if len(appid) != 0 && len(secret) != 0 {
			//获取用户token
			tokenobj, err := getWxToken(appid, secret)
			if err != nil {
				pollCheck.ErrCode = pollChecks[4].ErrCode
				pollCheck.ErrMsg = pollChecks[4].ErrMsg
				res_json, err := json.Marshal(pollCheck)
				if err != nil {
					beego.Error(err)
				} else {
					response_json = string(res_json)
				}
				c.Ctx.WriteString(response_json)
				return
			} else {
				if tokenobj.ErrCode == 0 {
					//获取用户信息
					userinfo, err := getWxUser(openid, tokenobj.AccessToken)
					if err != nil {
						pollCheck.ErrCode = pollChecks[5].ErrCode
						pollCheck.ErrMsg = pollChecks[5].ErrMsg
						res_json, err := json.Marshal(pollCheck)
						if err != nil {
							beego.Error(err)
						} else {
							response_json = string(res_json)
						}
						c.Ctx.WriteString(response_json)
						return
					} else {
						if userinfo.ErrCode == 0 {
							//判断用户是否注册
							beego.Debug("/wxqax/qualification subscribe:", userinfo.Subscribe)
							if userinfo.Subscribe == 1 { //1已经注册
								pollCheck.ErrCode = pollChecks[0].ErrCode
								pollCheck.ErrMsg = pollChecks[0].ErrMsg
								res_json, err := json.Marshal(pollCheck)
								if err != nil {
									beego.Error(err)
								} else {
									response_json = string(res_json)
								}
								beego.Debug("/wxqax/qualification response_json:", response_json)
								c.Ctx.WriteString(response_json)
								return
							} else {
								pollCheck.ErrCode = pollChecks[6].ErrCode
								pollCheck.ErrMsg = pollChecks[6].ErrMsg
								res_json, err := json.Marshal(pollCheck)
								if err != nil {
									beego.Error(err)
								} else {
									response_json = string(res_json)
								}
								beego.Debug("/wxqax/qualification response_json:", response_json)
								c.Ctx.WriteString(response_json)
								return
							}
						} else {
							pollCheck.ErrCode = pollChecks[5].ErrCode
							pollCheck.ErrMsg = pollChecks[5].ErrMsg
							res_json, err := json.Marshal(pollCheck)
							if err != nil {
								beego.Error(err)
							} else {
								response_json = string(res_json)
							}
							c.Ctx.WriteString(response_json)
							return
						}
					}
				} else {
					beego.Error(err)
					pollCheck.ErrCode = pollChecks[4].ErrCode
					pollCheck.ErrMsg = pollChecks[4].ErrMsg
					res_json, err := json.Marshal(pollCheck)
					if err != nil {
						beego.Error(err)
					} else {
						response_json = string(res_json)
					}
					c.Ctx.WriteString(response_json)
					return
				}
			}
		} else {
			pollCheck.ErrCode = pollChecks[3].ErrCode
			pollCheck.ErrMsg = pollChecks[3].ErrMsg
			res_json, err := json.Marshal(pollCheck)
			if err != nil {
				beego.Error(err)
			} else {
				response_json = string(res_json)
			}
			c.Ctx.WriteString(response_json)
			return
		}
	}

	c.Ctx.WriteString(response_json)
}

/**
获得token
*/
func getWxToken(appid string, secret string) (models.AccessTokenJson, error) {
	// https: //api.weixin.qq.com/cgi-bin/token?grant_type=client_credential&appid=APPID&secret=APPSECRET
	tokenobj := models.AccessTokenJson{}
	response_json := `{"errcode":1,"errmsg":"getWxAccessToken error"}`
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
		realm_name = "http://localhost:9090"
	} else {
		realm_name = "https://api.weixin.qq.com/cgi-bin/token"
	}
	realm_name = "https://api.weixin.qq.com/cgi-bin/token"
	wx_url = strings.Replace(wx_url, "[REALM]", realm_name, -1)
	wx_url = strings.Replace(wx_url, "[APPID]", appid, -1)
	wx_url = strings.Replace(wx_url, "[SECRET]", secret, -1)
	beego.Debug("wxqax getWxToken url :", wx_url)
	resp, err := http.Get(wx_url)
	if err != nil {
		beego.Error(err)
		return tokenobj, err
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		beego.Error(err)
		body = []byte(response_json)
	} else {
		beego.Debug("wxqax getWxToken boey :", string(body))
	}
	var atj models.AccessTokenJson
	if err := json.Unmarshal(body, &atj); err == nil {
		beego.Debug("get Token obj", atj)
		tokenobj = atj
	} else {
		beego.Error(err)
	}
	return tokenobj, err
}

/**
获得token授权
*/
func getWxTokenOauth(appid string, secret string, code string) (models.AccessTokenJson, error) {
	//https://api.weixin.qq.com/sns/oauth2/access_token?appid=APPID&secret=SECRET&code=CODE&grant_type=authorization_code
	tokenobj := models.AccessTokenJson{}
	response_json := `{"errcode":1,"errmsg":"getWxAccessToken error"}`
	wx_url := "[REALM]?appid=[APPID]&secret=[SECRET]&code=[CODE]&grant_type=authorization_code"
	realm_name := "https://api.weixin.qq.com/sns/oauth2/access_token"
	wx_url = strings.Replace(wx_url, "[REALM]", realm_name, -1)
	wx_url = strings.Replace(wx_url, "[APPID]", appid, -1)
	wx_url = strings.Replace(wx_url, "[SECRET]", secret, -1)
	wx_url = strings.Replace(wx_url, "[CODE]", code, -1)
	beego.Debug("getWxTokenFromCode url :", wx_url)
	resp, err := http.Get(wx_url)
	if err != nil {
		beego.Error(err)
		return tokenobj, err
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		beego.Error(err)
		body = []byte(response_json)
	} else {
		beego.Debug("getWxTokenFromCode body :", string(body))
	}
	var atj models.AccessTokenJson
	if err := json.Unmarshal(body, &atj); err == nil {
		beego.Debug("get Token obj", atj)
		tokenobj = atj
	} else {
		beego.Error(err)
	}
	return tokenobj, err
}

/**
获得用户信息授权
*/
func getWxUserOauth(openid string, access_token string) (models.WxOauthUser, error) {
	//https://api.weixin.qq.com/sns/userinfo?access_token=ACCESS_TOKEN&openid=OPENID&lang=zh_CN
	user := models.WxOauthUser{}
	response_json := `{"errcode":1,"errmsg":"getWxUser error"}`
	query_url := "[REALM]?access_token=[ACCESS_TOKEN]&openid=[OPENID]&lang=zh_CN"
	realm_name := "https://api.weixin.qq.com/sns/userinfo"
	query_url = strings.Replace(query_url, "[REALM]", realm_name, -1)
	query_url = strings.Replace(query_url, "[ACCESS_TOKEN]", access_token, -1)
	query_url = strings.Replace(query_url, "[OPENID]", openid, -1)
	beego.Debug("getWxUserOauth url:", query_url)

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
		beego.Debug("wxqax getWxUserOauth body", string(body))
	}
	var uij models.WxOauthUser
	if err := json.Unmarshal(body, &uij); err == nil {
		beego.Debug("wxqax getWxUserOauth json", uij)
		if uij.ErrCode == 0 {
			user = uij
		}

	} else {
		beego.Error(err)
	}
	return user, err
}

func getWxTicket(token string) (models.TicketJson, error) {
	// https://api.weixin.qq.com/cgi-bin/ticket/getticket?access_token=ACCESS_TOKEN&type=jsapi
	ticketobj := models.TicketJson{}
	response_json := `{"errcode":1,"errmsg":"getWxAccessToken error"}`
	wx_url := "[REALM]?access_token=[ACCESS_TOKEN]&type=jsapi"
	// if beego.AppConfig.Bool("qax580::isdebug") {
	realm_name := "https://api.weixin.qq.com/cgi-bin/ticket/getticket"
	wx_url = strings.Replace(wx_url, "[REALM]", realm_name, -1)
	wx_url = strings.Replace(wx_url, "[ACCESS_TOKEN]", token, -1)
	beego.Debug("wxqax getWxTicket url :", wx_url)
	resp, err := http.Get(wx_url)
	if err != nil {
		beego.Error(err)
		return ticketobj, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		beego.Error(err)
		body = []byte(response_json)
	} else {
		beego.Debug("wxqax getWxTicket boey :", string(body))
	}
	var atj models.TicketJson
	if err := json.Unmarshal(body, &atj); err == nil {
		beego.Debug("get getWxTicket obj", atj)
		ticketobj = atj
	} else {
		beego.Error(err)
	}
	return ticketobj, err
}

func getShare(appid string, secret string, share_url string) models.WxShare {
	ticket_cookie := ""
	tokenobj, err := getWxToken(appid, secret)
	if err != nil {
		beego.Error(err)
	}
	if tokenobj.ErrCode == 0 {
		ticket, err := getWxTicket(tokenobj.AccessToken)
		if err != nil {
			beego.Error(err)
		}
		if err != nil {
			beego.Error()
		}
		if ticket.ErrCode == 0 {
			ticket_cookie = ticket.Ticket
		}
	}

	wxShare := models.WxShare{}
	timestamp := time.Now().Unix()
	noncestr := getNonceStr(16, KC_RAND_KIND_ALL)
	basestr := "jsapi_ticket=[TICKET]&noncestr=[NONCESTR]&timestamp=[TIMESTAMP]&url=[URL]"
	basestr = strings.Replace(basestr, "[TICKET]", ticket_cookie, -1)
	basestr = strings.Replace(basestr, "[NONCESTR]", noncestr, -1)
	basestr = strings.Replace(basestr, "[TIMESTAMP]", fmt.Sprintf("%d", timestamp), -1)
	basestr = strings.Replace(basestr, "[URL]", share_url, -1)
	signaturestr := goWxJsSha1(basestr)
	beego.Debug(" getShare basestr", basestr)
	beego.Debug(" getShare ticket_cookie", ticket_cookie)
	beego.Debug(" getShare noncestr", noncestr)
	beego.Debug(" getShare timestamp", fmt.Sprintf("%d", timestamp))
	beego.Debug(" getShare signaturestr", signaturestr)
	beego.Debug(" getShare share_url", share_url)
	wxShare.AppId = appid
	wxShare.TimeStamp = timestamp
	wxShare.NonceStr = noncestr
	wxShare.Signature = signaturestr
	beego.Debug(" getShare WxShare", wxShare)
	return wxShare
}
