package orderService

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/ququgou-shop/modules/order/model"
	"github.com/ququgou-shop/modules/order/orderEnum"
	"strconv"
)

// get order list 获取基本订单信息
func GetOrderSmallInfoList(db *gorm.DB,
	q *GetOrderSmallInfoListModel,
	imgService string) (error,
	*[]SmallOrderInfModel,
	int) {
	var (
		err           error
		list          []SmallOrderInfModel
		count         int
		orderDetails  *[]model.OrderDetail
		orderProducts *[]model.OrderProduct
		ids           []uint64
	)
	err, list, count = GetOrderInfoList(db, q)
	if err != nil {
		return err, nil, count
	}

	loadOrderBusinessStatus(list)

	for _, v := range list {
		ids = append(ids, v.Id)
	}

	err, orderDetails, orderProducts = GetOrderProductInfo(db, ids)
	if err != nil {
		return err, nil, count
	}

	for i := 0; i < len(list); i++ {
		v := LoadOrderProductInfo(list[i].Id, orderDetails, orderProducts, imgService)
		list[i].Products = *v
		list[i].StatusText = orderEnum.OrderBusinessStatus(list[i].Status).Text()
	}

	return nil, &list, count
}

//
func GetOrderInfoList(db *gorm.DB, q *GetOrderSmallInfoListModel) (err error,
	list []SmallOrderInfModel,
	count int) {

	q.QueryParamsPage.PageSet()

	sqlSelectCount := " select count(o.id) as count "

	sqlSelect := `SELECT o.id,o.order_no as no ,o.mer_id,m.name as mer_name , o.order_amount_total,o.created_at as created_time,o.order_status,o.pay_status,o.delivery_status,o.type  `

	sql := ` from orders o LEFT JOIN merchants m on o.mer_id=m.id  where o.deleted_at is null`

	if q.UserId > 0 {
		sql += `  and  o.user_id=` + strconv.FormatUint(q.UserId, 10)
	}
	if q.MerId > 0 {
		sql += `  and  o.mer_id=` + strconv.FormatUint(q.MerId, 10)
	}

	if len(q.OrderNo) > 0 {
		sql += `  and  o.order_no like %` + q.OrderNo + `%`
	}

	if !q.All {
		//根据业务订单状态查询订单各个状态
		if len(q.BusinessStatus) > 1 {
			statusSql := ""
			for _, v := range q.BusinessStatus {
				s := loadQueryStatusSqlString(orderEnum.OrderBusinessStatus(v))
				if len(s) > 0 {
					statusSql += fmt.Sprintf("or %v", s)
				}
			}

			statusSql = string([]byte(statusSql)[2:])

			sql += fmt.Sprintf(" and ( %v ) ", statusSql)
		} else {
			sql += fmt.Sprintf(" and %v ", loadQueryStatusSqlString(orderEnum.OrderBusinessStatus(q.BusinessStatus[0])))
		}
	}

	sql += ` ORDER BY o.created_at desc `

	err = db.Raw(sqlSelectCount + sql).Count(&count).Error

	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, list, count
		}
		return err, list, count
	}

	err = db.Raw(sqlSelect + sql).Offset(q.Offset).Limit(q.Limit).Scan(&list).Error

	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, list, count
		}
		return err, list, count
	}

	return nil, list, count
}

func GetOrderProductInfo(db *gorm.DB, subIds []uint64) (error,
	*[]model.OrderDetail,
	*[]model.OrderProduct) {

	var (
		err           error
		orderDetails  []model.OrderDetail
		orderProducts []model.OrderProduct
	)
	//查询订单详情

	err = db.Where("order_id in (?)", subIds).Find(&orderDetails).Error
	if err != nil {
		return err, nil, nil
	}

	//查询订单商品
	err = db.Where("order_id in (?)", subIds).Find(&orderProducts).Error
	if err != nil {
		return err, nil, nil
	}

	//TODO:查询订单支付信息 (已支付情况)

	return nil, &orderDetails, &orderProducts

}

func LoadOrderProductInfo(subOrderId uint64,
	orderDetails *[]model.OrderDetail,
	orderProducts *[]model.OrderProduct,
	imgService string) *[]SmallOrderInfoProductModel {

	var (
		data []SmallOrderInfoProductModel
	)

	for _, o := range *orderDetails {

		if o.OrderId != subOrderId {
			continue
		}

		d := SmallOrderInfoProductModel{
			Status:      o.Status,
			No:          strconv.Itoa(int(o.ProductId)),
			UnitPrice:   o.ProductUnitPrice,
			Count:       o.ProductCount,
			SkuInfo:     o.ProductSkuAttributeInfo,
			AmountTotal: o.AmountTotal,
		}

		for _, p := range *orderProducts {
			if p.ProductId == o.ProductId && o.OrderId == p.OrderId {
				d.Name = p.ProductName
				d.Cover = imgService + p.CoverImage
				break
			}
		}
		data = append(data, d)
	}

	return &data
}

//设置订单业务状态
func loadOrderBusinessStatus(list []SmallOrderInfModel) {
	for i := 0; i < len(list); i++ {
		list[i].Status = string(GetOrderBusinessStatus(list[i].OrderStatus, list[i].PayStatus, list[i].DeliveryStatus, list[i].Type))
	}
}

func LoadQueryStatusSqlString(status orderEnum.OrderBusinessStatus) string {
	return loadQueryStatusSqlString(status)
}

func loadQueryStatusSqlString(status orderEnum.OrderBusinessStatus) string {
	sql := "  "
	switch status {
	case orderEnum.OrderBusinessStatusWaitPay:
		sql += fmt.Sprintf(" ( order_status = %v and pay_status= %v and delivery_status = %v and type = %v )",
			int(orderEnum.OrderCreate), int(orderEnum.PayWait), int(orderEnum.DeliveryDefaultStatus), int(orderEnum.OrderTypeOnline))
		break
	case orderEnum.OrderBusinessStatusWaitProcess:
		sql += fmt.Sprintf(" ( order_status = %v and pay_status= %v and delivery_status = %v and type = %v )",
			int(orderEnum.OrderCreate), int(orderEnum.PayWait), int(orderEnum.DeliveryDefaultStatus), int(orderEnum.OrderTypeOffline))
		break
	case orderEnum.OrderBusinessStatusPaySuccess:
		sql += fmt.Sprintf(" ( order_status = %v and pay_status= %v and delivery_status = %v )",
			int(orderEnum.OrderCreate), int(orderEnum.PaySuccess), int(orderEnum.DeliveryWait))
		break
	case orderEnum.OrderBusinessStatusShip:
		sql += fmt.Sprintf(" ( order_status = %v and pay_status= %v and delivery_status = %v )",
			int(orderEnum.OrderCreate), int(orderEnum.PaySuccess), int(orderEnum.DeliveryShip))
		break
	case orderEnum.OrderBusinessStatusOrderCancel:
		sql += fmt.Sprintf(" ( order_status = %v and pay_status= %v and delivery_status = %v )",
			int(orderEnum.OrderCancel), int(orderEnum.PayWait), int(orderEnum.DeliveryDefaultStatus))
		break
	case orderEnum.OrderBusinessStatusFinish: //and pay_status= %v and delivery_status = %v
		sql += fmt.Sprintf(" ( order_status = %v )",
			int(orderEnum.OrderSuccess))
		break
	default:
		return ""
	}

	return sql
}

//获取业务订单状态
func GetOrderBusinessStatus(orderStatus, payStatus, deliveryStatus, orderType int) orderEnum.OrderBusinessStatus {
	if orderStatus == int(orderEnum.OrderCreate) && payStatus == int(orderEnum.PayWait) && deliveryStatus == int(orderEnum.DeliveryDefaultStatus) && orderType == 1 {
		return orderEnum.OrderBusinessStatusWaitPay
	}

	if orderStatus == int(orderEnum.OrderCreate) && payStatus == int(orderEnum.PayFail) && orderType == 1 {
		return orderEnum.OrderBusinessStatusWaitPayFAIL
	}

	if orderStatus == int(orderEnum.OrderCreate) && payStatus == int(orderEnum.PayWait) && deliveryStatus == int(orderEnum.DeliveryDefaultStatus) && orderType == 0 {
		return orderEnum.OrderBusinessStatusWaitProcess
	}

	if orderStatus == int(orderEnum.OrderCreate) && payStatus == int(orderEnum.PaySuccess) && deliveryStatus == int(orderEnum.DeliveryWait) {
		return orderEnum.OrderBusinessStatusPaySuccess
	}

	if orderStatus == int(orderEnum.OrderCreate) && payStatus == int(orderEnum.PaySuccess) && deliveryStatus == int(orderEnum.DeliveryShip) {
		return orderEnum.OrderBusinessStatusShip
	}

	if orderStatus == int(orderEnum.OrderCreate) && payStatus == int(orderEnum.PaySuccess) && deliveryStatus == int(orderEnum.DeliveredSuccess) {
		return orderEnum.OrderBusinessStatusDelivered
	}

	//订单完成这里分两种情况，一种是需要支付类型的订单，一种是无需支付的类型
	if orderStatus == int(orderEnum.OrderSuccess) {
		return orderEnum.OrderBusinessStatusFinish
	}

	if orderStatus == int(orderEnum.OrderCancel) {
		return orderEnum.OrderBusinessStatusOrderCancel
	}

	if orderStatus == int(orderEnum.OrderCancel) && payStatus == int(orderEnum.PayWait) && deliveryStatus == int(orderEnum.DeliveryDefaultStatus) {
		return orderEnum.OrderBusinessStatusOrderCancel
	}

	return orderEnum.OrderBusinessStatus(0)
}
