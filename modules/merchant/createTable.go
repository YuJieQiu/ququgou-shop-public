package merchant

import (
	"github.com/ququgou-shop/modules/merchant/model"
)

func (m *ModuleMerchant) CreateTable() error {

	err := m.DB.AutoMigrate(
		&model.Merchant{}).Error

	err = m.DB.AutoMigrate(
		&model.MerchantAddress{}).Error

	err = m.DB.AutoMigrate(
		&model.MerchantApply{}).Error

	err = m.DB.AutoMigrate(
		&model.MerchantConfig{}).Error

	err = m.DB.AutoMigrate(
		&model.MerchantDetail{}).Error

	err = m.DB.AutoMigrate(
		&model.MerchantLabel{}).Error

	err = m.DB.AutoMigrate(
		&model.MerchantResources{}).Error

	err = m.DB.AutoMigrate(
		&model.MerchantType{}).Error

	err = m.DB.AutoMigrate(
		&model.MerchantUser{}).Error

	return err
}
