package model

import "github.com/ququgou-shop/library/base_model"

//标签
type Tags struct {
	base_model.IDAutoModel
	MerId uint64 `json:"merId" gorm:"column:mer_id"`          //商家Id
	Name  string `json:"name"  gorm:"column:name;index:name"` //标签名称
	base_model.TimeAllModel
}

// Set table name
func (Tags) TableName() string {
	return "tags"
}
