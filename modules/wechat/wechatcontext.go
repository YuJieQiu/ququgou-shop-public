package wechat

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"github.com/silenceper/wechat/cache"
	"io/ioutil"
	"net/http"
	"sync"
)

// Context struct
type WechatContext struct {
	AppID          string
	AppSecret      string
	Token          string
	EncodingAESKey string
	PayMchID       string
	PayNotifyURL   string
	PayKey         string

	Cache cache.Cache

	Writer  http.ResponseWriter
	Request *http.Request

	//accessTokenLock 读写锁 同一个AppID一个
	accessTokenLock *sync.RWMutex

	//jsAPITicket 读写锁 同一个AppID一个
	jsAPITicketLock *sync.RWMutex

	//accessTokenFunc 自定义获取 access token 的方法
	//accessTokenFunc GetAccessTokenFunc
}

//PostXML perform a HTTP/POST request with XML body
func PostXML(uri string, obj interface{}) ([]byte, error) {

	xmlData, err := xml.MarshalIndent(obj, "", "  ")
	if err != nil {
		return nil, err
	}

	//	testSrt := `<xml>
	//	<appid><![CDATA[wxe0e073c7992c2856]]></appid>
	//	<body><![CDATA[Task-jack]]></body>
	//	<mch_id><![CDATA[1501636421]]></mch_id>
	//	<nonce_str><![CDATA[L9qTalK3gqOvLIibODIoI2Jpd25pG1OP]]></nonce_str>
	//	<notify_url><![CDATA[https://main.ququgo.club/user/main/api/v1/pay]]></notify_url>
	//	<openid><![CDATA[oHqJO5S8FFidwgJUlxbodeSKwv54]]></openid>
	//	<out_trade_no><![CDATA[3191122642306957488d2]]></out_trade_no>
	//	<spbill_create_ip><![CDATA[123.12.12.123]]></spbill_create_ip>
	//	<total_fee><![CDATA[3899]]></total_fee>
	//	<trade_type><![CDATA[JSAPI]]></trade_type>
	//	<sign>386CEFDDFF91EB336D0DDC165564B290</sign>
	//</xml>`

	//bt := []byte(testSrt)

	fmt.Println(string(xmlData))

	body := bytes.NewBuffer(xmlData)

	response, err := http.Post(uri, "application/xml;charset=utf-8", body)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("http code error : uri=%v , statusCode=%v", uri, response.StatusCode)
	}
	return ioutil.ReadAll(response.Body)
}
