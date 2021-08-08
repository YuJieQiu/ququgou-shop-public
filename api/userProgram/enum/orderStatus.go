package enum

//订单前台显示状态
type OrderDisplayStatus int

const (
	OrderDisplayStatusAll OrderDisplayStatus = iota //0  全部
	_
	_
	OrderDisplayStatusWaitPay //3  待支付
	_
	OrderDisplayStatusWaitProcess //5  待完成
	_
	OrderDisplayStatusFinish //7  已完成
	_
	OrderDisplayStatusCancel //9  已取消

)
