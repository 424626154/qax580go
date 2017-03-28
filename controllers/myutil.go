package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/config"
	"qax580go/models"
	"time"
)

func getImageUrl() string {
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
	beego.Debug("111111", url)
	return url
}

func getOrderNumber(sizeid string, tempid string, photonum int) string {
	size_str := fmt.Sprintf("%03s", sizeid)
	temp_str := fmt.Sprintf("%03s", tempid)
	photonum_str := fmt.Sprintf("%02d", photonum)
	time_str := time.Now().Format("20060102")
	order_count, err := models.GetAllPorderNum()
	if err != nil {
		beego.Error(order_count)
	}
	order_count_str := fmt.Sprintf("%04d", order_count)
	order_number := fmt.Sprintf("%s%s%s%s%s", time_str, size_str, temp_str, photonum_str, order_count_str)
	return order_number
}
