package productService

import (
	"github.com/jinzhu/gorm"
	"github.com/ququgou-shop/modules/product/model"
)

//商品库存变化 TODO: 需要优化
func ProductStockUpdate(db *gorm.DB, m *ProductStockUpdateModel) error {
	var sku model.ProductSKU
	var err error

	if m.Type == -1 {
		err = db.Model(&sku).Where("id = ?", m.ProductSkuId).UpdateColumn("stock", gorm.Expr("stock - ?", m.Count)).Error
		if err != nil {
			return err
		}
	} else {
		err = db.Model(&sku).Where("id = ?", m.ProductSkuId).UpdateColumn("stock", gorm.Expr("stock + ?", m.Count)).Error
		if err != nil {
			return err
		}
	}

	return nil
}
