package model

import "github.com/ququgou-shop/library/base_model"

//订单支付(交易)表 发起支付请求成功时创建
type OrderTrade struct {
	base_model.IDAutoModel
	UserId        uint64  `json:"userId" gorm:"column:user_id"`
	OrderNo       string  `json:"orderNo" gorm:"column:order_no"`                 //订单编号  (这里暂时放两个和订单相关的列信息)
	OrderMasterNo string  `json:"orderMasterNo" gorm:"column:order_master_no;"`   //订单编号
	TradeNo       string  `json:"tradeNo" gorm:"column:trade_no;"`                //交易流水号
	PaymentTypeId uint64  `json:"paymentTypeId" gorm:"column:payment_type_id;"`   //支付方式 1、微信支付  ##PaymentType
	Amount        float64 `json:"amount" gorm:"column:amount;type:decimal(18,2)"` //支付金额
	Status        int     `json:"status" gorm:"column:status"`                    //0
	Remark        string  `json:"remark" gorm:"column:remark"`                    //备注
	base_model.TimeAllModel
}

// Set table name
func (OrderTrade) TableName() string {
	return "order_trades"
}
