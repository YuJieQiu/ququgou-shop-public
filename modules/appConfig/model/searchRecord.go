package model

import "github.com/ququgou-shop/library/base_model"

//搜索记录
type SearchRecord struct {
	base_model.IDAutoModel
	UserId uint64 `json:"userId" gorm:"column:user_id;"`
	Text   string `json:"text" gorm:"column:text;"` //搜索文本
	base_model.TimeAllModel
}

// Set table name
func (SearchRecord) TableName() string {
	return "search_records"
}
