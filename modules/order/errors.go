package order

import "errors"

var (
	//没有订单记录
	ErrOrderRecordNotFound = errors.New("order record not found")

	//订单状态无效
	ErrOrderStatusInvalid = errors.New("order status invalid")

	//订单超过最后支付时间
	ErrOrderOutExpireTime = errors.New("order out expire time")

	//订单支付金额不正确
	ErrOrderPayAmount = errors.New("expire order pay amount")
)
