package model

import "github.com/ququgou-shop/library/base_model"

//首页商品 配置
type HomeProductConfig struct {
	base_model.IDAutoModel
	ProductId   uint64 `json:"productId" gorm:"column:product_id;"`      //产品ID
	AppConfigId uint64 `json:"appConfigId" gorm:"column:app_config_id;"` //appconfig iD
	base_model.TimeAllModel
}

// Set table name
func (HomeProductConfig) TableName() string {
	return "home_product_configs"
}
