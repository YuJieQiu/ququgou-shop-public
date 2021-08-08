package paymentService

import (
	"strconv"

	"github.com/jinzhu/gorm"
	"github.com/ququgou-shop/modules/payment/model"
	"github.com/ququgou-shop/modules/payment/paymentEnum"
	"github.com/ququgou-shop/modules/wechat"
)

type WechatPayment struct {
	TradeNo      string  `json:"tradeNo"`
	Amount       float64 `json:"amount"`
	WechatOpenId string  `json:"wechatOpenId"`
	ClientIp     string  `json:"clientIp"`
}

//微信支付
func (q WechatPayment) OnlinePaymentCreate(db *gorm.DB) (error, *OnlinePaymentCreateResult) {
	var (
		err      error
		payConf  model.PaymentConfig
		payOrder wechat.PreOrder
		res      OnlinePaymentCreateResult
	)

	//defer func() {
	//	if r := recover(); r != nil {
	//		srt, _ := json.Marshal(r)
	//		err = errors.New(fmt.Sprintf("WechatOnlinePaymentCreate Error %v", srt))
	//	}
	//}()

	srt := string(paymentEnum.PaymentTypeWeChatPay)

	//获取支付配置
	err = db.Model(&payConf).Where("type = ?", srt).First(&payConf).Error
	if err != nil {
		return err, nil
	}

	//金额 单位 分
	totalFee := strconv.FormatFloat(q.Amount*100, 'g', -1, 64)

	payParams := wechat.PayParams{
		AppId:          payConf.AppId,
		MchId:          payConf.MchId,
		OutTradeNo:     q.TradeNo,
		Body:           "Task-jack",
		TotalFee:       totalFee,
		SpbillCreateIP: q.ClientIp,
		NotifyURL:      payConf.NotifyURL,
		TradeType:      "JSAPI",
		OpenID:         q.WechatOpenId,
		PayKey:         payConf.Key,
		SignType:       "MD5",
	}

	payOrder, err = wechat.PrePayOrder(&payParams)

	//微信的是返回支付参数(5个参数和Sign)

	if err != nil {
		res.Success = false
		res.Data = payOrder
		//TODO:创建失败日志
		return err, nil
	}

	if len(payOrder.ErrCode) > 0 {
		res.Success = false
		res.ErrCode = payOrder.ErrCode
		res.ErrCodeDes = payOrder.ErrCodeDes
		return err, &res
	}

	//需要进行签名返回
	//类别 小程序、网页、app 方式可能都不同，这里的方法只是小程序的 参考官方文档 https://pay.weixin.qq.com/wiki/doc/api/wxa/wxa_api.php?chapter=7_7&index=5

	prePayIdData, err := wechat.PrePayIdCreate(&payParams, payOrder)
	if err != nil {
		res.Success = false
		//TODO:创建失败日志
		return err, nil
	}

	res.Success = true
	res.Data = prePayIdData

	return nil, &res
}
