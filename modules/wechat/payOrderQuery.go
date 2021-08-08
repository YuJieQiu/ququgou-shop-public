package wechat

import (
	"encoding/xml"
	"errors"
	"fmt"
)

//支付订单查询
const payGatOrderInfo = "https://api.mch.weixin.qq.com/pay/orderquery"

//查询参数
type PayOrderQueryParams struct {
	XMLName       xml.Name `xml:"xml"`
	AppId         CDATA    `xml:"appid"`
	MchId         CDATA    `xml:"mch_id"`
	NonceStr      CDATA    `xml:"nonce_str"`                //随机字符串
	TransactionId *CDATA   `xml:"transaction_id,omitempty"` //微信订单号 二选一
	OutTradeNo    *CDATA   `xml:"out_trade_no,omitempty"`   //商户订单号 二选一
	Sign          string   `xml:"sign"`
}

//返回参数
type PayOrderQueryResponseParams struct {
	NotifyResult
	TradeState string `xml:"trade_state"` //SUCCESS—支付成功 REFUND—转入退款 NOTPAY—未支付 CLOSED—已关闭 REVOKED—已撤销（刷卡支付） USERPAYING--用户支付中 PAYERROR--支付失败(其他原因，如银行返回失败)
}

func PayOrderQuery(p *PayOrderQueryParams, payKey string) (error, *PayOrderQueryResponseParams) {
	var (
		err error
		res PayOrderQueryResponseParams
	)

	nonceStr := RandomStr(32)
	param := make(map[string]interface{})
	param["appid"] = p.AppId.Text
	param["mch_id"] = p.MchId.Text
	if p.TransactionId != nil {
		param["transaction_id"] = p.TransactionId.Text
	} else {
		param["out_trade_no"] = p.OutTradeNo.Text
	}

	param["nonce_str"] = nonceStr

	bizKey := "&key=" + payKey

	str := orderParam(param, bizKey)
	fmt.Println(str)
	sign := Md5Sum(str)
	fmt.Println(sign)

	p.NonceStr = CDATA{Text: nonceStr}
	p.Sign = sign

	p.XMLName = xml.Name{Space: "", Local: "xml"}

	rawRet, err := PostXML(payGatOrderInfo, p)
	if err != nil {
		return err, nil
	}
	err = xml.Unmarshal(rawRet, &res)
	if err != nil {
		return err, nil
	}

	if res.ReturnCode == "SUCCESS" {
		//pay success
		if res.ResultCode == "SUCCESS" {
			err = nil
			return err, &res
		}
		err = errors.New(res.ErrCode + res.ErrCodeDes)
		return err, &res
	}

	fmt.Println(string(rawRet))

	err = errors.New("[msg : xmlUnmarshalError] [rawReturn : " + string(rawRet) + "] [params : " + str + "] [sign : " + sign + "]")

	return err, nil
}
