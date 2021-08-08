package model

import "github.com/ququgou-shop/library/base_model"

//交易记录表
type TransactionRecord struct {
	base_model.IDAutoModel
	TransactionId uint64 `json:"transactionId" gorm:"column:transaction_id"`
	Events        string `json:"events" gorm:"column:events"` //事件详情
	Result        string `json:"result" gorm:"column:result;type:text;"`
	base_model.TimeAllModel
}

// Set table name
func (TransactionRecord) TableName() string {
	return "transaction_records"
}
