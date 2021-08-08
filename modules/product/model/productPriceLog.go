package model

import "github.com/ququgou-shop/library/base_model"

//产品价格表，如果发生变化会记录下来，默认最后一条记录的价格是最新的
type ProductPriceLog struct {
	base_model.IDAutoModel
	ProductId       uint64  `json:"productId" gorm:"column:product_id"`
	ProductSkuId    uint64  `json:"productSkuId" gorm:"column:product_sku_id"`
	Price           float64 `json:"price" gorm:"column:price"`
	OperatorAdminId uint64  `json:"operatorAdminId" gorm:"column:operator_admin_id"` //操作管理员用户编号
	base_model.TimeAllModel
}

// Set table name
func (ProductPriceLog) TableName() string {
	return "product_price_logs"
}
