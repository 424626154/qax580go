package controllers

/*
违章查询
*/
import (
	"encoding/json"
	"github.com/astaxie/beego"
	"io/ioutil"
	"net/http"
	"strings"
)

type Peccancy struct {
	Reason     string  `json:"reason"`
	ErrorCode  int64   `json:"error_code"`
	Resultcode string  `json:"resultcode"`
	PResult    PResult `json:"result"`
}

type PResult struct {
	Province  string      `json:"province"`
	City      string      `json:"city"`
	Hphm      string      `json:"hphm"`
	Hpzl      string      `json:"hpzl"`
	Peccancys []Peccancys `json:"lists"`
}
type Peccancys struct {
	Date    string `json:"date"`
	Area    string `json:"area"`
	Act     string `json:"act"`
	Code    string `json:"code"`
	Fen     string `json:"fen"`
	Money   string `json:"money"`
	Handled string `json:"handled"`
}

type QueryPeccancyController struct {
	beego.Controller
}

func (c *QueryPeccancyController) Get() {
	c.TplName = "querypeccancy.html"
	c.Data["IsShow"] = "false"
}

func (c *QueryPeccancyController) Post() {
	city := c.Input().Get("city")
	hphm := c.Input().Get("hphm")
	engineno := c.Input().Get("engineno")

	c.Data["IsShow"] = "true"
	if len(city) != 0 && len(hphm) != 0 && len(engineno) != 0 {
		queryPeccancy(city, hphm, engineno, c)
	}

	c.TplName = "querypeccancy.html"
}

func queryPeccancy(city string, hphm string, engineno string, c *QueryPeccancyController) {
	query_url := "[REALM]?city=[CITY]&hphm=[HPHM]&engineno=[ENGINENO]&key=[KEY]"
	query_url = strings.Replace(query_url, "[REALM]", "http://v.juhe.cn/wz/query", -1)
	query_url = strings.Replace(query_url, "[CITY]", city, -1)
	query_url = strings.Replace(query_url, "[HPHM]", hphm, -1)
	query_url = strings.Replace(query_url, "[ENGINENO]", engineno, -1)
	query_url = strings.Replace(query_url, "[KEY]", "6bd5c3ffa52bca14090b62833e5bfc05", -1)
	beego.Debug("signature_str:", query_url)
	resp, err := http.Get(query_url)
	if err != nil {
		beego.Debug(err)
	}

	defer resp.Body.Close()
	beego.Debug("----------------queryQutlets body--------------------")
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		beego.Debug(err)
	} else {
		beego.Debug(string(body))
	}
	// str := `{"resultcode":"101","reason":"error key.","result":null,"error_code":203603}`
	// str := `{
	// "resultcode":"200",
	// "reason":"查询成功",
	// "result":{
	// 	"province":"HB",
	// 	"city":"HB_HD",
	// 	"hphm":"冀DHL327",
	// 	"hpzl":"02",
	// 	"lists":[
	// 		{
	// 		"date":"2013-12-29 11:57:29",
	// 		"area":"316省道53KM+200M",
	// 		"act":"16362 : 驾驶中型以上载客载货汽车、校车、危险物品运输车辆以外的其他机动车在高速公路以外的道路上行驶超过规定时速20%以上未达50%的",
	// 		"code":"",
	// 		"fen":"6",
	// 		"money":"100",
	// 		"handled":"1"
	// 		}
	// 	]
	// }
	// }`
	// body := []byte(str)
	var obj Peccancy
	if err := json.Unmarshal(body, &obj); err == nil {
		beego.Debug("================json str 转struct==")
		beego.Debug(obj)
		beego.Debug("ErrorCode", obj.ErrorCode)
		c.Data["ErrorCode"] = obj.ErrorCode
		if obj.ErrorCode != 0 {
			c.Data["ErrorInfo"] = getPError(obj.ErrorCode)
		} else {
			c.Data["Hphm"] = obj.PResult.Hphm
			c.Data["Peccancys"] = obj.PResult.Peccancys
			Tips := ""
			if len(obj.PResult.Peccancys) > 0 {
				Tips = "以下为您的违章记录"
			} else {
				Tips = "暂无违章记录"
			}
			c.Data["Tips"] = Tips
		}
	} else {
		beego.Debug(err)
	}
}

func getPError(errorcode int64) string {
	error_indo := ""
	if errorcode == 203602 {
		error_indo = "车辆信息不存在"
	}
	if errorcode == 203603 {
		error_indo = "网络错误请重试"
	}
	if errorcode == 203604 {
		error_indo = "传递参数的格式不正确"
	}
	if errorcode == 203605 {
		error_indo = "没找到此城市代码或该城市正在维护"
	}
	if errorcode == 203606 {
		error_indo = "车辆信息错误,请确认输入的信息正确"
	}
	if errorcode == 203607 {
		error_indo = "交管局网络原因暂时无法查询"
	}
	if errorcode == 203608 {
		error_indo = "您好,你所查询的城市正在维护或未开通查询"
	}
	return error_indo
}
