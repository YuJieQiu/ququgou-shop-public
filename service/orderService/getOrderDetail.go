package orderService

import (
	"fmt"

	"github.com/jinzhu/gorm"
	"github.com/ququgou-shop/library/ext/ext_struct"
	MerchantModel "github.com/ququgou-shop/modules/merchant/model"
	"github.com/ququgou-shop/modules/order/model"
	"github.com/ququgou-shop/modules/order/orderEnum"
	UserModel "github.com/ququgou-shop/modules/user/model"
)

//获取订单详细信息
func GetUserOrderDetail(db *gorm.DB, q *GetUserOrderDetailModel, u *UserModel.User, imgServiceUrl string) (
	error,
	*UserOrderDetailModel) {
	var (
		err           error
		data          UserOrderDetailModel
		order         model.Order
		orderDetails  *[]model.OrderDetail
		orderProducts *[]model.OrderProduct
		mer           MerchantModel.Merchant
		ids           []uint64
	)

	err, order = getOrderInfo(db, q.OrderNo, u.ID, q.MerId)

	if err != nil {
		return err, nil
	}

	data = UserOrderDetailModel{
		NO:                   order.OrderNo,
		Id:                   order.ID,
		MerId:                order.MerId,
		OriginalAmountTotal:  order.ProductAmountTotal,
		DiscountsAmountTotal: order.DiscountsAmountTotal,
		OrderAmountTotal:     order.OrderAmountTotal,
		CreatedTime:          ext_struct.JsonTime(order.CreatedAt),
		Remark:               order.Remark,
		DeliveryType:         order.DeliveryType,
		Status:               string(GetOrderBusinessStatus(order.OrderStatus, order.PayStatus, order.DeliveryStatus, order.Type)),
	}

	if order.CancelTime != nil {
		data.CancelTime = ext_struct.JsonTime(*order.CancelTime)
	}

	data.StatusText = orderEnum.OrderBusinessStatus(data.Status).Text()

	//获取订单商品信息
	ids = append(ids, data.Id)
	err, orderDetails, orderProducts = GetOrderProductInfo(db, ids)
	if err != nil {
		return err, nil
	}
	data.Products = *LoadOrderProductInfo(data.Id, orderDetails, orderProducts, imgServiceUrl)

	//获取当前订单收货地址
	err, data.Address = getOrderAddressInfo(db, data.NO)
	if err != nil {
		return err, nil
	}

	//获取订单支付信息

	//获取订单售后信息

	//获取订单物流信息

	//获取商家信息
	err = db.Where("id =? ", order.MerId).First(&mer).Error
	if err == nil {
		data.MerName = mer.Name
	}

	return nil, &data
}

func getOrderAddressInfo(db *gorm.DB, subOrderNo string) (err error,
	data UserOrderAddressModel) {
	var (
		sql string
	)

	sql = fmt.Sprintf(` 
		SELECT a.city, a.region, a.town, a.address, a.phone
		, a.name
		FROM orders o
			INNER JOIN addresses a ON o.address_id = a.id
		WHERE o.order_no = ? 
			AND o.deleted_at IS NULL
 `)
	err = db.Raw(sql, subOrderNo).Scan(&data).Error

	if err != nil && err != gorm.ErrRecordNotFound {
		return err, data
	}

	return nil, data
}

//获取订单信息
//If merId>0 根据 MerId 获取 else 根据 userID获取
func getOrderInfo(db *gorm.DB, orderNo string, uId uint64, merId uint64) (err error,
	data model.Order) {

	//sqlSelect := `SELECT id,order_sub_no,mer_id,order_amount_total,created_at ,cancel_time, discounts_amount_total,product_amount_total,remark`
	//sql := ` from order_subs where user_id=? and order_sub_no=? and deleted_at is null `
	//db.Table("order_subs").Where("user_id=? and order_sub_no=?", uId, subOrderNo).
	//	Select("id,order_sub_no,mer_id,order_amount_total,created_at ,cancel_time, discounts_amount_total,product_amount_total,remark,os.order_status,pay_status,delivery_status,type").
	//	Scan(&data).

	w := model.Order{
		OrderNo: orderNo,
	}

	if merId > 0 {
		w.MerId = merId
	} else {
		w.UserId = uId
	}

	err = db.Where(&w).First(&data).Error

	if err != nil {
		return err, data
	}

	return nil, data

}
