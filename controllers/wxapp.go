package controllers

/**
*微信小程序服务器
*
**/
import (
	"encoding/json"
	// "fmt"
	"github.com/astaxie/beego"
	"github.com/satori/go.uuid"
	"path"
	"qax580go/models"
	"strconv"
)

type WxAppController struct {
	beego.Controller
}

//微信小程序
type ResponseJson struct {
	Data    string `json:"data"`
	ErrCode int64  `json:"errcode"`
	ErrMsg  string `json:"errmsg"`
}
type FidsJson struct {
	Data    []string `json:"data"`
	ErrCode int64    `json:"errcode"`
	ErrMsg  string   `json:"errmsg"`
}
type UploadJson struct {
	Data    FiledataJson `json:"data"`
	ErrCode int64        `json:"errcode"`
	ErrMsg  string       `json:"errmsg"`
}
type FiledataJson struct {
	Filename string `json:"filename"`
	Filetype string `json:"filetype"`
}

//获得UUID count=n [n0,n1,..]
func (c *WxAppController) Fids() {
	if c.Ctx.Input.IsGet() {
		// beego.Debug("Getuuid Get")
	}
	if c.Ctx.Input.IsPost() {
		// beego.Debug("Getuuid Post")
	}
	responseJson := FidsJson{}
	count := c.Input().Get("count")
	// beego.Debug("count :", count)
	if len(count) != 0 {
		count, err := strconv.Atoi(count)
		if err != nil {
			beego.Error(err)
		} else if count > 0 {
			uuids := []string{}
			for i := 0; i < count; i++ {
				_uuid := uuid.NewV4()
				uuids = append(uuids, _uuid.String())
			}
			beego.Debug("uuids :", uuids)
			responseJson.ErrCode = 0
			// data, err := json.Marshal(uuids)
			// if err != nil {
			// 	beego.Error(err)
			// }
			// files := fmt.Sprintf(`{"files":%s}`, uuids)
			responseJson.Data = uuids
		}
	} else {
		responseJson.ErrCode = WXAPP_ERR_CODE1000
		responseJson.ErrMsg = WXAPP_ERR_CODE1001_STR
	}
	body, err := json.Marshal(responseJson)
	if err != nil {
		beego.Error(err)
	}
	response_json := string(body)
	c.Ctx.WriteString(response_json)
	return
}

func (c *WxAppController) Upload() {
	if c.Ctx.Input.IsGet() {
		// beego.Debug("Upload Get")
	}
	if c.Ctx.Input.IsPost() {
		// beego.Debug("Upload Post")
	}
	responseJson := UploadJson{}
	filename := c.Input().Get("filename")
	filetype := c.Input().Get("filetype")
	beego.Debug("filename:", filename)
	beego.Debug("filetype:", filetype)
	beego.Debug("c:", c)
	if len(filename) != 0 && len(filetype) != 0 {
		// 获取附件
		_, fh, err := c.GetFile("file")
		if err != nil {
			beego.Error(err)
			responseJson.ErrCode = WXAPP_ERR_CODE1001
			responseJson.ErrMsg = WXAPP_ERR_CODE1001_STR
		}
		// } else {
		var attachment string
		if fh != nil {
			// 保存附件
			attachment = fh.Filename
			beego.Debug("attachment:", attachment) // 输出加密结果
			err = c.SaveToFile("file", path.Join("filehosting", filename))
			if err != nil {
				beego.Error(err)
			} else {
				err = models.AddFile(filename, filetype)
				if err != nil {
					beego.Error(err)
				} else {
					responseJson.ErrCode = 0
					filedataJson := FiledataJson{}
					filedataJson.Filename = filename
					filedataJson.Filetype = filetype
					responseJson.Data = filedataJson
				}
			}
		}
		// }
	} else {
		responseJson.ErrCode = WXAPP_ERR_CODE1000
		responseJson.ErrMsg = WXAPP_ERR_CODE1000_STR
	}

	body, err := json.Marshal(responseJson)
	if err != nil {
		beego.Error(err)
	}
	response_json := string(body)
	beego.Debug("response_json:", response_json)
	c.Ctx.WriteString(response_json)
	return
}

func (c *WxAppController) Addpost() {
	if c.Ctx.Input.IsGet() {
		beego.Debug("Addpost Get")
	}
	if c.Ctx.Input.IsPost() {
		beego.Debug("Addpost Post")
	}
	responseJson := ResponseJson{}
	responseJson.ErrCode = WXAPP_ERR_CODE1002
	responseJson.ErrMsg = WXAPP_ERR_CODE1002_STR
	title := c.Input().Get("title")
	content := c.Input().Get("content")
	images := c.Input().Get("images")
	beego.Debug("title:", title)
	beego.Debug("content:", content)
	beego.Debug("images:", images)
	if len(title) != 0 && len(content) != 0 {
		err := models.AddWxAppPost(title, content, images)
		if err != nil {
			beego.Error(err)
		} else {
			responseJson.ErrCode = 0
		}
	}
	body, err := json.Marshal(responseJson)
	if err != nil {
		beego.Error(err)
	}
	response_json := string(body)
	beego.Debug("response_json:", response_json)
	c.Ctx.WriteString(response_json)
	return
}
