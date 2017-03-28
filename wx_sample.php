<?php
define("TOKEN", "qax580");
require_once("wx_config.php");
$wechatObj = new wechatCallbackapiTest();
if (!isset($_GET['echostr'])) {
    $wechatObj->responseMsg();
} else {
    $wechatObj->valid();
}

class wechatCallbackapiTest
{
    //验证签名
    public function valid()
    {
        $echoStr = $_GET["echostr"];
        $signature = $_GET["signature"];
        $timestamp = $_GET["timestamp"];
        $nonce = $_GET["nonce"];
        $token = TOKEN;
        $tmpArr = array($token, $timestamp, $nonce);
        sort($tmpArr);
        $tmpStr = implode($tmpArr);
        $tmpStr = sha1($tmpStr);
        if ($tmpStr == $signature) {
            echo $echoStr;
            exit;
        }
    }

    //响应消息
    public function responseMsg()
    {
        $postStr = $GLOBALS["HTTP_RAW_POST_DATA"];
        if (!empty($postStr)) {
            $this->logger("R " . $postStr);
            $postObj = simplexml_load_string($postStr, 'SimpleXMLElement', LIBXML_NOCDATA);
            $RX_TYPE = trim($postObj->MsgType);

            //消息类型分离
            switch ($RX_TYPE) {
                case "event":
                    $result = $this->receiveEvent($postObj);
                    break;
                case "text":
                    $result = $this->receiveText($postObj);
                    break;
                case "image":
                    $result = $this->receiveImage($postObj);
                    break;
                case "location":
                    $result = $this->receiveLocation($postObj);
                    break;
                case "voice":
                    $result = $this->receiveVoice($postObj);
                    break;
                case "video":
                    $result = $this->receiveVideo($postObj);
                    break;
                case "link":
                    $result = $this->receiveLink($postObj);
                    break;
                default:
                    $result = "unknown msg type: " . $RX_TYPE;
                    break;
            }
            $this->logger("T " . $result);
            echo $result;
        } else {
            echo "";
            exit;
        }
    }

    //接收事件消息
    private function receiveEvent($object)
    {
        $content = "";
        switch ($object->Event) {
            case "subscribe":
                $content = subscribe_text;
                $content .= (!empty($object->EventKey)) ? ("\n来自二维码场景 " . str_replace("qrscene_", "", $object->EventKey)) : "";
                break;
            case "unsubscribe":
                $content = unsubscribe_text;
                break;
            case "SCAN":
                $content = "扫描场景 " . $object->EventKey;
                break;
            case "CLICK":
                switch ($object->EventKey) {
                    case "Recommend_Key"://推荐信息
                        $content = $this->recommendInfo();
                        break;
                    case "Contact_Us_Key"://联系我们
                        $content = Contact_Us_Key_text;
                        break;
                    case "Search_Voice_Key_text"://语音搜索
                        $content = Search_Voice_Key_text;
                        break;
                    default:
                        $content = "点击菜单：" . $object->EventKey;
                        break;
                }
                break;
            case "LOCATION":
                $content = "上传位置：纬度 " . $object->Latitude . ";经度 " . $object->Longitude;
                break;
            case "VIEW":
                $content = "跳转链接 " . $object->EventKey;
                break;
            case "MASSSENDJOBFINISH":
                $content = "消息ID：" . $object->MsgID . "，结果：" . $object->Status . "，粉丝数：" . $object->TotalCount . "，过滤：" . $object->FilterCount . "，发送成功：" . $object->SentCount . "，发送失败：" . $object->ErrorCount;
                break;
            default:
                $content = "receive a new event: " . $object->Event;
                break;
        }
        if (is_array($content)) {
            if (isset($content[0])) {
                $result = $this->transmitNews($object, $content);
            } else if (isset($content['MusicUrl'])) {
                $result = $this->transmitMusic($object, $content);
            }
        } else {
            $result = $this->transmitText($object, $content);
        }

        return $result;
    }

    //接收文本消息
    private function receiveText($object)
    {
        $keyword = trim($object->Content);
        //多客服人工回复模式
        if (strstr($keyword, "您好") || strstr($keyword, "你好") || strstr($keyword, "在吗")) {
            $result = $this->transmitService($object);
        } //自动回复模式
        else {
            if (strstr($keyword, "文本")) {
                $content = "这是个文本消息";
            } else if (strstr($keyword, "单图文")) {
                $content = array();
                $content[] = array("Title" => "单图文标题", "Description" => "单图文内容", "PicUrl" => "http://discuz.comli.com/weixin/weather/icon/cartoon.jpg", "Url" => "http://www.baoguangguang.cn/58qax");
            } else if (strstr($keyword, "图文") || strstr($keyword, "多图文")) {
                $content = array();
                $content[] = array("Title" => "多图文1标题", "Description" => "", "PicUrl" => "http://discuz.comli.com/weixin/weather/icon/cartoon.jpg", "Url" => "http://www.baoguangguang.cn/58qax");
                $content[] = array("Title" => "多图文2标题", "Description" => "", "PicUrl" => "http://d.hiphotos.bdimg.com/wisegame/pic/item/f3529822720e0cf3ac9f1ada0846f21fbe09aaa3.jpg", "Url" => "http://www.baoguangguang.cn/58qax");
                $content[] = array("Title" => "多图文3标题", "Description" => "", "PicUrl" => "http://g.hiphotos.bdimg.com/wisegame/pic/item/18cb0a46f21fbe090d338acc6a600c338644adfd.jpg", "Url" => "http://www.baoguangguang.cn/58qax");
            } else if (strstr($keyword, "音乐")) {
                $content = array();
                $content = array("Title" => "兄弟", "Description" => "歌手：旭日阳刚", "MusicUrl" => "http://www.baoguangguang.cn/58qax/res/test01.mp3", "HQMusicUrl" => "http://www.baoguangguang.cn/58qax/res/test01.mp3");
            } else if (strstr($keyword, "二手")) {
                $content = $this->searchText($keyword);
            } else {
//                $content = date("Y-m-d H:i:s", time()) . "\n技术支持 漂泊80 输入您所需要搜索的信息，我们会为您择优选择";
                $content = "输入您所需要搜索的信息，我们会为您择优选择,如 '二手'\n技术支持 漂泊80 \n" .date("Y-m-d H:i:s", time());
            }

            if (is_array($content)) {
                if (isset($content[0]['PicUrl'])) {
                    $result = $this->transmitNews($object, $content);
                } else if (isset($content['MusicUrl'])) {
                    $result = $this->transmitMusic($object, $content);
                }
            } else {
                $result = $this->transmitText($object, $content);
            }
        }

        return $result;
    }

    //接收图片消息
    private function receiveImage($object)
    {
        $content = array("MediaId" => $object->MediaId);
        $result = $this->transmitImage($object, $content);
        return $result;
    }

    //接收位置消息
    private function receiveLocation($object)
    {
        $content = "你发送的是位置，纬度为：" . $object->Location_X . "；经度为：" . $object->Location_Y . "；缩放级别为：" . $object->Scale . "；位置为：" . $object->Label;
        // $result = $this->transmitText($object, $content);
        $result = $this->transmitNewsMap($object);
        return $result;
    }

    //接收语音消息
    private function receiveVoice($object)
    {
        if (isset($object->Recognition) && !empty($object->Recognition)) {
            $content = "你刚才说的是：" . $object->Recognition;
            $result = $this->transmitText($object, $content);
        } else {
            $content = array("MediaId" => $object->MediaId);
            $result = $this->transmitVoice($object, $content);
        }

        return $result;
    }

    //接收视频消息
    private function receiveVideo($object)
    {
        $content = array("MediaId" => $object->MediaId, "ThumbMediaId" => $object->ThumbMediaId, "Title" => "", "Description" => "");
        $result = $this->transmitVideo($object, $content);
        return $result;
    }

    //接收链接消息
    private function receiveLink($object)
    {
        $content = "你发送的是链接，标题为：" . $object->Title . "；内容为：" . $object->Description . "；链接地址为：" . $object->Url;
        $result = $this->transmitText($object, $content);
        return $result;
    }

    //回复文本消息
    private function transmitText($object, $content)
    {
        $xmlTpl = "<xml>
		<ToUserName><![CDATA[%s]]></ToUserName>
		<FromUserName><![CDATA[%s]]></FromUserName>
		<CreateTime>%s</CreateTime>
		<MsgType><![CDATA[text]]></MsgType>
		<Content><![CDATA[%s]]></Content>
		</xml>";
        $result = sprintf($xmlTpl, $object->FromUserName, $object->ToUserName, time(), $content);
        return $result;
    }

    //回复图片消息
    private function transmitImage($object, $imageArray)
    {
        $itemTpl = "<Image>
			<MediaId><![CDATA[%s]]></MediaId>
			</Image>";

        $item_str = sprintf($itemTpl, $imageArray['MediaId']);

        $xmlTpl = "<xml>
			<ToUserName><![CDATA[%s]]></ToUserName>
			<FromUserName><![CDATA[%s]]></FromUserName>
			<CreateTime>%s</CreateTime>
			<MsgType><![CDATA[image]]></MsgType>
			$item_str
			</xml>";

        $result = sprintf($xmlTpl, $object->FromUserName, $object->ToUserName, time());
        return $result;
    }

    //回复语音消息
    private function transmitVoice($object, $voiceArray)
    {
        $itemTpl = "<Voice>
			<MediaId><![CDATA[%s]]></MediaId>
		</Voice>";

        $item_str = sprintf($itemTpl, $voiceArray['MediaId']);

        $xmlTpl = "<xml>
		<ToUserName><![CDATA[%s]]></ToUserName>
		<FromUserName><![CDATA[%s]]></FromUserName>
		<CreateTime>%s</CreateTime>
		<MsgType><![CDATA[voice]]></MsgType>
		$item_str
		</xml>";

        $result = sprintf($xmlTpl, $object->FromUserName, $object->ToUserName, time());
        return $result;
    }

    //回复视频消息
    private function transmitVideo($object, $videoArray)
    {
        $itemTpl = "<Video>
			<MediaId><![CDATA[%s]]></MediaId>
			<ThumbMediaId><![CDATA[%s]]></ThumbMediaId>
			<Title><![CDATA[%s]]></Title>
			<Description><![CDATA[%s]]></Description>
		</Video>";

        $item_str = sprintf($itemTpl, $videoArray['MediaId'], $videoArray['ThumbMediaId'], $videoArray['Title'], $videoArray['Description']);

        $xmlTpl = "<xml>
		<ToUserName><![CDATA[%s]]></ToUserName>
		<FromUserName><![CDATA[%s]]></FromUserName>
		<CreateTime>%s</CreateTime>
		<MsgType><![CDATA[video]]></MsgType>
		$item_str
		</xml>";

        $result = sprintf($xmlTpl, $object->FromUserName, $object->ToUserName, time());
        return $result;
    }

    //回复图文消息
    private function transmitNews($object, $newsArray)
    {
        if (!is_array($newsArray)) {
            return;
        }
        $itemTpl = "    <item>
        <Title><![CDATA[%s]]></Title>
        <Description><![CDATA[%s]]></Description>
        <PicUrl><![CDATA[%s]]></PicUrl>
        <Url><![CDATA[%s]]></Url>
    </item>
";
        $item_str = "";
        foreach ($newsArray as $item) {
            $item_str .= sprintf($itemTpl, $item['Title'], $item['Description'], $item['PicUrl'], $item['Url']);
        }
        $xmlTpl = "<xml>
		<ToUserName><![CDATA[%s]]></ToUserName>
		<FromUserName><![CDATA[%s]]></FromUserName>
		<CreateTime>%s</CreateTime>
		<MsgType><![CDATA[news]]></MsgType>
		<ArticleCount>%s</ArticleCount>
		<Articles>
		$item_str</Articles>
		</xml>";

        $result = sprintf($xmlTpl, $object->FromUserName, $object->ToUserName, time(), count($newsArray));
        return $result;
    }

    //回复音乐消息
    private function transmitMusic($object, $musicArray)
    {
        $itemTpl = "<Music>
		<Title><![CDATA[%s]]></Title>
		<Description><![CDATA[%s]]></Description>
		<MusicUrl><![CDATA[%s]]></MusicUrl>
		<HQMusicUrl><![CDATA[%s]]></HQMusicUrl>
	</Music>";

        $item_str = sprintf($itemTpl, $musicArray['Title'], $musicArray['Description'], $musicArray['MusicUrl'], $musicArray['HQMusicUrl']);

        $xmlTpl = "<xml>
		<ToUserName><![CDATA[%s]]></ToUserName>
		<FromUserName><![CDATA[%s]]></FromUserName>
		<CreateTime>%s</CreateTime>
		<MsgType><![CDATA[music]]></MsgType>
		$item_str
		</xml>";

        $result = sprintf($xmlTpl, $object->FromUserName, $object->ToUserName, time());
        return $result;
    }

    //回复多客服消息
    private function transmitService($object)
    {
        $xmlTpl = "<xml>
		<ToUserName><![CDATA[%s]]></ToUserName>
		<FromUserName><![CDATA[%s]]></FromUserName>
		<CreateTime>%s</CreateTime>
		<MsgType><![CDATA[transfer_customer_service]]></MsgType>
		</xml>";
        $result = sprintf($xmlTpl, $object->FromUserName, $object->ToUserName, time());
        return $result;
    }
    //回复地图信息
    private function transmitNewsMap($object)
    {
        $xmlTpl = "<xml>
                            <ToUserName><![CDATA[%s]]></ToUserName>
                            <FromUserName><![CDATA[%s]]></FromUserName>
                            <CreateTime>%s</CreateTime>
                             <MsgType><![CDATA[news]]></MsgType>
                             <ArticleCount>1</ArticleCount>
                             <Articles>
                             <item>
                             <Title><![CDATA[庆安兄弟微盟]]></Title> 
                             <Description><![CDATA[按照地图标注来到庆安线580]]></Description>
                             <PicUrl><![CDATA[%s]]></PicUrl>
                             <Url><![CDATA[%s]]></Url>
                             </item>
                             </Articles>
                             <FuncFlag>1</FuncFlag>
                            </xml>";
        $j=$object->Location_X;
        $w=$object->Location_Y;
        $picurl="http://api.map.baidu.com/staticimage?width=340&height=160&center=116.725467,23.368905&zoom=16&markers=116.724964,23.36781|{$w},{$j}&markerStyles=l,M,0xFF0000|l,Y,0x008000";
        $url="http://api.map.baidu.com/staticimage?width=640&height=320&center=116.725467,23.368905&zoom=16&markers=116.724964,23.36781|{$w},{$j}&markerStyles=l,M,0xFF0000|l,Y,0x008000";
        $result = sprintf($xmlTpl,$object->FromUserName, $object->ToUserName, time(),$picurl,$url);
        return $result;
    }
    //搜索
    private function searchText($keyword){
        $content = array();
        if(empty($content) == true){
            if(strstr($keyword, "二手")){
                $content[] = array(title_key => "自行车", description_key => "九成新自行车", picurl_key => "http://www.baoguangguang.cn/58qax/res/test1.jpg", url_key => "http://www.baoguangguang.cn/qax580/recommend.html");
                $content[] = array(title_key => "家具", description_key => "用了一年", picurl_key => "http://www.baoguangguang.cn/58qax/res/test2.jpg", url_key => "http://www.baoguangguang.cn/58qax/recommend.html");
                $content[] = array(title_key => "飞鸽自行车", description_key => "自己的很珍惜", picurl_key => "http://www.baoguangguang.cn/58qax/res/test4.jpg", url_key => "http://www.baoguangguang.cn/58qax/recommend.html");
            }
        }else{

        }
        return $content;
    }

    /**
     * 推荐信息
     */
    private function recommendInfo(){
        $content = array();
        $content[] = array(title_key => "自行车", description_key => "九成新自行车", picurl_key => "http://www.baoguangguang.cn/58qax/res/test1.jpg", url_key => "http://www.baoguangguang.cn/qax580/recommend.html");
        $content[] = array(title_key => "家具", description_key => "用了一年", picurl_key => "http://www.baoguangguang.cn/58qax/res/test2.jpg", url_key => "http://www.baoguangguang.cn/58qax/recommend.html");
        $content[] = array(title_key => "飞鸽自行车", description_key => "自己的很珍惜", picurl_key => "http://www.baoguangguang.cn/58qax/res/test4.jpg", url_key => "http://www.baoguangguang.cn/58qax/recommend.html");
        return $content;
    }
    //日志记录
    private function logger($log_content)
    {
        if (isset($_SERVER['HTTP_APPNAME'])) {   //SAE
            sae_set_display_errors(false);
            sae_debug($log_content);
            sae_set_display_errors(true);
        } else if ($_SERVER['REMOTE_ADDR'] != "127.0.0.1") { //LOCAL
            $max_size = 10000;
            $log_filename = "log.xml";
            if (file_exists($log_filename) and (abs(filesize($log_filename)) > $max_size)) {
                unlink($log_filename);
            }
            file_put_contents($log_filename, date('H:i:s') . " " . $log_content . "\r\n", FILE_APPEND);
        }
    }
}

?>