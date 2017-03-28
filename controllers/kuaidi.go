package controllers

/*
快递查询
*/
import (
	"encoding/json"
	"github.com/astaxie/beego"
	"io/ioutil"
	"net/http"
	"strings"
)

type Kuaidi struct {
	Reason     string  `json:"reason"`
	ErrorCode  int64   `json:"error_code"`
	Resultcode string  `json:"resultcode"`
	KResult    KResult `json:"result"`
}

type KResult struct {
	Company    string       `json:"company"`
	Com        string       `json:"com"`
	No         string       `json:"no"`
	KuaidiItem []KuaidiItem `json:"list"`
}
type KuaidiItem struct {
	Datetime string `json:"datetime"`
	Remark   string `json:"remark"`
	Zone     string `json:"zone"`
}

type KuaidiController struct {
	beego.Controller
}

func (c *KuaidiController) Get() {
	c.Data["IsShow"] = "false"
	c.Data["Test"] = "aaa"
	c.TplName = "kuaidi.html"
}

func (c *KuaidiController) Post() {
	c.Data["IsShow"] = "true"
	beego.Debug("KuaidiController------")
	com := c.Input().Get("com")
	no := c.Input().Get("no")
	if len(com) != 0 && len(no) != 0 {
		queryKuaidi(com, no, c)
	}
	c.TplName = "kuaidi.html"
}

func queryKuaidi(com string, no string, c *KuaidiController) {
	query_url := "[REALM]?com=[COM]&no=[NO]&key=[KEY]"
	query_url = strings.Replace(query_url, "[REALM]", "http://v.juhe.cn/exp/index", -1)
	query_url = strings.Replace(query_url, "[COM]", com, -1)
	query_url = strings.Replace(query_url, "[NO]", no, -1)
	query_url = strings.Replace(query_url, "[KEY]", "cc85926bb6f6c5b2eacdb5cc8bea3adb", -1)
	beego.Debug("signature_str:", query_url)
	resp, err := http.Get(query_url)
	if err != nil {
		beego.Debug(err)
	}

	defer resp.Body.Close()
	beego.Debug("----------------queryKuaidi body--------------------")
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		beego.Debug(err)
	} else {
		// beego.Debug(string(body))
	}
	var obj Kuaidi
	if err := json.Unmarshal(body, &obj); err == nil {
		beego.Debug("================json str 转struct==")
		beego.Debug(obj)
		beego.Debug("ErrorCode", obj.ErrorCode)
		c.Data["ErrorCode"] = obj.ErrorCode
		if obj.ErrorCode != 0 {
			c.Data["ErrorInfo"] = getKError(obj.ErrorCode)
		} else {
			c.Data["KResult"] = obj.KResult
			c.Data["KuaidiItems"] = obj.KResult.KuaidiItem
		}
	} else {
		beego.Debug(err)
	}
}

func getKError(errorcode int64) string {
	error_info := ""
	if errorcode == 204301 {
		error_info = "未被识别的快递公司"
	}
	if errorcode == 204302 {
		error_info = "请填写正确的运单号"
	}
	if errorcode == 204303 {
		error_info = "加载类库失败"
	}
	if errorcode == 204304 {
		error_info = "查询失败"
	}
	return error_info
}
