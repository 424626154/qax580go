package controllers

import (
	"encoding/base64"
	"github.com/astaxie/beego"
	// "image"
	"io/ioutil"
	// "os"
	// "bytes"
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"qax580go/models"
	"strings"
	"time"
)

type ImageController struct {
	beego.Controller
}

func (c *ImageController) Upload() {
	if c.Ctx.Input.IsGet() {
		beego.Debug("ImageController Upload Get")
	}
	if c.Ctx.Input.IsPost() {
		beego.Debug("ImageController Upload Post")
	}
	request_json := `{"errcode":1,"errmsg":"parameter error"}`
	// beego.Debug(c.Input())
	// beego.Debug(c.GetFile())
	img_data := c.Input().Get("img_data")
	index := c.Input().Get("index")
	img_data1 := img_data[strings.Index(img_data, ",")+1 : len(img_data)]
	// img_type := img_data[strings.Index(img_data, "/")+1 : strings.Index(img_data, ";")]
	// beego.Debug("img_data1", img_data1)
	// beego.Debug("img_type", img_type)
	img_buffer, err := base64.StdEncoding.DecodeString(img_data1) //成图片文件并把文件写入到buffer
	if err != nil {
		beego.Error(err)
		request_json = fmt.Sprintf(`{"errcode":1,"errmsg":"%s","index":"%s"}`, "base64 DecodeString error", index)
		c.Ctx.WriteString(request_json)
		return
	}
	image_name := createImgName("image")
	path := fmt.Sprintf("./imageserver/%s", image_name)
	err = ioutil.WriteFile(path, img_buffer, 0666) //buffer输出到jpg文件中（不做处理，直接写到文件）
	if err != nil {
		beego.Error(err)
		request_json = fmt.Sprintf(`{"errcode":1,"errmsg":"%s","index":"%s"}`, "WriteFile error", index)
		c.Ctx.WriteString(request_json)
		return
	}
	request_json = fmt.Sprintf(`{"errcode":0,"errmsg":"","name":"%s","index":"%s"}`, image_name, index)
	beego.Debug("request_json", request_json)
	c.Ctx.WriteString(request_json)
}

func (c *ImageController) AddImg() {
	if c.Ctx.Input.IsGet() {
		beego.Debug("ImageController AddImg Get")
	}
	if c.Ctx.Input.IsPost() {
		beego.Debug("ImageController AddImg Post")
	}
	request_json := `{"errcode":1,"errmsg":"parameter error"}`
	// beego.Debug(c.Input())
	// beego.Debug(c.GetFile())
	images := c.Input().Get("images")
	beego.Debug("images:", images)
	if len(images) != 0 {
		// openid, err := getPhotoOpenId(c.Ctx)
		//json 到 []string
		// var wo []string
		// if err := json.Unmarshal(images, &wo); err == nil {
		// 	fmt.Println("================json 到 []string==")
		// 	fmt.Println(wo)
		// }
		//array 到 json str
		// arr := []string{"hello", "apple", "python", "golang", "base", "peach", "pear"}
		// lang, err := json.Marshal(arr)
		// if err == nil {
		// 	fmt.Println("================array 到 json str==")
		// 	fmt.Println(string(lang))
		// }

		//json 到 []string
		var json_images []string
		err := json.Unmarshal([]byte(images), &json_images)
		if err != nil {
			beego.Error(err)
		}
		beego.Debug(len(json_images))
		group := createImageGroupId(len(json_images))
		openid, err := getPhotoOpenId(c.Ctx)
		if err != nil {
			beego.Error(err)
			request_json = fmt.Sprintf(`{"errcode":1,"errmsg":%s}`, "AddImg getPhotoOpenId error")
			// `{"errcode":1,"errmsg":%s}`
			c.Ctx.WriteString(request_json)
			return
		}
		// beego.Debug(group)
		// beego.Debug(openid)
		for i := 0; i < len(json_images); i++ {
			_, err := models.AddPhotos(openid, group, json_images[i])
			if err != nil {
				beego.Error(err)
				request_json = fmt.Sprintf(`{"errcode":1,"errmsg":%s}`, "AddImg AddPhotos error")
				// `{"errcode":1,"errmsg":%s}`
				c.Ctx.WriteString(request_json)
				return
			}
		}
	}
	request_json = `{"errcode":0,"errmsg":"","data":""}`
	c.Ctx.WriteString(request_json)
}

func createImgName(name string) string {
	noncestr := getNonceStr(16, KC_RAND_KIND_ALL)
	t := time.Now().Unix()
	str2 := fmt.Sprintf("%d", t)
	s := []string{name, str2, noncestr}
	h := md5.New()
	h.Write([]byte(strings.Join(s, ""))) // 需要加密的字符串
	image_name := hex.EncodeToString(h.Sum(nil))
	beego.Debug("image_name :", image_name)
	return image_name
}

func createImageGroupId(num int) string {
	noncestr := getNonceStr(16, KC_RAND_KIND_ALL)
	t := time.Now().Unix()
	str2 := fmt.Sprintf("%d", t)
	nums := fmt.Sprintf("%d", num)
	s := []string{nums, str2, noncestr}
	h := md5.New()
	h.Write([]byte(strings.Join(s, ""))) // 需要加密的字符串
	group := hex.EncodeToString(h.Sum(nil))
	beego.Debug("group :", group)
	return group
}
