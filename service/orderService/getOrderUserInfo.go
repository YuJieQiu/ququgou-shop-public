package orderService

import "github.com/jinzhu/gorm"

//获取订单用户相关信息

func MerGetOrderUserInfo(db *gorm.DB, m *MerGetOrderUserInfoModel) (err error,
	info MerOrderUserInfoModel) {

	err = db.Table("users").Select("users.user_name, users.mobile").
		Joins("inner join order_subs on order_subs.user_id = users.id").
		Where("order_subs.order_sub_no=? and order_subs.mer_id=?", m.OrderNo, m.MerId).
		Scan(&info).Error

	return err, info
}
