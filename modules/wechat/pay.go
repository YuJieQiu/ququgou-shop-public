package wechat

import (
	"bufio"
	"bytes"
	"crypto/hmac"
	"crypto/md5"
	"crypto/sha256"
	"encoding/hex"
	"encoding/xml"
	"errors"
	"fmt"
	"hash"
	"math/rand"
	"sort"
	"strconv"
	"strings"
	"time"
)

const payGateway = "https://api.mch.weixin.qq.com/pay/unifiedorder"

// Params was NEEDED when request unifiedorder
// 传入的参数，用于生成 prepay_id 的必需参数
type PayParams struct {
	AppId          string `xml:"appid,cdata"`
	MchId          string `xml:"mch_id,cdata"`
	DeviceInfo     string `xml:"device_info,cdata"` //设备号 非必填
	NonceStr       string `xml:"nonce_str,cdata"`   //随机字符串
	Sign           string `xml:"sign"`
	OutTradeNo     string `xml:"out_trade_no,cdata"`     //商户订单号
	Body           string `xml:"body,cdata"`             //String(128) 商品描述 商品简单描述 例如：腾讯充值中心-QQ会员充值
	TotalFee       string `xml:"total_fee,cdata"`        //	Int 88 标价金额 订单总金额，单位为分
	SpbillCreateIP string `xml:"spbill_create_ip,cdata"` //String(64) 终端IP 支持IPV4和IPV6两种格式的IP地址。调用微信支付API的机器IP
	NotifyURL      string `xml:"notify_url,cdata"`       //String(256)通知地址 异步接收微信支付结果通知的回调地址，通知url必须为外网可访问的url，不能携带参数。
	TradeType      string `xml:"trade_type,cdata"`       //交易类型 JSAPI--JSAPI支付（或小程序支付）、NATIVE--Native支付、APP--app支付，MWEB--H5支付
	ProductId      string `xml:"product_id,cdata"`       //String(32)商品ID	此参数为二维码中包含的商品ID，商户自行定义。  trade_type=NATIVE时，此参数必传 否则 非必填
	OpenID         string `xml:"open_id,cdata"`          //用户标识 trade_type=JSAPI，此参数必传，用户在商户appid下的唯一标识
	PayKey         string `xml:"-"`                      //支付key
	SignType       string `xml:"sign_type"`              //签名类型 非必填 签名类型，默认为MD5，支持HMAC-SHA256和MD5。
	//Detail         string `xml:"detail"`           //String(6000) 商品详情 非必填
	//Attach         string `xml:"attach"`           //String(127) 附加数据 非必填 例如：	深圳分店
	//FeeType        string `xml:"fee_type"`         //	String(16) 标价币种 非必填 符合ISO 4217标准的三位字母代码，默认人民币：CNY
	//TimeStart      string `xml:"time_start"`       //交易起始时间  非必填 格式为yyyyMMddHHmmss
	//TimeExpire     string `xml:"time_expire"`      //交易结束时间  非必填 格式为yyyyMMddHHmmss
	//GoodsTag       string `xml:"goods_tag"`        //String(32)订单优惠标记 非必填
	//LimitPay       string `xml:"limit_pay"`        //指定支付方式  非必填 上传此参数no_credit--可限制用户不能使用信用卡支付
	//Receipt        string `xml:"receipt"`          //电子发票入口开放标识 非必填 Y，传入Y时，支付成功消息和支付详情页将出现开票入口。需要在微信支付商户平台或微信公众平台开通电子发票功能，传此字段才可生效
	//SceneInfo      string `xml:"scene_info"`       //场景信息  非必填 该字段常用于线下活动时的场景信息上报，支持上报实际门店信息，商户也可以按需求自己上报相关信息。该字段为JSON对象数据，对象格式为{"store_info":{"id": "门店ID","name": "名称","area_code": "编码","address": "地址" }}
}

// Config 是传出用于 jsdk 用的参数
type Config struct {
	Timestamp int64
	NonceStr  string
	PrePayID  string
	SignType  string
	Sign      string
}

// PreOrder 是 unifie order 接口的返回
type PreOrder struct {
	ReturnCode string `xml:"return_code"`
	ReturnMsg  string `xml:"return_msg"`
	AppID      string `xml:"appid,omitempty"`
	MchID      string `xml:"mch_id,omitempty"`
	NonceStr   string `xml:"nonce_str,omitempty"`
	Sign       string `xml:"sign,omitempty"`
	ResultCode string `xml:"result_code,omitempty"`
	TradeType  string `xml:"trade_type,omitempty"`
	PrePayID   string `xml:"prepay_id,omitempty"` //预支付交易会话标识
	CodeURL    string `xml:"code_url,omitempty"`
	ErrCode    string `xml:"err_code,omitempty"`
	ErrCodeDes string `xml:"err_code_des,omitempty"`
}

type CDATA struct {
	Text string `xml:",cdata"`
}

//payRequest 接口请求参数
type payRequest struct {
	XMLName        xml.Name `xml:"xml"`
	AppId          CDATA    `xml:"appid"`
	Body           CDATA    `xml:"body"` //String(128) 商品描述 商品简单描述 例如：腾讯充值中心-QQ会员充值
	MchId          CDATA    `xml:"mch_id"`
	NonceStr       CDATA    `xml:"nonce_str"`        //随机字符串
	NotifyURL      CDATA    `xml:"notify_url"`       //String(256)通知地址 异步接收微信支付结果通知的回调地址，通知url必须为外网可访问的url，不能携带参数。
	OpenID         CDATA    `xml:"openid"`           //用户标识 trade_type=JSAPI，此参数必传，用户在商户appid下的唯一标识
	OutTradeNo     CDATA    `xml:"out_trade_no"`     //商户订单号
	SpbillCreateIP CDATA    `xml:"spbill_create_ip"` //String(64) 终端IP 支持IPV4和IPV6两种格式的IP地址。调用微信支付API的机器IP
	TotalFee       CDATA    `xml:"total_fee"`        //	Int 88 标价金额 订单总金额，单位为分
	TradeType      CDATA    `xml:"trade_type"`       //交易类型
	Sign           string   `xml:"sign"`

	//DeviceInfo     CDATA    `xml:"device_info,omitempty"` //设备号 非必填
	//SignType       CDATA    `xml:"sign_type,omitempty"`   //签名类型 非必填 签名类型，默认为MD5，支持HMAC-SHA256和MD5。
	//Detail         CDATA    `xml:"detail,omitempty"`      //String(6000) 商品详情 非必填
	//Attach         CDATA    `xml:"attach,omitempty"`      //String(127) 附加数据 非必填 例如：	深圳分店
	//FeeType        CDATA    `xml:"fee_type,omitempty"`    //	String(16) 标价币种 非必填 符合ISO 4217标准的三位字母代码，默认人民币：CNY
	//TimeStart      CDATA    `xml:"time_start,omitempty"`  //交易起始时间  非必填 格式为yyyyMMddHHmmss
	//TimeExpire     CDATA    `xml:"time_expire,omitempty"` //交易结束时间  非必填 格式为yyyyMMddHHmmss
	//GoodsTag       CDATA    `xml:"goods_tag,omitempty"`   //String(32)订单优惠标记 非必填
	//ProductId      CDATA    `xml:"product_id,omitempty"`  //String(32)商品ID	此参数为二维码中包含的商品ID，商户自行定义。  trade_type=NATIVE时，此参数必传 否则 非必填
	//LimitPay       CDATA    `xml:"limit_pay,omitempty"`   //指定支付方式  非必填 上传此参数no_credit--可限制用户不能使用信用卡支付
	//Receipt        CDATA    `xml:"receipt,omitempty"`     //电子发票入口开放标识 非必填 Y，传入Y时，支付成功消息和支付详情页将出现开票入口。需要在微信支付商户平台或微信公众平台开通电子发票功能，传此字段才可生效
	//SceneInfo      CDATA    `xml:"scene_info,omitempty"`  //场景信息  非必填 该字段常用于线下活动时的场景信息上报，支持上报实际门店信息，商户也可以按需求自己上报相关信息。该字段为JSON对象数据，对象格式为{"store_info":{"id": "门店ID","name": "名称","area_code": "编码","address": "地址" }}

}

// PrePayOrder return data for invoke wechat payment
func PrePayOrder(p *PayParams) (payOrder PreOrder, err error) {
	nonceStr := RandomStr(32)
	param := make(map[string]interface{})
	param["appid"] = p.AppId
	param["nonce_str"] = nonceStr
	param["body"] = p.Body
	param["notify_url"] = p.NotifyURL
	param["mch_id"] = p.MchId
	param["out_trade_no"] = p.OutTradeNo
	param["spbill_create_ip"] = p.SpbillCreateIP
	param["total_fee"] = p.TotalFee
	param["trade_type"] = p.TradeType
	param["openid"] = p.OpenID

	bizKey := "&key=" + p.PayKey

	str := orderParam(param, bizKey)

	fmt.Println(str)

	sign := Md5Sum(str)

	fmt.Println(sign)

	request := payRequest{
		AppId:          CDATA{Text: p.AppId},
		Body:           CDATA{Text: p.Body},
		MchId:          CDATA{Text: p.MchId},
		NonceStr:       CDATA{Text: nonceStr},
		NotifyURL:      CDATA{Text: p.NotifyURL},
		OutTradeNo:     CDATA{Text: p.OutTradeNo},
		SpbillCreateIP: CDATA{Text: p.SpbillCreateIP},
		TotalFee:       CDATA{Text: p.TotalFee},
		TradeType:      CDATA{Text: p.TradeType},
		OpenID:         CDATA{Text: p.OpenID},
		Sign:           sign,
	}
	request.XMLName = xml.Name{Space: "", Local: "xml"}

	rawRet, err := PostXML(payGateway, request)
	if err != nil {
		return
	}

	err = xml.Unmarshal(rawRet, &payOrder)
	if err != nil {
		return
	}

	if payOrder.ReturnCode == "SUCCESS" {
		//pay success
		if payOrder.ResultCode == "SUCCESS" {
			err = nil
			return
		}
		err = errors.New(payOrder.ErrCode + payOrder.ErrCodeDes)
		return
	}

	fmt.Println(string(rawRet))

	err = errors.New("[msg : xmlUnmarshalError] [rawReturn : " + string(rawRet) + "] [params : " + str + "] [sign : " + sign + "]")
	return
}

// Config 是传出用于 js sdk 用的参数
type PrePayIdCreateConfig struct {
	Timestamp string `json:"timestamp"`
	NonceStr  string `json:"nonceStr"`
	PrePayID  string `json:"prePayId"`
	SignType  string `json:"signType"`
	Package   string `json:"package"`
	PaySign   string `json:"paySign"`
}

//  生成prepay_id 请求
func PrePayIdCreate(p *PayParams, payOrder PreOrder) (cfg PrePayIdCreateConfig, err error) {
	var (
		buffer    strings.Builder
		h         hash.Hash
		timestamp = strconv.FormatInt(time.Now().Unix(), 10)
	)

	buffer.WriteString("appId=")
	buffer.WriteString(payOrder.AppID)
	buffer.WriteString("&nonceStr=")
	buffer.WriteString(payOrder.NonceStr)
	buffer.WriteString("&package=")
	buffer.WriteString("prepay_id=" + payOrder.PrePayID)
	buffer.WriteString("&signType=")
	buffer.WriteString(p.SignType)
	buffer.WriteString("&timeStamp=")
	buffer.WriteString(timestamp)
	buffer.WriteString("&key=")
	buffer.WriteString(p.PayKey)

	if p.SignType == "MD5" {
		h = md5.New()
	} else {
		h = hmac.New(sha256.New, []byte(p.PayKey))
	}
	h.Write([]byte(buffer.String()))

	// 签名
	cfg.PaySign = strings.ToUpper(hex.EncodeToString(h.Sum(nil)))
	cfg.NonceStr = payOrder.NonceStr
	cfg.Timestamp = timestamp
	cfg.PrePayID = payOrder.PrePayID
	cfg.SignType = p.SignType
	cfg.Package = "prepay_id=" + payOrder.PrePayID

	return
}

// order params
func orderParam(source interface{}, bizKey string) (returnStr string) {
	switch v := source.(type) {
	case map[string]string:
		keys := make([]string, 0, len(v))
		for k := range v {
			if k == "sign" {
				continue
			}
			keys = append(keys, k)
		}
		sort.Strings(keys)
		var buf bytes.Buffer
		for _, k := range keys {
			if v[k] == "" {
				continue
			}
			if buf.Len() > 0 {
				buf.WriteByte('&')
			}
			buf.WriteString(k)
			buf.WriteByte('=')
			buf.WriteString(v[k])
		}
		buf.WriteString(bizKey)
		returnStr = buf.String()
	case map[string]interface{}:
		keys := make([]string, 0, len(v))
		for k := range v {
			if k == "sign" {
				continue
			}
			keys = append(keys, k)
		}
		sort.Strings(keys)
		var buf bytes.Buffer
		for _, k := range keys {
			if v[k] == "" {
				continue
			}
			if buf.Len() > 0 {
				buf.WriteByte('&')
			}
			buf.WriteString(k)
			buf.WriteByte('=')
			switch vv := v[k].(type) {
			case string:
				buf.WriteString(vv)
			case int:
				buf.WriteString(strconv.FormatInt(int64(vv), 10))
			default:
				panic("params type not supported")
			}
		}
		buf.WriteString(bizKey)
		returnStr = buf.String()
	}
	return
}

//RandomStr 随机生成字符串
func RandomStr(length int) string {
	str := "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	bytes := []byte(str)
	result := []byte{}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < length; i++ {
		result = append(result, bytes[r.Intn(len(bytes))])
	}
	return string(result)
}

// MD5Sum 计算 32 位长度的 MD5 sum
func Md5Sum(txt string) (sum string) {
	h := md5.New()
	buf := bufio.NewWriterSize(h, 128)
	buf.WriteString(txt)
	buf.Flush()
	sign := make([]byte, hex.EncodedLen(h.Size()))
	hex.Encode(sign, h.Sum(nil))
	sum = string(bytes.ToUpper(sign))
	return
}
