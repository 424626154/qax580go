package admin

/*
后台修改商户内容
*/
import (
	"qax580go/models"
	"qax580go/qutil"

	"github.com/astaxie/beego"
)

type AdminUpShangHuInfoController struct {
	beego.Controller
}

func (c *AdminUpShangHuInfoController) Get() {
	id := c.Input().Get("id")
	obj, err := models.GetOneShanghu(id)
	if err != nil {
		beego.Error(err)
	}
	c.Data["Obj"] = obj
	c.TplName = "adminupshanghuinfo.html"
}
func (c *AdminUpShangHuInfoController) Post() {
	id := c.Input().Get("id")
	title := c.Input().Get("title")
	info := c.Input().Get("info")
	mtype := c.Input().Get("type")
	// sh_type := c.Input().Get("type")
	if len(id) != 0 && len(title) != 0 && len(info) != 0 {
		mtype_i := int16(0)
		typemap := make(map[string]int16)
		typemap[qutil.SH_CANYIN] = qutil.CANYIN_TYPE
		typemap[qutil.SH_QICHE] = qutil.QICHE_TYPE
		typemap[qutil.SH_WEIXIU] = qutil.WEIXIU_TYPE
		typemap[qutil.SH_PEIXUN] = qutil.PEIXUN_TYPE
		if v, ok := typemap[mtype]; ok {
			mtype_i = v
		} else {

		}
		err := models.UpdateShangHuInfo(id, title, info, mtype_i)
		if err != nil {
			beego.Error(err)
		}
		c.Redirect("/admin/shanghus", 302)
	}
	c.TplName = "adminupshanghuinfo.html"
}
