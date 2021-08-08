package paymentEnum

//交易状态枚举
type TradeType int

const (
	TradeTypeOrderPay TradeType = 1 //订单支付

	TradeTypeWithdraw = 10 //提现
)
