package controllers

/*
金钱详情
*/
import (
	"github.com/astaxie/beego"
)

type MoneyHelpController struct {
	beego.Controller
}

func (c *MoneyHelpController) Get() {
	c.TplName = "moneyhelp.html"
}
func (c *MoneyHelpController) Post() {
	c.TplName = "moneyhelp.html"
}
