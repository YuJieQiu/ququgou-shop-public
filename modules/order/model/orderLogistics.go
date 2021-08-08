package model

import (
	"github.com/ququgou-shop/library/base_model"
	"time"
)

//订单 快递 物流

type OrderLogistics struct {
	base_model.IDAutoModel
	OrderId                 uint64    `json:"orderId" gorm:"column:order_id"`
	ExpressNo               string    `json:"expressNo" gorm:"column:express_no"`                              //物流订单号
	AddressId               uint64    `json:"addressId" gorm:"column:address_id"`                              //收货地址
	ConsigneeRealName       string    `json:"consigneeRealName" gorm:"column:consignee_real_name"`             //收货人姓名
	ConsigneeTelPhone       string    `json:"consigneeTelPhone" gorm:"column:consignee_tel_phone"`             //收货人手机号
	ConsigneeAddress        string    `json:"consigneeAddress" gorm:"column:consignee_address"`                //收货地址
	ConsigneeZip            string    `json:"consigneeZip"  gorm:"column:consignee_zip"`                       //邮政编码
	LogisticsType           int       `json:"logisticsType" gorm:"column:logistics_type"`                      //物流类型 (应该是物流公司！ 暂留)
	LogisticsFee            float64   `json:"logisticsFee" gorm:"column:logistics_fee"`                        //运费
	AgencyFee               float64   `json:"agencyFee" gorm:"column:agency_fee"`                              //快递公司代收费 (暂留)
	DeliveryAmount          float64   `json:"deliveryAmount" gorm:"column:delivery_amount"`                    //运费 成本 (付给物流公司都价格)
	Status                  int       `json:"status" gorm:"column:status"`                                     //物流状态
	LogisticsResult         string    `json:"logisticsResult" gorm:"column:logistics_result"`                  //描述信息 (暂留)
	LogisticsCreateTime     time.Time `json:"logisticsCreateTime" gorm:"column:logistics_create_time"`         //物流创建时间  (暂留)
	LogisticsUpdateTime     time.Time `json:"logisticsUpdateTime" gorm:"column:logistics_update_time"`         //物流更新时间 (暂留)
	LogisticsSettlementTime time.Time `json:"logisticsSettlementTime" gorm:"column:logistics_settlement_time"` //物流结算时间 (暂留)
	base_model.TimeAllModel
}

// Set table name
func (OrderLogistics) TableName() string {
	return "order_logistics"
}
