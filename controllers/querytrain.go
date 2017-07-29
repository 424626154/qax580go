package controllers

/*
车次查询
*/
import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/astaxie/beego"
)

type Tation struct {
	Reason    string  `json:"reason"`
	ErrorCode int64   `json:"error_code"`
	Result1   Result1 `json:"result"`
}

type Result1 struct {
	TrainInfo   TrainInfo     `json:"train_info"`
	StationList []StationList `json:"station_list"`
}

type TrainInfo struct {
	Name      string `json:"name"`
	Start     string `json:"start"`
	End       string `json:"end"`
	Starttime string `json:"starttime"`
	Endtime   string `json:"endtime"`
	Mileage   string `json:"mileage"`
}

type StationList struct {
	TrainId     string `json:"train_id"`
	StationName string `json:"station_name"`
	ArrivedTime string `json:"arrived_time"`
	LeaveTime   string `json:"leave_time"`
	Stay        string `json:"stay"`
	Mileage     string `json:"mileage"`
	SsoftSeat   string `json:"ssoftSeat"`
	FsoftSeat   string `json:"fsoftSeat"`
	HardSead    string `json:"hardSead"`
	SoftSeat    string `json:"softSeat"`
	HardSleep   string `json:"hardSleep"`
	SoftSleep   string `json:"softSleep"`
	Wuzuo       string `json:"wuzuo"`
	Swz         string `json:"swz"`
	Tdz         string `json:"tdz"`
	Gjrw        string `json:"gjrw"`
}

type QueryTrainController struct {
	beego.Controller
}

func (c *QueryTrainController) Get() {
	c.Data["IsShow"] = "false"
	c.TplName = "querytrain.html"
}

func (c *QueryTrainController) Post() {
	c.Data["IsShow"] = "true"
	name := c.Input().Get("name")
	if len(name) != 0 {
		queryTrain(name, c)
	}
	c.TplName = "querytrain.html"

}
func queryTrain(name string, c *QueryTrainController) {
	url := "[REALM]?name=[NAME]&key=[KEY]"
	url = strings.Replace(url, "[REALM]", "http://apis.juhe.cn/train/s", -1)
	url = strings.Replace(url, "[NAME]", name, -1)
	url = strings.Replace(url, "[KEY]", "3d074d491176f4e6d3309746460b1ef9", -1)
	beego.Debug("url:", url)
	resp, err := http.Get(url)
	if err != nil {
		beego.Debug(err)
	}

	defer resp.Body.Close()
	beego.Debug("----------------queryTrain body--------------------")
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		beego.Debug(err)
	} else {
		// beego.Debug(string(body))
	}

	var tation Tation
	if err := json.Unmarshal(body, &tation); err == nil {
		beego.Debug("================json str 转struct==")
		// beego.Debug(tation)
		beego.Debug("ErrorCode", tation.ErrorCode)
		c.Data["ErrorCode"] = tation.ErrorCode
		if tation.ErrorCode != 0 {
			c.Data["ErrorInfo"] = GetError(tation.ErrorCode)
			beego.Debug("Errorinfo", GetError(tation.ErrorCode))
		} else {
			c.Data["TrainInfo"] = tation.Result1.TrainInfo
			c.Data["StationList"] = tation.Result1.StationList
			beego.Debug("TrainInfo", tation.Result1.TrainInfo)
			beego.Debug("StationList", tation.Result1.StationList)
		}
	} else {
		beego.Debug(err)
	}

}

// func getError1(errorcode int64) string {
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
