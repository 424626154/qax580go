package controllers

import (
	"github.com/astaxie/beego"
	"strconv"
	"strings"
)

/**
js string转byte post 传 之后转 str
*/
func jsTostr(js string) string {
	strs := strings.Split(js, ",")
	slice2 := []byte{}
	for i := 0; i < len(strs); i++ {
		is, err := strconv.Atoi(strs[i])
		if err != nil {
			beego.Error(err)
		}
		slice2 = append(slice2, byte(is))
	}
	return string(slice2)
}
