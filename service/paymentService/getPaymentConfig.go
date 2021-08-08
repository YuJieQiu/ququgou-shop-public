package paymentService

import (
	"github.com/jinzhu/gorm"
	"github.com/ququgou-shop/modules/payment/model"
)

//获取支持的支付类型列表
func GetPaymentTypeList(db *gorm.DB, q *GetPaymentTypeListModel) (error, *[]model.PaymentType) {
	var (
		err  error
		list []model.PaymentType
	)

	err = db.Where("status=?", 0).Order("sort").Find(&list).Error

	if err != nil && err != gorm.ErrRecordNotFound {
		return err, nil
	}

	return nil, &list
}
