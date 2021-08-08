package paymentService

import (
	"github.com/jinzhu/gorm"
)

//创建支付

type OnlinePaymentCreateResult struct {
	Success    bool        `json:"success"`
	Msg        string      `json:"msg"`
	TradeNo    string      `json:"tradeNo"`
	Data       interface{} `json:"data"`
	ErrCode    string      `json:"errCode"`
	ErrCodeDes string      `json:"errCodeDes"`
}

//在线支付
type OnlinePayment interface {
	//调用生成方法
	OnlinePaymentCreate(db *gorm.DB) (error, *OnlinePaymentCreateResult)
}
