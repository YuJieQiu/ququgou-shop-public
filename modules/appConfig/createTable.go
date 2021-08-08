package appConfig

import (
	"github.com/ququgou-shop/modules/appConfig/model"
)

func (m *ModuleAppConfig) CreateTable() error {

	var err error

	err = m.DB.AutoMigrate(
		&model.AppConfig{}).Error
	err = m.DB.AutoMigrate(
		&model.AppConfigType{}).Error
	err = m.DB.AutoMigrate(
		&model.HomeProductConfig{}).Error
	err = m.DB.AutoMigrate(
		&model.HotSearchConfig{}).Error
	err = m.DB.AutoMigrate(
		&model.SearchRecord{}).Error

	return err
}
