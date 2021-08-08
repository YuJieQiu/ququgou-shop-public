package productService

import (
	"github.com/jinzhu/gorm"
)

func GetProductPaymentTypeList(db *gorm.DB, q *GetProductPaymentTypeListModel) (error,
	*[]ProductPaymentTypeModel) {
	var (
		err  error
		list []ProductPaymentTypeModel
	)

	sqlSelect := ""

	err = db.Table("payment_types").Select(sqlSelect).
		Joins("inner join product_payment_type on product_payment_type.payment_type_id = payment_types.id").
		Where("payment_types.deleted_at is null and product_payment_type.deleted_at is null  and payment_types.status=? and product_payment_type.product_id=?", 0, q.ProductId).
		Scan(&list).Error

	if err != nil && err != gorm.ErrRecordNotFound {
		return err, nil
	}

	return nil, &list
}
