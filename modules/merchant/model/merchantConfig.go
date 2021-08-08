package model

import "github.com/ququgou-shop/library/base_model"

type MerchantConfig struct {
	base_model.IDAutoModel
	MerchantId    uint64 `json:"merchantId" gorm:"column:merchant_id"`
	ActiveGoods   bool   `json:"activeGoods" gorm:"column:active_goods"`     //启用商品
	ActiveService bool   `json:"activeService" gorm:"column:active_service"` //启用服务
	ActiveComment bool   `json:"activeComment" gorm:"column:active_comment"` //启用评论
	base_model.TimeAllModel
}

// Set table name
func (MerchantConfig) TableName() string {
	return "merchant_configs"
}
