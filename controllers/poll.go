package controllers

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"path"
	"qax580go/models"
	"qax580go/qutil"
	"strings"
	"time"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/config"
)

type PollController struct {
	beego.Controller
}

/**
投票后台
*/
func (c *PollController) Adminpolls() {
	bool, username := qutil.ChackAccount(c.Ctx)
	if bool {
		c.Data["isUser"] = bool
		c.Data["User"] = username
	} else {
		c.Redirect("/admin", 302)
		return
	}
	if c.Ctx.Input.IsGet() {
		beego.Debug("Adminpolls Get")
	}
	if c.Ctx.Input.IsPost() {
		beego.Debug("Adminpolls Post")
	}

	op := c.Input().Get("op")
	id := c.Input().Get("id")
	beego.Debug("op :", op)
	switch op {
	case "state":
		err := models.UpPollsState(id, 1)
		if err != nil {
			beego.Error(err)
		}
		beego.Debug("state0 id :", id)
		c.Redirect("/poll/adminpolls", 302)
		return
		return
	case "state1":
		err := models.UpPollsState(id, 0)
		if err != nil {
			beego.Error(err)
		}
		c.Redirect("/poll/adminpolls", 302)
		return
	case "del":
		err := models.DeletePolls(id)
		if err != nil {
			beego.Error(err)
		}
		c.Redirect("/poll/adminpolls", 302)
		return
	}

	objs, err := models.GetAllPolls()
	if err != nil {
		beego.Error(err)
	}
	c.Data["Objs"] = objs
	c.TplName = "adminpolls.html"
}

func (c *PollController) AdminUppollsInfo() {
	bool, username := qutil.ChackAccount(c.Ctx)
	if bool {
		c.Data["isUser"] = bool
		c.Data["User"] = username
	} else {
		c.Redirect("/admin", 302)
		return
	}
	// beego.Debug("c.Input() :", c.Input())
	if c.Ctx.Input.IsGet() {
		beego.Debug("AdminUppollsInfo Get")
	}
	if c.Ctx.Input.IsPost() {
		beego.Debug("AdminUppollsInfo Post")
		id := c.Input().Get("id")
		title := c.Input().Get("title")
		info := c.Input().Get("info")
		more := c.Input().Get("more")
		endtime := c.Input().Get("endtime")
		appid := c.Input().Get("appid")
		secret := c.Input().Get("secret")
		prize := c.Input().Get("prize")
		ext := c.Input().Get("ext")
		// test := []byte("http://mp.weixin.qq.com/s?__biz=MzA5MTQ2NjQ2MA==&mid=401382373&idx=1&sn=f585506b1c883712bc3a2b47ae6f09b7#rd")
		// beego.Debug("AdminUppollsInfo Post test:", test)
		// beego.Debug("AdminUppollsInfo Post more:", more)
		// if len(more) != 0 {
		// 	more = string(fmt.Sprintf("[%s]", more))
		// }
		// beego.Debug("AdminUppollsInfo Post test str:", string(test))
		// beego.Debug("AdminUppollsInfo Post more str:", more)
		// beego.Debug("more byte", strings.Split(more, ","))
		// strs := strings.Split(more, ",")
		// slice2 := []byte{}
		// for i := 0; i < len(strs); i++ {
		// 	is, err := strconv.Atoi(strs[i])
		// 	if err != nil {
		// 		beego.Error(err)
		// 	}
		// 	slice2 = append(slice2, byte(is))
		// }
		// more = string(slice2)
		more = jsTostr(more)
		beego.Debug("AdminUppollsInfo Post more ", more)
		// beego.Debug("AdminUppollsInfo Post test str:", string(strings.Split(more, ",")))
		if len(title) != 0 && len(info) != 0 && len(more) != 0 && len(appid) != 0 && len(secret) != 0 {
			beego.Debug("endtime", endtime)
			//获取本地location
			toBeCharge := endtime                                           //待转化为时间戳的字符串 注意 这里的小时和分钟还要秒必须写 因为是跟着模板走的 修改模板的话也可以不写
			timeLayout := "2006-01-02 15:04"                                //转化所需模板
			loc, _ := time.LoadLocation("Local")                            //重要：获取时区
			theTime, _ := time.ParseInLocation(timeLayout, toBeCharge, loc) //使用模板在对应时区转化为time.time类型
			endtimelong := theTime.Unix()                                   //转化为时间戳 类型是int64
			beego.Debug("theTime", theTime)                                 //打印输出theTime 2015-01-01 15:15:00 +0800 CST
			beego.Debug("endtimelong ", endtimelong)                        //打印输出时间戳 1420041600
			t := time.Now().Unix()
			beego.Debug("local time", t)
			if endtimelong < t {
				beego.Error("select end time error")
			} else {
				err := models.UpPollsInfo(id, title, info, more, endtimelong, appid, secret, prize, ext)
				if err != nil {
					beego.Error(err)
				}
			}
			url := "/poll/adminpolls"
			c.Redirect(url, 302)
			return
		}
	}

	id := c.Input().Get("id")
	obj, err := models.GetOnePolls(id)
	if err != nil {
		beego.Error(err)
	}
	beego.Debug("id :", id)
	beego.Debug("obj :", obj)
	c.Data["Id"] = id
	c.Data["Obj"] = obj
	c.TplName = "adminuppollsinfo.html"
}

func (c *PollController) AdminUppollsImg() {
	bool, username := qutil.ChackAccount(c.Ctx)
	if bool {
		c.Data["isUser"] = bool
		c.Data["User"] = username
	} else {
		c.Redirect("/admin", 302)
		return
	}
	c.Data["isUser"] = bool
	c.Data["User"] = username
	if c.Ctx.Input.IsGet() {
		beego.Debug("AdminUppollsImg Get")
	}
	if c.Ctx.Input.IsPost() {
		beego.Debug("AdminUppollsImg Post")

		id := c.Input().Get("id")
		originalimg := c.Input().Get("originalimg")
		image_name := originalimg
		if len(id) != 0 {
			// 获取附件
			_, fh, err := c.GetFile("image")
			// beego.Debug("上传图片:", fh)
			if err != nil {
				beego.Error(err)
			}
			var attachment string
			if fh != nil {
				// 保存附件
				attachment = fh.Filename
				t := time.Now().Unix()
				str2 := fmt.Sprintf("%d", t)
				s := []string{attachment, str2}
				h := md5.New()
				h.Write([]byte(strings.Join(s, ""))) // 需要加密的字符串
				image_name = hex.EncodeToString(h.Sum(nil))
				beego.Info(image_name) // 输出加密结果
				err = c.SaveToFile("image", path.Join("imagehosting", image_name))
				if err != nil {
					beego.Error(err)
					image_name = originalimg
				}
			}

			beego.Debug("上传前图片", originalimg, "上传后图片", image_name)
			if len(image_name) != 0 {
				err := models.UpPollsImg(id, image_name)
				if err != nil {
					beego.Error(err)
				} else {
					c.Redirect("/poll/adminpolls", 302)
					return
				}
			}
		}
	}
	id := c.Input().Get("id")
	obj, err := models.GetOnePolls(id)
	if err != nil {
		beego.Error(err)
	}
	c.Data["Obj"] = obj
	c.TplName = "adminuppollsimg.html"
}

func (c *PollController) Adminpollscon() {
	bool, username := qutil.ChackAccount(c.Ctx)
	if bool {
		c.Data["isUser"] = bool
		c.Data["User"] = username
	} else {
		c.Redirect("/admin", 302)
		return
	}
	pollsid := c.Input().Get("pollsid")
	if c.Ctx.Input.IsGet() {
		beego.Debug("Adminpollscon Get")
		op := c.Input().Get("op")
		beego.Debug("op:", op)
		switch op {
		case "state":
			id := c.Input().Get("id")
			if len(id) == 0 {
				break
			}
			id = c.Input().Get("id")
			err := models.UpdatePollState(pollsid, id, 1)
			if err != nil {
				beego.Error(err)
			}
			url := fmt.Sprintf("/poll/adminpollscon?pollsid=%s", pollsid)
			c.Redirect(url, 302)
			return
		case "state1":
			id := c.Input().Get("id")
			if len(id) == 0 {
				break
			}
			id = c.Input().Get("id")
			err := models.UpdatePollState(pollsid, id, 0)
			if err != nil {
				beego.Error(err)
			}
			url := fmt.Sprintf("/poll/adminpollscon?pollsid=%s", pollsid)
			c.Redirect(url, 302)
			return
		}
	}

	if c.Ctx.Input.IsPost() {
		beego.Debug("Adminpollscon Post")
	}
	objs, err := models.GetAllPoll(pollsid)
	if err != nil {
		beego.Error(err)
	}
	for i := 0; i < len(objs); i++ {
		num, err := models.GetVoteNum(pollsid, objs[i].Id)
		if err != nil {
			beego.Error(err)
		}
		objs[i].VoteNum = num
	}
	beego.Debug("pollsid:", pollsid)
	beego.Debug("objs:", objs)
	c.Data["Objs"] = objs
	c.Data["PollsId"] = pollsid
	c.TplName = "adminpollscon.html"
}
func (c *PollController) AdminpollVote() {
	bool, username := qutil.ChackAccount(c.Ctx)
	if bool {
		c.Data["isUser"] = bool
		c.Data["User"] = username
	} else {
		c.Redirect("/admin", 302)
		return
	}
	pollsid := c.Input().Get("pollsid")
	pollid := c.Input().Get("pollid")
	if c.Ctx.Input.IsGet() {
		beego.Debug("Adminpollscon Get")
	}

	if c.Ctx.Input.IsPost() {
		beego.Debug("Adminpollscon Post")
	}
	objs, err := models.GetAllVote(pollsid, pollid)
	if err != nil {
		beego.Error(err)
	}
	beego.Debug("pollsid:", pollsid)
	beego.Debug("pollid:", pollid)
	beego.Debug("objs:", objs)
	c.Data["Objs"] = objs
	c.Data["PollsId"] = pollsid
	c.TplName = "adminpollvote.html"
}

/**
投票后台添加新投票
*/
func (c *PollController) Adminaddpoll() {
	bool, username := qutil.ChackAccount(c.Ctx)
	if bool {
		c.Data["isUser"] = bool
		c.Data["User"] = username
	} else {
		c.Redirect("/admin", 302)
		return
	}
	if c.Ctx.Input.IsGet() {
		beego.Debug("Adminaddpoll Get")
	}
	if c.Ctx.Input.IsPost() {
		beego.Debug("Adminaddpoll Post")
		image_name := ""
		title := c.Input().Get("title")
		info := c.Input().Get("info")
		more := c.Input().Get("more")
		endtime := c.Input().Get("endtime")
		appid := c.Input().Get("appid")
		secret := c.Input().Get("secret")
		prize := c.Input().Get("prize")
		ext := c.Input().Get("ext")
		if len(title) != 0 && len(info) != 0 && len(more) != 0 && len(appid) != 0 && len(secret) != 0 {
			// 获取附件
			_, fh, err := c.GetFile("image")
			// beego.Debug("上传图片:", fh)
			if err != nil {
				beego.Error(err)
			}
			var attachment string
			if fh != nil {
				// 保存附件
				attachment = fh.Filename
				t := time.Now().Unix()
				str2 := fmt.Sprintf("%d", t)
				s := []string{attachment, str2}
				h := md5.New()
				h.Write([]byte(strings.Join(s, ""))) // 需要加密的字符串
				image_name = hex.EncodeToString(h.Sum(nil))
				beego.Info(image_name) // 输出加密结果
				err = c.SaveToFile("image", path.Join("imagehosting", image_name))
				if err != nil {
					beego.Error(err)
					image_name = ""
				}
			}
			beego.Debug("endtime", endtime)
			//获取本地location
			toBeCharge := endtime                                           //待转化为时间戳的字符串 注意 这里的小时和分钟还要秒必须写 因为是跟着模板走的 修改模板的话也可以不写
			timeLayout := "2006-01-02 15:04"                                //转化所需模板
			loc, _ := time.LoadLocation("Local")                            //重要：获取时区
			theTime, _ := time.ParseInLocation(timeLayout, toBeCharge, loc) //使用模板在对应时区转化为time.time类型
			endtimelong := theTime.Unix()                                   //转化为时间戳 类型是int64
			beego.Debug("theTime", theTime)                                 //打印输出theTime 2015-01-01 15:15:00 +0800 CST
			beego.Debug("endtimelong ", endtimelong)                        //打印输出时间戳 1420041600
			t := time.Now().Unix()
			beego.Debug("local time", t)
			if endtimelong < t {
				beego.Error("select end time error")
			} else {
				err = models.AddPolls(title, info, image_name, more, endtimelong, appid, secret, prize, ext)
				if err != nil {
					beego.Error(err)
				}
			}
			url := "/poll/adminpolls"
			c.Redirect(url, 302)
			return
		}
	}
	c.TplName = "adminaddpoll.html"
}

/**
投票主页
*/
func (c *PollController) PollHome() {
	if c.Ctx.Input.IsGet() {
		beego.Debug("PollHome Get")
	}
	if c.Ctx.Input.IsPost() {
		beego.Debug("PollHome Post")
	}
	pollsid := c.Input().Get("pollsid")
	beego.Debug("pollsid :", pollsid)
	state := c.Input().Get("state")
	code := c.Input().Get("code")
	beego.Debug("/poll/pollhem state :", state)
	beego.Debug("/poll/pollhem code :", code)
	share_url := "http://www.baoguangguang.cn/poll/pollhome"
	if len(code) != 0 && len(state) != 0 {
		pollsid = state
		_, err := getPollWxOpenId(c, pollsid, code)
		if err != nil {
			beego.Error(err)
		}
		share_url = fmt.Sprintf("http://www.baoguangguang.cn/poll/pollhome?code=%s&state=%s", code, state)
	}
	beego.Debug("/poll/pollhem pollsid :", pollsid)
	openid := getPollCookie(c)
	//测试openid
	isdebug := "false"
	iniconf, err := config.NewConfig("json", "conf/myconfig.json")
	if err != nil {
		beego.Debug(err)
	} else {
		isdebug = iniconf.String("qax580::isdebug")
	}
	if isdebug == "true" {
		openid = "o3AhEuB_wdTELvlErL4F1Em4Nck4"
		c.Data["OpenId"] = openid
	}

	op := c.Input().Get("op")
	beego.Debug("op :", op)
	switch op {
	case "vote":
		id := c.Input().Get("id")
		err := models.AddVote(openid, pollsid, id)
		if err != nil {
			beego.Error(err)
		}
		url := fmt.Sprintf("/poll/pollhome?pollsid=%s", pollsid)
		c.Redirect(url, 302)
		return
	}
	c.Data["Time"] = int64(0)
	if len(pollsid) != 0 {
		err := models.AddPollsPv(pollsid)
		if err != nil {
			beego.Error(err)
		}
		polls, err := models.GetOnePolls(pollsid)
		if err != nil {
			beego.Error(err)
		}
		beego.Debug("polls", polls)
		c.Data["Time"] = polls.EndTimeLong
		c.Data["Polls"] = polls
		pv, err := models.GetPollsPv(pollsid)
		c.Data["PV"] = pv
		pollnum, err := models.GetPollAllNum(pollsid)
		c.Data["PollNum"] = pollnum
		votenum, err := models.GetVoteAllNum(pollsid)
		c.Data["VoteNum"] = votenum

		endtime := polls.EndTimeLong
		curtime := time.Now().Unix()
		timestr := "活动已过期"
		if endtime-curtime > 0 {
			t := time.Unix(endtime, 0)

			_, mon, day := t.Date()
			_, cmon, cday := time.Now().Date()
			hour, min, _ := t.Clock()
			chour, cmin, _ := time.Now().Clock()
			timestr = fmt.Sprintf("%d月%d天%02d小时%02d分", mon-cmon, day-cday, hour-chour, min-cmin)
			// beego.Debug(timestr)
		}
		c.Data["TimeStr"] = timestr

		objs, err := models.GetAllPollState(pollsid, 1)
		if err != nil {
			beego.Debug(err)
		}
		for i := 0; i < len(objs); i++ {
			num, err := models.GetVoteNum(pollsid, objs[i].Id)
			if err != nil {
				beego.Error(err)
			}
			objs[i].VoteNum = num
		}
		// beego.Debug("objs :", objs)
		c.Data["Objs"] = objs
		wxShareCon := models.WxShareCon{}
		wxShareCon.Title = polls.Title
		wxShareCon.Link = fmt.Sprintf("http://www.baoguangguang.cn/poll/pollwx?id=%s", pollsid)
		wxShareCon.ImgUrl = fmt.Sprintf("http://182.92.167.29:8080/imagehosting/%s", polls.Image)
		getPollShare(polls.Appid, polls.Secret, share_url, wxShareCon, c)
	}
	c.Data["PollsId"] = pollsid
	c.TplName = "pollhome.html"
}

/**
投票详情
*/
func (c *PollController) PollHomeCon() {
	if c.Ctx.Input.IsGet() {
		beego.Debug("PollHome Get")
	}
	if c.Ctx.Input.IsPost() {
		beego.Debug("PollHome Post")
	}
	openid := getPollCookie(c)
	pollsid := c.Input().Get("pollsid")
	pollid := c.Input().Get("pollid")
	beego.Debug("pollsid:", pollsid)
	beego.Debug("pollid:", pollid)
	c.Data["PollsId"] = pollsid
	c.Data["PollId"] = pollid
	op := c.Input().Get("op")
	beego.Debug("op:", op)
	switch op {
	case "vote":
		err := models.AddVote(openid, pollsid, pollid)
		if err != nil {
			beego.Debug(err)
		}
		url := fmt.Sprintf("/poll/pollhomecon?pollsid=%s&pollid=%s", pollsid, pollid)
		beego.Debug("url:", url)
		c.Redirect(url, 302)
		return
	}

	obj, err := models.GetOnePoll(pollsid, pollid)
	if err != nil {
		beego.Error(err)
	}
	num, err := models.GetVoteNum1(pollsid, pollid)
	if err != nil {
		beego.Error(err)
	}
	polls, err := models.GetOnePolls(pollsid)
	if err != nil {
		beego.Error(err)
	}

	wxShareCon := models.WxShareCon{}
	wxShareCon.Title = obj.Title
	wxShareCon.Link = fmt.Sprintf("http://www.baoguangguang.cn/poll/pollwx?id=%s", pollsid)
	wxShareCon.ImgUrl = fmt.Sprintf("http://182.92.167.29:8080/imagehosting/%s", obj.Image)
	shaer_url := fmt.Sprintf("http://www.baoguangguang.cn/poll/pollhomecon?pollsid=%s&pollid=%s", pollsid, pollid)
	getPollShare(polls.Appid, polls.Secret, shaer_url, wxShareCon, c)
	beego.Debug("VoteNum", num)
	obj.VoteNum = num
	c.Data["Time"] = polls.EndTimeLong
	c.Data["Polls"] = polls
	c.Data["Obj"] = obj
	c.TplName = "pollhomecon.html"
}

/**
投票搜索详情
*/
func (c *PollController) PollHomeSearch() {
	if c.Ctx.Input.IsGet() {
		beego.Debug("PollHomeSearch Get")
	}
	if c.Ctx.Input.IsPost() {
		beego.Debug("PollHomeSearch Post")
	}
	openid := getPollCookie(c)
	pollsid := c.Input().Get("pollsid")
	search := c.Input().Get("search")
	beego.Debug("pollsid:", pollsid)
	beego.Debug("search:", search)
	c.Data["PollsId"] = pollsid
	c.Data["Search"] = search
	op := c.Input().Get("op")
	beego.Debug("op:", op)
	switch op {
	case "vote":
		pollid := c.Input().Get("pollid")
		err := models.AddVote(openid, pollsid, pollid)
		if err != nil {
			beego.Debug(err)
		}
		url := fmt.Sprintf("/poll/pollhomesearch?pollsid=%s&search=%s", pollsid, search)
		beego.Debug("url:", url)
		c.Redirect(url, 302)
		return
	}
	objs, err := models.GetAllPollOr(search)
	if err != nil {
		beego.Debug(err)
	}
	for i := 0; i < len(objs); i++ {
		num, err := models.GetVoteNum(pollsid, objs[i].Id)
		if err != nil {
			beego.Error(err)
		}
		objs[i].VoteNum = num
	}
	polls, err := models.GetOnePolls(pollsid)
	if err != nil {
		beego.Error(err)
	}
	c.Data["Time"] = polls.EndTimeLong
	c.Data["Polls"] = polls
	c.Data["Objs"] = objs
	c.TplName = "pollhomesearch.html"
}

/**
查看排名
*/
func (c *PollController) PollHomeRanking() {
	if c.Ctx.Input.IsGet() {
		beego.Debug("PollHomeRanking Get")
	}
	if c.Ctx.Input.IsPost() {
		beego.Debug("PollHomeRanking Post")
	}
	pollsid := c.Input().Get("pollsid")
	beego.Debug("pollsid:", pollsid)
	c.Data["PollsId"] = pollsid
	if len(pollsid) != 0 {
		objs, err := models.GetAllPollState(pollsid, 1)
		if err != nil {
			beego.Debug(err)
		}
		for i := 0; i < len(objs); i++ {
			num, err := models.GetVoteNum(pollsid, objs[i].Id)
			if err != nil {
				beego.Error(err)
			}
			objs[i].VoteNum = num
		}
		for i := 0; i < len(objs); i++ {
			for j := 0; j < len(objs)-i-1; j++ {
				if objs[j].VoteNum < objs[j+1].VoteNum {
					objs[j], objs[j+1] = objs[j+1], objs[j]
				}
			}
		}
		for i := 0; i < len(objs); i++ {
			objs[i].Ranking = int32(i)
		}
		beego.Debug("objs :", objs)
		c.Data["Objs"] = objs
	}
	c.TplName = "pollhomeranking.html"
}

/**
添加投票
*/
func (c *PollController) AddPoll() {
	openid := getPollCookie(c)
	pollsid := c.Input().Get("pollsid")
	op := c.Input().Get("op")
	switch op {
	case "del":
		id := c.Input().Get("id")
		err := models.DelPoll(pollsid, id)
		if err != nil {
			beego.Error(err)
		}
		beego.Debug("add poll del :", id)
		url := fmt.Sprintf("/poll/addpoll?pollsid=%s", pollsid)
		c.Redirect(url, 302)
		return
	}
	poll, err := models.GetMyPoll(pollsid, openid)
	if err != nil {
		beego.Error(err)
	}
	if len(poll) > 0 {
		c.Data["IsAdd"] = true //是否已经添加过
		mypoll := poll[0]
		num, err := models.GetVoteNum(pollsid, mypoll.Id)
		if err != nil {
			beego.Error(err)
		}
		mypoll.VoteNum = num
		c.Data["Poll"] = mypoll
	} else {
		c.Data["IsAdd"] = false
	}
	if c.Ctx.Input.IsGet() {
		beego.Debug("PollHome Get")
	}
	if c.Ctx.Input.IsPost() {
		beego.Debug("PollHome Post")
		image_name := ""
		title := c.Input().Get("title")
		info := c.Input().Get("info")
		contactway := c.Input().Get("contactway")
		if len(pollsid) != 0 && len(title) != 0 && len(info) != 0 && len(contactway) != 0 {
			// 获取附件
			_, fh, err := c.GetFile("image")
			// beego.Debug("上传图片:", fh)
			if err != nil {
				beego.Error(err)
			}
			var attachment string
			if fh != nil {
				// 保存附件
				attachment = fh.Filename
				t := time.Now().Unix()
				str2 := fmt.Sprintf("%d", t)
				s := []string{attachment, str2}
				h := md5.New()
				h.Write([]byte(strings.Join(s, ""))) // 需要加密的字符串
				image_name = hex.EncodeToString(h.Sum(nil))
				beego.Info(image_name) // 输出加密结果
				err = c.SaveToFile("image", path.Join("imagehosting", image_name))
				if err != nil {
					beego.Error(err)
					image_name = ""
				}
			}
			err = models.AddPoll(openid, pollsid, title, info, image_name, contactway)
			if err != nil {
				beego.Error(err)
			}
			url := fmt.Sprintf("/poll/pollhome?pollsid=%s", pollsid)
			c.Redirect(url, 302)
			return
		}
	}
	c.Data["PollsId"] = pollsid
	c.TplName = "addpoll.html"
}

func (c *PollController) PollWx() {
	if c.Ctx.Input.IsGet() {
		beego.Debug("PollWx Get")
	}
	if c.Ctx.Input.IsPost() {
		beego.Debug("PollWx Post")
	}
	id := c.Input().Get("id")
	if len(id) != 0 {
		polls, err := models.GetOnePolls(id)
		if err != nil {
			beego.Error(err)
		} else {
			appid := polls.Appid
			secret := polls.Secret
			if len(appid) != 0 && len(secret) != 0 {
				isdebug := "true"
				iniconf, err := config.NewConfig("json", "conf/myconfig.json")
				if err != nil {
					beego.Debug(err)
				} else {
					isdebug = iniconf.String("qax580::isdebug")
				}
				wx_url := "[REALM]?appid=[APPID]&redirect_uri=[REDIRECT_URI]&response_type=code&scope=snsapi_base&state=[STATE]#wechat_redirect"
				realm_name := ""
				if isdebug == "true" {
					realm_name = "http://localhost:9095"
				} else {
					realm_name = "https://open.weixin.qq.com/connect/oauth2/authorize"
				}
				redirect_uri := "http%3a%2f%2fwww.baoguangguang.cn%2fpoll%2fpollhome"
				state := id
				wx_url = strings.Replace(wx_url, "[REALM]", realm_name, -1)
				wx_url = strings.Replace(wx_url, "[APPID]", appid, -1)
				wx_url = strings.Replace(wx_url, "[REDIRECT_URI]", redirect_uri, -1)
				wx_url = strings.Replace(wx_url, "[STATE]", state, -1)
				beego.Debug("/poll/pollwx autho url :", wx_url)
				c.Redirect(wx_url, 302)
				return
			}
		}
	}
	c.TplName = "pollwx.html"
}

//校验是否可以投票
func (c *PollController) PollCheckVote() {
	if c.Ctx.Input.IsGet() {
		beego.Debug("PollCheckVote Get")
	}
	if c.Ctx.Input.IsPost() {
		beego.Debug("PollCheckVote Post")
	}
	request_json := `{"errcode":1,"errmsg":"pollcheckvote error"}`
	openid := c.Input().Get("openid")
	pollsid := c.Input().Get("pollsid")
	pollid := c.Input().Get("pollid")
	beego.Debug("PollCheckVote openid", openid)
	beego.Debug("PollCheckVote pollsid", pollsid)
	beego.Debug("PollCheckVote pollid", pollid)
	if len(openid) != 0 && len(pollsid) != 0 && len(pollid) != 0 {
		//判断今天是否投票
		votes, err := models.GetAllVote1(openid, pollsid, pollid)
		if err != nil {
			beego.Error(err)
		}
		if len(votes) != 0 {
			t := time.Unix(votes[0].CreateTime, 0)
			beego.Debug("PollCheckVote votes[0].CreateTime;", votes)
			beego.Debug("PollCheckVote time.Now().Unix()", time.Now().Unix())
			_, _, day := t.Date()
			_, _, cday := time.Now().Date()
			beego.Debug("PollCheckVote day;", day, "cday:", cday)
			if day != cday {
				request_json = `{"errcode":0,"errmsg":"pollcheckvote ok"}`
			}
		} else {
			request_json = `{"errcode":0,"errmsg":"pollcheckvote ok"}`
		}
	}
	beego.Debug("PollCheckVote request_json;", request_json)
	c.Ctx.WriteString(request_json)
}

func getPollWxOpenId(c *PollController, pollsid string, code string) (string, error) {
	polls, err := models.GetOnePolls(pollsid)
	if err != nil {
		beego.Error(err)
	} else {
		tokenobj, err := getWxAutoToken(polls.Appid, polls.Secret, code)
		if err != nil {
			beego.Error(err)
		} else {
			maxAge := 1<<31 - 1
			c.Ctx.SetCookie(qutil.COOKIE_WX_OPENID, tokenobj.OpenID, maxAge, "/")
			return tokenobj.OpenID, nil
		}
	}
	return "", err
}

/**
获得授权token
*/
func getWxAutoToken(appid string, secret string, code string) (models.AccessTokenJson, error) {
	// ?appid=APPID&secret=SECRET&code=CODE&grant_type=authorization_code
	tokenobj := models.AccessTokenJson{}
	response_json := `{"errcode":1,"errmsg":"getWxAccessToken error"}`
	wx_url := "[REALM]?appid=[APPID]&secret=[SECRET]&&code=[CODE]&grant_type=authorization_code"
	realm_name := "https://api.weixin.qq.com/sns/oauth2/access_token"
	wx_url = strings.Replace(wx_url, "[REALM]", realm_name, -1)
	wx_url = strings.Replace(wx_url, "[APPID]", appid, -1)
	wx_url = strings.Replace(wx_url, "[SECRET]", secret, -1)
	wx_url = strings.Replace(wx_url, "[CODE]", code, -1)
	beego.Debug("/poll getWxAutoToken url :", wx_url)
	resp, err := http.Get(wx_url)
	if err != nil {
		beego.Error(err)
		return tokenobj, err
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		beego.Error(err)
		body = []byte(response_json)
	} else {
		beego.Debug("wxqax getWxToken boey :", string(body))
	}
	var atj models.AccessTokenJson
	if err := json.Unmarshal(body, &atj); err == nil {
		beego.Debug("get Token obj", atj)
		tokenobj = atj
	} else {
		beego.Error(err)
	}
	return tokenobj, err
}

func getPollCookie(c *PollController) string {
	openid := c.Ctx.GetCookie(qutil.COOKIE_WX_OPENID)
	c.Data["OpenId"] = openid
	beego.Debug("/poll getPollCookie openid:", openid)
	return openid
}

func getPollShare(appid string, secret string, share_url string, wxShareCon models.WxShareCon, c *PollController) {

	// ticket_cookie := c.Ctx.GetCookie(COOKIE_WX_TICKET)
	// if len(ticket_cookie) != 0 {

	// } else {
	ticket_cookie := ""
	tokenobj, err := getWxToken(appid, secret)
	if err != nil {
		beego.Error(err)
	}
	if tokenobj.ErrCode == 0 {
		ticket, err := getWxTicket(tokenobj.AccessToken)
		if err != nil {
			beego.Error(err)
		}
		if err != nil {
			beego.Error()
		}
		if ticket.ErrCode == 0 {
			c.Ctx.SetCookie(qutil.COOKIE_WX_TICKET, ticket.Ticket, ticket.ExpiresIn, "/")
			ticket_cookie = ticket.Ticket
		}
	}
	// }
	wxShare := models.WxShare{}
	timestamp := time.Now().Unix()
	noncestr := getNonceStr(16, KC_RAND_KIND_ALL)
	basestr := "jsapi_ticket=[TICKET]&noncestr=[NONCESTR]&timestamp=[TIMESTAMP]&url=[URL]"
	basestr = strings.Replace(basestr, "[TICKET]", ticket_cookie, -1)
	basestr = strings.Replace(basestr, "[NONCESTR]", noncestr, -1)
	basestr = strings.Replace(basestr, "[TIMESTAMP]", fmt.Sprintf("%d", timestamp), -1)
	basestr = strings.Replace(basestr, "[URL]", share_url, -1)
	signaturestr := goWxJsSha1(basestr)
	beego.Debug(" getPollShare basestr", basestr)
	beego.Debug(" getPollShare ticket_cookie", ticket_cookie)
	beego.Debug(" getPollShare noncestr", noncestr)
	beego.Debug(" getPollShare timestamp", fmt.Sprintf("%d", timestamp))
	beego.Debug(" getPollShare signaturestr", signaturestr)
	beego.Debug(" getPollShare share_url", share_url)
	wxShare.AppId = appid
	wxShare.TimeStamp = timestamp
	wxShare.NonceStr = noncestr
	wxShare.Signature = signaturestr
	// beego.Debug(" getPollShare WxShare", wxShare)
	c.Data["WxShare"] = wxShare
	c.Data["WxShareCon"] = wxShareCon
	beego.Debug(" getPollShare wxShareCon", wxShareCon)

}
