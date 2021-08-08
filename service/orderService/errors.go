package orderService

import "errors"

var (
	//
	ErrProductPriceChange = errors.New("product price change ")

	//商品库存不够
	ErrProductStockInvalid = errors.New("product stock invalid ")

	//无效的产品
	ErrProductInvalid = errors.New("product invalid ")

	//商品总价无效
	ErrProductAmountPriceInvalid = errors.New("product amount price invalid ")

	//商品优惠总金额无效
	ErrDiscountsAmountPriceInvalid = errors.New("Discounts amount price invalid ")

	//订单总价无效
	ErrOrderAmountPriceInvalid = errors.New("order amount price invalid ")

	//商家无效
	ErrMerchantsInvalid = errors.New("merchants invalid ")

	//订单状态无效
	ErrOrderStatusInvalid = errors.New("order status invalid ")

	//订单支付方式无效
	ErrOrderPaymentTypeInvalid = errors.New("order payment type invalid ")
)
