package controllers

/*
微信公众号服务器
*/
import (
	"crypto/sha1"
	"encoding/hex"
	"encoding/xml"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/config"
	"io/ioutil"
	"net/http"
	"qax580go/models"
	"sort"
	"strings"
	"time"
)

const (
	token          = "qax580"
	qax580_name    = "qax580"
	error_info     = "我们正在努力成为一个有情怀的免费信息网站"
	null_info0     = "未搜索到相关信息"
	null_info1     = "未搜索到有关&的信息"
	subscribe_info = "欢迎关注咱这580，我们正在努力成为一个有情怀的免费信息发布平台，为大家服务"
	about_info     = "【客服服务】\n关注我们\n公众号:qax580\n微信:qax580kf\n腾讯微博:庆安兄弟微盟\nQQ : 2063883729\n邮箱：qaxiongdiweimeng@163.com"
	content_url    = "http://www.baoguangguang.cn/content?op=con&id=s%"
	jieshao_info   = "【帮助】\n你好，咱这580是免费的信息发布平台，在这里您可以发布信息也可以搜索相关信息，相关功能在功能菜单中"
	function_info  = "【帮助】\n发布信息－》更多－》发布信息，意见反馈－》更多－》意见反馈"
)

//接收文本消息
type TextRequestBody struct {
	XMLName      xml.Name `xml:"xml"`
	ToUserName   string
	FromUserName string
	CreateTime   time.Duration
	MsgType      string
	Content      string
	MsgId        int
}

//接收音频消息
type VoiceRequestBody struct {
	XMLName      xml.Name `xml:"xml"`
	ToUserName   string
	FromUserName string
	CreateTime   time.Duration
	MsgType      string
	MediaId      string
	Format       string
	Recognition  string
	MsgId        int
}

//点击消息
type EventResponseBody struct {
	XMLName      xml.Name `xml:"xml"`
	ToUserName   string
	FromUserName string
	CreateTime   time.Duration
	MsgType      string
	Event        string
	EventKey     string
}

//消息类型
type TypeResponseBody struct {
	XMLName xml.Name `xml:"xml"`
	MsgType string
}
type CDATAText struct {
	Text string `xml:",innerxml"`
}

//文本消息
type TextResponseBody struct {
	XMLName      xml.Name `xml:"xml"`
	ToUserName   CDATAText
	FromUserName CDATAText
	CreateTime   time.Duration
	MsgType      CDATAText
	Content      CDATAText
}

func value2CDATA(v string) CDATAText {
	//return CDATAText{[]byte("<![CDATA[" + v + "]]>")}
	return CDATAText{"<![CDATA[" + v + "]]>"}
}

//图文消息
type ImageTextResponseBody struct {
	XMLName      xml.Name `xml:"xml"`
	ToUserName   CDATAText
	FromUserName CDATAText
	CreateTime   time.Duration
	MsgType      CDATAText
	ArticleCount int
	Articles     []ImageTextResponseItem `xml:"Articles>item"`
}

//图文消息元素
type ImageTextResponseItem struct {
	XMLName     xml.Name `xml:"item"`
	Title       CDATAText
	Description CDATAText
	PicUrl      CDATAText
	Url         CDATAText
}

//多客服
type CustomServiceResponseBody struct {
	XMLName      xml.Name `xml:"xml"`
	ToUserName   CDATAText
	FromUserName CDATAText
	CreateTime   time.Duration
	MsgType      CDATAText
}
type WXController struct {
	beego.Controller
}

func (c *WXController) Get() {
	verification(c)
}
func (c *WXController) Post() {
	responseMsg(c)
}

//验证签名
func verification(c *WXController) {
	beego.Debug("valid ")
	echoStr := c.Input().Get("echostr")     //随机字符串
	signature := c.Input().Get("signature") //微信加密签名，signature结合了开发者填写的token参数和请求中的timestamp参数、nonce参数。
	timestamp := c.Input().Get("timestamp") //时间戳
	nonce := c.Input().Get("nonce")         //随机数
	token := "qax580"
	tmpArr := []string{token, timestamp, nonce}
	sort.Strings(tmpArr)
	tmpStr := ""
	for i := 0; i < len(tmpArr); i++ {
		tmpStr += tmpArr[i]
	}
	tmpStrSha1 := goSha1(tmpStr)
	respnse := ""
	if strings.EqualFold(tmpStrSha1, signature) {
		respnse = echoStr
	}
	beego.Debug(respnse)
	c.Ctx.WriteString(respnse)
	return
}

//对字符串进行SHA1哈希
func goSha1(str string) string {
	s := sha1.New()
	s.Write([]byte(str))
	return hex.EncodeToString(s.Sum(nil))
}

//响应消息
func responseMsg(c *WXController) {
	body := c.Ctx.Input.RequestBody
	requestType := &TypeResponseBody{}
	err := xml.Unmarshal(body, requestType)
	response_xml := ""
	if err != nil {
		beego.Debug(err.Error())
	} else {
		beego.Debug(requestType.MsgType)
		response_xml = responseTypeMsg(body, requestType.MsgType)
	}
	beego.Debug(response_xml)
	c.Ctx.WriteString(response_xml)
	return
}

//根据类型解析消息
func responseTypeMsg(body []byte, msgType string) string {
	response_xml := ""
	switch msgType {
	//文字
	case "text":
		requestBody := &TextRequestBody{}
		err := xml.Unmarshal(body, requestBody)
		if err != nil {
			beego.Debug(err.Error())
			response_xml = responseTextMsg(requestBody.FromUserName, error_info, qax580_name)
		} else {
			requestBody := &TextRequestBody{}
			err := xml.Unmarshal(body, requestBody)
			if err != nil {
				response_xml = responseTextMsg(requestBody.FromUserName, error_info, qax580_name)
			} else if requestBody.Content == "你好" || requestBody.Content == "您好" {
				response_xml = responseTextMsg(requestBody.FromUserName, jieshao_info, qax580_name)
			} else if strings.Index(requestBody.Content, "发布") >= 0 {
				response_xml = responseTextMsg(requestBody.FromUserName, function_info, qax580_name)
			} else {
				//是否存在关键字
				key_count := int32(0)
				count, err := models.GetKeywordsCount(requestBody.Content)
				if err != nil {
					beego.Error(err)
				} else {
					key_count = count
				}
				if key_count > 0 {
					obj, err := models.GetOneKeywords(requestBody.Content)
					if err != nil {
						beego.Error(err)
					} else {
						objs, err := models.QueryFuzzyLimitKeyobj(obj.Id, 5)
						if err != nil {
							beego.Error(err)
						}
						if len(objs) > 0 {
							response_xml = responseKeyXML(requestBody.FromUserName, requestBody.Content, objs, qax580_name)
						} else {
							response_xml = responseCustomerService(requestBody.FromUserName, requestBody.ToUserName)
						}
					}
				} else {
					//信息查询
					beego.Debug(requestBody.Content)
					posts, err := models.QueryFuzzyLimitPost(requestBody.Content, 5)
					if err != nil {
						beego.Error(err)
					}
					// beego.Debug(requestBody.FromUserName)
					// beego.Debug(requestBody.ToUserName)
					if len(posts) > 0 {
						response_xml = responseImageTextXML(requestBody.FromUserName, requestBody.Content, posts, qax580_name)
					} else {
						response_xml = responseCustomerService(requestBody.FromUserName, requestBody.ToUserName)
					}
				}
			}
		}
		//音频
	case "voice":
		requestBody := &VoiceRequestBody{}
		err := xml.Unmarshal(body, requestBody)
		if err != nil {
			response_xml = responseTextMsg(requestBody.FromUserName, error_info, qax580_name)
		} else {
			beego.Debug(requestBody.Recognition)
			posts, err := models.QueryFuzzyLimitPost(requestBody.Recognition, 5)
			if err != nil {
				beego.Error(err)
			}
			response_xml = responseImageTextXML(requestBody.FromUserName, requestBody.Recognition, posts, qax580_name)

		}
		//点击
	case "event":
		requestBody := &EventResponseBody{}
		err := xml.Unmarshal(body, requestBody)
		if err != nil {
			beego.Debug(err.Error())
			response_xml = responseTextMsg(requestBody.FromUserName, error_info, qax580_name)
		} else {
			beego.Debug("qax 580 Event:", requestBody.Event, "EventKey", requestBody.EventKey)
			//自定义点击事件
			if requestBody.Event == "CLICK" {
				//推荐
				if requestBody.EventKey == "recommend" {
					// posts, err := models.QueryLimitPost(5)
					// if err != nil {
					// 	beego.Error(err)
					// }
					// response_xml = responseImageTextXML(requestBody.FromUserName, "", posts)
					index, err := models.GetQueryIndex(requestBody.FromUserName)
					if err != nil {
						beego.Error(err)
					}
					posts, err := models.QueryPagePost(index, 5)
					if err != nil {
						beego.Error(err)
					}
					beego.Debug("recommend count :", len(posts))
					if len(posts) != 0 {
						response_xml = responseImageTextXML(requestBody.FromUserName, "", posts, qax580_name)
					} else {
						response_xml = responseTextMsg(requestBody.FromUserName, "今日已无更多推荐信息", qax580_name)
					}
					//关于
				} else if requestBody.EventKey == "about" {
					response_xml = responseAbout(requestBody.FromUserName, qax580_name, about_info)
				} else if requestBody.EventKey == "today" { //今日580
					response_xml = responseToday(requestBody.FromUserName, qax580_name)
				} else {

				}
				//关注
			} else if requestBody.Event == "subscribe" {

				response_xml = responseTextMsg(requestBody.FromUserName, subscribe_info, qax580_name)
				beego.Debug("subscribe qax580")
				eventSubscribe(requestBody, qa_appid, qa_secret)
			} else {
				//其他类型
				response_xml = responseTextMsg(requestBody.FromUserName, error_info, qax580_name)
			}

		}

	default:
		beego.Debug(msgType)
	}
	beego.Debug(response_xml)
	return response_xml
}

func analysisNull(toUserName string, content string, default_null_info string) string {
	null_info := default_null_info
	if len(content) > 0 {
		null_info = strings.Replace(null_info1, "&", content, -1)
	}
	textResponseBody := &TextResponseBody{}
	textResponseBody.FromUserName = value2CDATA(qax580_name)
	textResponseBody.ToUserName = value2CDATA(toUserName)
	textResponseBody.MsgType = value2CDATA("text")
	textResponseBody.Content = value2CDATA(null_info)
	textResponseBody.CreateTime = time.Duration(time.Now().Unix())
	error_info, _ := xml.MarshalIndent(textResponseBody, " ", "  ")
	return string(error_info)
}

//返回文本信息
//textMsg 文本信息
func responseTextMsg(toUserName string, textMsg string, from_name string) string {
	textResponseBody := &TextResponseBody{}
	textResponseBody.FromUserName = value2CDATA(from_name)
	textResponseBody.ToUserName = value2CDATA(toUserName)
	textResponseBody.MsgType = value2CDATA("text")
	textResponseBody.Content = value2CDATA(textMsg)
	textResponseBody.CreateTime = time.Duration(time.Now().Unix())
	about_xml, _ := xml.MarshalIndent(textResponseBody, " ", "  ")
	return string(about_xml)
}

func responseAbout(toUserName string, from_name string, about_info string) string {
	textResponseBody := &TextResponseBody{}
	textResponseBody.FromUserName = value2CDATA(from_name)
	textResponseBody.ToUserName = value2CDATA(toUserName)
	textResponseBody.MsgType = value2CDATA("text")
	textResponseBody.Content = value2CDATA(about_info)
	textResponseBody.CreateTime = time.Duration(time.Now().Unix())
	about_xml, _ := xml.MarshalIndent(textResponseBody, " ", "  ")
	return string(about_xml)
}

func responseImageTextXML(toUserName string, content string, posts []models.Post, from_name string) string {
	articles := ""
	if posts != nil && len(posts) > 0 {
		imageTextResponseItems := []ImageTextResponseItem{}
		for i := 0; i < len(posts); i++ {
			imageTextResponseItem := ImageTextResponseItem{}
			imageTextResponseItem.Title = value2CDATA(posts[i].Title)
			imageTextResponseItem.Description = value2CDATA(posts[i].Info)
			imageTextResponseItem.PicUrl = value2CDATA(getWxImageUrl(posts[i].Image))
			imageTextResponseItem.Url = value2CDATA(strings.Replace(content_url, "s%", fmt.Sprintf("%d", posts[i].Id), -1))
			imageTextResponseItems = append(imageTextResponseItems, imageTextResponseItem)
		}
		textResponseBody := &ImageTextResponseBody{}
		textResponseBody.MsgType = value2CDATA("news")
		textResponseBody.ArticleCount = len(imageTextResponseItems)
		textResponseBody.Articles = imageTextResponseItems
		textResponseBody.FromUserName = value2CDATA(from_name)
		textResponseBody.ToUserName = value2CDATA(toUserName)
		res, err := xml.MarshalIndent(textResponseBody, " ", "  ")

		if err != nil {
			beego.Debug(err.Error())
		} else {
			articles = string(res)
		}
	} else {
		articles = analysisNull(toUserName, content, null_info0)
	}

	return articles
}

//关键字
func responseKeyXML(toUserName string, content string, objs []models.Keyobj, from_name string) string {
	articles := ""
	if objs != nil && len(objs) > 0 {
		imageTextResponseItems := []ImageTextResponseItem{}
		for i := 0; i < len(objs); i++ {
			imageTextResponseItem := ImageTextResponseItem{}
			imageTextResponseItem.Title = value2CDATA(objs[i].Title)
			imageTextResponseItem.Description = value2CDATA(objs[i].Info)
			imageTextResponseItem.PicUrl = value2CDATA(getWxImageUrl(objs[i].Image))
			imageTextResponseItem.Url = value2CDATA(objs[i].Url)
			imageTextResponseItems = append(imageTextResponseItems, imageTextResponseItem)
		}
		textResponseBody := &ImageTextResponseBody{}
		textResponseBody.MsgType = value2CDATA("news")
		textResponseBody.ArticleCount = len(imageTextResponseItems)
		textResponseBody.Articles = imageTextResponseItems
		textResponseBody.FromUserName = value2CDATA(from_name)
		textResponseBody.ToUserName = value2CDATA(toUserName)
		res, err := xml.MarshalIndent(textResponseBody, " ", "  ")

		if err != nil {
			beego.Debug(err.Error())
		} else {
			articles = string(res)
		}
	} else {
		articles = analysisNull(toUserName, content, null_info0)
	}

	return articles
}

func responseToday(toUserName string, from_name string) string {
	articles := ""
	imageTextResponseItems := []ImageTextResponseItem{}
	//添加历史今天
	imageTextResponseItem := ImageTextResponseItem{}
	imageTextResponseItem.Title = value2CDATA("历史今天")
	imageTextResponseItem.Description = value2CDATA("回顾历史的长河，历史是生活的一面镜子")
	imageTextResponseItem.PicUrl = value2CDATA("http://182.92.167.29:8080/static/img/lishi.png")
	imageTextResponseItem.Url = value2CDATA("http://www.baoguangguang.cn/history")
	imageTextResponseItems = append(imageTextResponseItems, imageTextResponseItem)

	//添加老黄历
	imageTextResponseItem.Title = value2CDATA("老黄历")
	imageTextResponseItem.Description = value2CDATA("每日吉凶宜忌")
	imageTextResponseItem.PicUrl = value2CDATA("http://182.92.167.29:8080/static/img/laohuangli.png")
	imageTextResponseItem.Url = value2CDATA("http://www.baoguangguang.cn/laohuangli")
	imageTextResponseItems = append(imageTextResponseItems, imageTextResponseItem)

	//周边Wi-Fi
	imageTextResponseItem.Title = value2CDATA("周边Wi-Fi")
	imageTextResponseItem.Description = value2CDATA("周边免费的WIFI热点分布")
	imageTextResponseItem.PicUrl = value2CDATA("http://182.92.167.29:8080/static/img/wifi.png")
	imageTextResponseItem.Url = value2CDATA("http://www.baoguangguang.cn/zhoubianwifiwx")
	imageTextResponseItems = append(imageTextResponseItems, imageTextResponseItem)

	//天气预报
	imageTextResponseItem.Title = value2CDATA("天气预报")
	imageTextResponseItem.Description = value2CDATA("阴晴冷暖早知道")
	imageTextResponseItem.PicUrl = value2CDATA("http://182.92.167.29:8080/static/img/tianqi.png")
	imageTextResponseItem.Url = value2CDATA("http://www.baoguangguang.cn/tianqiwx")
	imageTextResponseItems = append(imageTextResponseItems, imageTextResponseItem)

	textResponseBody := &ImageTextResponseBody{}
	textResponseBody.MsgType = value2CDATA("news")
	textResponseBody.ArticleCount = len(imageTextResponseItems)
	textResponseBody.Articles = imageTextResponseItems
	textResponseBody.FromUserName = value2CDATA(from_name)
	textResponseBody.ToUserName = value2CDATA(toUserName)
	res, err := xml.MarshalIndent(textResponseBody, " ", "  ")

	if err != nil {
		beego.Debug(err.Error())
	} else {
		articles = string(res)
	}

	return articles
}

//多客服消息
func responseCustomerService(fromUserName string, toUserName string) string {
	body := &CustomServiceResponseBody{}
	body.FromUserName = value2CDATA(toUserName)
	body.ToUserName = value2CDATA(fromUserName)
	body.CreateTime = time.Duration(time.Now().Unix())
	body.MsgType = value2CDATA("transfer_customer_service")
	about_xml, _ := xml.MarshalIndent(body, " ", "  ")
	return string(about_xml)
}
func getWxImageUrl(url string) string {
	new_url := "http://182.92.167.29:8080/static/img/type0.jpg"
	if len(url) != 0 {
		new_url = fmt.Sprintf("%s%s", "http://182.92.167.29:8080/imagehosting/", url)
	}
	return new_url
}

/*
关注事件
*/
func eventSubscribe(requestBody *EventResponseBody, appid string, secret string) string {
	if strings.Contains(requestBody.EventKey, "qrscene_") { //扫码

	} else {

	}
	//wxqax/sunscribe
	//http请求参数 subscribe_type from_openid to_user create_time
	response_json := `{"errcode":1,"errmsg":"getWxAccessToken error"}`
	isdebug := "true"
	isurl := ""
	iniconf, err := config.NewConfig("json", "conf/myconfig.json")
	if err != nil {
		beego.Debug(err)
	} else {
		isdebug = iniconf.String("qax580::isdebug")
		isurl = iniconf.String("qax580::url")
	}
	wx_url := "[REALM]/wxqax/sunscribe?subscribe_type=[SUNSCRIBE]&from_openid=[FROMOPENID]&to_user=[TOUSER]&create_time=[CREATETIME]&appid=[APPID]&secret=[SECRET]"
	// if beego.AppConfig.Bool("qax580::isdebug") {
	realm_name := ""
	if isdebug == "true" {
		realm_name = "http://localhost:9090"
	} else {
		realm_name = isurl
	}
	wx_url = strings.Replace(wx_url, "[REALM]", realm_name, -1)
	wx_url = strings.Replace(wx_url, "[SUNSCRIBE]", requestBody.Event, -1)
	wx_url = strings.Replace(wx_url, "[FROMOPENID]", requestBody.FromUserName, -1)
	wx_url = strings.Replace(wx_url, "[TOUSER]", requestBody.ToUserName, -1)
	wx_url = strings.Replace(wx_url, "[APPID]", appid, -1)
	wx_url = strings.Replace(wx_url, "[SECRET]", secret, -1)
	create_time := fmt.Sprintf("%d", requestBody.CreateTime)
	wx_url = strings.Replace(wx_url, "[CREATETIME]", create_time, -1)
	beego.Debug("subscribe url", wx_url)
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
	response_json = string(body)
	return response_json
}
