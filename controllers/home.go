package controllers

/*
主页
*/
import (
	"qax580go/models"
	"qax580go/qutil"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/config"
)

type HomeController struct {
	beego.Controller
}

func (c *HomeController) Get() {
	openid := getCookie(c)
	c.Data["FromType"] = getHomeFromType(c)
	setUrl(c)
	// city := getSelectCity(c)
	// CurrentPage := int32(1)
	// count, err := models.GetPostCount()
	// if strings.Contains(city, qutil.CITY_ALL) {
	// } else {
	// 	count, err = models.GetCityPostCount(city)
	// }
	// NumberofPages := int32(10)
	// temp := count / NumberofPages
	// if (count % NumberofPages) != 0 {
	// 	temp = temp + 1
	// }
	// CotalPages := temp
	// pagetype := c.Input().Get("type")
	// page := c.Input().Get("page")
	// beego.Debug("pagetype:", pagetype)

	guanggaos, err := models.GetAllGuanggaosState1()
	if err != nil {
		beego.Error(err)
	}
	// c.Data["City"] = city
	c.Data["Guanggaos"] = guanggaos

	// if len(pagetype) != 0 && len(page) != 0 {
	// 	switch pagetype {
	// 	case "first": //首页
	// 		CurrentPage = 1
	// 	case "prev": //上一页
	// 		pageint, error := strconv.Atoi(page)
	// 		if error != nil {
	// 			beego.Error(error)
	// 		}
	// 		CurrentPage = int32(pageint)
	// 	case "next": //下一页
	// 		pageint, error := strconv.Atoi(page)
	// 		if error != nil {
	// 			beego.Error(error)
	// 		}
	// 		CurrentPage = int32(pageint)
	// 	case "last": //尾页
	// 		CurrentPage = CotalPages
	// 	case "page": //页码
	// 		pageint, error := strconv.Atoi(page)
	// 		if error != nil {
	// 			beego.Error(error)
	// 		}
	// 		CurrentPage = int32(pageint)
	// 	}
	// }
	// c.Data["CurrentPage"] = CurrentPage
	// c.Data["CotalPages"] = CotalPages
	// c.Data["NumberofPages"] = NumberofPages
	// if strings.Contains(city, qutil.CITY_ALL) {
	// 	posts, err := models.QueryPagePost(CurrentPage-1, NumberofPages)
	// 	if err != nil {
	// 		beego.Error(err)
	// 	}
	// 	c.Data["Posts"] = posts
	// } else {
	// 	posts, err := models.QueryCityPagePost(CurrentPage-1, NumberofPages, city)
	// 	if err != nil {
	// 		beego.Error(err)
	// 	}
	// 	c.Data["Posts"] = posts
	// }
	posts, err := models.QueryHomePost()
	if err != nil {
		beego.Debug(err)
	}
	c.Data["Posts"] = posts

	wxofficials, err := models.GetHomeWxOfficials()
	if err != nil {
		beego.Error(err)
	}
	c.Data["WxOfficials"] = wxofficials

	wxplatforms, err := models.GetWxPlatformTJ(1)
	if err != nil {
		beego.Error(err)
	}
	c.Data["WxPlatforms"] = wxplatforms

	wechats, err := models.GetHomeWeChats()
	if err != nil {
		beego.Error(err)
	}
	c.Data["WeChats"] = wechats

	// beego.Debug(posts)
	c.TplName = "home/home.html"

	isdebug := "true"
	iscanting := "false"
	iniconf, err := config.NewConfig("json", "conf/myconfig.json")
	if err != nil {
		beego.Error(err)
	} else {
		isdebug = iniconf.String("qax580::isdebug")
		iscanting = iniconf.String("qax580::iscanting")
	}
	// beego.Debug(isdebug)
	beego.Debug("IsCanting", iscanting)
	c.Data["IsDebug"] = isdebug
	c.Data["IsCanting"] = iscanting
	notice_num, err := models.GetUserNoticeNum(openid)
	if err != nil {
		beego.Error(err)
	}
	beego.Debug("notice_num :", notice_num)
	c.Data["NoticeNum"] = notice_num
	op := c.Input().Get("op")
	// beego.Debug("op------", op)
	switch op {
	case "del":
		id := c.Input().Get("id")
		if len(id) == 0 {
			break
		}
		id = c.Input().Get("id")
		err := models.DeletePost(id)
		if err != nil {
			beego.Error(err)
		}
		// beego.Debug("is del " + id)
		c.Redirect("/", 302)
		return
	case "back":
		// beego.Debug("退出登陆------")
		maxAge := 0
		c.Ctx.SetCookie(qutil.COOKIE_UID, "", maxAge, "/")
		c.Redirect("/", 302)
		return
	}
}

func (c *HomeController) Post() {
	op := c.Input().Get("op")
	switch op {
	case "city":
		city := c.Input().Get("city")
		maxAge := 1<<31 - 1
		citys := [14]City{City{qutil.CITY_QA, qutil.PY_QA}, City{qutil.CITY_SH, qutil.PY_SH}, City{qutil.CITY_BL, qutil.PY_BL}, City{qutil.CITY_AD, qutil.PY_AD},
			City{qutil.CITY_SD, qutil.PY_SD}, City{qutil.CITY_HL, qutil.PY_HL}, City{qutil.CITY_WK, qutil.PY_WK}, City{qutil.CITY_LX, qutil.PY_LX},
			City{qutil.CITY_QG, qutil.PY_QG}, City{qutil.CITY_MS, qutil.PY_MS}, City{qutil.CITY_SL, qutil.PY_SL}, City{qutil.CITY_ALL, qutil.PY_ALL},
			City{qutil.CITY_TL, qutil.PY_TL}, City{qutil.CITY_MX, qutil.PY_MX}}
		city_name := ""
		for i := 0; i < len(citys); i++ {
			if citys[i].City == city {
				city_name = citys[i].Name
			}
		}
		if len(city_name) != 0 {
			c.Ctx.SetCookie(qutil.COOKIE_CITY, city_name, maxAge, "/")
		}
		c.Redirect("/", 302)
		return
	default:
		c.Redirect("/", 302)
		return
	}
}

func getCookie(c *HomeController) string {
	isUser := false
	// openid := c.Ctx.GetCookie(COOKIE_WX_OPENID)
	openid := c.Ctx.GetCookie(qutil.COOKIE_UID)
	// beego.Debug("------------openid--------")
	// beego.Debug(openid)
	if len(openid) != 0 {
		// wxuser, err := models.GetOneWxUserInfo(openid)
		wxuser, err := models.GetOneUserUid(openid)
		if err != nil {
			beego.Error(err)
		} else {
			isUser = true
			// beego.Debug("--------------wxuser----------")
			// beego.Debug(wxuser)
			c.Data["WxUser"] = wxuser
		}
	}
	c.Data["isUser"] = isUser
	return openid
}

func setUrl(c *HomeController) {
	c.Data["ImgUrlPath"] = getImageUrl()
}

type City struct {
	City string
	Name string
}

// func getSelectCity(c *HomeController) string {
// 	citys := [14]City{City{qutil.CITY_QA, qutil.PY_QA}, City{qutil.CITY_SH, qutil.PY_SH}, City{qutil.CITY_BL, qutil.PY_BL}, City{qutil.CITY_AD, qutil.PY_AD},
// 		City{qutil.CITY_SD, qutil.PY_SD}, City{qutil.CITY_HL, qutil.PY_HL}, City{qutil.CITY_WK, qutil.PY_WK}, City{qutil.CITY_LX, qutil.PY_LX},
// 		City{qutil.CITY_QG, qutil.PY_QG}, City{qutil.CITY_MS, qutil.PY_MS}, City{qutil.CITY_SL, qutil.PY_SL}, City{qutil.CITY_ALL, qutil.PY_ALL},
// 		City{qutil.CITY_TL, qutil.PY_TL}, City{qutil.CITY_MX, qutil.PY_MX}}
// 	city_default := qutil.CITY_ALL
// 	city := ""
// 	city_name := c.Ctx.GetCookie(qutil.COOKIE_CITY)
// 	if len(city_name) == 0 { //未默认选择
// 		//判断是否有来源庆安县580:from_qingan 铁力580:from_tieli 茂县580:from_maoxian
// 		fromtype := getHomeFromType(c)
// 		if fromtype == "from_qingan" {
// 			city = qutil.CITY_QA
// 		} else if fromtype == "from_tieli" {
// 			city = qutil.CITY_TL
// 		} else if fromtype == "from_maoxian" {
// 			city = qutil.CITY_MX
// 		}
// 		if len(city) != 0 {
// 			return city
// 		}
// 		for i := 0; i < len(citys); i++ {
// 			if citys[i].Name == city_name {
// 				city = citys[i].City
// 			}
// 		}
// 	}
//
// 	for i := 0; i < len(citys); i++ {
// 		if citys[i].Name == city_name {
// 			city = citys[i].City
// 		}
// 	}
// 	if len(city) == 0 {
// 		city = city_default
// 	}
// 	// beego.Debug("getSelectCity", city)
// 	return city
// }

/**
*来源类型
 */
func getHomeFromType(c *HomeController) string {
	from_type := c.Ctx.GetCookie(qutil.COOKIE_FROM_TYPE)
	beego.Debug("show COOKIE_FROM_TYPE:", from_type)
	if len(from_type) == 0 {
		from_type = qutil.COOKIE_FROM_ALL
	}
	return from_type
}
