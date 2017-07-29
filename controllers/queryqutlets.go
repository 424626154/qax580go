package controllers

/*
代售点
*/
import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/astaxie/beego"
)

type Qutlets struct {
	Reason    string    `json:"reason"`
	ErrorCode int64     `json:"error_code"`
	QResult   []QResult `json:"result"`
}
type QResult struct {
	Province    string `json:"province"`
	City        string `json:"city"`
	County      string `json:"county"`
	AgencyName  string `json:"agency_name"`
	Address     string `json:"address"`
	PhoneNo     string `json:"phone_no"`
	StartTimeAm string `json:"start_time_am"`
	StopTimeAm  string `json:"stop_time_am"`
	StartTimePm string `json:"start_time_pm"`
	StopTimePm  string `json:"stop_time_pm"`
}

type QueryQutletsController struct {
	beego.Controller
}

func (c *QueryQutletsController) Get() {
	c.Data["IsShow"] = "false"
	c.TplName = "queryqutlets.html"
}

func (c *QueryQutletsController) Post() {
	province := c.Input().Get("province")
	city := c.Input().Get("city")
	county := c.Input().Get("county")
	c.Data["IsShow"] = "true"
	if len(province) != 0 && len(city) != 0 && len(county) != 0 {
		queryQutlets(province, city, county, c)
	}

	c.TplName = "queryqutlets.html"
}

func queryQutlets(province string, city string, county string, c *QueryQutletsController) {
	query_url := "[REALM]?province=[PROVINCE]&city=[CITY]&county=[COUNTY]&key=[KEY]"
	query_url = strings.Replace(query_url, "[REALM]", "http://apis.juhe.cn/train/dsd", -1)
	query_url = strings.Replace(query_url, "[PROVINCE]", province, -1)
	query_url = strings.Replace(query_url, "[CITY]", city, -1)
	query_url = strings.Replace(query_url, "[COUNTY]", county, -1)
	query_url = strings.Replace(query_url, "[KEY]", "3d074d491176f4e6d3309746460b1ef9", -1)
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

	var obj Qutlets
	if err := json.Unmarshal(body, &obj); err == nil {
		beego.Debug("================json str 转struct==")
		beego.Debug(obj)
		beego.Debug("ErrorCode", obj.ErrorCode)
		c.Data["ErrorCode"] = obj.ErrorCode
		if obj.ErrorCode != 0 {
			c.Data["ErrorInfo"] = GetError(obj.ErrorCode)
		} else {
			c.Data["QResults"] = obj.QResult
		}
	} else {
		beego.Debug(err)
	}
}

// func getQError(errorcode int64) string {
// 	error_indo := ""
// 	if errorcode == 202201 {
// 		error_indo = "车次不能为空"
// 	}
// 	if errorcode == 202202 {
// 		error_indo = "查询不到车次的相关信息"
// 	}
// 	if errorcode == 202203 {
// 		error_indo = "出发站或终点站不能为空"
// 	}
// 	if errorcode == 202204 {
// 		error_indo = "查询不到结果"
// 	}
// 	if errorcode == 202205 {
// 		error_indo = "错误的出发站名称"
// 	}
// 	if errorcode == 202206 {
// 		error_indo = "错误的到达站名称"
// 	}
// 	if errorcode == 202207 {
// 		error_indo = "查询不到余票相关数据哦"
// 	}
// 	if errorcode == 202208 {
// 		error_indo = "错误的请求，请确认传递的参数正确"
// 	}
// 	if errorcode == 202209 {
// 		error_indo = "请求12306网络错误,请重试"
// 	}
// 	if errorcode == 202210 {
// 		error_indo = "12306账号密码错误"
// 	}
// 	if errorcode == 202211 {
// 		error_indo = "邮箱不存在"
// 	}
// 	if errorcode == 202212 {
// 		error_indo = "查询出错"
// 	}
// 	if errorcode == 202213 {
// 		error_indo = "提交订单超时，请重试"
// 	}
// 	if errorcode == 202214 {
// 		error_indo = "出票失败"
// 	}
// 	if errorcode == 202215 {
// 		error_indo = "排队失败"
// 	}
// 	if errorcode == 202216 {
// 		error_indo = "该车次无法预定"
// 	}
// 	if errorcode == 202217 {
// 		error_indo = "不合法的座位类型"
// 	}
// 	return error_indo
// }
