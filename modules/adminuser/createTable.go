package adminuser

import (
	"github.com/ququgou-shop/modules/adminuser/model"
)

func (m *ModuleAdminUser) CreateTable() error {
	err := m.DB.AutoMigrate(
		&model.AdminUser{}).Error
	return err
}
