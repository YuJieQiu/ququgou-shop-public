package label

import (
	"github.com/ququgou-shop/modules/label/model"
)

func (m *ModuleLabel) CreateTable() error {

	var err error

	err = m.DB.AutoMigrate(
		&model.Label{}).Error

	return err
}
