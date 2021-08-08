package productService

import (
	"github.com/jinzhu/gorm"
	"github.com/ququgou-shop/modules/product/model"
)

//商品销量变化
func ProductSalesUpdate(db *gorm.DB, m *ProductSalesUpdateModel) error {
	var sku model.ProductSKU
	var p model.Product
	var err error

	if m.Type == -1 {
		err = db.Model(&sku).Where("id = ?", m.ProductSkuId).UpdateColumn("sales", gorm.Expr("sales - ?", m.Count)).Error
		if err != nil {
			return err
		}
	} else {
		err = db.Model(&sku).Where("id = ?", m.ProductSkuId).UpdateColumn("sales", gorm.Expr("sales + ?", m.Count)).Error
		if err != nil {
			return err
		}
	}

	err = db.Model(&p).Where("id = ?", m.ProductId).UpdateColumn("sales", gorm.Expr("sales + ?", m.Count)).Error
	if err != nil {
		return err
	}

	return nil
}
