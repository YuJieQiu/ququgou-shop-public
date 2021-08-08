package user

import (
	"github.com/ququgou-shop/modules/user/model"
)

func (m *ModuleUser) CreateTable() error {

	err := m.DB.AutoMigrate(
		&model.User{}).Error

	err = m.DB.AutoMigrate(
		&model.UserWeChatInfo{}).Error

	err = m.DB.AutoMigrate(
		&model.UserLoginLog{}).Error

	return err
}
