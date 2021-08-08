package paymentEnum

//交易状态枚举
type TradeStatus int

const (
	TradeStatusProcessed TradeStatus = 0  // 		处理中 默认
	TradeStatusWaitPay               = 3  // 	    等待支付
	TradeStatusSucceed               = 5  //		交易成功
	TradeStatusFail                  = -1 //		交易失败
	TradeStatusCancel                = -3 //		交易取消
)
