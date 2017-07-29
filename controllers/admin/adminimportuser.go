package admin

/*
后台导入微信用户
*/
import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"qax580go/models"
	"qax580go/qutil"
	"strings"

	"github.com/astaxie/beego"
)

type ImportUser struct {
	OpenId   string
	IsImport bool
}

// {"total":2,"count":2,"data":{"openid":["","OPENID1","OPENID2"]},"next_openid":"NEXT_OPENID"}
// {"errcode":40013,"errmsg":"invalid appid"}
// 参数	说明
// total	关注该公众账号的总用户数
// count	拉取的OPENID个数，最大值为10000
// data	列表数据，OPENID的列表
// next_openid	拉取列表的最后一个用户的OPENID
type OpenIds struct {
	Total      int64      `json:"total"`
	Count      int64      `json:"count"`
	ErrorCode  int64      `json:"errcode"`
	ErrorMsg   string     `json:"errmsg"`
	Openiddata Openiddata `json:"data"`
}

type Openiddata struct {
	Openid []string `json:"openid"`
}

type AdminImportUserController struct {
	beego.Controller
}

func (c *AdminImportUserController) Get() {
	op := c.Input().Get("op")
	if op == "import" {
		beego.Debug("AdminImportUserController get import")
		openid := c.Input().Get("openid")
		errcode, token := qutil.GetToken()
		if errcode == 0 && len(openid) != 0 {
			user, err := qutil.GetWxUser(openid, token)
			if err == nil {
				if user.ErrCode == 0 {
					err = models.AddWxUserInfo(user)
					if err != nil {
						beego.Error(err)
					} else {
						err = models.AddWxUserMoney(user.OpenId, 4)
						if err != nil {
							beego.Error(err)
						} else {
							_, err = models.AddUserMoneyRecord(user.OpenId, qutil.MONEY_SUBSCRIBE_SUM, qutil.MONEY_SUBSCRIBE)
						}
					}

				}
			}
		}
		url := "/admin/importuser"
		c.Redirect(url, 302)
		return
	}
	wxusers, err := models.GetAllWxUsers()
	if err != nil {
		beego.Error("err :", err)
	}
	// var wxusermap map[string]int64
	wxusermap := make(map[string]int64)
	for i := 0; i < len(wxusers); i++ {
		wxusermap[wxusers[i].OpenId] = wxusers[i].Id
	}

	errcode, token := qutil.GetToken()
	beego.Debug("errcode :", errcode, "token :", token)
	if errcode == 0 {
		openids := getWxUserList(token, c)
		importUsers := make([]*ImportUser, len(openids))
		for i := 0; i < len(openids); i++ {
			importUser := &ImportUser{OpenId: openids[i]}
			if _, ok := wxusermap[openids[i]]; ok {
				//存在
				importUser.IsImport = true
			}
			// beego.Debug(" openids ", i, importUser)
			importUsers[i] = importUser
		}
		c.Data["ImportUsers"] = importUsers
	}
	c.TplName = "adminimportuser.html"
}
func (c *AdminImportUserController) Post() {

	c.TplName = "adminimportuser.html"
}

func getWxUserList(access_token string, c *AdminImportUserController) []string {
	var openids []string
	// https://api.weixin.qq.com/cgi-bin/user/get?access_token=ACCESS_TOKEN&next_openid=NEXT_OPENID
	query_url := "[REALM]?access_token=[ACCESS_TOKEN]"
	query_url = strings.Replace(query_url, "[REALM]", "https://api.weixin.qq.com/cgi-bin/user/get", -1)
	query_url = strings.Replace(query_url, "[ACCESS_TOKEN]", access_token, -1)
	beego.Debug("getWxUserList url:", query_url)
	resp, err := http.Get(query_url)
	if err != nil {
		beego.Debug(err)
	}

	defer resp.Body.Close()
	beego.Debug("----------------getWxUserList body--------------------")
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		beego.Debug(err)
	} else {
		// beego.Debug(string(body))
	}

	var station OpenIds
	if err := json.Unmarshal(body, &station); err == nil {
		beego.Debug("================json str 转struct==")
		// beego.Debug(station)
		beego.Debug("ErrorCode", station.ErrorCode)
		c.Data["ErrorCode"] = station.ErrorCode
		if station.ErrorCode != 0 {
			c.Data["ErrorInfo"] = qutil.GetError(station.ErrorCode)
		} else {
			openids = station.Openiddata.Openid
			// beego.Debug(station.Openiddata)
		}
	} else {
		beego.Debug(err)
	}
	return openids
}
