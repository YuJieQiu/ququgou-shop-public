package model

import (
	"github.com/ququgou-shop/library/base_model"
	"github.com/ququgou-shop/modules/product/shop_ext_struct"
)

//订单详情表

type OrderDetail struct {
	base_model.IDAutoModel
	UserId                  uint64                                       `json:"userId" gorm:"column:user_id"`
	OrderId                 uint64                                       `json:"orderId" gorm:"column:order_id"`
	ProductId               uint64                                       `json:"productId" gorm:"column:product_id"`
	ProductSkuId            uint64                                       `json:"productSkuId" gorm:"column:product_sku_id"`
	ProductSkuAttributeInfo shop_ext_struct.SkuAttributeValuesArrayModel `json:"productSkuAttributeInfo" gorm:"column:product_sku_attribute_info;type:text(1000);"` //规格属性信息 ，冗余字段
	ProductCount            int                                          `json:"productCount" gorm:"column:product_count"`                                          //产品数量
	ProductUnitPrice        float64                                      `json:"productUnitPrice" gorm:"column:product_unit_price"`                                 //产品单价
	OriginalAmountTotal     float64                                      `json:"originalAmountTotal" gorm:"type:decimal(18,2);column:original_amount_total;"`       // 原始金额(优惠前)
	AmountTotal             float64                                      `json:"amountTotal" gorm:"type:decimal(18,2);column:amount_total;" `                       //付款金额 (优惠后)
	DiscountsAmountTotal    float64                                      `json:"discountsAmountTotal" gorm:"type:decimal(18,2);column:discounts_amount_total;"`     //分摊到的优惠金额
	Status                  int                                          `json:"status" gorm:"column:status;"`                                                      //状态 默认 0 暂留
	Type                    int                                          `json:"type" gorm:"column:type;"`                                                          //类型 默认 0 暂留
	base_model.TimeAllModel
}

// Set table name
func (OrderDetail) TableName() string {
	return "order_details"
}
