package orderService

import (
	"fmt"
	"strconv"

	"github.com/jinzhu/gorm"
	"github.com/ququgou-shop/modules/order/model"
	"github.com/ququgou-shop/modules/order/orderEnum"
)

//获取订单各个状态的数量
func GetOrderCountByStatus(db *gorm.DB, status orderEnum.OrderBusinessStatus,
	merId uint64,
	userId uint64) (
	err error, count int) {

	w := ""

	if merId > 0 {
		w += fmt.Sprintf("mer_id=%v and ", strconv.FormatUint(merId, 10))
	} else {
		w += fmt.Sprintf("user_id=%v and ", strconv.FormatUint(userId, 10))
	}

	w += LoadQueryStatusSqlString(status)

	err = db.Model(&model.Order{}).Where(w).Count(&count).Error

	return
}
