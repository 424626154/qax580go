package qutil

import (
	"qax580go/models"
	"strings"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/config"
	"github.com/astaxie/beego/context"
)

// const COOKIE_UID = "zz580_uid"

func GetCookeUid() string {
	return COOKIE_UID
}

func ChackAccount(ctx *context.Context) (bool, string) {
	ck, err := ctx.Request.Cookie("username")
	if err != nil {
		return false, ""
	}

	username := ck.Value

	ck, err = ctx.Request.Cookie("password")
	if err != nil {
		return false, ""
	}

	password := ck.Value

	admin, err := models.GetOneAdmin(username)
	if err != nil {
		return false, ""
	}
	if admin != nil && strings.EqualFold(username, admin.Username) && strings.EqualFold(password, admin.Password) {
		beego.Debug(" cookie username ", username)
		return true, username
	} else {
		return false, username
	}
}

func GetImageUrl() string {
	url := ""
	isdebug := "flase"
	iniconf, err := config.NewConfig("json", "conf/myconfig.json")
	if err != nil {
		beego.Error(err)
	} else {
		isdebug = iniconf.String("qax580::isdebug")
		if isdebug == "true" {
			url = iniconf.String("qax580::imgurltest")
		} else {
			url = iniconf.String("qax580::imgurl")
		}

	}
	return url
}
