package controllers

/*
注册
*/
import (
	"crypto/md5"
	"encoding/hex"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/config"
	"github.com/satori/go.uuid"
	"gopkg.in/gomail.v2"
	"qax580go/models"
	"strings"
)

type RegisterController struct {
	beego.Controller
}

func (c *RegisterController) Get() {

	c.Data["isTips"] = ""
	c.Data["Alert"] = ""
	c.TplName = "register.html"
}
func (c *RegisterController) Post() {
	beego.Debug("post register")
	c.Data["isTips"] = ""
	c.Data["Alert"] = ""
	email := c.Input().Get("email")
	pwd := c.Input().Get("pwd")
	if len(email) != 0 && len(pwd) != 0 {
		user, err := models.GetOneUser(email)
		if err != nil {
			beego.Error(err)
		}
		if len(user.Email) == 0 {
			verify := getVerify(email, pwd)
			uid, username := getUserName(email, pwd)
			user, err = models.AddUser(email, pwd, uid, username, verify)
			if err != nil {
				beego.Error(err)
			}
			url := ""
			isdebug := "flase"
			iniconf, err := config.NewConfig("json", "conf/myconfig.json")
			if err != nil {
				beego.Error(err)
			} else {
				isdebug = iniconf.String("qax580::isdebug")
				if isdebug == "true" {
					url = iniconf.String("qax580::emailurltest") + "registerverify"
				} else {
					url = iniconf.String("qax580::emailurl") + "registerverify"
				}

			}
			beego.Debug("emailurl:", url)
			verifurl := url + "?verify=" + verify
			sendEmail(email, username, verifurl)
			c.Data["isTips"] = "请登录邮箱验证"
		} else {
			c.Data["Alert"] = "邮箱已注册"
		}
		beego.Error("user:", user)
	}
	c.TplName = "register.html"
}
func sendEmail(email string, username string, verifurl string) {
	to_email := email
	to_user := username
	subject := "咱这580账号注册成功"
	text := verifurl
	html := `
				<html>
				<body>
				<h3>
				点击链接完成您的注册:
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

func getVerify(email string, pwd string) string {
	verify := []string{email + "_" + pwd}
	h := md5.New()
	h.Write([]byte(strings.Join(verify, ""))) // 需要加密的字符串
	verifymd5 := hex.EncodeToString(h.Sum(nil))
	beego.Debug("verifymd5:", verifymd5)
	return verifymd5
}
func getUserName(email string, pwd string) (string, string) {
	_uuid := uuid.NewV4()
	uid := _uuid.String()
	username := "zz580_" + uid
	beego.Debug("username:", username)
	return uid, username
}
