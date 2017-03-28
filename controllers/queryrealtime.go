package controllers

/*
实时查询
*/
import (
	"encoding/json"
	"github.com/astaxie/beego"
	"io/ioutil"
	"net/http"
	"strings"
)

type RealTime struct {
	Reason    string     `json:"reason"`
	ErrorCode int64      `json:"error_code"`
	RTResult  []RTResult `json:"result"`
}
type RTResult struct {
	TrainNo          string `json:"train_no"`
	StartStationName string `json:"start_station_name"`
	EndStationName   string `json:"end_station_name"`
	FromStationName  string `json:"from_station_name"`
	ToStationName    string `json:"to_station_name"`
	StartTime        string `json:"start_time"`
	ArriveTime       string `json:"arrive_time"`
	TrainClassName   string `json:"train_class_name"`
	DayDifference    string `json:"day_difference"`
	Lishi            string `json:"lishi"`
	GrNum            string `json:"gr_num"`
	QtNum            string `json:"qt_num"`
	RwNum            string `json:"rw_num"`
	RzNum            string `json:"rz_num"`
	TzNum            string `json:"tz_num"`
	WzNum            string `json:"wz_num"`
	YwNum            string `json:"yw_num"`
	YzNum            string `json:"yz_num"`
	ZeNum            string `json:"ze_num"`
	ZyNum            string `json:"zy_num"`
	SwzNum           string `json:"swz_num"`
}

type QueryRealTimeController struct {
	beego.Controller
}

func (c *QueryRealTimeController) Get() {
	c.Data["IsShow"] = "false"
	c.TplName = "queryrealtime.html"
}

func (c *QueryRealTimeController) Post() {
	start := c.Input().Get("start")
	end := c.Input().Get("end")
	data := c.Input().Get("data")
	if len(start) != 0 && len(end) != 0 && len(data) != 0 {
		queryRealTime(start, end, data, c)
	}
	c.Data["IsShow"] = "true"
	c.TplName = "queryrealtime.html"

}
func queryRealTime(start string, end string, date string, c *QueryRealTimeController) {
	query_url := "[REALM]?key=[KEY]&dtype=json&from=[FROM]&to=[TO]&date=[DATE]&tt="
	query_url = strings.Replace(query_url, "[REALM]", "http://apis.juhe.cn/train/yp", -1)
	query_url = strings.Replace(query_url, "[FROM]", start, -1)
	query_url = strings.Replace(query_url, "[TO]", end, -1)
	query_url = strings.Replace(query_url, "[DATE]", date, -1)
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

	var obj RealTime
	if err := json.Unmarshal(body, &obj); err == nil {
		beego.Debug("================json str 转struct==")
		beego.Debug(obj)
		beego.Debug("ErrorCode", obj.ErrorCode)
		c.Data["ErrorCode"] = obj.ErrorCode
		if obj.ErrorCode != 0 {
			c.Data["ErrorInfo"] = getError(obj.ErrorCode)
		} else {
			c.Data["Results"] = obj.RTResult
		}
	} else {
		beego.Debug(err)
	}
}

// func getRTError(errorcode int64) string {
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
