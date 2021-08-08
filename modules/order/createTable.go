package order

import (
	"github.com/ququgou-shop/modules/order/model"
)

//创建模块下所有表
func (m *ModuleOrder) CreateTable() error {
	var err error

	err = m.DB.AutoMigrate(
		&model.OrderMaster{}).Error

	err = m.DB.AutoMigrate(
		&model.Order{}).Error

	err = m.DB.AutoMigrate(
		&model.OrderDetail{}).Error

	err = m.DB.AutoMigrate(
		&model.OrderLogistics{}).Error

	err = m.DB.AutoMigrate(
		&model.OrderLogs{}).Error

	err = m.DB.AutoMigrate(
		&model.OrderPaymentLogs{}).Error

	err = m.DB.AutoMigrate(
		&model.OrderProduct{}).Error

	err = m.DB.AutoMigrate(
		&model.OrderReturns{}).Error

	err = m.DB.AutoMigrate(
		&model.OrderReturnsApply{}).Error

	err = m.DB.AutoMigrate(
		&model.OrderTrade{}).Error

	return err
}
