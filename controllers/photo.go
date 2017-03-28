package controllers

/**
洗相
*/
import (
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
	"net/url"
	"qax580go/models"
	"strings"
)

type PhotoController struct {
	beego.Controller
}

/*********home***********/
func (c *PhotoController) Home() {
	if c.Ctx.Input.IsGet() {
		beego.Debug("Home Get")
	}
	if c.Ctx.Input.IsPost() {
		beego.Debug("Home Post")
	}
	// state := c.Input().Get("state")
	// code := c.Input().Get("code")
	// // appid :=
	// // secret :=
	// beego.Debug("/poll/pollhem state :", state)
	// beego.Debug("/poll/pollhem code :", code)
	// c.Data["Parameter"] = false
	// if len(code) != 0 && len(state) != 0 {
	// 	obj, err := models.GetPoauthFromId(state)
	// 	if err != nil {
	// 		beego.Error(err)
	// 	} else if obj == nil {

	// 	} else {
	// 		obj0, err := getWxTokenOauth(obj.Appid, obj.Secret, code)
	// 		if err != nil {
	// 			beego.Error(err)
	// 		} else {
	// 			beego.Debug("getWxTokenOauth :", obj0)
	// 			if obj0.ErrCode == 0 {
	// 				obj1, err := getWxUserOauth(obj0.OpenID, obj0.AccessToken)
	// 				if err != nil {
	// 					beego.Error(err)
	// 				} else {
	// 					obj2, err := models.AddPuser(obj1, obj.Appid, obj.Secret)
	// 					if err != nil {
	// 						beego.Error(err)
	// 					} else {
	// 						//授权成功
	// 						c.Data["Parameter"] = true
	// 						c.Data["User"] = obj2
	// 					}

	// 				}
	// 			} else {

	// 			}
	// 		}
	// 	}
	// }
	savePhotoOpenId(c, "o3AhEuBQDmU1BE77UQREd8Z-9F44")
	openid, err := getPhotoOpenId(c.Ctx)
	obj2, err := models.GetPuserFromOpenId(openid)
	if err != nil {
		beego.Error(err)
	}
	c.Data["Parameter"] = true
	c.Data["User"] = obj2

	objs, err := models.GetAllPhotos(openid)
	if err != nil {
		beego.Error(err)
	}
	// beego.Debug("objs:", len(objs))
	// photo4 := models.Photos4{}
	// photo4s := []models.Photos4{}
	// for i := 0; i < len(objs); i++ {
	// 	if i%4 == 0 {
	// 		photo4 = models.Photos4{}
	// 	}
	// 	if i%4 == 0 {
	// 		photo4.Id0 = objs[i].Id
	// 		photo4.Image0 = objs[i].Image
	// 	}
	// 	if i%4 == 1 {
	// 		photo4.Id1 = objs[i].Id
	// 		photo4.Image1 = objs[i].Image
	// 	}
	// 	if i%4 == 2 {
	// 		photo4.Id2 = objs[i].Id
	// 		photo4.Image2 = objs[i].Image
	// 	}
	// 	if i%4 == 3 {
	// 		photo4.Id3 = objs[i].Id
	// 		photo4.Image3 = objs[i].Image
	// 	}
	// 	if i%4 == 3 {
	// 		photo4s = append(photo4s, photo4)
	// 	}
	// 	if i == len(objs)-1 && i%4 != 3 {
	// 		photo4s = append(photo4s, photo4)
	// 	}
	// }
	// beego.Debug("len photo4s", len(photo4s))
	c.Data["Objs"] = objs
	beego.Debug(c.Data["Parameter"])
	c.TplName = "phome.html"
}
func (c *PhotoController) Select() {
	if c.Ctx.Input.IsGet() {
		beego.Debug("Home Get")
	}
	if c.Ctx.Input.IsPost() {
		beego.Debug("Home Post")
	}
	openid, err := getPhotoOpenId(c.Ctx)
	if err != nil {
		beego.Error(err)
	}
	objs, err := models.GetAllPhotos(openid)
	if err != nil {
		beego.Error(err)
	}
	sphotos := []models.SPhotos{}
	for i := 0; i < len(objs); i++ {
		sp := models.SPhotos{}
		sp.Id = objs[i].Id
		sp.OpenId = objs[i].OpenId
		sp.Image = objs[i].Image
		sp.Select = false
		sphotos = append(sphotos, sp)
	}
	c.Data["Objs"] = sphotos
	c.Data["Parameter"] = true

	sizes, err := models.GetAllPsizeState1()
	if err != nil {
		beego.Error(err)
	}
	c.Data["Sizes"] = sizes
	temps, err := models.GetAllPtempState1()
	if err != nil {
		beego.Error(err)
	}
	c.Data["Temps"] = temps
	c.TplName = "pselect.html"
}
func (c *PhotoController) Upload() {
	if c.Ctx.Input.IsGet() {
		beego.Debug("Upload Get")
	}
	if c.Ctx.Input.IsPost() {
		beego.Debug("Upload Post")
	}
	c.TplName = "pupload.html"
}

func (c *PhotoController) My() {
	if c.Ctx.Input.IsGet() {
		beego.Debug("My Get")
	}
	if c.Ctx.Input.IsPost() {
		beego.Debug("My Post")
	}

	obj2, err := models.GetPuserFromOpenId("o3AhEuBQDmU1BE77UQREd8Z-9F44")
	if err != nil {
		beego.Error(err)
	}
	c.Data["User"] = obj2
	c.TplName = "pmy.html"
}

func (c *PhotoController) Order() {
	if c.Ctx.Input.IsGet() {
		beego.Debug("Order Get")
	}
	if c.Ctx.Input.IsPost() {
		beego.Debug("Order Post")
	}
	openid, err := getPhotoOpenId(c.Ctx)
	if err != nil {
		beego.Error(err)
	}
	objs, err := models.GetMyAllPorder(openid)
	c.Data["Objs"] = objs
	c.TplName = "porder.html"
}

func (c *PhotoController) OrderDet() {
	if c.Ctx.Input.IsGet() {
		beego.Debug("OrderDet Get")
	}
	if c.Ctx.Input.IsPost() {
		beego.Debug("OrderDet Post")
	}
	id := c.Input().Get("id")
	beego.Debug("id:", id)
	if len(id) != 0 {
		porder, err := models.GetPorder(id)
		if err != nil {
			beego.Error(err)
		}
		c.Data["Porder"] = porder
		beego.Debug("Porder:", porder)
		pdetails, err := models.GetPdetails(porder.OpenId, porder.Pnumber)
		if err != nil {
			beego.Error(err)
		}
		c.Data["Pdetails"] = pdetails
		beego.Debug("Pdetails:", pdetails)
	}

	c.TplName = "porderdet.html"
}

/*********admin***********/
func (c *PhotoController) AdminHome() {
	bool, username := chackAccount(c.Ctx)
	if bool {

	} else {
		c.Redirect("/admin", 302)
		return
	}
	c.Data["isUser"] = bool
	c.Data["User"] = username
	c.TplName = "padmin.html"
}

func (c *PhotoController) AdminUrl() {
	bool, username := chackAccount(c.Ctx)
	if bool {

	} else {
		c.Redirect("/admin", 302)
		return
	}
	c.Data["isUser"] = bool
	c.Data["User"] = username
	c.TplName = "padminurl.html"
}

func (c *PhotoController) AdminOauth() {
	bool, username := chackAccount(c.Ctx)
	if bool {

	} else {
		c.Redirect("/admin", 302)
		return
	}
	c.Data["isUser"] = bool
	c.Data["User"] = username
	objs, err := models.GetPoauth()
	if err != nil {
		beego.Error(err)
	}

	c.Data["Objs"] = objs
	beego.Debug("objs:", objs)
	c.TplName = "padminoauth.html"
}

func (c *PhotoController) Poauthtest() {
	state := c.Input().Get("state")
	code := c.Input().Get("code")
	c.Data["Code"] = code
	c.Data["State"] = state
	if c.Ctx.Input.IsGet() {
		beego.Debug("PhotoController Get")
	}
	if c.Ctx.Input.IsPost() {
		beego.Debug("PhotoController Post")
	}
	c.TplName = "poauthtest.html"
}

func (c *PhotoController) AdminSize() {
	bool, username := chackAccount(c.Ctx)
	if bool {

	} else {
		c.Redirect("/admin", 302)
		return
	}
	c.Data["isUser"] = bool
	c.Data["User"] = username
	objs, err := models.GetAllPsize()
	if err != nil {
		beego.Error(err)
	}

	c.Data["Objs"] = objs
	beego.Debug("objs:", objs)
	c.TplName = "padminsize.html"
}
func (c *PhotoController) AdminTemp() {
	bool, username := chackAccount(c.Ctx)
	if bool {

	} else {
		c.Redirect("/admin", 302)
		return
	}
	c.Data["isUser"] = bool
	c.Data["User"] = username
	objs, err := models.GetAllPtemp()
	if err != nil {
		beego.Error(err)
	}

	c.Data["Objs"] = objs
	beego.Debug("objs:", objs)
	c.TplName = "padmintemp.html"
}

func (c *PhotoController) AdminOrders() {
	bool, username := chackAccount(c.Ctx)
	if bool {

	} else {
		c.Redirect("/admin", 302)
		return
	}
	c.Data["isUser"] = bool
	c.Data["User"] = username
	objes, err := models.GetAdminPorder()
	if err != nil {
		beego.Error(err)
	}
	c.Data["Objs"] = objes
	c.TplName = "padminorders.html"
}

func (c *PhotoController) AdminDetails() {
	bool, username := chackAccount(c.Ctx)
	if bool {

	} else {
		c.Redirect("/admin", 302)
		return
	}
	c.Data["isUser"] = bool
	c.Data["User"] = username
	id := c.Input().Get("id")
	beego.Debug("id:", id)
	if len(id) != 0 {
		porder, err := models.GetPorder(id)
		if err != nil {
			beego.Error(err)
		}
		c.Data["Porder"] = porder
		beego.Debug("Porder:", porder)
		pdetails, err := models.GetPdetails(porder.OpenId, porder.Pnumber)
		if err != nil {
			beego.Error(err)
		}
		c.Data["Pdetails"] = pdetails
		psize, err := models.GetOnePsize(porder.PsizeId)
		if err != nil {
			beego.Error(err)
		}
		c.Data["Psize"] = psize

		ptemp, err := models.GetOnePtemp(porder.PtempId)
		if err != nil {
			beego.Error(err)
		}
		c.Data["Ptemp"] = ptemp
		beego.Debug("Ptemp:", ptemp)
		var photos_i []int64
		// beego.Debug(photos)
		err = json.Unmarshal([]byte(porder.Photos), &photos_i)
		if err != nil {
			beego.Error(err)
		}
		var photoss []*models.Photos
		for i := 0; i < len(photos_i); i++ {
			photos, err := models.GetOnePhotos(photos_i[i])
			if err != nil {
				beego.Error(err)

			}
			photoss = append(photoss, photos)
		}
		c.Data["Photos"] = photoss

	}
	c.TplName = "padmindetails.html"
}

/************post请求************/
func (c *PhotoController) PostUrl() {
	beego.Debug(c.Input())
	request_json := `{"errcode":1,"errmsg":"request_json error"}`
	if c.Ctx.Input.IsGet() {
		beego.Debug("Adminpolls Get")
	}
	if c.Ctx.Input.IsPost() {
		beego.Debug("PostAdmin Post")
		appid := c.Input().Get("appid")
		redirect_uri := c.Input().Get("redirect_uri")
		if len(appid) != 0 && len(redirect_uri) != 0 {

			obj, err := models.GetPoauthFromAppid(appid)
			if err != nil {
				beego.Error(err)
				request_json = `{"errcode":1,"errmsg":"getpoauth error"}`
			} else if obj == nil {
				request_json = `{"errcode":1,"errmsg":"getpoauth nil"}`
			} else {
				url := "[REALM]?appid=[APPID]&redirect_uri=[REDIRECT_URI]&response_type=code&scope=snsapi_userinfo&state=[STATE]#wechat_redirect"
				realm_name := "https://open.weixin.qq.com/connect/oauth2/authorize"
				redirect_uri = urlEncode(redirect_uri)
				url = strings.Replace(url, "[REALM]", realm_name, -1)
				url = strings.Replace(url, "[APPID]", appid, -1)
				url = strings.Replace(url, "[REDIRECT_URI]", redirect_uri, -1)
				url = strings.Replace(url, "[STATE]", fmt.Sprintf("%d", obj.Id), -1)
				request_json = fmt.Sprintf(`{"errcode":0,"errmsg":"","data":"%s"}`, url)
			}

		} else {
			request_json = `{"errcode":1,"errmsg":"parameter error"}`
		}
		c.Ctx.WriteString(request_json)
		return
	}

	c.Ctx.WriteString(request_json)
}

func (c *PhotoController) PostOauth() {
	beego.Debug(c.Input())
	request_json := `{"errcode":1,"errmsg":"request_json error"}`
	if c.Ctx.Input.IsGet() {
		beego.Debug("PostOauth Get")
	}
	if c.Ctx.Input.IsPost() {
		beego.Debug("PostOauth Post")
		op := c.Input().Get("op")
		beego.Debug("PostOauth op", op)
		if len(op) != 0 {
			switch op {
			case "add":
				appid := c.Input().Get("appid")
				secret := c.Input().Get("secret")
				if len(appid) != 0 && len(secret) != 0 {
					id, err := models.AddPoauth(appid, secret)
					if err != nil {
						beego.Debug(err)
						request_json = `{"errcode":1,"errmsg":"addpoauth error"}`
					} else {
						request_json = fmt.Sprintf(`{"errcode":0,"errmsg":"","data":"%d"}`, id)
					}
				} else {
					request_json = `{"errcode":1,"errmsg":"addpoauth parameter error"}`
				}

				break
			case "del":
				id := c.Input().Get("id")
				if len(id) != 0 {
					err := models.DelPoauth(id)
					if err != nil {
						beego.Debug(err)
						request_json = `{"errcode":1,"errmsg":"addpoauth error"}`
					} else {
						request_json = fmt.Sprintf(`{"errcode":0,"errmsg":"","data":"%d"}`, id)
					}
				} else {
					request_json = `{"errcode":1,"errmsg":"delpoauth parameter error"}`
				}

				break
			}

		} else {
			request_json = `{"errcode":1,"errmsg":"parameter error"}`
		}
		c.Ctx.WriteString(request_json)
		return
	}

	c.Ctx.WriteString(request_json)
}

func (c *PhotoController) PostPsize() {
	beego.Debug(c.Input())
	request_json := `{"errcode":1,"errmsg":"request_json error"}`
	if c.Ctx.Input.IsGet() {
		beego.Debug("PostPsize Get")
	}
	if c.Ctx.Input.IsPost() {
		beego.Debug("PostPsize Post")
		op := c.Input().Get("op")
		// beego.Debug("PostAdmin op", op)
		if len(op) != 0 {
			switch op {
			case "add":
				title := c.Input().Get("title")
				money := c.Input().Get("money")
				if len(title) != 0 && len(money) != 0 {
					id, err := models.AddPsize(title, money)
					if err != nil {
						beego.Debug(err)
						request_json = `{"errcode":1,"errmsg":"addpoauth error"}`
					} else {
						request_json = fmt.Sprintf(`{"errcode":0,"errmsg":"","data":"%d"}`, id)
					}
				} else {
					request_json = `{"errcode":1,"errmsg":"addpoauth parameter error"}`
				}

				break
			case "del":
				id := c.Input().Get("id")
				if len(id) != 0 {
					err := models.DelPsize(id)
					if err != nil {
						beego.Debug(err)
						request_json = `{"errcode":1,"errmsg":"addpoauth error"}`
					} else {
						request_json = fmt.Sprintf(`{"errcode":0,"errmsg":"","data":"%d"}`, id)
					}
				} else {
					request_json = `{"errcode":1,"errmsg":"delpoauth parameter error"}`
				}

				break
			case "state0":
				id := c.Input().Get("id")
				beego.Debug("id", id)
				if len(id) != 0 {
					err := models.UpdatePsize(id, 1)
					if err != nil {
						beego.Debug(err)
						request_json = `{"errcode":1,"errmsg":"addpoauth error"}`
					} else {
						request_json = fmt.Sprintf(`{"errcode":0,"errmsg":"","data":"%d"}`, id)
					}
				} else {
					request_json = `{"errcode":1,"errmsg":"delpoauth parameter error"}`
				}

				break
			case "state1":
				id := c.Input().Get("id")
				if len(id) != 0 {
					err := models.UpdatePsize(id, 0)
					if err != nil {
						beego.Debug(err)
						request_json = `{"errcode":1,"errmsg":"addpoauth error"}`
					} else {
						request_json = fmt.Sprintf(`{"errcode":0,"errmsg":"","data":"%d"}`, id)
					}
				} else {
					request_json = `{"errcode":1,"errmsg":"delpoauth parameter error"}`
				}

				break
			}

		} else {
			request_json = `{"errcode":1,"errmsg":"parameter error"}`
		}
		c.Ctx.WriteString(request_json)
		return
	}

	c.Ctx.WriteString(request_json)
}

func (c *PhotoController) PostPtemp() {
	// beego.Debug(c.Input())
	request_json := `{"errcode":1,"errmsg":"request_json error"}`
	if c.Ctx.Input.IsGet() {
		beego.Debug("PostPtemp Get")
	}
	if c.Ctx.Input.IsPost() {
		beego.Debug("PostPtemp Post")
		op := c.Input().Get("op")
		// beego.Debug("PostAdmin op", op)
		if len(op) != 0 {
			switch op {
			case "add":
				title := c.Input().Get("title")
				money := c.Input().Get("money")
				image := c.Input().Get("image")
				if len(title) != 0 && len(money) != 0 {
					id, err := models.AddPtemp(title, image, money)
					if err != nil {
						beego.Debug(err)
						request_json = `{"errcode":1,"errmsg":"addpoauth error"}`
					} else {
						request_json = fmt.Sprintf(`{"errcode":0,"errmsg":"","data":"%d"}`, id)
					}
				} else {
					request_json = `{"errcode":1,"errmsg":"addpoauth parameter error"}`
				}

				break
			case "del":
				id := c.Input().Get("id")
				if len(id) != 0 {
					err := models.DelPtemp(id)
					if err != nil {
						beego.Debug(err)
						request_json = `{"errcode":1,"errmsg":"addpoauth error"}`
					} else {
						request_json = fmt.Sprintf(`{"errcode":0,"errmsg":"","data":"%d"}`, id)
					}
				} else {
					request_json = `{"errcode":1,"errmsg":"delpoauth parameter error"}`
				}

				break
			case "state0":
				id := c.Input().Get("id")
				beego.Debug("id", id)
				if len(id) != 0 {
					err := models.UpdatePtemp(id, 1)
					if err != nil {
						beego.Debug(err)
						request_json = `{"errcode":1,"errmsg":"addpoauth error"}`
					} else {
						request_json = fmt.Sprintf(`{"errcode":0,"errmsg":"","data":"%d"}`, id)
					}
				} else {
					request_json = `{"errcode":1,"errmsg":"delpoauth parameter error"}`
				}

				break
			case "state1":
				id := c.Input().Get("id")
				if len(id) != 0 {
					err := models.UpdatePtemp(id, 0)
					if err != nil {
						beego.Debug(err)
						request_json = `{"errcode":1,"errmsg":"addpoauth error"}`
					} else {
						request_json = fmt.Sprintf(`{"errcode":0,"errmsg":"","data":"%d"}`, id)
					}
				} else {
					request_json = `{"errcode":1,"errmsg":"delpoauth parameter error"}`
				}

				break
			}

		} else {
			request_json = `{"errcode":1,"errmsg":"parameter error"}`
		}
		c.Ctx.WriteString(request_json)
		return
	}

	c.Ctx.WriteString(request_json)
}

func (c *PhotoController) PostAddOrder() {
	beego.Debug(c.Input())
	request_json := `{"errcode":1,"errmsg":"request_json error"}`
	if c.Ctx.Input.IsGet() {
		beego.Debug("PostAddOrder Get")
	}
	if c.Ctx.Input.IsPost() {
		beego.Debug("PostAddOrder Post")
		size_id := c.Input().Get("size_id")
		temp_id := c.Input().Get("temp_id")
		photos := c.Input().Get("photos")
		//json 到 []string
		var photos_i []int64
		beego.Debug(photos)
		err := json.Unmarshal([]byte(photos), &photos_i)
		if err != nil {
			beego.Error(photos_i)
			request_json = `{"errcode":1,"errmsg":"to json error"}`
			c.Ctx.WriteString(request_json)
			return
		}
		order_number := getOrderNumber(size_id, temp_id, len(photos_i))
		beego.Debug("order_number:", order_number)
		openid, err := getPhotoOpenId(c.Ctx)
		if err != nil {
			beego.Error(err)
			request_json = `{"errcode":1,"errmsg":"getPhotoOpenId error"}`
			c.Ctx.WriteString(request_json)
			return
		}
		_, err = models.AddPorder(openid, order_number, photos, temp_id, size_id)
		if err != nil {
			beego.Error(err)
			request_json = `{"errcode":1,"errmsg":"AddPorder error"}`
			c.Ctx.WriteString(request_json)
			return
		}
		_, err = models.AddPdetails(openid, order_number, 1, "订单已经生成，等待客服处理")
		if err != nil {
			beego.Error(err)
			request_json = `{"errcode":1,"errmsg":"AddPorder error"}`
			c.Ctx.WriteString(request_json)
			return
		}
		request_json = `{"errcode":0,"errmsg":"","data":""}`
	}
	c.Ctx.WriteString(request_json)
}

func (c *PhotoController) PostAdminUpState() {
	beego.Debug(c.Input())
	request_json := `{"errcode":1,"errmsg":"request_json error"}`
	if c.Ctx.Input.IsGet() {
		beego.Debug("PostAadminUpState Get")
	}
	if c.Ctx.Input.IsPost() {
		beego.Debug("PostAadminUpState Post")
		id := c.Input().Get("id")
		state := c.Input().Get("state")
		if len(id) == 0 || len(state) == 0 {
			request_json = `{"errcode":1,"errmsg":"id or state error"}`
			c.Ctx.WriteString(request_json)
			return
		}
		porder, err := models.GetPorder(id)
		if err != nil {
			beego.Error(err)
			request_json = `{"errcode":1,"errmsg":"GetPorder error"}`
			c.Ctx.WriteString(request_json)
			return
		}

		otype := int8(0)
		det := "类型错误"
		switch state {
		case "1":
			otype = int8(1)
			det = "订单已经生成，等待客服处理"
			break
		case "2":
			otype = int8(2)
			det = "客服确认订单"
			break
		case "3":
			otype = int8(3)
			det = "客服发送物流"
			break
		case "4":
			otype = int8(4)
			det = "订单完成"
			break
		}
		if otype == int8(0) {
			request_json = `{"errcode":1,"errmsg":"otype error"}`
			c.Ctx.WriteString(request_json)
			return
		}
		err = models.UpPorderState(id, state)
		if err != nil {
			beego.Error(err)
			request_json = `{"errcode":1,"errmsg":"UpPorderState error"}`
			c.Ctx.WriteString(request_json)
			return
		}
		_, err = models.AddPdetails(porder.OpenId, porder.Pnumber, otype, det)
		if err != nil {
			beego.Error(err)
			request_json = `{"errcode":1,"errmsg":"AddPorder error"}`
			c.Ctx.WriteString(request_json)
			return
		}
		request_json = fmt.Sprintf(`{"errcode":0,"errmsg":"","data":%d}`, otype)
		beego.Debug("request_json :", request_json)
	}
	c.Ctx.WriteString(request_json)
}

func urlEncode(urlin string) string {
	return url.QueryEscape(urlin)
}

func getPhotoOpenId(ctx *context.Context) (string, error) {
	ck, err := ctx.Request.Cookie("photo_open_id")
	if err != nil {
		return "", err
	}
	open_id := ck.Value
	return open_id, nil
}

func savePhotoOpenId(c *PhotoController, openid string) {
	maxAge := 1<<31 - 1
	c.Ctx.SetCookie("photo_open_id", openid, maxAge, "/")
}
