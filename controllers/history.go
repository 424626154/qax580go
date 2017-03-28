package controllers

/*
历史上的今天
*/
import (
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

type History struct {
	Reason    string    `json:"reason"`
	ErrorCode int64     `json:"error_code"`
	HResult   []HResult `json:"result"`
}

type HResult struct {
	Day   int32  `json:"day"`
	Des   string `json:"des"`
	Id    string `json:"_id"`
	Lunar string `json:"lunar"`
	Month int32  `json:"month"`
	Pic   string `json:"pic"`
	Title string `json:"title"`
	Year  int32  `json:"year"`
}

type HistoryController struct {
	beego.Controller
}

func (c *HistoryController) Get() {
	queryHistory(c)
	c.TplName = "history.html"
}

func (c *HistoryController) Post() {
	c.TplName = "history.html"
}

func queryHistory(c *HistoryController) {
	now := time.Now()
	_, mon, day := now.Date()
	mon_str := fmt.Sprintf("%d", mon)
	day_str := fmt.Sprintf("%d", day)
	query_url := "[REALM]?&v=1.0&month=[MONTH]&day=[DAY]&key=[KEY]"
	query_url = strings.Replace(query_url, "[REALM]", "http://api.juheapi.com/japi/toh", -1)
	query_url = strings.Replace(query_url, "[MONTH]", mon_str, -1)
	query_url = strings.Replace(query_url, "[DAY]", day_str, -1)
	query_url = strings.Replace(query_url, "[KEY]", "434d026ae8f9d813163bb31f470ceca7", -1)
	beego.Debug("query_url:", query_url)
	resp, err := http.Get(query_url)
	if err != nil {
		beego.Debug(err)
	}

	defer resp.Body.Close()
	beego.Debug("----------------queryHistory body--------------------")
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		beego.Debug(err)
	} else {
		beego.Debug(string(body))
	}

	var obj History
	if err := json.Unmarshal(body, &obj); err == nil {
		beego.Debug("================json str 转struct==")
		// beego.Debug(obj)
		beego.Debug("ErrorCode", obj.ErrorCode)
		c.Data["ErrorCode"] = obj.ErrorCode
		if obj.ErrorCode != 0 {
			c.Data["ErrorInfo"] = getLishiError(obj.ErrorCode)
		} else {
			c.Data["HResults"] = obj.HResult
		}
	} else {
		beego.Debug(err)
	}
}

func getLishiError(errorcode int64) string {
	error_info := ""
	if errorcode == 10019 {
		error_info = "错误的版本号"
	}
	if errorcode == 206301 {
		error_info = "错误的请求参数"
	}
	if errorcode == 206302 {
		error_info = "无相关数据"
	}
	return error_info
}
