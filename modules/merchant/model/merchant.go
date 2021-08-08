package model

import "github.com/ququgou-shop/library/base_model"

//商户信息
type Merchant struct {
	base_model.IDAutoModel
	Guid   		string `json:"guid" gorm:"unique_index;not null;unique;column:guid;"` //商户guid
	Name        string `json:"name" gorm:"column:name"` //名称
	Description string `json:"description" gorm:"column:description"` //描述
	Status      int    `json:"status" gorm:"column:status"`           //0 状态 ( 0 审核通过)    -1 审核中
	TypeId      uint64 `json:"typeId" gorm:"column:type_id"`          //类型 默认 0  MerchantTypeId
	Active      bool   `json:"active" gorm:"column:active"`           //激活
	base_model.TimeAllModel
}

// Set table name
func (Merchant) TableName() string {
	return "merchants"
}
