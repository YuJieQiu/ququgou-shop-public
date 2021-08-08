package model

import (
	"github.com/ququgou-shop/library/base_model"
	"time"
)

//订单表 (根据 OrderMaster拆分出来的 订单表)
type Order struct {
	base_model.IDAutoModel
	UserId                uint64     `json:"userId" gorm:"column:user_id"`                                                  //用户编号
	OrderNo               string     `json:"orderNo" gorm:"column:order_no"`                                                //订单编号  (这里暂时放两个和订单相关的列信息)
	OrderMasterNo         string     `json:"orderMasterNo" gorm:"column:order_master_no;"`                                  //订单编号
	OrderMasterId         uint64     `json:"orderMasterId" gorm:"column:order_master_id;"`                                  //订单原始表id
	MerId                 uint64     `json:"merId" gorm:"column:mer_id"`                                                    //商户id
	OrderStatus           int        `json:"orderStatus" gorm:"column:order_status;"`                                       //子订单状态 (前面都是和订单主表状态同步，可能 发货和售后一块是和订单主表不同步的)  0未付款,1已付款,3已发货,5已签收,-1退货申请,-3退货中,-5已退货,-7取消交易 -9撤销申请',
	PayStatus             int        `json:"payStatus" gorm:"column:pay_status;"`                                           //支付状态
	DeliveryStatus        int        `json:"deliveryStatus" gorm:"column:delivery_status;"`                                 //物流/配送 状态
	ProductAmountTotal    float64    `json:"productAmountTotal"gorm:"column:product_amount_total;type:decimal(18,2)"`       //商品总价(未优惠前)
	OrderAmountTotal      float64    `json:"orderAmountTotal" gorm:"column:order_amount_total;type:decimal(18,2)"`          //实际付款金额(优惠后)
	DiscountsAmountTotal  float64    `json:"discountsAmountTotal"  gorm:"column:discounts_amount_total;type:decimal(18,2)"` //分摊到的优惠金额
	OrderSettlementStatus int8       `json:"orderSettlementStatus" gorm:"column:order_settlement_status"`                   //订单结算状态 0未结算 1 结算中 3 已结算
	OrderSettlementTime   *time.Time `json:"orderSettlementTime"  gorm:"column:order_settlement_time"`                      //结算时间 TODO：时间是转化成本地时间
	Remark                string     `json:"remark" gorm:"column:remark;varchar(5000)"`                                     //备注
	DeliveryType          int        `json:"deliveryType" gorm:"column:delivery_type"`                                      //配送方式
	AddressId             uint64     `json:"addressId" gorm:"column:address_id"`                                            //收货地址id
	AfterStatus           int        `json:"afterStatus" gorm:"column:after_status"`                                        //售后状态 默认0 (暂留)
	Type                  int        `json:"type" gorm:"column:type"`                                                       //类型 #OrderType
	CancelTime            *time.Time `json:"cancelTime" gorm:"column:cancel_time"`                                          //订单取消时间
	base_model.TimeAllModel
}

// Set table name
func (Order) TableName() string {
	return "orders"
}
