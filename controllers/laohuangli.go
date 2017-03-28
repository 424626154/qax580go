package controllers

/*
老黄历
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

type Rili struct {
	Reason    string   `json:"reason"`
	ErrorCode int64    `json:"error_code"`
	RLResult  RLResult `json:"result"`
}
type ShiChen struct {
	Reason    string     `json:"reason"`
	ErrorCode int64      `json:"error_code"`
	SCResult  []SCResult `json:"result"`
}

type RLResult struct {
	Id        string `json:"id"`
	Yangli    string `json:"yangli"`
	Yinli     string `json:"yinli"`
	Wuxing    string `json:"wuxing"`
	Chongsha  string `json:"chongsha"`
	Baiji     string `json:"baiji"`
	Jishen    string `json:"jishen"`
	Yi        string `json:"yi"`
	Xiongshen string `json:"xiongshen"`
	Ji        string `json:"ji"`
}

type SCResult struct {
	Yangli string `json:"yangli"`
	Hours  string `json:"hours"`
	Des    string `json:"des"`
	Yi     string `json:"yi"`
	Ji     string `json:"ji"`
}

type LaohuangliController struct {
	beego.Controller
}

func (c *LaohuangliController) Get() {
	queryLaohuangliRili(c)
	queryLaohuangliShichen(c)
	c.TplName = "laohuangli.html"
}

func (c *LaohuangliController) Post() {
	c.TplName = "laohuangli.html"
}

//老黄历日历
func queryLaohuangliRili(c *LaohuangliController) {
	now := time.Now()
	year, mon, day := now.Date()
	date_str := fmt.Sprintf("%d-%d-%d", year, mon, day)
	query_url := "[REALM]?date=[DATE]&key=[KEY]"
	query_url = strings.Replace(query_url, "[REALM]", "http://v.juhe.cn/laohuangli/d", -1)
	query_url = strings.Replace(query_url, "[DATE]", date_str, -1)
	query_url = strings.Replace(query_url, "[KEY]", "e463fa7d86b928766354c1d4d9cb906a", -1)
	beego.Debug("query_url:", query_url)
	resp, err := http.Get(query_url)
	if err != nil {
		beego.Debug(err)
	}

	defer resp.Body.Close()
	beego.Debug("----------------queryLaohuangliRili body--------------------")
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		beego.Debug(err)
	} else {
		beego.Debug(string(body))
	}

	var obj Rili
	if err := json.Unmarshal(body, &obj); err == nil {
		beego.Debug("================json str 转struct==")
		// beego.Debug(obj)
		beego.Debug("RLErrorCode", obj.ErrorCode)
		c.Data["RLErrorCode"] = obj.ErrorCode
		if obj.ErrorCode != 0 {
			c.Data["RLErrorInfo"] = getLaouangliError(obj.ErrorCode)
		} else {
			c.Data["RLResult"] = obj.RLResult
		}
	} else {
		beego.Debug(err)
	}
}

//老黄历时辰
func queryLaohuangliShichen(c *LaohuangliController) {
	now := time.Now()
	year, mon, day := now.Date()
	date_str := fmt.Sprintf("%d-%d-%d", year, mon, day)
	query_url := "[REALM]?date=[DATE]&key=[KEY]"
	query_url = strings.Replace(query_url, "[REALM]", "http://v.juhe.cn/laohuangli/h", -1)
	query_url = strings.Replace(query_url, "[DATE]", date_str, -1)
	query_url = strings.Replace(query_url, "[KEY]", "e463fa7d86b928766354c1d4d9cb906a", -1)
	beego.Debug("query_url:", query_url)
	resp, err := http.Get(query_url)
	if err != nil {
		beego.Debug(err)
	}

	defer resp.Body.Close()
	beego.Debug("----------------queryLaohuangliShichen body--------------------")
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		beego.Debug(err)
	} else {
		beego.Debug(string(body))
	}

	var obj ShiChen
	if err := json.Unmarshal(body, &obj); err == nil {
		beego.Debug("================json str 转struct==")
		// beego.Debug(obj)
		beego.Debug("SCErrorCode", obj.ErrorCode)
		c.Data["SCErrorCode"] = obj.ErrorCode
		if obj.ErrorCode != 0 {
			c.Data["SCErrorInfo"] = getLaouangliError(obj.ErrorCode)
		} else {
			c.Data["SCResults"] = obj.SCResult
		}
	} else {
		beego.Debug(err)
	}
}

func getLaouangliError(errorcode int64) string {
	error_info := ""
	if errorcode == 206501 {
		error_info = "日期不能为空"
	}
	return error_info
}
