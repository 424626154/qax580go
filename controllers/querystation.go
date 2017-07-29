package controllers

/*
起点终点查询
*/
import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/astaxie/beego"
)

type Station struct {
	Reason    string `json:"reason"`
	ErrorCode int64  `json:"error_code"`
	Result    Result `json:"result"`
}

type Result struct {
	Data []Data `json:"list"`
}

type Data struct {
	TrainNo          string  `json:"train_no"`           //车次
	TrainType        string  `json:"train_type"`         //类型
	StartStaion      string  `json:"start_station"`      //发车站
	StartStationType string  `json:"start_station_type"` //发车类型
	EndStation       string  `json:"end_station"`        //停止站
	EndStationType   string  `json:"end_station_type"`   //停车类型
	StartTime        string  `json:"start_time"`         //发车时间
	EndTime          string  `json:"end_time"`           //停车时间
	RunTime          string  `json:"run_time"`           //行驶时间
	RunDistance      string  `json:"run_distance"`       //里程
	Price            []Price `json:"price_list"`
}

type Price struct {
	PriceType string `json:"price_type"`
	Price     string `json:"price"`
}
type QueryStationController struct {
	beego.Controller
}

func (c *QueryStationController) Get() {
	c.Data["IsShow"] = "false"
	c.TplName = "querystation.html"
}

func (c *QueryStationController) Post() {
	start := c.Input().Get("start")
	end := c.Input().Get("end")

	c.Data["IsShow"] = "true"
	if len(start) != 0 && len(end) != 0 {
		queryStation(start, end, c)
	}

	c.TplName = "querystation.html"
}

func queryStation(start string, end string, c *QueryStationController) {
	query_url := "[REALM]?start=[START]&end=[END]&key=[KEY]"
	query_url = strings.Replace(query_url, "[REALM]", "http://apis.juhe.cn/train/s2swithprice", -1)
	query_url = strings.Replace(query_url, "[START]", start, -1)
	query_url = strings.Replace(query_url, "[END]", end, -1)
	query_url = strings.Replace(query_url, "[KEY]", "3d074d491176f4e6d3309746460b1ef9", -1)
	beego.Debug("signature_str:", query_url)
	resp, err := http.Get(query_url)
	if err != nil {
		beego.Debug(err)
	}

	defer resp.Body.Close()
	beego.Debug("----------------queryStation body--------------------")
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		beego.Debug(err)
	} else {
		// beego.Debug(string(body))
	}

	var station Station
	if err := json.Unmarshal(body, &station); err == nil {
		beego.Debug("================json str 转struct==")
		// beego.Debug(station)
		beego.Debug("ErrorCode", station.ErrorCode)
		c.Data["ErrorCode"] = station.ErrorCode
		if station.ErrorCode != 0 {
			c.Data["ErrorInfo"] = GetError(station.ErrorCode)
		} else {
			c.Data["Datas"] = station.Result.Data
			beego.Debug(station.Result.Data)
		}
	} else {
		beego.Debug(err)
	}
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
