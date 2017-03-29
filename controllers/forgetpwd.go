package controllers

/**
*忘记密码
 */

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/config"
	"gopkg.in/gomail.v2"
	"qax580go/models"
	"strings"
)

type ForgetpwdController struct {
	beego.Controller
}

func (c *ForgetpwdController) Get() {
	beego.Debug("get register")
	c.Data["Alert"] = ""
	c.Data["isTips"] = ""
	c.TplName = "forgetpwd.html"
}
func (c *ForgetpwdController) Post() {
	beego.Debug("post register")
	c.Data["isTips"] = ""
	c.Data["Alert"] = ""
	email := c.Input().Get("email")
	if len(email) != 0 {
		user, err := models.GetOneUser(email)
		if err != nil {
			beego.Error(err)
		}
		if len(user.Email) == 0 {
			c.Data["Alert"] = "邮箱未注册"
		} else {
			username := user.Name
			verify := user.Verify
			url := ""
			isdebug := "flase"
			iniconf, err := config.NewConfig("json", "conf/myconfig.json")
			if err != nil {
				beego.Error(err)
			} else {
				isdebug = iniconf.String("qax580::isdebug")
				if isdebug == "true" {
					url = iniconf.String("qax580::emailurltest") + "forgetpwdverify"
				} else {
					url = iniconf.String("qax580::emailurl") + "forgetpwdverify"
				}

			}
			beego.Debug("emailurl:", url)
			verifurl := url + "?verify=" + verify
			sendForgetEmail(email, username, verifurl)
			c.Data["isTips"] = "请登录邮箱修改密码"
		}
	}
	c.TplName = "forgetpwd.html"
}

func sendForgetEmail(email string, username string, verifurl string) {
	beego.Debug("sendForgetEmail email:", email, "username:", username, "verifurl:", verifurl)
	to_email := email
	to_user := username
	subject := "咱这580找回密码"
	text := verifurl
	html := `
				<html>
				<body>
				<h3>
				<a href="[TEXT]">[TEXT]</a>
				</h3>
				</body>
				</html>
				`
	html = strings.Replace(html, "[TEXT]", text, -1)
	beego.Debug("send email html:", html)
	m := gomail.NewMessage()
	m.SetAddressHeader("From", "13671172337@163.com", "咱这580") // 发件人
	m.SetHeader("To",                                          // 收件人
		m.FormatAddress(to_email, to_user),
	)
	m.SetHeader("Subject", subject) // 主题
	m.SetBody("text/html", html)    // 正文

	d := gomail.NewPlainDialer("smtp.163.com", 465, "13671172337@163.com", "s123456") // 发送邮件服务器、端口、发件人账号、发件人密码
	if err := d.DialAndSend(m); err != nil {
		beego.Error(err)
	}
}
