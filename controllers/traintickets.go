package controllers

/*
火车票查询
*/
import (
	"github.com/astaxie/beego"
)

type TrainTicketsController struct {
	beego.Controller
}

func (c *TrainTicketsController) Get() {
	c.TplName = "traintickets.html"
}

func (c *TrainTicketsController) Post() {

}
