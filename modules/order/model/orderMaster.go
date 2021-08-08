package model

import (
	"github.com/ququgou-shop/library/base_model"
	"time"
)

//订单 原始表(主订单 原始表，可能拆分为多个 Order)
type OrderMaster struct {
	base_model.IDAutoModel
	OrderMasterNo         string     `json:"orderMasterNo" gorm:"column:order_master_no;"`                                  //订单编号
	UserId                uint64     `json:"userId" gorm:"column:user_id"`                                                  //用户编号
	UserGuid              string     `json:"userGuid" gorm:"column:user_guid"`                                              //用户编号
	ProductAmountTotal    float64    `json:"productAmountTotal" gorm:"type:decimal(18,2);column:product_amount_total;"`     //商品总价(未优惠前)
	OrderAmountTotal      float64    `json:"orderAmountTotal" gorm:"type:decimal(18,2);column:order_amount_total"`          //实际付款金额(优惠后)
	DiscountsAmountTotal  float64    `json:"discountsAmountTotal" gorm:"type:decimal(18,2);column:discounts_amount_total;"` //优惠金额
	DeliveryFee           float64    `json:"deliveryFee" gorm:"type:decimal(18,2);column:delivery_fee;"`                    //配送运费
	DeliveryType          int        `json:"deliveryType" gorm:"column:delivery_type"`                                      //配送方式  线下自提 、送货上门
	AddressId             uint64     `json:"addressId" gorm:"column:address_id"`                                            //收货
	OrderSettlementStatus int8       `json:"orderSettlementStatus" gorm:"column:order_settlement_status"`                   //订单结算状态 0未结算 1 结算中 3 已结算
	OrderSettlementTime   *time.Time `json:"orderSettlementTime" gorm:"column:order_settlement_time" `                      //结算时间 TODO：时间是转化成本地时间
	Remark                string     `json:"remark" gorm:"varchar(5000);column:remark;"`
	PaymentTypeId         uint64     `json:"paymentTypeId" gorm:"column:payment_type_id"`      //支付类型id 暂无
	PaymentTypeCode       string     `json:"paymentTypeCode" gorm:"column:payment_type_code;"` //支付方式Code
	SourceType            int        `json:"sourceType" gorm:"column:source_type"`             //订单来源 默认 0
	Type                  int        `json:"type" gorm:"column:type"`                          //订单类型  #OrderType
	base_model.TimeAllModel
}

// Set table name
func (OrderMaster) TableName() string {
	return "order_master"
}
