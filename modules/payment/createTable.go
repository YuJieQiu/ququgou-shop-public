package payment

import (
	"github.com/ququgou-shop/modules/payment/model"
)

func (m *ModulePayment) CreateTable() error {
	var err error

	err = m.DB.AutoMigrate(
		&model.ClientPay{}).Error

	err = m.DB.AutoMigrate(
		&model.PaymentConfig{}).Error

	err = m.DB.AutoMigrate(
		&model.PaymentLog{}).Error

	err = m.DB.AutoMigrate(
		&model.PaymentOnlineRecord{}).Error

	err = m.DB.AutoMigrate(
		&model.PaymentType{}).Error

	err = m.DB.AutoMigrate(
		&model.Transaction{}).Error

	err = m.DB.AutoMigrate(
		&model.TransactionDetail{}).Error

	err = m.DB.AutoMigrate(
		&model.TransactionRecord{}).Error

	return err
}
