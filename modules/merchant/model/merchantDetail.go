package model

import "github.com/ququgou-shop/library/base_model"

type MerchantDetail struct {
	base_model.IDAutoModel
	MerchantId        uint64 `json:"merchantId" gorm:"column:merchant_id"`                //商户ID
	BusinessStartTime string `json:"businessStartTime" gorm:"column:business_start_time"` //营业开始时间
	BusinessEndTime   string `json:"businessEndTime" gorm:"column:business_end_time"`     //营业结束时间
	Logo              string `json:"logo" gorm:"column:logo"`                             //商标
	Phones            string `json:"phones" gorm:"column:phones"`
	Notice            string `json:"notice" gorm:"type:varchar(2000);column:notice"` //公告
	base_model.TimeAllModel
}

// Set table name
func (MerchantDetail) TableName() string {
	return "merchant_details"
}
