package models

import (
	"strconv"
	"time"

	"github.com/astaxie/beego/orm"
)

type WxOfficial struct {
	Id        int64
	Title     string `orm:"size(100)"`
	Introduce string `orm:"size(1000)"`
	Number    string `orm:"size(100)"`
	Evaluate  string `orm:"size(1000)"` //评价
	Image     string
	IsHome    bool
	Time      int64
	State     int8 //0未上线  1 已上线
}

//添加微信公众号

func AddWxOfficial(title string, info string, num string, evaluate string, image string) error {
	o := orm.NewOrm()
	my_time := time.Now().Unix()
	cate := &WxOfficial{Title: title, Introduce: info, Number: num, Evaluate: evaluate, Image: image, Time: my_time, State: int8(0)}

	// 查询数据
	qs := o.QueryTable("wx_official")
	err := qs.Filter("title", title).One(cate)
	if err == nil {
		return err
	}

	// 插入数据
	_, err = o.Insert(cate)
	if err != nil {
		return err
	}

	return nil
}
func DeleteWxOfficial(id string) error {
	cid, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return err
	}
	o := orm.NewOrm()
	cate := &WxOfficial{Id: cid}
	_, err = o.Delete(cate)
	return err
}

func UpdateWxOfficial(id string, state int8) error {
	cid, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return err
	}
	o := orm.NewOrm()
	obj := &WxOfficial{Id: cid}
	obj.State = state
	_, err = o.Update(obj, "state")
	return err
}

func UpdateWxOfficialIsHome(id string, show bool) error {
	cid, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return err
	}
	o := orm.NewOrm()
	obj := &WxOfficial{Id: cid}
	obj.IsHome = show
	_, err = o.Update(obj, "is_home")
	return err
}

func GetAdminWxOfficials() ([]WxOfficial, error) {
	o := orm.NewOrm()
	var wxnums []WxOfficial
	_, err := o.Raw("SELECT * FROM wx_official  ORDER BY id DESC").QueryRows(&wxnums)
	return wxnums, err
}
func GetWxOfficials() ([]WxOfficial, error) {
	o := orm.NewOrm()
	var objs []WxOfficial
	_, err := o.Raw("SELECT * FROM wx_official  WHERE state = ? ORDER BY id DESC", 1).QueryRows(&objs)
	return objs, err
}

func GetHomeWxOfficials() ([]WxOfficial, error) {
	o := orm.NewOrm()
	var objs []WxOfficial
	_, err := o.Raw("SELECT * FROM wx_official  WHERE state = ? AND is_home = true ORDER BY id DESC", 1).QueryRows(&objs)
	return objs, err
}

func GetOneWxOfficial(id string) (*WxOfficial, error) {
	cid, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return nil, err
	}
	o := orm.NewOrm()
	obj := &WxOfficial{Id: cid}
	err = o.Read(obj)
	return obj, err
}

func UpdateWxOfficialnfo(id string, title string, info string, number string, evaluate string) error {
	cid, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return err
	}
	o := orm.NewOrm()
	cate := &WxOfficial{Id: cid}
	cate.Title = title
	cate.Introduce = info
	cate.Number = number
	cate.Evaluate = evaluate
	_, err = o.Update(cate, "title", "introduce", "number", "evaluate")
	return err
}

func UpdateWxOfficialImg(id string, img string) error {
	cid, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return err
	}
	o := orm.NewOrm()
	cate := &WxOfficial{Id: cid}
	cate.Image = img
	_, err = o.Update(cate, "image")
	return err
}
