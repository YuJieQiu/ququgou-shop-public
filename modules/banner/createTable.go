package banner

import (
	"github.com/ququgou-shop/modules/banner/model"
)

func (m *ModuleBanner) CreateTable() error {

	err := m.DB.AutoMigrate(
		&model.Banner{}).Error

	return err
}
