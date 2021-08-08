package common

import (
	"strconv"
	"time"
)

//生成订单编号
func CreateOrderNo(id uint64) string {
	// x - xxxxxx - xxxxx -xxx -xxx

	// x 业务编号 1 下单
	// xxxxxx  年月日
	// xxxxx 时间戳 最后 10位数
	// xxx 用户guid后两位id

	no := "1"

	t := time.Now()

	nt := string(t.Format("060102")) // xxxxxx  年月日
	timeUnixNano := int(t.UnixNano())
	tunS := strconv.Itoa(timeUnixNano)[6:16]

	no += nt + strconv.Itoa(int(id)) + tunS

	return no
}

//生成子订单编号
func CreateSubOrderNo(merId uint64, uId uint64) string {
	// x - xxxxxx - xxxxx -xxx -xxx

	// x 业务编号 3 子订单
	// xxxxxx  年月日
	// xxxxx 时间戳 最后 10位数
	// xxx 用户guid后两位id

	no := "3"

	t := time.Now()

	nt := string(t.Format("060102")) // xxxxxx  年月日
	timeUnixNano := int(t.UnixNano())
	tunS := strconv.Itoa(timeUnixNano)[6:16]

	no += nt + strconv.Itoa(int(merId)) + strconv.Itoa(int(uId)) + tunS

	return no
}
