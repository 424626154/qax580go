package models

import (
	"strconv"
	"time"

	"github.com/astaxie/beego/orm"
)

type WeChat struct {
	Id        int64
	Name      string `orm:"size(100)"`  //名称
	Introduce string `orm:"size(1000)"` //简介
	Number    string `orm:"size(100)"`  //好吗
	Evaluate  string `orm:"size(1000)"` //评价
	Image     string
	IsHome    bool
	Time      int64
	State     int8 //0未上线  1 已上线
}

// 添加微信号
func AddWeChat(name, introduce, number, evaluate, image string) error {
	o := orm.NewOrm()
	my_time := time.Now().Unix()
	cate := &WeChat{Name: name, Introduce: introduce, Number: number, Evaluate: evaluate, Image: image, Time: my_time}
	// 查询数据
	qs := o.QueryTable("we_chat")
	err := qs.Filter("number", number).One(cate)
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

// 删除微信号
func DeleteWeChat(id string) error {
	cid, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return err
	}
	o := orm.NewOrm()
	cate := &WeChat{Id: cid}
	_, err = o.Delete(cate)
	return err
}

// 修改微信状态
func UpdateWeChatState(id string, state int8) error {
	cid, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return err
	}
	o := orm.NewOrm()
	obj := &WeChat{Id: cid}
	obj.State = state
	_, err = o.Update(obj, "state")
	return err
}

// 修改微信是否主页显示
func UpdateWeChatIsHome(id string, show bool) error {
	cid, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return err
	}
	o := orm.NewOrm()
	obj := &WeChat{Id: cid}
	obj.IsHome = show
	_, err = o.Update(obj, "is_home")
	return err
}

// 获得微信
func GetAdminWeChats() ([]WeChat, error) {
	o := orm.NewOrm()
	var wxnums []WeChat
	_, err := o.Raw("SELECT * FROM we_chat  ORDER BY id DESC").QueryRows(&wxnums)
	return wxnums, err
}

// 获得发布微信
func GetWeChats() ([]WeChat, error) {
	o := orm.NewOrm()
	var objs []WeChat
	_, err := o.Raw("SELECT * FROM we_chat  WHERE state = ? ORDER BY id DESC", 1).QueryRows(&objs)
	return objs, err
}

// 获得主页微信
func GetHomeWeChats() ([]WeChat, error) {
	o := orm.NewOrm()
	var objs []WeChat
	_, err := o.Raw("SELECT * FROM we_chat  WHERE state = ? AND is_home  = true ORDER BY id DESC", 1).QueryRows(&objs)
	return objs, err
}

// 根据id获得微信
func GetOneWeChat(id string) (*WeChat, error) {
	cid, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return nil, err
	}
	o := orm.NewOrm()
	obj := &WeChat{Id: cid}
	err = o.Read(obj)
	return obj, err
}

// 修改微信信息
func UpdateWeChatInfo(id string, name string, introduce string, number string, evaluate string) error {
	cid, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return err
	}
	o := orm.NewOrm()
	cate := &WeChat{Id: cid}
	cate.Name = name
	cate.Introduce = introduce
	cate.Number = number
	cate.Evaluate = evaluate
	_, err = o.Update(cate, "name", "introduce", "number", "evaluate")
	return err
}

// 修改微信图片
func UpdateWeChatImg(id string, img string) error {
	cid, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return err
	}
	o := orm.NewOrm()
	cate := &WeChat{Id: cid}
	cate.Image = img
	_, err = o.Update(cate, "image")
	return err
}
