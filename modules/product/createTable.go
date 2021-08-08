package product

import (
	"github.com/ququgou-shop/modules/product/model"
)

func (m *ModuleProduct) CreateTable() error {

	var err error

	err = m.DB.AutoMigrate(
		&model.Attribute{}).Error

	err = m.DB.AutoMigrate(
		&model.AttributeGroup{}).Error

	err = m.DB.AutoMigrate(
		&model.AttributeGroupSKU{}).Error

	err = m.DB.AutoMigrate(
		&model.AttributeValue{}).Error

	err = m.DB.AutoMigrate(
		&model.AttributeValueGroup{}).Error

	err = m.DB.AutoMigrate(
		&model.AttributeValueSKU{}).Error

	err = m.DB.AutoMigrate(
		&model.Brand{}).Error

	err = m.DB.AutoMigrate(
		&model.Category{}).Error

	err = m.DB.AutoMigrate(
		&model.CategoryProduct{}).Error

	err = m.DB.AutoMigrate(
		&model.Product{}).Error

	err = m.DB.AutoMigrate(
		&model.ProductCollection{}).Error

	err = m.DB.AutoMigrate(
		&model.ProductContent{}).Error

	err = m.DB.AutoMigrate(
		&model.ProductDeliveryType{}).Error

	err = m.DB.AutoMigrate(
		&model.ProductPaymentType{}).Error

	err = m.DB.AutoMigrate(
		&model.ProductPriceLog{}).Error

	err = m.DB.AutoMigrate(
		&model.ProductResource{}).Error

	err = m.DB.AutoMigrate(
		&model.ProductSKU{}).Error

	err = m.DB.AutoMigrate(
		&model.ProductTags{}).Error

	err = m.DB.AutoMigrate(
		&model.Property{}).Error

	err = m.DB.AutoMigrate(
		&model.PropertyValue{}).Error

	err = m.DB.AutoMigrate(
		&model.Tags{}).Error

	err = m.DB.Create(&model.Attribute{Name: "规格", CategoryId: 0}).Error

	return err
}
