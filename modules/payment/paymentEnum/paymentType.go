package paymentEnum

//支付类型枚举
type PaymentType string

//状态 0 未开始 1 进行中  2 已结束\
const (
	PaymentTypeWeChatPay  PaymentType = "WeChatPay"  //微信支付
	PaymentTypeOfflinePay             = "OfflinePay" //线下支付
)
