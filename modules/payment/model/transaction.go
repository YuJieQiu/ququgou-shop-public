package model

import (
	"github.com/ququgou-shop/library/base_model"
	"time"
)

//交易流水
type Transaction struct {
	base_model.IDAutoModel
	UserId           uint64     `json:"userId" gorm:"column:user_id"`                       //交易用户Id
	BusinessNo       string     `json:"businessNo" gorm:"column:business_no"`               //业务编号 (可以是订单系统中的订单编号等....)
	BusinessType     int        `json:"businessType" gorm:"column:business_type;"`          //业务类型 (主订单、子订单 类型等)
	BusinessTypeCode string     `json:"businessTypeCode" gorm:"column:business_type_code;"` //业务类型编码
	TradeNo          string     `json:"tradeNo" gorm:"column:trade_no"`                     //交易号
	TradeType        int        `json:"tradeType" gorm:"column:trade_type"`                 //交易类型   1、订单支付  ##TradeType
	OutTradeNo       string     `json:"outTradeNo" gorm:"column:out_trade_no"`              //第三方交易号 (第三方支付系统返回)
	Amount           float64    `json:"amount" gorm:"column:amount"`                        //交易金额
	PaymentTypeId    uint64     `json:"paymentTypeId" gorm:"column:payment_type_id;"`       //支付类型Id
	PaymentTypeCode  string     `json:"paymentTypeCode" gorm:"column:payment_type_code;"`   //支付方式Code
	Source           int        `json:"source" gorm:"column:source"`                        //来源 默认0
	Status           int        `json:"status" gorm:"column:status"`                        //状态 // 0 处理中	 1已完成  -1：失败 -3:取消  ##TradeStatus
	CompletionTime   *time.Time `json:"completionTime" gorm:"column:completion_time"`       //交易完成时间
	Note             string     `json:"note" gorm:"column:note"`                            //记录内容
	base_model.TimeAllModel
}

// Set table name
func (Transaction) TableName() string {
	return "transactions"
}
