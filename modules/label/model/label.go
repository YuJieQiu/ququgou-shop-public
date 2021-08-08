package model

import "github.com/ququgou-shop/library/base_model"

//标签
type Label struct {
	base_model.IDAutoModel
	Text     string `json:"text" gorm:"column:text"`
	IsSystem bool   `json:"isSystem" gorm:"column:is_system"` //是否系统标签
	Type     int    `json:"type" gorm:"column:type"`          //1 merchant(商户)
	base_model.TimeAllModel
}

// Set table name
func (Label) TableName() string {
	return "labels"
}
