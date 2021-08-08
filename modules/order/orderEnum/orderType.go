package orderEnum

//订单类型
type OrderType int

const (
	OrderTypeOnline OrderType = 0 //线上订单

	OrderTypeOffline = 1 //线下订单

)
