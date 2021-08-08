package paymentEnum

//交易表中的业务类型
type BusinessType int

const (
	BusinessTypeOrderMaster BusinessType = 1 //原始订单
	BusinessTypeOrder                    = 2 //订单
)

func (s BusinessType) String() string {
	switch s {
	case BusinessTypeOrderMaster:
		return "OrderMaster"
	case BusinessTypeOrder:
		return "Order"
	default:
		return ""
	}

}
