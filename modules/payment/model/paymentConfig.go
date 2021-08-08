package model

import (
	"github.com/jinzhu/gorm"
	"github.com/ququgou-shop/library/base_model"
	"github.com/ququgou-shop/modules/payment/paymentEnum"
)

//支付配置 其它字段需要再加
type PaymentConfig struct {
	base_model.IDAutoModel
	Name       string `json:"name" gorm:"column:name"`
	AppId      string `json:"appId" gorm:"column:app_id"`
	MchId      string `json:"mchId" gorm:"column:mch_id"`
	Type       string `json:"type" gorm:"column:type"` //1 微信支付  ##PaymentType
	Key        string `json:"key" gorm:"column:key"`   //支付key 第三方系统
	NotifyURL  string `json:"notifyUrl" gorm:"column:notify_url"`
	ExpireTime uint64 `json:"expireTime" gorm:"column:expire_time;"` //支付过期时间 单位 秒
	base_model.TimeAllModel
}

// Set table name
func (PaymentConfig) TableName() string {
	return "payment_configs"
}

//获取支付配置信息
func (p *PaymentConfig) GetPaymentConfig(db *gorm.DB, configType paymentEnum.PaymentType) error {

	err := db.Where("type=?", string(configType)).First(p).Error
	return err
}
