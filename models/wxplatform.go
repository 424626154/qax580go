package models

import (
	"strconv"
	"time"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

type WxPlatform struct {
	Id         int64
	Title      string `orm:"size(1000)"` //标题
	Introduce  string `orm:"size(1000)"` //内容
	Wid        string //微信号
	Qrcode     string //二维码
	WRange     string //服务范围
	State      int16  //0 未上线 1上线
	Tuijian    int16  //推荐
	CreateTime int64  //创建时间
}

/***微平台数据***/
func AddWxPlatform(title string, info string, wid string, wrange string, qrcode string) error {
	o := orm.NewOrm()
	my_time := time.Now().Unix()
	obj := &WxPlatform{Title: title, Introduce: info, Wid: wid, Qrcode: qrcode, WRange: wrange, CreateTime: my_time}
	// 插入数据
	_, err := o.Insert(obj)
	if err != nil {
		return err
	}
	return nil
}

/**
获得所有
*/
func GetAdminWxPlatforms() ([]WxPlatform, error) {
	o := orm.NewOrm()
	var objs []WxPlatform
	_, err := o.QueryTable("wx_platform").OrderBy("-id").All(&objs)
	if err != nil {
		beego.Error(err)
	}
	return objs, err
}

/**
删除平台
*/
func DelWxPlatform(id string) error {
	cid, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return err
	}
	o := orm.NewOrm()
	cate := &WxPlatform{Id: cid}
	_, err = o.Delete(cate)
	return err
}

/**
修改平台状态
*/
func UpWxPlatformState(id string, state int16) error {
	cid, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return err
	}
	o := orm.NewOrm()
	cate := &WxPlatform{Id: cid}
	cate.State = state
	_, err = o.Update(cate, "state")
	if err != nil {
		beego.Error(err)
	}
	return err
}

/**
修改平台推荐
*/
func UpWxPlatformTuijian(id string, tuijian int16) error {
	cid, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return err
	}
	o := orm.NewOrm()
	cate := &WxPlatform{Id: cid}
	cate.Tuijian = tuijian
	_, err = o.Update(cate, "tuijian")
	if err != nil {
		beego.Error(err)
	}
	return err
}

/**
修改图片
*/
func UpWxPlatformImg(id string, qrcode string) error {
	cid, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return err
	}
	o := orm.NewOrm()
	cate := &WxPlatform{Id: cid}
	cate.Qrcode = qrcode
	_, err = o.Update(cate, "qrcode")
	return err
}

/**
获得平台
*/
func GetOneWxPlatform(id string) (*WxPlatform, error) {
	cid, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return nil, err
	}
	o := orm.NewOrm()
	obj := &WxPlatform{}
	err = o.QueryTable("wx_platform").Filter("id", cid).One(obj)
	if err != nil {
		beego.Error(err)
		return nil, err
	}
	return obj, err
}

/**
修改平台内容
*/
func UpWxPlatformInfo(id string, title string, info string, wid string, wrange string) error {
	cid, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return err
	}
	o := orm.NewOrm()
	cate := &WxPlatform{Id: cid}
	cate.Title = title
	cate.Introduce = info
	cate.Wid = wid
	cate.WRange = wrange
	_, err = o.Update(cate, "title", "introduce", "wid", "w_range")
	return err
}

/**
返回微信平台
*/
func GetWxPlatformTJ(tuijian int16) ([]WxPlatform, error) {
	o := orm.NewOrm()
	var objs []WxPlatform
	_, err := o.Raw("SELECT * FROM wx_platform WHERE state = 1  AND tuijian = ? ORDER BY id DESC", tuijian).QueryRows(&objs)
	return objs, err
}

/**
返回微信平台
*/
func GetWxPlatforms() ([]WxPlatform, error) {
	o := orm.NewOrm()
	var objs []WxPlatform
	_, err := o.Raw("SELECT * FROM wx_platform WHERE state = 1 ORDER BY id DESC").QueryRows(&objs)
	return objs, err
}

/**
返回微信平台 关键字
*/
func GetWxPlatformstLike(like string) ([]WxPlatform, error) {
	o := orm.NewOrm()
	var objs []WxPlatform
	_, err := o.Raw("SELECT * FROM wx_platform WHERE title LIKE ? OR wid = ? ORDER BY id DESC ", "%"+like+"%", "%"+like+"%").QueryRows(&objs)
	return objs, err
}
