package controllers

/*
历史上的今天
*/
import (
	"encoding/json"
	"github.com/astaxie/beego"
	"io/ioutil"
	"net/http"
	"strings"
)

type HistoryCon struct {
	Reason    string   `json:"reason"`
	ErrorCode int64    `json:"error_code"`
	HCResult  HCResult `json:"result"`
}

type HCResult struct {
	Content string `json:"content"`
	Day     int32  `json:"day"`
	Des     string `json:"des"`
	Id      string `json:"_id"`
	Lunar   string `json:"lunar"`
	Month   int32  `json:"month"`
	Pic     string `json:"pic"`
	Title   string `json:"title"`
	Year    int32  `json:"year"`
}

type HistoryConController struct {
	beego.Controller
}

func (c *HistoryConController) Get() {
	id := c.Input().Get("id")
	if len(id) != 0 {
		queryHistoryCon(id, c)
	}
	c.TplName = "historycon.html"
}

func (c *HistoryConController) Post() {
	c.TplName = "historycon.html"
}

func queryHistoryCon(id string, c *HistoryConController) {
	query_url := "[REALM]?&v=1.0&id=[ID]&key=[KEY]"
	query_url = strings.Replace(query_url, "[REALM]", "http://api.juheapi.com/japi/tohdet", -1)
	query_url = strings.Replace(query_url, "[ID]", id, -1)
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

	var obj HistoryCon
	if err := json.Unmarshal(body, &obj); err == nil {
		beego.Debug("================json str 转struct==")
		// beego.Debug(obj)
		beego.Debug("ErrorCode", obj.ErrorCode)
		c.Data["ErrorCode"] = obj.ErrorCode
		if obj.ErrorCode != 0 {
			c.Data["ErrorInfo"] = getLishiError(obj.ErrorCode)
		} else {
			c.Data["HCResult"] = obj.HCResult
		}
	} else {
		beego.Debug(err)
	}
}
