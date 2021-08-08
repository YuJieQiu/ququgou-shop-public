package shopcart

import (
	"github.com/ququgou-shop/modules/shopcart/model"
)

func (m *ModuleShopCart) CreateTable() error {

	var err error

	err = m.DB.AutoMigrate(
		&model.ShopCart{}).Error

	return err
}
