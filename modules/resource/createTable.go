package resource

import (
	"github.com/ququgou-shop/modules/resource/model"
)

func (m *ModuleResource) CreateTable() error {

	err := m.DB.AutoMigrate(
		&model.Resource{}).Error

	err = m.DB.AutoMigrate(
		&model.ResourceConfig{}).Error

	return err
}
