package model

import "github.com/ququgou-shop/library/base_model"

//退款信息
type OrderReturns struct {
	base_model.IDAutoModel
	UserId            uint64 `json:"userId" gorm:"column:user_id"`
	OrderId           uint64 `json:"orderId" gorm:"column:order_id"`
	OrderMasterId     uint64 `json:"orderMasterId" gorm:"column:order_master_id;"`
	OrderDetailId     uint64 `json:"orderDetailId" gorm:"column:order_detail_id"`
	ReturnNo          string `json:"returnNo" gorm:"column:return_no"`                    //售后流水
	MerId             uint64 `json:"merId" gorm:"column:mer_id"`                          //商家id
	ExpressNo         string `json:"expressNo" gorm:"column:express_no"`                  //物流订单号
	AddressId         uint64 `json:"addressId" gorm:"column:address_id"`                  //收货地址
	ConsigneeRealName string `json:"consigneeRealName" gorm:"column:consignee_real_name"` //收货人姓名
	ConsigneeTelPhone string `json:"consigneeTelPhone" gorm:"column:consignee_tel_phone"` //收货人手机号
	ConsigneeAddress  string `json:"consigneeAddress" gorm:"column:consignee_address"`    //收货地址
	ConsigneeZip      string `json:"consigneeZip" gorm:"column:consignee_zip"`            //邮政编码
	ReturnsType       int    `json:"returnsType" gorm:"column:returns_type"`              //'0全部退单 1部分退单'
	ReturnsAmount     int    `json:"returnsAmount" gorm:"column:returns_amount"`          //退款金额
	Status            int    `json:"status" gorm:"column:status"`
	base_model.TimeAllModel
}

// Set table name
func (OrderReturns) TableName() string {
	return "order_returns"
}
