package routers

import (
	"os"
	"qax580go/controllers"
	"qax580go/controllers/admin"
	"qax580go/controllers/home"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.HomeController{})           //主页
	beego.Router("/uplode", &controllers.UplodeController{})   //发布消息
	beego.Router("/content", &controllers.ContentController{}) //消息详情
	beego.Router("/content/*", &controllers.ContentController{})
	beego.Router("/login", &controllers.LoginController{}) //登录
	beego.Router("/register", &controllers.RegisterController{})
	beego.Router("/registerverify", &controllers.RegisterVerifyController{}) //注册
	beego.Router("/usercenter", &controllers.UserCenterController{})
	beego.Router("/upuser", &controllers.UpUserController{})
	beego.Router("/forgetpwd", &controllers.ForgetpwdController{})
	beego.Router("/forgetpwdverify", &controllers.ForgetpwdVerifyController{})
	beego.Router("/wxofficials", &home.WxOfficial{}, "*:WxOfficials")                   //微推荐公众号列表
	beego.Router("/wx", &controllers.WXController{})                                    //微信公众号服务器
	beego.Router("/wxtl", &controllers.WXTlController{})                                //微信公众号服务器铁力
	beego.Router("/feedback", &controllers.FeedbackController{})                        //意见反馈
	beego.Router("/wxautho", &controllers.WxAuthoController{})                          //微信用户信息
	beego.Router("/wxhome", &controllers.WxHomeController{})                            //微信回调主页
	beego.Router("/wxuplode", &controllers.WxUplodeController{})                        //微信回调上传
	beego.Router("/wxfeedback", &controllers.WxFeedbackController{})                    //微信回调意见反馈
	beego.Router("/weather", &controllers.WeatherController{})                          //天气预报
	beego.Router("/traintickets", &controllers.TrainTicketsController{})                //火车票
	beego.Router("/querystation", &controllers.QueryStationController{})                //火车票起点终点查询
	beego.Router("/querytrain", &controllers.QueryTrainController{})                    //火车票车次查询
	beego.Router("/queryrealtime", &controllers.QueryRealTimeController{})              //火车票实时查询
	beego.Router("/queryqutlets", &controllers.QueryQutletsController{})                //火车票代售点查询
	beego.Router("/querypeccancy", &controllers.QueryPeccancyController{})              //违章查询
	beego.Router("/history", &controllers.HistoryController{})                          //历史今天
	beego.Router("/historycon", &controllers.HistoryConController{})                    //历史今天
	beego.Router("/laohuangli", &controllers.LaohuangliController{})                    //老黄历
	beego.Router("/zhoubianwifiwx", &controllers.ZhouBianWifiWXController{})            //周边Wi-Fi
	beego.Router("/kuaidi", &controllers.KuaidiController{})                            //快递查询
	beego.Router("/tianqiwx", &controllers.TianqiWXController{})                        //天气查询
	beego.Router("/recommend", &controllers.RecommendController{})                      //推荐
	beego.Router("/contactus", &controllers.ContactusController{})                      //联系我们
	beego.Router("updatelog", &controllers.UpdateLogController{})                       //更新日志
	beego.Router("/guanggaocontent", &controllers.GuanggaoContentController{})          //广告详情
	beego.Router("/waimailist", &controllers.WaimaiListController{})                    //外卖订餐
	beego.Router("/caidans", &controllers.CaidansController{})                          //菜单
	beego.Router("/wechats", &home.WeChat{}, "*:WeChats")                               //推荐微信号
	beego.Router("/about", &controllers.AboutController{})                              //关于
	beego.Router("/wxgame", &controllers.WeixinGameController{})                        //微信游戏
	beego.Router("/mymessage", &controllers.MyMessageController{})                      //我的发布
	beego.Router("/wxmymessage", &controllers.WxMyMessageController{})                  //微信回调我的发布
	beego.Router("/myextension", &controllers.MyExtensionController{})                  //我的推广
	beego.Router("/mymoney", &controllers.MymoneyController{})                          //我的帮帮币
	beego.Router("/mmyextensionresponse", &controllers.MyExtensionResponseController{}) //我的推广返回测试
	beego.Router("/moneyinfo", &controllers.MoneyInfoController{})                      //我的金钱详情
	beego.Router("/moneyhelp", &controllers.MoneyHelpController{})                      //我的金钱帮助
	beego.Router("/mall", &controllers.MallController{})                                //商城
	beego.Router("/exchange", &controllers.ExchangeController{})                        //兑换
	beego.Router("/subsribe", &controllers.SubsribeController{})                        //关注与取消关注
	beego.Router("/shanghus", &controllers.ShangHusController{})                        //商户列表
	beego.Router("/shanghulist", &controllers.ShangHuListController{})                  //商户子列表
	beego.Router("/mynotice", &controllers.MynoticeController{})                        //我的消息
	beego.Router("/posts", &home.Posts{}, "*:Posts")

	beego.AutoRouter(&controllers.WxqaxController{}) //微信http自动匹配

	// 附件处理
	os.Mkdir("imagehosting", os.ModePerm)
	beego.Router("/imagehosting/:all", &controllers.ImageHostingController{})

	os.Mkdir("imageserver", os.ModePerm)
	beego.Router("/imageserver/:all", &controllers.ImageHostingController{})

	os.Mkdir("filehosting", os.ModePerm)
	beego.Router("/filehosting/:all", &controllers.FileHostingController{})

	beego.Router("/admin/home", &admin.AdminHomeController{})                       //后台主页
	beego.Router("/admin/modify", &admin.AdminModifyController{})                   //修改信息
	beego.Router("/admin/uplode", &admin.AdminUplodeController{})                   //后台上传
	beego.Router("/admin/feedbacklist", &admin.AdminFeedbackListController{})       //意见反馈列表
	beego.Router("/admin/feedbackcontent", &admin.AdminFeedbackContentController{}) //意见反馈内容
	beego.Router("/admin", &admin.AdminLoginController{})                           //后台登陆
	beego.Router("/admin/adminlist", &admin.AdminAdminListController{})             //后台可登录用户列表
	beego.Router("/admin/userlist", &admin.AdminUserListController{})
	beego.Router("/admin/adduser", &admin.AdminAddUserController{})                 //添加后台用户
	beego.Router("/admin/content", &admin.AdminContentController{})                 //后台消息内容
	beego.Router("/admin/wxuserlist", &admin.WxUserListController{})                //后台统计公众号登录用户列表
	beego.Router("/admin/juhe", &admin.AdminJuheController{})                       //聚合数据主页
	beego.Router("/admin/newskey", &admin.AdminNewsKeyController{})                 //新闻关键词
	beego.Router("/admin/addguanggao", &admin.AdminaAddGuanggaoController{})        //后台添加广告
	beego.Router("/admin/guanggaos", &admin.AdminGuanggaosController{})             //后台广告列表
	beego.Router("/admin/guanggaocontent", &admin.AdminGuanggaoContentController{}) //后台广告内容
	beego.Router("/admin/upguanggaoinfo", &admin.AdminUpGuanggaoInfoController{})   //后台修改广告内容
	beego.Router("/admin/upguanggaoimg", &admin.AdminUpGuanggaoImgController{})     //后台修改广告图片
	beego.Router("/admin/waimailist", &admin.AdminWaimaiListController{})           //外卖列表
	beego.Router("/admin/addwaimai", &admin.AdminAddWaimaiController{})             //后台添加外卖
	beego.Router("/admin/caidans", &admin.AdminCaidansController{})                 //后台菜单列表
	beego.Router("/admin/addcaidan", &admin.AdminAddCaidanController{})             //后台添加菜单
	// admin推荐微信号
	beego.Router("/admin/wechats", &admin.AdminWeChat{}, "*:WeChats")       //后台推荐微信号列表
	beego.Router("/admin/addwechat", &admin.AdminWeChat{}, "*:Add")         //后台添加推荐微信号
	beego.Router("/admin/upwechatinfo", &admin.AdminWeChat{}, "*:UpInfo")   //后台修改微信号内容
	beego.Router("/admin/upwechatimg", &admin.AdminWeChat{}, "*:UpImg")     //后台修改微信号图片
	beego.Router("/admin/upusermoney", &admin.AdminUpUserMoneyController{}) //后台用户金钱
	// admin 推荐公众号
	beego.Router("/admin/wxofficials", &admin.AdminWxOfficial{}, "*:WxOfficials") //公众号列表
	beego.Router("/admin/addwxofficial", &admin.AdminWxOfficial{}, "*:Add")       //添加微信公众号
	beego.Router("/admin/upwxofficialinfo", &admin.AdminWxOfficial{}, "*:UpInfo") //修改微信公众号内容
	beego.Router("/admin/upwxofficialimg", &admin.AdminWxOfficial{}, "*:UpImg")   //修改微信公众号图片

	beego.Router("/admin/moneyinfo", &admin.AdminMoneyInfoController{})            //后台用户金钱记录
	beego.Router("admin/importuser", &admin.AdminImportUserController{})           //后台导入微信用户
	beego.Router("admin/upwxuserinfo", &admin.AdminUpWxuserInfoController{})       //后台导入微信用户
	beego.Router("admin/mall", &admin.AdminMallController{})                       //后台商城
	beego.Router("admin/addcommodity", &admin.AdminaAddCommodityController{})      //添加商品
	beego.Router("admin/upcommodityinfo", &admin.AdminUpCommodityInfoController{}) //修改商品信息
	beego.Router("admin/upcommodityimg", &admin.AdminUpCommodityImgController{})   //修改商品图片
	beego.Router("admin/exchange", &admin.AdminExchangeController{})               //用户兑换
	beego.Router("admin/shanghus", &admin.AdminShanghusController{})               //后台商户
	beego.Router("admin/addshanghu", &admin.AdminAddShanghuController{})           //添加商户
	beego.Router("admin/upshanghuinfo", &admin.AdminUpShangHuInfoController{})     //修改商户信息
	beego.Router("admin/upshanghuimg", &admin.AdminUpShangHuImgController{})       //修改商户图片
	beego.Router("admin/keywords", &admin.AdminKeywordsController{})               //关键字列表
	beego.Router("admin/addkeywords", &admin.AdminaAddKeywordsController{})        //添加关键字
	beego.Router("admin/keyobj", &admin.AdminKeyobjController{})                   //关键字内容
	beego.Router("admin/addkeyobj", &admin.AdminaAddKeyobjController{})            //添加关键字内容
	beego.Router("admin/wxtest", &admin.AdminWxTestController{})                   //添加关键字内容
	beego.Router("admin/updatelog", &admin.AdminUpdateLogController{})             //后台更新日志
	beego.Router("admin/notice", &admin.AdminNoticeController{})                   //后台通知管理

	beego.AutoRouter(&controllers.PollController{})  //投票系统
	beego.AutoRouter(&controllers.RinseController{}) //冲洗系统
	beego.AutoRouter(&controllers.WptController{})   //微信平台
	beego.AutoRouter(&admin.AdminPostController{})   //后台提交post
	beego.AutoRouter(&controllers.PhotoController{}) //洗相系统

	beego.AutoRouter(&controllers.ImageController{}) //图床

	beego.AutoRouter(&controllers.WeiZhanController{}) //微站
	beego.AutoRouter(&controllers.DqsjController{})    //大签世界

	beego.Router("/admin/dqsj", &admin.AdminDqsjUserListController{})
	beego.Router("/admin/adddqsjuser", &admin.AdminAddDqsjUserController{})

	beego.AutoRouter(&controllers.WxAppController{}) //微信小程序

	beego.AutoRouter(&controllers.BeerMapController{}) //大签世界

}
