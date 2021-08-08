package orderEnum

//订单业务状态枚举 (注意这里是订单业务状态，不等于前台显示的状态)
type OrderBusinessStatus string

//订单业务状态 规则  xxxx 四位数字
//第一位 OrderBusinessStatus
//第二位 PayStatus
//第三位 DeliveredStatus
//第四位 自定义业务 如:待付款 和 待处理 前三位状态一样， 第四位判断(一般type 值)

const (
	OrderBusinessStatusWaitPay     OrderBusinessStatus = "0000" // 待支付/待付款/已提交......
	OrderBusinessStatusWaitProcess                     = "0001" //			未支付 待完成(线下订单)
	OrderBusinessStatusPaySuccess                      = "0910" //			付款成功
	OrderBusinessStatusShip                            = "0930" //			已发货
	OrderBusinessStatusDelivered                       = "0990" //			已签收(已完成)
	OrderBusinessStatusFinish                          = "9990" //订单已完成
	//OrderBusinessStatusNotPayFinish                     = "9000"  //订单交易完成(无需支付的订单类型)
	OrderBusinessStatusOrderCancel = "-1000" //			取消交易(已取消)
	OrderBusinessStatusWaitPayFAIL = "-0200" //			支付失败
	//OrderBusinessStatusRefundApply   = -1   //			推销
	//OrderBusinessStatusRefund        = -3   //			退货中
	//OrderBusinessStatusRefundSuccess = -5   //			已退货
	//OrderBusinessStatusApplyCancel   = -9   //			撤销申请
)

func (s OrderBusinessStatus) String() string {
	switch s {
	case OrderBusinessStatusWaitPay:
		return "WaitPay"
	case OrderBusinessStatusWaitProcess:
		return "WaitProcess"
	case OrderBusinessStatusPaySuccess:
		return "PaySuccess"
	case OrderBusinessStatusShip:
		return "Ship"
	case OrderBusinessStatusDelivered:
		return "Delivered"
	case OrderBusinessStatusFinish:
		return "Finish"
	//case OrderBusinessStatusNotPayFinish:
	//	return "NotPayOrderFinish"
	case OrderBusinessStatusOrderCancel:
		return "OrderCancel"
	case OrderBusinessStatusWaitPayFAIL:
		return "PayFAIL"

	//case OrderBusinessStatusRefundApply:
	//	return "RefundApply"
	//case OrderBusinessStatusRefund:
	//	return "Refund"
	//case OrderBusinessStatusRefundSuccess:
	//	return "RefundSuccess"
	//case OrderBusinessStatusApplyCancel:
	//	return "ApplyCancel"
	default:
		return ""
	}
	//return [...]string{"WaitPay", "PaySuccess", "Ship", "Delivered", "RefundApply", "Refund", "RefundSuccess", "OrderCancel", "ApplyCancel"}[s]
}

func (s OrderBusinessStatus) Text() string {
	switch s {
	case OrderBusinessStatusWaitPay:
		return "待付款"
	case OrderBusinessStatusWaitProcess:
		return "待完成"
	case OrderBusinessStatusPaySuccess:
		return "付款成功"
	case OrderBusinessStatusShip:
		return "已发货"
	case OrderBusinessStatusDelivered:
		return "已签收"
	case OrderBusinessStatusFinish:
		return "已完成"
	//case OrderBusinessStatusNotPayFinish:
	//	return "已完成" //无需支付类型的订单
	case OrderBusinessStatusOrderCancel:
		return "取消交易"
	case OrderBusinessStatusWaitPayFAIL:
		return "支付失败"

	//case OrderBusinessStatusRefundApply:
	//	return "退货申请"
	//case OrderBusinessStatusRefund:
	//	return "退货中"
	//case OrderBusinessStatusRefundSuccess:
	//	return "已退货"
	//case OrderBusinessStatusApplyCancel:
	//	return "撤销申请"
	default:
		return ""
	}

	//return [...]string{"待付款", "付款成功", "已发货", "已签收", "退货申请", "退货中", "已退货", "取消交易", "撤销申请"}[s]
}
