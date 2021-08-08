package model

import "github.com/ququgou-shop/library/base_model"

//交易详情 ß
type TransactionDetail struct {
	base_model.IDAutoModel
	TransactionId uint64 `json:"transactionId" gorm:"column:transaction_id"`
	OutTradeNo    string `json:"outTradeNo" gorm:"column:out_trade_no"` //第三方交易号 (第三方支付系统返回)
	BankType      string `json:"bankType" gorm:"column:bank_type;"`     //付款银行
	base_model.TimeAllModel
}

// Set table name
func (TransactionDetail) TableName() string {
	return "transaction_details"
}

//交易流水详情
//type TransactionDetail struct {
//	base_model.IDAutoModel
//	UserId         uint64    `json:"userId" gorm:"column:user_id"`                 //交易用户Id
//	BusinessNo     string    `json:"businessNo" gorm:"column:business_no"`         //业务编号 (可以是订单系统中的订单编号等....)
//	TradeNo        string    `json:"tradeNo" gorm:"column:trade_no"`               //交易号
//	TradeType      int       `json:"tradeType" gorm:"column:trade_type"`           //交易类型 0 普通支付 (默认) 1、提现  ##TradeType
//	OutTradeNo     string    `json:"outTradeNo" gorm:"column:out_trade_no"`        //第三方交易号 (第三方支付系统返回)
//	Amount         float64   `json:"amount" gorm:"column:amount"`                  //交易金额
//	PayType        int       `json:"payType" gorm:"column:pay_type"`               //支付类型 0:余额 1:微信 2:支付宝 ...', 暂时默认 1 微信支付 ##PaymentType
//	Source         int       `json:"source" gorm:"column:source"`                  //来源 默认0
//	Status         int       `json:"status" gorm:"column:status"`                  //状态 // 0 处理中	 1已完成  -1：失败 -3:取消  ##TradeStatus
//	CompletionTime time.Time `json:"completionTime" gorm:"column:completion_time"` //交易完成时间
//	Note           string    `json:"note" gorm:"column:note"`                      //记录内容
//	base_model.TimeAllModel
//}

//thirdParty

//OrderNo string `json:"orderNo" gorm:"column:order_no"` //交易订单号
