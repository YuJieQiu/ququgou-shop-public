package orderEnum

//订单状态枚举
type OrderStatus int

//支付状态
type PayStatus int

//物流/配送/交付 状态
type DeliveryStatus int

const (
	OrderCreate OrderStatus = 0 //从0 开始 增长 订单提交/创建

	OrderSuccess = 9 //订单完成

	OrderCancel = -1 //订单取消

	OrderServiceApply = -3 //售后服务申请

	OrderServiceSuccess = -9 //售后服务完成

)

const (
	PayWait PayStatus = 0 //从0 开始 增长 等待支付

	PayProcess = 1 //支付处理中

	PaySuccess = 9 //支付成功

	PayCancel = -1 //取消支付

	PayFail = -2 //支付失败

	PayRefundApply = -5 // 退款申请/处理中

	PayRefundFail = -6 // 退款失败

	PayRefundSuccess = -9 //退款成功
)

const (
	DeliveryDefaultStatus DeliveryStatus = 0 //从0 开始 增长 默认状态 无意义

	DeliveryWait = 1 //等待发货

	DeliveryShip = 3 //已发货

	DeliveredSuccess = 9 //已签收/已完成
)
